package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func GetCategories(c *gin.Context) {

	var categories []models.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"failed to fetch categories",
		})
		return
	}

	c.JSON(http.StatusOK,categories)
}

func CreateCategory(c *gin.Context)  {
	var input struct{
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category name"})
		return
	}

	category := models.Category{
		Name: input.Name,
	}
	if err := database.DB.Create(&category).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"failed to create category",
		})
		return
	}
	c.JSON(http.StatusCreated,category)
}
