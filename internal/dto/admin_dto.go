package dto

type UpdateAdminRequest struct {
	Name     string `json:"admin"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Photo    string `json:"photo"`
	IsActive bool   `json:"is_active"`
}
