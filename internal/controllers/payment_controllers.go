package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func CreatePayment(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req struct {
		OrderID uint `json:"order_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var order models.Order
	err := database.DB.
		Where("id = ? AND user_id = ?", req.OrderID, userID).
		First(&order).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	if order.Status != "pending" {
		c.JSON(400, gin.H{"error": "Order already processed"})
		return
	}

	// FAKE payment intent
	c.JSON(200, gin.H{
		"message":         "Payment initiated",
		"order_id":        order.ID,
		"amount":          order.TotalAmount,
		"fake_payment_id": "PAY_FAKE_123456",
	})
}


func VerifyPayment(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req struct {
		OrderID uint `json:"order_id"`
		Success bool `json:"success"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var order models.Order
	err := database.DB.
		Where("id = ? AND user_id = ?", req.OrderID, userID).
		First(&order).Error

	if err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	if order.Status != "pending" {
		c.JSON(400, gin.H{"error": "Order already processed"})
		return
	}

	if req.Success {
		order.Status = "paid"
	} else {
		order.Status = "failed"
	}

	if err := database.DB.Save(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Payment processed",
		"status":  order.Status,
	})
}
