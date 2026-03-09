package middlewares

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/cache"
	utils "github.com/muhammedfazall/go-ecommerce/utils/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Read token from cookie
		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		// Fallback: Authorization header
		if tokenString == "" {
			authHeader := c.GetHeader("Authorization")
			if token, ok := strings.CutPrefix(authHeader, "Bearer "); ok {
				tokenString = token
			}
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		// Check JWT blacklist (Redis)
		tokenHash := fmt.Sprintf("%x", sha256.Sum256([]byte(tokenString)))
		blacklistKey := "blacklist:" + tokenHash

		if val := cache.Client.Exists(cache.Ctx, blacklistKey).Val(); val > 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token revoked, please login again",
			})
			return
		}

		// Validate JWT
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		// normalize data (JWT numeric claims decode as float64 in Go)
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token payload",
			})
			return
		}
		userID := uint(userIDFloat)

		// Store token string for logout blacklisting
		c.Set("token_string", tokenString)
		c.Set("user_id", userID)
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])

		c.Next()
	}
}