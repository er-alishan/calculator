package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a router
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/add", addHandler)
	router.GET("/sub", subHandler)
	router.GET("/mul", mulHandler)
	router.GET("/div", divHandler)
	return router
}

// Separate handlers for testing
func addHandler(c *gin.Context) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter a"})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter b"})
		return
	}

	result := add(a, b)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func subHandler(c *gin.Context) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter a"})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter b"})
		return
	}

	result := sub(a, b)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func mulHandler(c *gin.Context) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter a"})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter b"})
		return
	}

	result := mul(a, b)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func divHandler(c *gin.Context) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter a"})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter b"})
		return
	}

	result := div(a, b)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func TestAdd(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add?a=5&b=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"result": 8}`, w.Body.String())
}

func TestSub(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sub?a=5&b=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"result": 2}`, w.Body.String())
}

func TestMul(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mul?a=5&b=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"result": 15}`, w.Body.String())
}

func TestDiv(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/div?a=6&b=3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"result": 2}`, w.Body.String())
}
