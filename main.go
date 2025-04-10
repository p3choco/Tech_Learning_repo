package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	products   []Product
	nextProdID = 1
)

func main() {
	e := echo.New()

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProductByID)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start(":1323"))
}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, p := range products {
		if p.ID == id {
			return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
}

func createProduct(c echo.Context) error {
	var newProduct Product
	if err := c.Bind(&newProduct); err != nil {
		return err
	}
	newProduct.ID = nextProdID
	nextProdID++
	products = append(products, newProduct)

	return c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var updated Product
	if err := c.Bind(&updated); err != nil {
		return err
	}

	for i, p := range products {
		if p.ID == id {
			products[i].Name = updated.Name
			return c.JSON(http.StatusOK, products[i])
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
}
