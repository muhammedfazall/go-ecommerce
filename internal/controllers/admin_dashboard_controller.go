package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func AdminDashboard(c *gin.Context) {

	// -----------------------
	// Stats
	// -----------------------
	var totalUsers int64
	var totalOrders int64
	var totalProducts int64
	var totalRevenue float64

	database.DB.Model(&models.User{}).Count(&totalUsers)
	database.DB.Model(&models.Order{}).Count(&totalOrders)
	database.DB.Model(&models.Sneaker{}).Count(&totalProducts)

	database.DB.
		Model(&models.Order{}).
		Select("COALESCE(SUM(total_amount), 0)").
		Where("status = ?", "paid").
		Scan(&totalRevenue)

	// -----------------------
	// Recent Orders
	// -----------------------
	var recentOrders []models.Order
	database.DB.
		Preload("User").
		Order("created_at DESC").
		Limit(5).
		Find(&recentOrders)

	// -----------------------
	// Sales Chart (Last 7 Days)
	// -----------------------
	type SalesData struct {
		Date  string  `json:"Date"`
		Total float64 `json:"Total"`
	}

	var sales []SalesData

	database.DB.
		Model(&models.Order{}).
		Select(`
			TO_CHAR(created_at::date, 'DD Mon') AS date,
			SUM(total_amount) AS total
		`).
		Where("status = ?", "paid").
		Group("created_at::date").
		Order("created_at::date ASC").
		Limit(7).
		Scan(&sales)

	// Convert to JSON for safe JS embedding
	salesJSON, err := json.Marshal(sales)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "dashboard.html", gin.H{
			"Error": "Failed to load sales data",
		})
		return
	}

	// -----------------------
	// Render Dashboard
	// -----------------------
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"CurrentSection": "dashboard",

		"TotalUsers":    totalUsers,
		"TotalOrders":   totalOrders,
		"TotalProducts": totalProducts,
		"TotalRevenue":  totalRevenue,
		"RecentOrders":  recentOrders,

		// 👇 SAFE JSON for Chart.js
		"SalesJSON": template.JS(salesJSON),
	})
}
