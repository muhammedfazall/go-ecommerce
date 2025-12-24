package models

type Cart struct {
	BaseModel

	UserID    uint `gorm:"not null;index"`
	SneakerID uint `gorm:"not null;index"`
	Quantity  int  `gorm:"not null;check:quantity > 0"`

	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Sneaker Sneaker `gorm:"foreignKey:SneakerID;constraint:OnDelete:CASCADE"`
}
