package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"

	"github.com/rs/zerolog/log"
)

type SupplierService interface {
	GetAll() ([]*domain.Supplier, error)
	GetByID(id uint) (*domain.Supplier, error)
	Create(req dto.CreateSupplierRequest) error
	Update(id uint, req dto.UpdateSupplierRequest) error
	Delete(id uint) error
}

type supplierService struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierService{repo}
}

func (s *supplierService) GetAll() ([]*domain.Supplier, error) {
	suppliers, err := s.repo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to get all suppliers data")
		return nil, err
	}
	return suppliers, err
}

func (s *supplierService) GetByID(id uint) (*domain.Supplier, error) {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Msg("failed to get supplier id")
		return nil, err
	}

	if supplier == nil {
		log.Warn().Uint("id", id).Msg("supplier id not found")
		return nil, errors.New("supplier id not found")
	}
	return supplier, nil
}

func (s *supplierService) Create(req dto.CreateSupplierRequest) error {
	supplier := &domain.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	if err := s.repo.Create(supplier); err != nil {
		log.Error().Err(err).Msg("failed to create new supplier data")
		return err
	}

	log.Info().Str("name", supplier.Name).Msg("supplier data successfully created")
	return nil
}

func (s *supplierService) Update(id uint, req dto.UpdateSupplierRequest) error {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to get supplier by id")
		return err
	}

	if supplier == nil {
		log.Warn().Uint("id", id).Msg("supplier data not found")
		return errors.New("supplier data not found")
	}

	supplier.Name = req.Name
	supplier.Email = req.Email
	supplier.Phone = req.Phone
	supplier.Address = req.Address

	if err := s.repo.Update(supplier); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to update supplier data")
		return err
	}
	log.Info().Uint("id", id).Msg("supplier data successfully updated")
	return nil
}

func (s *supplierService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed delete supplier data")
	}
	log.Info().Uint("id", id).Msg("supplier data successfully deleted")
	return nil
}
