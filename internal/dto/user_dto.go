package dto

import "pg_pritani/backend/internal/domain"

type UpdateUserRequest struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Role     domain.Role `json:"role"`
}
