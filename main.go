package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Cart struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	CartValue float64 `json:"cart_value"`
}

func main() {
	initDB()

	e := echo.New()

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProductByID)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.GET("/carts", getCarts)
	e.POST("/carts", createCart)

	e.Logger.Fatal(e.Start(":1323"))
}

func initDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("example.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Product{}, &Cart{})
}

func getProducts(c echo.Context) error {
	var products []Product
	result := DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func getProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product Product
	result := DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	var newProduct Product
	if err := c.Bind(&newProduct); err != nil {
		return err
	}
	result := DB.Create(&newProduct)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
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
