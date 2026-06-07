package domain

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	PaymentMethodCash     PaymentMethod = "cash"
	PaymentMethodTransfer PaymentMethod = "transfer"
)

type Payment struct {
	gorm.Model
	SalesID uint          `json:"sales_id" gorm:"not null"`
	Sales   Sales         `json:"sales" gorm:"foreignKey:SalesID"`
	Amount  float64       `json:"amount" gorm:"type:numeric(10,2)"`
	Method  PaymentMethod `json:"method" gorm:"varchar(20);not null"`
	Date    time.Time     `json:"date" gorm:"not null"`
	Note    string        `json:"note"`
}
