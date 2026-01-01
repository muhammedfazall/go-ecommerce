package services

import (
	"errors"

	"github.com/muhammedfazall/go-ecommerce/database"
	"github.com/muhammedfazall/go-ecommerce/models"
	"github.com/muhammedfazall/go-ecommerce/helpers"
	"gorm.io/gorm"
)

func RegisterUser(name, email, password string) error {
	var existingUser models.User

	err := database.DB.Where("email = ?", email).First(&existingUser).Error

	if err == nil {
		return errors.New("user already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword,err := helpers.HashPassword(password)
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

func LoginUser(email,password string) (*models.User,error) {
	var user models.User

	err := database.DB.Where("email = ?",email).First(&user).Error
	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return nil,errors.New("invalid email or password")
		}
		return nil,err
	}

	if err := helpers.CheckPassword(user.Password,password);err != nil{
		return nil, errors.New("invalid email or password")
	}
	return &user, nil
}