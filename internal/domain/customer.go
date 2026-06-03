package domain

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone" gorm:"unique"`
	Address string `json:"address"`
	Company string `json:"company"`
}
