package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register",controllers.Register)
		auth.POST("/login",controllers.Login)
	}
}
