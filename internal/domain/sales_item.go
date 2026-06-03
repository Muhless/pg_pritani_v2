package domain

import "gorm.io/gorm"

type SalesItem struct {
	gorm.Model
	SalesID   uint    `json:"sales_id" gorm:"not null"`
	Sales     Sales   `json:"sales" gorm:"foreqignKey:SalesID"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null;type:numeric(10,2)"`
	SubTotal  float64 `json:"sub_total" gorm:"not null;type:numeric(10,2)"`
}
