package dto

type SalesItemRequest struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type CreateSalesRequest struct {
	CustomerID *uint              `json:"customer_id"`
	PaidAmount float64            `json:"paid_amount"`
	Date       string             `json:"date"`
	Items      []SalesItemRequest `json:"items"`
}

type UpdateSalesStatusRequest struct {
	PaidAmount float64 `json:"padi_amount"`
}
