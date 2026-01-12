package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

// show categories
func Categories(c *gin.Context) {
	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"categories": categories,
	})
}

// products by category
func ProductsByCategory(c *gin.Context) {
	categoryID := c.Param("id")

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "12")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	var products []models.Sneaker

	query := database.DB.
		Where("category_id = ?", categoryID)

	// sorting
	switch c.Query("sort") {
	case "price_asc":
		query = query.Order("price asc")
	case "price_desc":
		query = query.Order("price desc")
	default:
		query = query.Order("created_at desc")
	}

	if err := query.
		Limit(limitInt).
		Offset(offset).
		Find(&products).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to fetch products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"page":    pageInt,
		"limit":   limitInt,
		"count":   len(products),
		"data":    products,
	})
}
