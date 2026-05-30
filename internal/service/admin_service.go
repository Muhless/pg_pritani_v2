package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"
)

type AdminService interface {
	GetAll() ([]*domain.Admin, error)
	GetByID(id uint) (*domain.Admin, error)
	Update(id uint, req dto.UpdateAdminRequest) error
	Delete(id uint) error
}

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{repo}
}

func (s *adminService) GetAll() ([]*domain.Admin, error) {
	return s.repo.FindAll()
}

func (s *adminService) GetByID(id uint) (*domain.Admin, error) {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if admin == nil {
		return nil, errors.New("admin not found")
	}
	return admin, err
}

func (s *adminService) Update(id uint, req dto.UpdateAdminRequest) error {
	admin, err := s.repo.FindByID(id)
	if err != nil {
		return nil
	}
	if admin == nil {
		return errors.New("id not found")
	}

	admin.Name = req.Name
	admin.Email = req.Email
	admin.Phone = req.Phone
	admin.Photo = req.Photo
	admin.IsActive = req.IsActive

	return s.repo.Update(admin)
}

func (s *adminService) Delete(id uint) error {
	return s.repo.Delete(id)
}
