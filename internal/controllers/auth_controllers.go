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

	// generates token
	accesstoken, err := utils.GenerateAccessToken(
		user.ID,
		user.Email,
		user.Role,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate token",
		})
		return
	}

	// set token in cookies
	c.SetCookie("access_token", accesstoken, 600, "/", "", false, true)

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
		c.HTML(http.StatusBadRequest, "admin_login.html", gin.H{
			"error": "Email and password are required",
		})
		return
	}

	var admin models.User

	// Find user by email
	err := database.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		c.HTML(http.StatusUnauthorized, "admin_login.html", gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// Check role
	if admin.Role != "admin" {
		c.HTML(http.StatusForbidden, "admin_login.html", gin.H{
			"error": "Admin access only",
		})
		return
	}

	// Check account flags
	if admin.IsBlocked {
		c.HTML(http.StatusForbidden, "admin_login.html", gin.H{
			"error": "Account is blocked",
		})
		return
	}

	if admin.Status != "active" {
		c.HTML(http.StatusForbidden, "admin_login.html", gin.H{
			"error": "Account is not active",
		})
		return
	}

	// Verify password

	if err := helpers.CheckPassword(admin.Password, password); err != nil{
		c.HTML(http.StatusUnauthorized,"admin_login.html",gin.H{
			"error":"invalid credentials",
		})
		return
	}

	c.Redirect(http.StatusFound,"/admin/dashboard")

}
