package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"
)

type UserService interface {
	GetByID(id uint) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Update(id uint, req dto.UpdateUserRequest) error
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetByID(id uint) (*domain.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return nil, err
}

func (s *userService) GetAll() ([]*domain.User, error) {
	return s.repo.FindAll()
}

func (s *userService) Update(id uint, req dto.UpdateUserRequest) error {
	user, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	user.Username = req.Username
	user.Role = req.Role

	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
