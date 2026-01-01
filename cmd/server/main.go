package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/config"
	"github.com/muhammedfazall/go-ecommerce/database"
	"github.com/muhammedfazall/go-ecommerce/routes"
)

func main() {
	config.LoadEnv() //load environment variables
	database.Connect() //connecting to database
	database.Migrate() //auto migrations

	r := gin.Default()


	routes.RegisterRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
