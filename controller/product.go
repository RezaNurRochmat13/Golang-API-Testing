package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllProducts func does get all products
func GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": 1,
		"data":  "All products",
		"count": 1})
}

// GetDetailProducts
func GetDetailProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Detail products"})
}
