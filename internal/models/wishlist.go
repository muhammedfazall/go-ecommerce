package models

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model

	UserID    uint `gorm:"not null;uniqueIndex:idx_user_sneaker"`
	SneakerID uint `gorm:"not null;uniqueIndex:idx_user_sneaker"`

	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Sneaker Sneaker `gorm:"foreignKey:SneakerID;constraint:OnDelete:CASCADE"`
}
