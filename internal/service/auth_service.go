package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (*domain.User, error)
}

type authService struct {
	db   *gorm.DB
	repo repository.UserRepository
}

func NewAuthService(db *gorm.DB, repo repository.UserRepository) AuthService {
	return &authService{db, repo}
}

func (s *authService) Register(req dto.RegisterRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil
		}

		user := &domain.User{
			Username: req.Username,
			Password: string(hashed),
			Role:     req.Role,
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		switch user.Role {
		case domain.RoleAdmin:
			return tx.Create(&domain.Admin{UserID: user.ID, Name: req.Username}).Error
		}
		return nil
	})
}

func (s *authService) Login(req dto.LoginRequest) (*domain.User, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	return user, nil
}
