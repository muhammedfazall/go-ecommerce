package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/config"
	"github.com/muhammedfazall/go-ecommerce/internal/cache"
	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/routes"
)

func main() {
	config.LoadEnv()    // load environment variables
	database.Connect() // connect to PostgreSQL
	database.Migrate() // auto migrations
	database.SeedAdmin(database.DB) // seed default admin
	cache.ConnectRedis() // connect to Redis

	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}