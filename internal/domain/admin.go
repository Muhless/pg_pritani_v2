package domain

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	UserID   uint        `json:"user_id" gorm:"not null"`
	User     domain.User `json:"user" gorm:"foreignKey:UserID"`
	Name     string      `json:"name" gorm:"not null"`
	Email    string      `json:"email" gorm:"uniqueIndex"`
	Phone    string      `json:"phone" gorm:"type:varchar(15);uniqueIndex;not null"`
	Photo    string      `json:"photo"`
	IsActive bool        `json:"is_active" gorm:"not null;default:true"`
}
