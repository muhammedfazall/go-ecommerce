package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

// add item to cart
func AddToCart(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Quantity < 1 {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// validate product exists
	var product models.Sneaker
	if err := database.DB.First(&product, req.ProductID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	var cart models.Cart
	if err := database.DB.FirstOrCreate(&cart, models.Cart{UserID: userID}).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create cart"})
		return
	}

	var item models.CartItem
	err := database.DB.
		Where("cart_id = ? AND sneaker_id = ?", cart.ID, req.ProductID).
		First(&item).Error

	if err == nil {
		item.Quantity += req.Quantity
		database.DB.Save(&item)
	} else {
		database.DB.Create(&models.CartItem{
			CartID:    cart.ID,
			SneakerID: req.ProductID,
			Quantity:  req.Quantity,
		})
	}

	c.JSON(200, gin.H{"message": "Added to cart"})
}

// view cart
func GetCart(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var cart models.Cart
	err := database.DB.
		Preload("Items.Sneaker").
		Where("user_id = ?", userID).
		First(&cart).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{"items": []interface{}{}, "total": 0})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to load cart"})
		return
	}

	var total float64
	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Sneaker.Price
	}

	c.JSON(200, gin.H{
		"cart_id": cart.ID,
		"items":   cart.Items,
		"total":   total,
	})
}

// update cart item
func UpdateCartItem(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Quantity < 1 {
		c.JSON(400, gin.H{"error": "Invalid quantity"})
		return
	}

	var cart models.Cart
	database.DB.Where("user_id = ?", userID).First(&cart)

	result := database.DB.Model(&models.CartItem{}).
		Where("cart_id = ? AND sneaker_id = ?", cart.ID, req.ProductID).
		Update("quantity", req.Quantity)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Item not found in cart"})
		return
	}

	c.JSON(200, gin.H{"message": "Cart updated"})
}

// remove from cart
func RemoveFromCart(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	productID := c.Param("productId")

	var cart models.Cart
	database.DB.Where("user_id = ?", userID).First(&cart)

	result := database.DB.
		Where("cart_id = ? AND sneaker_id = ?", cart.ID, productID).
		Delete(&models.CartItem{})

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Item removed"})
}

// clear cart
func ClearCart(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var cart models.Cart
	database.DB.Where("user_id = ?", userID).First(&cart)

	if err := database.DB.Where("cart_id = ?", cart.ID).
		Delete(&models.CartItem{}).Error; err != nil {

		c.JSON(500, gin.H{"error": "Failed to clear cart"})
		return
	}

	c.JSON(200, gin.H{"message": "Cart cleared"})
}
