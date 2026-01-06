package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model

	OrderID   uint    `gorm:"not null;index"`
	SneakerID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`

	Order   Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE"`
	Sneaker Sneaker `gorm:"foreignKey:SneakerID;references:ID"`
}
