package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// place order
func PlaceOrder(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	// 1. Load cart with items
	var cart models.Cart
	err := database.DB.
		Preload("Items.Sneaker").
		Where("user_id = ?", userID).
		First(&cart).Error

	if err != nil || len(cart.Items) == 0 {
		c.JSON(400, gin.H{"error": "Cart is empty"})
		return
	}

	// 2. Transaction (MANDATORY)
	err = database.DB.Transaction(func(tx *gorm.DB) error {

		var total float64

		// 3. Validate & reduce stock
		for _, item := range cart.Items {

			var sneaker models.Sneaker

			// lock row (important)
			if err := tx.
				Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&sneaker, item.SneakerID).Error; err != nil {
				return err
			}

			// check stock
			if sneaker.Stock < item.Quantity {
				return fmt.Errorf(
					"insufficient stock for %s (available %d)",
					sneaker.Name,
					sneaker.Stock,
				)
			}

			// reduce stock
			sneaker.Stock -= item.Quantity
			if err := tx.Save(&sneaker).Error; err != nil {
				return err
			}

			total += float64(item.Quantity) * sneaker.Price
		}

		// 4. Create order
		order := models.Order{
			UserID:      userID,
			TotalAmount: total,
			Status:      "pending",
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// 5. Create order items
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

		// 6. Clear cart
		if err := tx.Where("cart_id = ?", cart.ID).
			Delete(&models.CartItem{}).Error; err != nil {
			return err
		}

		return nil
	})

	// 7. Handle transaction failure
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Order placed successfully",
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
