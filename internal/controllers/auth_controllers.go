package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/email"
	"github.com/muhammedfazall/go-ecommerce/internal/helpers"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"github.com/muhammedfazall/go-ecommerce/internal/services"
	"github.com/muhammedfazall/go-ecommerce/utils/jwt"
	"github.com/muhammedfazall/go-ecommerce/utils/otp"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, email and password are required"})
		return
	}

	// create user with status=pending
	err := services.RegisterUser(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate otp
	OTP,err := otp.GenerateOTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to generate OTP"})
		return
	}

	// store otp on redis
	err = otp.StoreOTP(req.Email ,OTP)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to Store OTP"})
		return
	}

	//  send OTP mail
	if err := email.SendOTPEmail(req.Email,OTP); err != nil{
		c.JSON(http.StatusCreated,gin.H{"message":"registered successfully but failed to send OTP email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "OTP sent to your email, please verify to activate your account",
	})
}

// activates the user account after OTP check
func VerifyOTP(c *gin.Context)  {
	var req struct{
		Email string `json:"email"`
		OTP string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.OTP == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and otp are required"})
		return
	}

	// Verify OTP from Redis
	if !otp.VerifyOTP(req.Email, req.OTP) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or expired OTP"})
		return
	}

	// Activate user
	result := database.DB.Model(&models.User{}).
		Where("email = ?", req.Email).Updates(map[string]interface{}{
			"status":      "active",
			"is_verified": true,
		})
 
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to activate account"})
		return
	}
 
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// send welcome email (non-blocking)
	var user models.User
	database.DB.Where("email = ?", req.Email).First(&user)
	go email.SendWelcomeEmail(req.Email, user.Username)
 
	c.JSON(http.StatusOK, gin.H{"message": "email verified successfully, you can now login"})
}

// ResendOTP generates and sends a fresh OTP
func ResendOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
 
	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
 
	// Check user exists and is still pending
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
 
	if user.IsVerified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account is already verified"})
		return
	}
 
	// Generate and store new OTP
	OTP, err := otp.GenerateOTP()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate OTP"})
		return
	}
 
	if err := otp.StoreOTP(req.Email, OTP); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store OTP"})
		return
	}
 
	if err := email.SendOTPEmail(req.Email, OTP); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send OTP email"})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{"message": "OTP resent successfully"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	//calls login service
	user, err := services.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// generates access token
	accesstoken, err := jwt.GenerateAccessToken(
		user.ID,
		user.Email,
		user.Role,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate access token"})
		return
	}

	// generate refresh token
	refreshtoken, err := jwt.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate refresh token"})
		return
	}

	// store refresh token in Redis
	if err := jwt.StoreRefreshToken(user.ID, refreshtoken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not store refresh token"})
		return
	}

	// set access token cookie (15 min)
	c.SetCookie(
		"access_token",
		accesstoken,
		int(jwt.AccessTokenDuration.Seconds()),
		"/",
		"",
		false,
		true,
	)

	// set refresh token cookie (7 days, scoped to /auth/refresh)
	c.SetCookie(
		"refresh_token",
		refreshtoken,
		int(jwt.RefreshTokenDuration.Seconds()),
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
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": "Email and password are required"})
		return
	}

	var admin models.User

	// Find user by email
	err := database.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "invalid credentials"})
		return
	}

	// Check role
	if admin.Role != "admin" {
		c.HTML(http.StatusForbidden, "login.html", gin.H{"Error": "Admin access only"})
		return
	}

	// Check account flags
	if admin.IsBlocked {
		c.HTML(http.StatusForbidden, "login.html", gin.H{"Error": "Account is blocked"})
		return
	}

	if admin.Status != "active" {
		c.HTML(http.StatusForbidden, "login.html", gin.H{"Error": "Account is not active"})
		return
	}

	// Verify password

	if err := helpers.CheckPassword(admin.Password, password); err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "invalid credentials"})
		return
	}

	// Generate JWT
	accesstoken, err := jwt.GenerateAccessToken(
		admin.ID,
		admin.Email,
		admin.Role,
	)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": "Could not login. Please try again."})
		return
	}

	// Generate refresh token
	refreshtoken, err := jwt.GenerateRefreshToken()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": "Could not login. Please try again."})
		return
	}

	// Store refresh token in Redis
	if err := jwt.StoreRefreshToken(admin.ID, refreshtoken); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": "Could not login. Please try again."})
		return
	}

	// Set access token cookie (15 min)
	c.SetCookie(
		"access_token",
		accesstoken,
		int(jwt.AccessTokenDuration.Seconds()),
		"/",
		"",
		false,
		true,
	)

	// Set refresh token cookie (7 days)
	c.SetCookie(
		"refresh_token",
		refreshtoken,
		int(jwt.RefreshTokenDuration.Seconds()),
		"/auth/refresh",
		"",
		false,
		true,
	)

	c.Redirect(http.StatusFound, "/admin/dashboard")

}
