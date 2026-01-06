package routes

import (
	"time"

	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
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

	r.HTMLRender = createMyRender("templates")

	// --------------------
	// Auth APIs
	// --------------------
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// auth1 := r.Group("/auth")
	// auth1.Use(middlewares.AuthMiddleware())
	// auth1.GET("/user", func(c *gin.Context) {
	// 	role := c.MustGet("role").(string)
	// 	c.JSON(200, gin.H{
	// 		"role": role,
	// 	})
	// })

	// --------------------
	// Admin Public Routes
	// --------------------
	adminPublic := r.Group("/admin")
	{
		//login page
		adminPublic.GET("/login", func(c *gin.Context) {
			c.HTML(200, "login.html", nil)
		})
		adminPublic.POST("/login", controllers.AdminLogin)
	}

	// --------------------
	// Admin Protected Routes
	// --------------------
	adminProtected := r.Group("/admin")
	adminProtected.Use(middlewares.AuthMiddleware())
	adminProtected.Use(middlewares.AdminMiddleware())
	{

		//DASHBOARD
		adminProtected.GET("/dashboard", func(c *gin.Context) {
			c.HTML(200, "dashboard.html", nil)
		})

		//PRODUCTS
		adminProtected.GET("/products", controllers.GetAllSneakers)
		// Add product (form)
		adminProtected.GET("/products/new", controllers.AddSneakerPage)
		// Add product (submit)
		adminProtected.POST("/products/new", controllers.AddSneaker)

		//view product
		adminProtected.GET("/products/:id",controllers.ViewSneaker)
		//edit product (form)
		adminProtected.GET("/products/:id/edit", controllers.EditSneaker)
		//update product (submit)
		adminProtected.POST("/products/:id/edit", controllers.UpdateSneaker)
		adminProtected.POST("/products/:id/delete", controllers.DeleteSneaker)

		//USERS

		adminProtected.GET("/users", controllers.GetUsers)

		//CATEGORIES
		adminProtected.GET("/categories", controllers.GetCategories)
		adminProtected.POST("/categories", controllers.CreateCategory)

		// adminProtected.GET("/ping", func(ctx *gin.Context) {
		// 	ctx.JSON(200, gin.H{
		// 		"message": "admin access granted",
		// 	})
		// })
	}
}

func createMyRender(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// load all html files
	files, err := filepath.Glob(filepath.Join(templatesDir, "*.html"))
	if err != nil {
		panic(err)
	}

	adminBase := filepath.Join(templatesDir, "admin_base.html")
	// adminLoginBase := filepath.Join(templatesDir, "admin_login_base.html")

	for _, file := range files {
		name := filepath.Base(file)

		switch name {

		// -------------------------
		// Admin Login Page
		// -------------------------
		case "login.html":
			// OPTION 1: login without base
			r.AddFromFiles(name, file)

			// OPTION 2 (if you want a login base)
			// r.AddFromFiles(name, adminLoginBase, file)

		// -------------------------
		// Base templates (skip)
		// -------------------------
		case "admin_base.html", "admin_login_base.html":
			continue

		// -------------------------
		// All other admin pages
		// -------------------------
		default:
			r.AddFromFiles(name, adminBase, file)
		}
	}

	return r
}
