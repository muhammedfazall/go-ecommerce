package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden,gin.H{
				"error":"role not found",
			})
			return 
		}

		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden,gin.H{
				"error":"admin access required",
			})
			return
		}

		c.Next()
	}
}