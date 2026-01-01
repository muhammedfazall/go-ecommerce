package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Load .env at app startup
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Environment variables loaded")
}
