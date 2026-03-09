package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/cache"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	utils "github.com/muhammedfazall/go-ecommerce/utils/jwt"
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
	// Blacklist the access token in Redis so it can't be reused
	tokenString, exists := c.Get("token_string")
	if exists && tokenString != "" {
		tokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(tokenString.(string))))
		cache.Client.Set(cache.Ctx, "blacklist:"+tokenHash, "1", 60*time.Minute)
	}

	// Delete refresh token from Redis
	userID, exists := c.Get("user_id")
	if exists {
		utils.DeleteRefreshToken(userID.(uint))
	}

	// Clear the access token cookie
	c.SetCookie(
		"access_token",
		"",
		-1, // expire immediately
		"/",
		"",
		false,
		true,
	)

	// Clear the refresh token cookie
	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/auth/refresh",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
