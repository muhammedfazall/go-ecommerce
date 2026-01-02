package services

import (
	"errors"

	"github.com/muhammedfazall/go-ecommerce/internal/database"
	"github.com/muhammedfazall/go-ecommerce/internal/helpers"
	"github.com/muhammedfazall/go-ecommerce/internal/models"
	"gorm.io/gorm"
)

func RegisterUser(username, email, password string) error {
	var existingUser models.User

	err := database.DB.Where("email = ?", email).First(&existingUser).Error

	if err == nil {
		return errors.New("user already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Username:     username,
		Email:    email,
		Password: hashedPassword,
	}

	return database.DB.Create(&user).Error
}

func LoginUser(email, password string) (*models.User, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if err := helpers.CheckPassword(user.Password, password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.IsBlocked {
		return nil, errors.New("account is blocked")
	}

	if user.Status != "active" {
		return nil, errors.New("account is not active")
	}
	return &user, nil
}
