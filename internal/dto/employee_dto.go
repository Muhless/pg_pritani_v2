package dto

type UpdateEmployeeRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Photo    string `json:"photo"`
	IsActive bool   `json:"is_active"`
}
