package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muhammedfazall/go-ecommerce/config"
	"github.com/muhammedfazall/go-ecommerce/database"
)

func main() {
	config.LoadEnv() //load environment variables at application startup using a config layer
	database.Connect()

	r := gin.Default()

	r.GET("health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
			"db":     "connected",
		})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}
