package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

var googleOauthConfig *oauth2.Config
var githubOauthConfig *oauth2.Config

var jwtSecret = []byte("")

type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"-"`
	Name         string `json:"name"`
	OAuthID      string `json:"oauth_id" gorm:"column:o_auth_id"`
	OAuthType    string `json:"oauth_type" gorm:"column:o_auth_type"` // "google", "github", or empty for local
	RefreshToken string `json:"-"`
}

type UserSession struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	Token     string    `json:"token" gorm:"unique"`
	ExpiresAt time.Time `json:"expires_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type Product struct {
	gorm.Model
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
}

type Cart struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	CartValue float64 `json:"cart_value"`
}

type Payment struct {
	gorm.Model
	CustomerName  string        `json:"name"`
	CustomerEmail string        `json:"email"`
	Total         float64       `json:"total"`
	Items         []PaymentItem `json:"items" gorm:"foreignKey:PaymentID"`
}

type PaymentItem struct {
	gorm.Model
	PaymentID uint    `json:"-"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Qty       uint    `json:"qty"`
}

func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("example.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Category{}, &Product{}, &Cart{}, &Payment{}, &PaymentItem{}, &User{}, &UserSession{})
}

func initOauth() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	githubOauthConfig = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:8080/auth/github/callback",
		Scopes:       []string{"user:email", "read:user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
}

func main() {
	initDB()
	initOauth()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	auth := e.Group("")
	auth.Use(authMiddleware)

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProductByID)
	e.POST("/products", createProduct)
	auth.PUT("/products/:id", updateProduct)
	auth.DELETE("/products/:id", deleteProduct)
	e.GET("/products/filter", filterProducts)

	auth.GET("/carts", getCarts)
	auth.POST("/carts", createCart)

	e.GET("/categories", getCategories)
	e.GET("/categories/:id", getCategoryByID)
	e.POST("/categories", createCategory)

	auth.POST("/payments", createPayment)

	e.POST("/login", login)
	e.POST("/register", register)
	auth.GET("/user", getCurrentUser)

	e.GET("/auth/google/login", googleLogin)
	e.GET("/auth/google/callback", googleCallback)
	e.GET("/auth/github/login", githubLogin)
	e.GET("/auth/github/callback", githubCallback)

	e.Logger.Fatal(e.Start(":8080"))
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		fmt.Println("Auth header received:", tokenString)

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing authorization token"})
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		fmt.Println("Token after Bearer removal:", tokenString[:15]+"...")

		var session UserSession
		if err := DB.Where("token = ? AND expires_at > ?", tokenString, time.Now()).First(&session).Error; err != nil {
			fmt.Println("Error finding session:", err.Error())
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired token"})
		}

		fmt.Println("Valid session found for user ID:", session.UserID)
		c.Set("userID", session.UserID)
		return next(c)
	}
}

func googleLogin(c echo.Context) error {
	stateToken := generateStateToken()
	url := googleOauthConfig.AuthCodeURL(stateToken)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func googleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	fmt.Println("Google callback received with code:", code != "")

	token, err := googleOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		fmt.Println("Error exchanging code for token:", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to exchange code for token"})
	}

	fmt.Println("Successfully obtained OAuth token")

	userInfo, err := getUserInfoFromGoogle(token.AccessToken)
	if err != nil {
		fmt.Println("Error getting user info:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get user info from Google"})
	}

	fmt.Println("User info obtained:", userInfo["email"], userInfo["name"])

	var user User
	if err := DB.Where("o_auth_id = ? AND o_auth_type = ?", userInfo["id"], "google").First(&user).Error; err != nil {
		var existingUser User
		if err := DB.Where("email = ?", userInfo["email"]).First(&existingUser).Error; err == nil {
			fmt.Println("Existing user found with email:", existingUser.Email)
			existingUser.OAuthID = userInfo["id"].(string)
			existingUser.OAuthType = "google"
			existingUser.RefreshToken = token.RefreshToken
			if err := DB.Save(&existingUser).Error; err != nil {
				fmt.Println("Error updating user:", err)
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update user: " + err.Error()})
			}
			user = existingUser
		} else {
			fmt.Println("Creating new user with email:", userInfo["email"])
			user = User{
				Email:        userInfo["email"].(string),
				Name:         userInfo["name"].(string),
				OAuthID:      userInfo["id"].(string),
				OAuthType:    "google",
				RefreshToken: token.RefreshToken,
			}

			if err := DB.Create(&user).Error; err != nil {
				fmt.Println("Error creating user:", err)
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user: " + err.Error()})
			}
		}
	} else {
		fmt.Println("Existing OAuth user found:", user.Email)
	}

	sessionToken, err := generateToken(user.ID)
	if err != nil {
		fmt.Println("Error generating session token:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate session token"})
	}

	fmt.Println("Generated session token, redirecting to frontend")

	redirectURL := fmt.Sprintf("http://localhost:3000/auth/callback?token=%s", url.QueryEscape(sessionToken))
	fmt.Println("Redirecting to:", redirectURL)
	return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func githubLogin(c echo.Context) error {
	stateToken := generateStateToken()
	url := githubOauthConfig.AuthCodeURL(stateToken)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func githubCallback(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := githubOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Failed to exchange code for token"})
	}

	userInfo, err := getUserInfoFromGithub(token.AccessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get user info from Github"})
	}

	var user User
	if err := DB.Where("o_auth_id = ? AND o_auth_type = ?", userInfo["id"], "github").First(&user).Error; err != nil {
		var existingUser User
		if err := DB.Where("email = ?", userInfo["email"]).First(&existingUser).Error; err == nil {
			existingUser.OAuthID = fmt.Sprintf("%v", userInfo["id"])
			existingUser.OAuthType = "github"
			existingUser.RefreshToken = token.RefreshToken
			if err := DB.Save(&existingUser).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update user: " + err.Error()})
			}
			user = existingUser
		} else {
			user = User{
				Email:        userInfo["email"].(string),
				Name:         userInfo["name"].(string),
				OAuthID:      fmt.Sprintf("%v", userInfo["id"]),
				OAuthType:    "github",
				RefreshToken: token.RefreshToken,
			}

			if err := DB.Create(&user).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user: " + err.Error()})
			}
		}
	}

	sessionToken, err := generateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate session token"})
	}

	return c.Redirect(http.StatusTemporaryRedirect,
		fmt.Sprintf("http://localhost:3000/auth/callback?token=%s", url.QueryEscape(sessionToken)))
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func getUserInfoFromGoogle(accessToken string) (map[string]interface{}, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func getUserInfoFromGithub(accessToken string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	if _, ok := userInfo["email"]; !ok || userInfo["email"] == nil {
		emailReq, err := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
		if err == nil {
			emailReq.Header.Set("Authorization", "token "+accessToken)
			emailResp, err := client.Do(emailReq)
			if err == nil {
				defer emailResp.Body.Close()
				var emails []map[string]interface{}
				if err := json.NewDecoder(emailResp.Body).Decode(&emails); err == nil {
					for _, email := range emails {
						if primary, ok := email["primary"].(bool); ok && primary {
							userInfo["email"] = email["email"]
							break
						}
					}
				}
			}
		}
	}

	return userInfo, nil
}

func getProducts(c echo.Context) error {
	var products []Product
	if err := DB.Preload("Category").Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func getProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	if err := DB.Preload("Category").First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	var newProduct Product
	if err := c.Bind(&newProduct); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := DB.Create(&newProduct).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	if err := DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func filterProducts(c echo.Context) error {
	var products []Product
	minPriceStr := c.QueryParam("minPrice")
	categoryIDStr := c.QueryParam("categoryID")

	dbQuery := DB.Preload("Category").Model(&Product{})
	if minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			dbQuery = dbQuery.Scopes(ScopeMinPrice(minPrice))
		}
	}
	if categoryIDStr != "" {
		if catID, err := strconv.ParseUint(categoryIDStr, 10, 64); err == nil {
			dbQuery = dbQuery.Scopes(ScopeCategoryID(uint(catID)))
		}
	}
	if err := dbQuery.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func getCarts(c echo.Context) error {
	var carts []Cart
	if err := DB.Find(&carts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, carts)
}

func createCart(c echo.Context) error {
	var cart Cart
	if err := c.Bind(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := DB.Create(&cart).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, cart)
}

func getCategories(c echo.Context) error {
	var categories []Category
	if err := DB.Preload("Products").Find(&categories).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, categories)
}

func getCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category Category
	if err := DB.Preload("Products").First(&category, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
	}
	return c.JSON(http.StatusOK, category)
}

func createCategory(c echo.Context) error {
	var category Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, category)
}

func createPayment(c echo.Context) error {
	var payload struct {
		Customer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"customer"`
		Items []struct {
			ID    uint    `json:"product_id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
			Qty   uint    `json:"qty"`
		} `json:"items"`
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Niepoprawny format danych"})
	}

	var total float64
	for _, it := range payload.Items {
		total += it.Price * float64(it.Qty)
	}

	payment := Payment{
		CustomerName:  payload.Customer.Name,
		CustomerEmail: payload.Customer.Email,
		Total:         total,
	}
	for _, it := range payload.Items {
		payment.Items = append(payment.Items, PaymentItem{
			ProductID: it.ID,
			Name:      it.Name,
			Price:     it.Price,
			Qty:       it.Qty,
		})
	}

	if err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Płatność przyjęta",
		"payment": payment,
	})
}

func ScopeMinPrice(min float64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price >= ?", min)
	}
}

func ScopeCategoryID(catID uint) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", catID)
	}
}

func login(c echo.Context) error {
	var request LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Niepoprawny format danych"})
	}

	var user User
	if err := DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Nieprawidłowy email lub hasło"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Nieprawidłowy email lub hasło"})
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Błąd generowania tokenu"})
	}

	return c.JSON(http.StatusOK, AuthResponse{
		Token: token,
		User:  user,
	})
}

func register(c echo.Context) error {
	var request RegisterRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Niepoprawny format danych"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Błąd hashowania hasła"})
	}

	user := User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Błąd rejestracji"})
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Błąd generowania tokenu"})
	}

	return c.JSON(http.StatusCreated, AuthResponse{
		Token: token,
		User:  user,
	})
}

func generateToken(userID uint) (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	tokenStr := base64.StdEncoding.EncodeToString(token)

	session := UserSession{
		UserID:    userID,
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := DB.Create(&session).Error; err != nil {
		return "", err
	}

	return tokenStr, nil
}

func getCurrentUser(c echo.Context) error {
	userID := c.Get("userID").(uint)

	var user User
	if err := DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}
