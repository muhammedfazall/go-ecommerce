package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	UserID    uint `gorm:"not null;index"`
	SneakerID uint `gorm:"not null;index"`
	Quantity  int  `gorm:"not null;check:(quantity > 0)"`

	User    User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Sneaker Sneaker `gorm:"foreignKey:SneakerID;references:ID;constraint:OnDelete:CASCADE"`
}
