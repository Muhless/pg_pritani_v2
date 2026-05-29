package domain

import "gorm.io/gorm"

type Sales struct {
	gorm.Model
	Admin       Admin   `json:"admin" gorm:"not null"`
	PaidAmount  float64 `json:"paid_amount" gorm:"not null"`
	TotalAmount float64 `json:"total_amount" gorm:"not null"`
}
