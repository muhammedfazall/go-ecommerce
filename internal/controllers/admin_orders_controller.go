package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func GetAllOrders(c *gin.Context) {

	status := c.Query("status")
	fromDate := c.Query("from")
	toDate := c.Query("to")

	query := database.DB.
		Preload("User").
		Order("created_at desc")

	// filter by status
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// filter by date range
	if fromDate != "" {
		query = query.Where("created_at >= ?", fromDate)
	}
	if toDate != "" {
		query = query.Where("created_at <= ?", toDate)
	}

	var orders []models.Order
	if err := query.Find(&orders).Error; err != nil {
		c.HTML(500, "orders.html", gin.H{
			"Error": "Failed to load orders",
		})
		return
	}

	c.HTML(200, "orders.html", gin.H{
		"Orders":         orders,

		// pass filters back to template
		"FilterStatus": status,
		"FromDate":     fromDate,
		"ToDate":       toDate,
	})
}


type OrderItemView struct {
	SneakerName string
	Price       float64
	Quantity    int
	Subtotal    float64
}

func ViewOrderDetails(c *gin.Context) {
	id := c.Param("id")

	var order models.Order
	err := database.DB.
		Preload("User").
		Preload("OrderItems.Sneaker").
		First(&order, id).Error

	if err != nil {
		c.HTML(404, "order_details.html", gin.H{
			"Error": "Order not found",
		})
		return
	}

	var items []OrderItemView
	for _, item := range order.OrderItems {
		items = append(items, OrderItemView{
			SneakerName: item.Sneaker.Name,
			Price:       item.Price,
			Quantity:    item.Quantity,
			Subtotal:    item.Price * float64(item.Quantity),
		})
	}

	c.HTML(200, "order_details.html", gin.H{
		"Order":          order,
		"Items":          items,
	})
}

func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	status := c.PostForm("status")

	allowed := map[string]bool{
		"pending":   true,
		"paid":      true,
		"shipped":   true,
		"delivered": true,
		"cancelled": true,
	}

	if !allowed[status] {
		c.Redirect(302, "/admin/orders/"+id)
		return
	}

	if err := database.DB.
		Model(&models.Order{}).
		Where("id = ?", id).
		Update("status", status).Error; err != nil {

		c.Redirect(302, "/admin/orders/"+id)
		return
	}

	c.Redirect(302, "/admin/orders/"+id)
}
