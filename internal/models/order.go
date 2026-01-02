package models

type Order struct {
	BaseModel

	UserID      uint    `gorm:"not null;index"`
	TotalAmount float64 `gorm:"not null"`
	Status      string  `gorm:"size:30;default:pending"`

	User       User `gorm:"foreignKey:UserID"`
	OrderItems []OrderItem
}
