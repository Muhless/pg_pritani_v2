package dto

type CreateDiscountRequest struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Name      string  `json:"name" binding:"required,min=2,max=100"`
	Type      string  `json:"type" binding:"required,oneof=percentage nominal"`
	Value     float64 `json:"value" binding:"required,gt=0"`
	StartDate string  `json:"start_date" binding:"required"`
	EndDate   string  `json:"end_date" binding:"required"`
	IsActive  bool    `json:"is_active"`
}

type UpdateDiscountRequest struct {
	Name      string  `json:"name" binding:"required,min=2,max=100"`
	Type      string  `json:"type" binding:"required,oneof=percentage nominal"`
	Value     float64 `json:"value" binding:"required,gt=0"`
	StartDate string  `json:"start_date" binding:"required"`
	EndDate   string  `json:"end_date" binding:"required"`
	IsActive  bool    `json:"is_active"`
}
