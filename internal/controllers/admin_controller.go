package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/cache"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/helpers"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	utils "github.com/muhammedfazall/go-ecommerce/utils/jwt"
)

// get admin profile
func AdminProfile(c *gin.Context) {
	adminID, exists := c.Get("user_id")
	if !exists {
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	var admin models.User

	if err := database.DB.First(&admin, adminID).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.HTML(http.StatusOK, "admin_profile.html", gin.H{
		"User": admin,
	})
}

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

// Edit
func EditAdminProfile(c *gin.Context) {
	adminID, _ := c.Get("user_id")

	var admin models.User
	database.DB.First(&admin, adminID)

	c.HTML(http.StatusOK, "admin_profile_edit.html", gin.H{
		"User": admin,
	})
}

func UpdateAdminProfile(c *gin.Context) {
	adminID, _ := c.Get("user_id")

	username := c.PostForm("username")
	email := c.PostForm("email")

	var admin models.User
	if err := database.DB.First(&admin, adminID).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if username == "" || email == "" {
		c.HTML(http.StatusBadRequest, "admin_profile_edit.html", gin.H{
			"error": "Username and Email are required",
			"User":  admin,
		})
		return
	}

	admin.Username = username
	admin.Email = email

	if err := database.DB.Save(&admin).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "admin_profile_edit.html", gin.H{
			"error": "Failed to update profile",
			"User":  admin,
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/profile")
}

func AdminChangePasswordPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_change_password.html", nil)
}

func AdminChangePassword(c *gin.Context) {
	adminID, _ := c.Get("user_id")

	currentPassword := c.PostForm("current_password")
	newPassword := c.PostForm("new_password")
	confirmPassword := c.PostForm("confirm_password")

	if newPassword != confirmPassword {
		c.HTML(http.StatusBadRequest, "admin_change_password.html", gin.H{
			"error": "Passwords do not match",
		})
		return
	}

	var admin models.User
	database.DB.First(&admin, adminID)

	// Compare old password
	if err := helpers.CheckPassword(admin.Password, currentPassword); err != nil {
		c.HTML(http.StatusUnauthorized, "admin_change_password.html", gin.H{
			"error": "Current password is incorrect",
		})
		return
	}

	hashedPassword, _ := helpers.HashPassword(newPassword)
	admin.Password = hashedPassword

	database.DB.Save(&admin)

	c.Redirect(http.StatusFound, "/admin/profile")
}

func AdminLogout(c *gin.Context) {
	// Blacklist the token in Redis so it can't be reused
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

	// Clear the cookie
	c.SetCookie(
		"access_token",
		"",
		-1,
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

	c.Redirect(http.StatusFound, "/admin/login")
}
