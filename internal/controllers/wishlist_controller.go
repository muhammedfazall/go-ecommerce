package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

// add to wishlist
func AddToWishlist(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req struct {
		ProductID uint `json:"product_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// validate product exists
	var sneaker models.Sneaker
	if err := database.DB.First(&sneaker, req.ProductID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	wishlist := models.Wishlist{
		UserID:    userID,
		SneakerID: req.ProductID,
	}

	if err := database.DB.Create(&wishlist).Error; err != nil {
		c.JSON(400, gin.H{"error": "Already in wishlist"})
		return
	}

	c.JSON(200, gin.H{"message": "Added to wishlist"})
}

//View Wishlist
func GetWishlist(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var wishlist []models.Wishlist
	if err := database.DB.
		Preload("Sneaker").
		Where("user_id = ?", userID).
		Find(&wishlist).Error; err != nil {

		c.JSON(500, gin.H{"error": "Failed to fetch wishlist"})
		return
	}

	c.JSON(200, gin.H{
		"items": wishlist,
	})
}

// remove from wishlist
func RemoveFromWishlist(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	productID := c.Param("productId")

	result := database.DB.
		Where("user_id = ? AND sneaker_id = ?", userID, productID).
		Delete(&models.Wishlist{})

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Item not found in wishlist"})
		return
	}

	c.JSON(200, gin.H{"message": "Removed from wishlist"})
}
