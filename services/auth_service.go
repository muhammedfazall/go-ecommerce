package services

import (
	"errors"

	"github.com/muhammedfazall/go-ecommerce/database"
	"github.com/muhammedfazall/go-ecommerce/models"
	"github.com/muhammedfazall/go-ecommerce/utils"
	"gorm.io/gorm"
)

func RegisterUser(name, email, password string) error {
	var existingUser models.User

	err := database.DB.Where("email = ?", email).First(&existingUser).Error

	if err == nil {
		return errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword,err := utils.HashPassword(password)
	if err != nil{
		return err
	}

	user := models.User{
		Name: name,
		Email: email,
		Password: hashedPassword,
		Role:"user",
	}

	return database.DB.Create(&user).Error
}
