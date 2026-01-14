package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
)

func UserProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// never expose password
	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"role":      user.Role,
		"status":    user.Status,
		"isBlocked": user.IsBlocked,
		"createdAt": user.CreatedAt,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"access_token",
		"",
		-1, // expire immediately
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
