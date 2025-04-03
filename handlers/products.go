package handlers

import (
	"golang_auth/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang_auth/database"
)

// GetProducts retrieves all products from the database
func GetProducts(c *gin.Context) {
	var products []models.Product

	// Fetch all products from the database
	if err := db.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}

	// Return the list of products
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}
