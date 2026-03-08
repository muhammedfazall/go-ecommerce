package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	//Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	//Auth
	JWTSecret string

	//Redis
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

// Load .env at app startup
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	} else {
		log.Println("Environment variables loaded from .env")
	}
}

func Load() *Config {
	return &Config{
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		// Auth
		JWTSecret: getEnv("JWT_SECRET", ""),

		// Redis
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),
	}
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultVal
}
