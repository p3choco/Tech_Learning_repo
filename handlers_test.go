package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup() *echo.Echo {
	DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	DB.AutoMigrate(&Category{}, &Product{}, &Cart{})

	e := echo.New()
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

	return e
}

func TestProductHandlers(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := getProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]", strings.TrimSpace(rec.Body.String()))

	p := Product{Name: "X", Price: 9.99, CategoryID: 0}
	b, _ := json.Marshal(p)
	req = httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = createProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var outP Product
	json.Unmarshal(rec.Body.Bytes(), &outP)
	assert.Equal(t, "X", outP.Name)
	assert.Equal(t, 9.99, outP.Price)

	req = httptest.NewRequest(http.MethodPost, "/products", strings.NewReader("notjson"))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = createProduct(c)
	he, ok := err.(*echo.HTTPError)
	if assert.True(t, ok, "create powinien zwracać *echo.HTTPError przy złym JSON") {
		assert.Equal(t, http.StatusBadRequest, he.Code)
	}

	req = httptest.NewRequest(http.MethodGet, "/products/"+strconv.Itoa(int(outP.ID)), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(outP.ID)))
	err = getProductByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	outP.Name = "Y"
	b, _ = json.Marshal(outP)
	req = httptest.NewRequest(http.MethodPut, "/products/"+strconv.Itoa(int(outP.ID)), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(outP.ID)))
	err = updateProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	var upP Product
	json.Unmarshal(rec.Body.Bytes(), &upP)
	assert.Equal(t, "Y", upP.Name)

	req = httptest.NewRequest(http.MethodPut, "/products/"+strconv.Itoa(int(outP.ID)), strings.NewReader("notjson"))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(outP.ID)))
	err = updateProduct(c)
	he, ok = err.(*echo.HTTPError)
	if assert.True(t, ok, "update powinien zwracać *echo.HTTPError przy złym JSON") {
		assert.Equal(t, http.StatusBadRequest, he.Code)
	}

	req = httptest.NewRequest(http.MethodPut, "/products/9999", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9999")
	err = updateProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	req = httptest.NewRequest(http.MethodGet, "/products/filter?minPrice=5", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = filterProducts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	req = httptest.NewRequest(http.MethodDelete, "/products/"+strconv.Itoa(int(outP.ID)), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(outP.ID)))
	err = deleteProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)

	req = httptest.NewRequest(http.MethodDelete, "/products/9999", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9999")
	err = deleteProduct(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCartHandlers(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/carts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := getCarts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]", strings.TrimSpace(rec.Body.String()))

	ct := Cart{UserID: 42, CartValue: 123.45}
	b, _ := json.Marshal(ct)
	req = httptest.NewRequest(http.MethodPost, "/carts", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = createCart(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var outC Cart
	json.Unmarshal(rec.Body.Bytes(), &outC)
	assert.Equal(t, uint(42), outC.UserID)
	assert.Equal(t, 123.45, outC.CartValue)

	req = httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader("notjson"))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = createCart(c)
	he, ok := err.(*echo.HTTPError)
	if assert.True(t, ok, "ostatni błąd powinien być *echo.HTTPError") {
		assert.Equal(t, http.StatusBadRequest, he.Code)
	}
}

func TestCategoryHandlers(t *testing.T) {
	e := setup()

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := getCategories(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]", strings.TrimSpace(rec.Body.String()))

	cat := Category{Name: "Foo"}
	b, _ := json.Marshal(cat)
	req = httptest.NewRequest(http.MethodPost, "/categories", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = createCategory(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var outCat Category
	json.Unmarshal(rec.Body.Bytes(), &outCat)
	assert.Equal(t, "Foo", outCat.Name)

	req = httptest.NewRequest(http.MethodGet, "/categories/"+strconv.Itoa(int(outCat.ID)), nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(outCat.ID)))
	err = getCategoryByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	req = httptest.NewRequest(http.MethodGet, "/categories/9999", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9999")
	err = getCategoryByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	req = httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader("bad"))
	req.Header.Set("Content-Type", "application/json")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	err = createCategory(c)
	he, ok := err.(*echo.HTTPError)
	if assert.True(t, ok, "ostatni błąd powinien być *echo.HTTPError") {
		assert.Equal(t, http.StatusBadRequest, he.Code)
	}
}

func TestScopes(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Product{})
	db.Create(&Product{Name: "Cheap", Price: 1})
	db.Create(&Product{Name: "Expensive", Price: 100})
	var out []Product
	db.Scopes(ScopeMinPrice(50)).Find(&out)
	assert.Len(t, out, 1)
	assert.Equal(t, "Expensive", out[0].Name)

	db.Create(&Product{Name: "CatA", Price: 5, CategoryID: 7})
	db.Create(&Product{Name: "CatB", Price: 5, CategoryID: 8})
	out = []Product{}
	db.Scopes(ScopeCategoryID(7)).Find(&out)
	assert.Len(t, out, 1)
	assert.Equal(t, uint(7), out[0].CategoryID)
}