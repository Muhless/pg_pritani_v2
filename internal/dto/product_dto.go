package dto

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
	Photo string  `json:"photo"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
	Photo string  `json:"photo"`
}
