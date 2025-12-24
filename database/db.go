package database

import (
	"fmt"
	"log"
	"os"

	"github.com/muhammedfazall/go-ecommerce/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// Read DB config from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL database")
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{}, 
		&models.Sneaker{},
		&models.Cart{}, 
		&models.Order{}, 
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Migration Failed", err)
	}

	log.Println("Database migration completed")
}
