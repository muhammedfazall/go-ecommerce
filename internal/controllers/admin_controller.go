package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

// get admins
func GetAdmins(c *gin.Context) {
	var admins []models.User

	if err := database.DB.
		Where("role = ?", "admin").
		Find(&admins).Error; err != nil {

		c.HTML(500, "admins.html", gin.H{
			"Error": "Failed to load admins",
		})
		return
	}

	c.HTML(200, "admins.html", gin.H{
		"Admins": admins,
	})
}
