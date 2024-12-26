package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}

func main() {
	router := gin.Default()
	//Add
	router.GET("/add", func(c *gin.Context) {
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
	})
	//subtract
	router.GET("/sub", func(c *gin.Context) {
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
	})
	// multiply
	router.GET("/mul", func(c *gin.Context) {
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
	})
	// divide
	router.GET("/div", func(c *gin.Context) {
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
	})
	router.Run(":8080")
}
