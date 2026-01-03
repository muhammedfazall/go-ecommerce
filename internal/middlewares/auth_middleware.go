package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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

		//Validate JWT
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}
		
		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])

		c.Next()
	}
}
