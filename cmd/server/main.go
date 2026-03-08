package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/config"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/routes"
)

func main() {
	config.LoadEnv() //load environment variables
	database.Connect() //connecting to database
	database.Migrate() //auto migrations
	database.SeedAdmin(database.DB) 

	r := gin.Default()
	// r.LoadHTMLGlob("templates/**/*.html")
	routes.RegisterRoutes(r)
	log.Println("Server running on :8080")
	r.Run(":8080")
}
