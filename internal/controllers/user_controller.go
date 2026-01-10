package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func Collections(c *gin.Context) {
	var sneakers []models.Sneaker

	result := database.DB.Find(&sneakers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching products",
		})
		return
	}

	// Optional: handle empty table (not an error)
	if len(sneakers) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":     "No products found",
			"collections": []models.Sneaker{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"collections": sneakers,
	})
}
