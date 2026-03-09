package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/muhammedfazall/go-ecommerce/utils/jwt"
)

// RefreshToken handles POST /auth/refresh.
// It reads the refresh token from a cookie, validates it against Redis,
// then issues a new access token and rotates the refresh token.
func RefreshToken(c *gin.Context) {
	// 1. Read refresh token from cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "refresh token not found",
		})
		return
	}

	// 2. Read the (potentially expired) access token to extract user claims
	accessToken, _ := c.Cookie("access_token")
	if accessToken == "" {
		// Fallback: check Authorization header
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			accessToken = authHeader[7:]
		}
	}

	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "access token required for refresh",
		})
		return
	}

	// 3. Parse claims from the expired access token (signature is still validated)
	claims, err := utils.ParseTokenUnverifiedClaims(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid access token",
		})
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token payload",
		})
		return
	}
	userID := uint(userIDFloat)

	email, _ := claims["email"].(string)
	role, _ := claims["role"].(string)

	// 4. Validate the refresh token against Redis
	if err := utils.ValidateRefreshToken(userID, refreshToken); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid or expired refresh token",
		})
		return
	}

	// 5. Issue new access token
	newAccessToken, err := utils.GenerateAccessToken(userID, email, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate access token",
		})
		return
	}

	// 6. Rotate refresh token (delete old, create+store new)
	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate refresh token",
		})
		return
	}

	if err := utils.StoreRefreshToken(userID, newRefreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not store refresh token",
		})
		return
	}

	// 7. Set cookies
	c.SetCookie(
		"access_token",
		newAccessToken,
		int(utils.AccessTokenDuration.Seconds()), // 900s = 15 min
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		int(utils.RefreshTokenDuration.Seconds()), // 7 days
		"/auth/refresh",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message":      "token refreshed successfully",
		"access_token": newAccessToken,
	})
}
