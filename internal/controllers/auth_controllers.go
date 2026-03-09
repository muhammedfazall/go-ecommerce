package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/helpers"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"github.com/muhammedfazall/go-ecommerce/internal/services"
	utils "github.com/muhammedfazall/go-ecommerce/utils/jwt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name, email and password are required",
		})
		return
	}

	//calls register service
	err := services.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	//calls login service
	user, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// generates access token
	accesstoken, err := utils.GenerateAccessToken(
		user.ID,
		user.Email,
		user.Role,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate access token",
		})
		return
	}

	// generate refresh token
	refreshtoken, err := utils.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate refresh token",
		})
		return
	}

	// store refresh token in Redis
	if err := utils.StoreRefreshToken(user.ID, refreshtoken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not store refresh token",
		})
		return
	}

	// set access token cookie (15 min)
	c.SetCookie(
		"access_token",
		accesstoken,
		int(utils.AccessTokenDuration.Seconds()),
		"/",
		"",
		false,
		true,
	)

	// set refresh token cookie (7 days, scoped to /auth/refresh)
	c.SetCookie(
		"refresh_token",
		refreshtoken,
		int(utils.RefreshTokenDuration.Seconds()),
		"/auth/refresh",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message":      "login successfull",
		"access_token": accesstoken,
		"user":         user.Username,
	})
}

func AdminLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"Error": "Email and password are required",
		})
		return
	}

	var admin models.User

	// Find user by email
	err := database.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "invalid credentials",
		})
		return
	}

	// Check role
	if admin.Role != "admin" {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"Error": "Admin access only",
		})
		return
	}

	// Check account flags
	if admin.IsBlocked {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"Error": "Account is blocked",
		})
		return
	}

	if admin.Status != "active" {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"Error": "Account is not active",
		})
		return
	}

	// Verify password

	if err := helpers.CheckPassword(admin.Password, password); err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "invalid credentials",
		})
		return
	}

	// Generate JWT
	accesstoken, err := utils.GenerateAccessToken(
		admin.ID,
		admin.Email,
		admin.Role,
	)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"Error": "Could not login. Please try again.",
		})
		return
	}

	// Generate refresh token
	refreshtoken, err := utils.GenerateRefreshToken()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"Error": "Could not login. Please try again.",
		})
		return
	}

	// Store refresh token in Redis
	if err := utils.StoreRefreshToken(admin.ID, refreshtoken); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"Error": "Could not login. Please try again.",
		})
		return
	}

	// Set access token cookie (15 min)
	c.SetCookie(
		"access_token",
		accesstoken,
		int(utils.AccessTokenDuration.Seconds()),
		"/",
		"",
		false,
		true,
	)

	// Set refresh token cookie (7 days)
	c.SetCookie(
		"refresh_token",
		refreshtoken,
		int(utils.RefreshTokenDuration.Seconds()),
		"/auth/refresh",
		"",
		false,
		true,
	)

	c.Redirect(http.StatusFound, "/admin/dashboard")

}
