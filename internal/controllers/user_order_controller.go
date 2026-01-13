package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

// place order
func PlaceOrder(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	// 1. Load cart
	var cart models.Cart
	err := database.DB.
		Preload("Items.Sneaker").
		Where("user_id = ?", userID).
		First(&cart).Error

	if err != nil || len(cart.Items) == 0 {
		c.JSON(400, gin.H{"error": "Cart is empty"})
		return
	}

	// 2. Calculate total
	var total float64
	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Sneaker.Price
	}

	// 3. Transaction (VERY IMPORTANT)
	err = database.DB.Transaction(func(tx *gorm.DB) error {

		// create order
		order := models.Order{
			UserID:      userID,
			TotalAmount: total,
			Status:      "pending",
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// create order items
		for _, item := range cart.Items {
			orderItem := models.OrderItem{
				OrderID:   order.ID,
				SneakerID: item.SneakerID,
				Quantity:  item.Quantity,
				Price:     item.Sneaker.Price,
			}

			if err := tx.Create(&orderItem).Error; err != nil {
				return err
			}
		}

		// clear cart
		if err := tx.Where("cart_id = ?", cart.ID).
			Delete(&models.CartItem{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to place order"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Order placed successfully",
		"total":   total,
	})
}

// get my orders
func GetMyOrders(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var orders []models.Order
	if err := database.DB.
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&orders).Error; err != nil {

		c.JSON(500, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(200, gin.H{"orders": orders})
}

// order details
func GetOrderDetails(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	orderID := c.Param("id")

	var order models.Order
	err := database.DB.
		Preload("OrderItems.Sneaker").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, gin.H{"order": order})
}
