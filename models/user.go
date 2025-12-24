package models

type User struct {
	BaseModel

	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:150;uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"size:20;default:user"`
}
