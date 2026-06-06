package domain

import (
	"time"

	"gorm.io/gorm"
)

type SalesStatus string

const (
	SalesStatusPending   SalesStatus = "pending"
	SalesStatusPartial   SalesStatus = "partial"
	SalesStatusPaid      SalesStatus = "paid"
	SalesStatusCancelled SalesStatus = "cancelled"
)

type Sales struct {
	gorm.Model
	CustomerID      *uint       `json:"customer_id"`
	Customer        *Customer   `json:"customer" gorm:"foreignKey:CustomerID"`
	TotalPrice      float64     `json:"total_price" gorm:"not null;type:numeric(10,2)"`
	PaidAmount      float64     `json:"paid_amount" gorm:"not null;default:0"`
	RemainingAmount float64     `json:"remaining_amount" gorm:"type:numeric(10,2)"`
	Status          SalesStatus `json:"status" gorm:"not null;default:'pending'"`
	Date            time.Time   `json:"date" gorm:"not null"`
	Items           []SalesItem `json:"items" gorm:"foreignKey:SalesID"`
}
