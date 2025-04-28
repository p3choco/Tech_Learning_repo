package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

func main() {
	initDB()
	e := echo.New()

    e.GET("/", func(c echo.Context) error {
    html := `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="dummy" content="<!doctype html>">
  <title>Example Domain</title>
  <script>
    if (window.history && window.history.replaceState) {
      var path = window.location.pathname;
      if (path !== '/') {
        var newPath = path.replace(/\/+$/, '');
        history.replaceState({}, document.title, window.location.origin + newPath);
      }
    }
  </script>
  <style>
      body { background: red; }
  #container { background-color: transparent; }
  </style>
</head>
<body>
  <div id="container">
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents.</p>
    <a href="https://www.iana.org/domains/example?flag=.example">More information...</a>
  </div>
</body>
</html>`
    return c.HTML(http.StatusOK, html)
    })

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProductByID)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.GET("/products/filter", filterProducts)

	e.GET("/carts", getCarts)
	e.POST("/carts", createCart)

	e.GET("/categories", getCategories)
	e.GET("/categories/:id", getCategoryByID)
	e.POST("/categories", createCategory)

	e.Logger.Fatal(e.Start(":1323"))
}

func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("example.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Category{}, &Product{}, &Cart{})
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
		return err
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
		return err
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
		return err
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
		return err
	}
	if err := DB.Create(&category).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, category)
}
