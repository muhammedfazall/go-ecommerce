package models

type OrderItem struct {
	BaseModel

	OrderID   uint    `gorm:"not null;index"`
	SneakerID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`

	Order   Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE"`
	Sneaker Sneaker `gorm:"foreignKey:SneakerID;references:ID"`
}
