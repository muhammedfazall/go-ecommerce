package models

type Sneaker struct {
	BaseModel

	Name        string  `gorm:"not null"`
	Brand       string  `gorm:"not null"`
	CategoryID  uint    `gorm:"not null"`
	Gender      string  `gorm:"size:20;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	ImageURL    string
	IsActive    bool `gorm:"default:true"`
}
