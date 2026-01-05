package models

type Sneaker struct {
	BaseModel

	Name        string  `gorm:"not null" json:"name"`
	Brand       string  `gorm:"not null" json:"brand"`
	CategoryID  uint    `gorm:"not null" json:"category_id"`
	Gender      string  `gorm:"size:20;not null" json:"gender"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"not null" json:"stock"`
	ImageURL    string `json:"image_url"`
	IsActive    bool `gorm:"default:true" json:"is_active"`
}
