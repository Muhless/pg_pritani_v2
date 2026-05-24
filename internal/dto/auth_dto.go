package dto

import "pg_pritani/backend/internal/domain"

type RegisterRequest struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     domain.Role `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
