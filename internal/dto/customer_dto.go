package dto

type CreateCustomerRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone" gorm:"unique"`
	Address string `json:"address"`
	Company string `json:"company"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone" gorm:"unique"`
	Address string `json:"address"`
	Company string `json:"company"`
}