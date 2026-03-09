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

	r.GET("/collections", controllers.AllProducts)
	r.GET("/products/:id", controllers.ProductDetails)
	r.GET("/categories", controllers.Categories)
	r.GET("/categories/:id/products", controllers.ProductsByCategory)
	r.GET("/products/search", controllers.SearchProducts)

	// --------------------
	// Auth APIs
	// --------------------
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", controllers.Logout)
		auth.POST("/refresh", controllers.RefreshToken)
	}

	// --------------------
	// Wishlist APIs (User Protected)
	// --------------------

	wishlist := r.Group("/wishlist")
	wishlist.Use(middlewares.AuthMiddleware())
	{
		wishlist.POST("/add", controllers.AddToWishlist)
		wishlist.GET("", controllers.GetWishlist)
		wishlist.DELETE("/remove/:productId", controllers.RemoveFromWishlist)
	}

	// --------------------
	// Cart APIs (User Protected)
	// --------------------
	cart := r.Group("/cart")
	cart.Use(middlewares.AuthMiddleware())
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("", controllers.GetCart)
		cart.PUT("/update", controllers.UpdateCartItem)
		cart.DELETE("/remove/:productId", controllers.RemoveFromCart)
		cart.DELETE("/clear", controllers.ClearCart)
	}

	// --------------------
	// Order APIs (User Protected)
	// --------------------

	orders := r.Group("/orders")
	orders.Use(middlewares.AuthMiddleware())
	{
		orders.POST("", controllers.PlaceOrder)
		orders.GET("", controllers.GetMyOrders)
		orders.GET("/:id", controllers.GetOrderDetails)
	}

	// --------------------
	// Payment APIs (User Protected)
	// --------------------

	payments := r.Group("/payments")
	payments.Use(middlewares.AuthMiddleware())
	{
		payments.POST("/create", controllers.CreatePayment)
		payments.POST("/verify", controllers.VerifyPayment)
	}

	// --------------------
	// User profile/Logout (User Protected)
	// --------------------

	user := r.Group("/user")
	user.Use(middlewares.AuthMiddleware())
	{
		user.GET("/profile", controllers.UserProfile)
	}

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
		adminProtected.GET("/dashboard", controllers.AdminDashboard)

		//PRODUCTS
		adminProtected.GET("/products", controllers.GetAllSneakers)
		// Add product (form)
		adminProtected.GET("/products/new", controllers.AddSneakerPage)
		// Add product (submit)
		adminProtected.POST("/products/new", controllers.AddSneaker)

		//view product
		adminProtected.GET("/products/:id", controllers.ViewSneaker)
		//edit product (form)
		adminProtected.GET("/products/:id/edit", controllers.EditSneaker)
		//update product (submit)
		adminProtected.POST("/products/:id/edit", controllers.UpdateSneaker)
		adminProtected.POST("/products/:id/delete", controllers.DeleteSneaker)

		//USERS

		adminProtected.GET("/users", controllers.GetUsers)
		//block/unbock user
		adminProtected.POST("/users/:id/block", controllers.ToggleBlockUser)
		//change role
		adminProtected.POST("/users/:id/role", controllers.ChangeUserRole)

		//CATEGORIES
		adminProtected.GET("/categories", controllers.GetCategories)
		adminProtected.POST("/categories", controllers.CreateCategory)

		// ORDERS
		adminProtected.GET("/orders", controllers.GetAllOrders)
		adminProtected.GET("/orders/:id", controllers.ViewOrderDetails)
		adminProtected.POST("/orders/:id/status", controllers.UpdateOrderStatus)

		//Admins
		adminProtected.GET("/admins", controllers.GetAdmins)
		adminProtected.GET("/profile", controllers.AdminProfile)
		adminProtected.GET("/profile/edit", controllers.EditAdminProfile)
		adminProtected.POST("/profile/update", controllers.UpdateAdminProfile)

		adminProtected.GET("/password-change", controllers.AdminChangePasswordPage)
		adminProtected.POST("/password-change", controllers.AdminChangePassword)

		adminProtected.GET("/logout", controllers.AdminLogout)

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
