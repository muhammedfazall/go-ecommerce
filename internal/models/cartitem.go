package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model

	CartID    uint `gorm:"not null;index"`
	SneakerID uint `gorm:"not null;index"`
	Quantity  int  `gorm:"not null;check:(quantity > 0)"`

	Sneaker Sneaker `gorm:"foreignKey:SneakerID;constraint:OnDelete:CASCADE"`
}
