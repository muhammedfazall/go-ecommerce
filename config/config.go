// Load .env at app startup

// Fail fast if .env is missing

package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Environment variables loaded")
}