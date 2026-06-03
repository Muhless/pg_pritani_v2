package domain

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null;type:varchar(50)"`
	Email   string `json:"email" gorm:"unique"`
	Phone   string `json:"phone" gorm:"not null,unique;type:varchar(15)"`
	Address string `json:"address"`
}
