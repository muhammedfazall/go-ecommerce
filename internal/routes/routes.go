package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/internal/controllers"
	"github.com/muhammedfazall/go-ecommerce/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {

	// --------------------
	// Global Middleware
	// --------------------
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --------------------
	// Auth APIs
	// --------------------
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	auth1 := r.Group("/auth")
	auth1.Use(middlewares.AuthMiddleware())
	auth1.GET("/user",func(ctx *gin.Context) {

		role := ctx.MustGet("role").(string)
		ctx.JSON(200,gin.H{
			"role":role,
		})
	})



	// --------------------
	// Admin Public Routes
	// --------------------
	adminPublic := r.Group("/admin")
	{
		adminPublic.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(200, "admin_login.html", nil)
		})

		adminPublic.POST("/login",controllers.AdminLogin)
	}

	// --------------------
	// Admin Protected Routes
	// --------------------
	adminProtected := r.Group("/admin")
	adminProtected.Use(middlewares.AuthMiddleware())
	adminProtected.Use(middlewares.AdminMiddleware())
	{
		adminProtected.GET("/dashboard", func(ctx *gin.Context) {
			ctx.HTML(200, "dashboard.html", nil)
		})

		adminProtected.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "admin access granted",
			})
		})
	}
}
