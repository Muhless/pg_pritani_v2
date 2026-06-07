package dto

import "time"

type CreatePaymentRequest struct {
	SalesID uint      `json:"sales_id"`
	Amount  float64   `json:"amount"`
	Method  string    `json:"method"`
	Date    time.Time `json:"time"`
	Note    string    `json:"note"`
}
