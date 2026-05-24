package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(username, password string, role domain.Role) error
	Login(username, password string) (*domain.User, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(username, password string, role domain.Role) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	user := &domain.User{
		Username: username,
		Password: string(hashed),
		Role:     role,
	}

	return s.repo.Register(user)
}

func (s *authService) Login(username, password string) (*domain.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return user, nil
}
