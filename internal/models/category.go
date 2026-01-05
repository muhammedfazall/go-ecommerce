package models

type Category struct {
	BaseModel
	Name string `gorm:"size:100;unique;not null"`
}
