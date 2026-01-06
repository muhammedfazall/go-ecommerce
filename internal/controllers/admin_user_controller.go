package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.HTML(500, "users.html", gin.H{
			"Error": "Failed to load users",
		})
		return
	}

	c.HTML(200, "users.html", gin.H{
		"Users": users,
	})
}
