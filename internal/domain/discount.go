package domain

import (
	"time"

	"gorm.io/gorm"
)

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "percentage"
	DiscountTypeNominal    DiscountType = "nominal"
)

type Discount struct {
	gorm.Model
	ProductID uint         `json:"product_id" gorm:"not null"`
	Product   Product      `json:"product" gorm:"foreignKey:ProductID"`
	Name      string       `json:"name" gorm:"not null;type:varchar(50)"`
	Type      DiscountType `json:"type" gorm:"not null;type:varchar(50)"`
	Value     float64      `json:"value" gorm:"not null;numeric(10,2)"`
	StartDate time.Time    `json:"start_date" gorm:"not null"`
	EndDate   time.Time    `json:"end_date" gorm:"not null"`
	IsActive  bool         `json:"is_active" gorm:"not null;default:true"`
}
