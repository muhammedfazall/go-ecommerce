package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"gorm.io/gorm"
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

func ProductDetails(c *gin.Context) {
	id := c.Param("id")

	var sneaker models.Sneaker

	result := database.DB.First(&sneaker, id)

	if result.Error != nil {

		// record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
			return
		}

		// other DB error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": sneaker,
	})
}
