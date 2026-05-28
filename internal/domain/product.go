package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name" gorm:"not null"`
	Type  string  `json:"type" gorm:"not null"`
	Stock int     `json:"stock" gorm:"not null;default:0"`
	Price float64 `json:"price" gorm:"not null;type:numeric(10,2)"`
	Photo string  `json:"photo" gorm:"type:text"`
}
