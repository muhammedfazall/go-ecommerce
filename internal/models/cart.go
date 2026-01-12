package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	UserID uint `gorm:"uniqueIndex;not null"`

	User  User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Items []CartItem `gorm:"foreignKey:CartID"`
}

