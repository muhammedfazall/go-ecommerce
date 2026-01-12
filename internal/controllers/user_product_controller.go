package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

// show all products
func AllProducts(c *gin.Context) {
	var sneakers []models.Sneaker

	// pagination defaults
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "12")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if pageInt < 1 {
		pageInt = 1
	}
	if limitInt > 50 {
		limitInt = 50
	}

	offset := (pageInt - 1) * limitInt

	query := database.DB.Model(&models.Sneaker{})

	// filtering
	if brand := c.Query("brand"); brand != "" {
		query = query.Where("brand = ?", brand)
	}

	if gender := c.Query("gender"); gender != "" {
		query = query.Where("gender = ?", gender)
	}

	if min := c.Query("minPrice"); min != "" {
		query = query.Where("price >= ?", min)
	}

	if max := c.Query("maxPrice"); max != "" {
		query = query.Where("price <= ?", max)
	}

	// sorting (safe whitelist)
	switch c.Query("sort") {
	case "price_asc":
		query = query.Order("price asc")
	case "price_desc":
		query = query.Order("price desc")
	case "latest":
		query = query.Order("created_at desc")
	default:
		query = query.Order("id desc")
	}

	// execute query
	if err := query.
		Limit(limitInt).
		Offset(offset).
		Find(&sneakers).Error; err != nil {

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
		"count":   len(sneakers),
		"data":    sneakers,
	})
}

// view product by id
func ProductDetails(c *gin.Context) {
	id := c.Param("id")

	var sneaker models.Sneaker

	result := database.DB.First(&sneaker, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch product",
		})
		return
	}

	c.JSON(http.StatusOK, sneaker)
}


func SearchProducts(c *gin.Context) {
	queryText := c.Query("q")

	if queryText == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Search query (q) is required",
		})
		return
	}

	// pagination defaults
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "12")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if pageInt < 1 {
		pageInt = 1
	}
	if limitInt > 50 {
		limitInt = 50
	}

	offset := (pageInt - 1) * limitInt

	var products []models.Sneaker

	dbQuery := database.DB.Model(&models.Sneaker{}).
		Where(
			"name ILIKE ? OR brand ILIKE ? OR description ILIKE ?",
			"%"+queryText+"%",
			"%"+queryText+"%",
			"%"+queryText+"%",
		)

	// Filters
	if brand := c.Query("brand"); brand != "" {
		dbQuery = dbQuery.Where("brand = ?", brand)
	}

	if gender := c.Query("gender"); gender != "" {
		dbQuery = dbQuery.Where("gender = ?", gender)
	}

	if min := c.Query("minPrice"); min != "" {
		dbQuery = dbQuery.Where("price >= ?", min)
	}

	if max := c.Query("maxPrice"); max != "" {
		dbQuery = dbQuery.Where("price <= ?", max)
	}

	// Sorting (whitelisted)
	switch c.Query("sort") {
	case "price_asc":
		dbQuery = dbQuery.Order("price asc")
	case "price_desc":
		dbQuery = dbQuery.Order("price desc")
	case "latest":
		dbQuery = dbQuery.Order("created_at desc")
	default:
		dbQuery = dbQuery.Order("id desc")
	}

	// Execute
	if err := dbQuery.
		Limit(limitInt).
		Offset(offset).
		Find(&products).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Search failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"query":   queryText,
		"page":    pageInt,
		"limit":   limitInt,
		"count":   len(products),
		"data":    products,
	})
}
