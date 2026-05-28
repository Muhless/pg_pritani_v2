package domain

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"type:varchar(20);uniqueIndex"`
	Phone    string `json:"phone" gorm:"type:varchar(15);uniqueIndex"`
	Address  string `json:"address" gorm:"type:varchar(50)"`
	Photo    string `json:"photo"`
	IsActive bool   `json:"is_active" gorm:"not null;default:true"`
}
