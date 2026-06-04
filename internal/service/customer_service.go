package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"

	"github.com/rs/zerolog/log"
)

type CustomerService interface {
	GetAll() []*domain.Customer
	GetByID(id uint) (*domain.Customer, error)
	Create(req dto.CreateCustomerRequest) error
	Update(req dto.UpdateCustomerRequest) error
	Delete(id uint)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo}
}

func (s *customerService) GetAll() ([]*domain.Customer, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to get all customers data")
		return nil, err
	}
	return customers, err
}

func (s *customerService) GetByID(id uint) (*domain.Customer, error) {
	customer, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to get customer data by id")
		return nil, err
	}
	if customer == nil {
		log.Warn().Uint("id", id).Msg("customer data not found")
		return nil, errors.New("customer data not found")
	}

	return customer, nil
}

