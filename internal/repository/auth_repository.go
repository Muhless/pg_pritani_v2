package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Register(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User

	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
