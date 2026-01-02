package models

type User struct {
	BaseModel

	Username  string `gorm:"size:100;not null"`
	Email     string `gorm:"size:150;uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"size:20;default:user"`
	Status    string `gorm:"size:20;default:active"`
	IsBlocked bool   `gorm:"default:false"`
}
