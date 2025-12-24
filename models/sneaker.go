package models

type Sneaker struct {
	BaseModel

	Name     string  `gorm:"size:150;not null"`
	Brand    string  `gorm:"size:100;not null;index"`
	Price    float64 `gorm:"not null"`
	Size     int     `gorm:"not null"`
	Stock    int     `gorm:"not null"`
	ImageURL string
}