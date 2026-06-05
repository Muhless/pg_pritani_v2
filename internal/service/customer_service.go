package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"

	"github.com/rs/zerolog/log"
)

type CustomerService interface {
	GetAll() ([]*domain.Customer, error)
	GetByID(id uint) (*domain.Customer, error)
	Create(req dto.CreateCustomerRequest) error
	Update(id uint, req dto.UpdateCustomerRequest) error
	Delete(id uint) error
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

func (s *customerService) Create(req dto.CreateCustomerRequest) error {
	customer := &domain.Customer{
		Name:    req.Name,
		Company: req.Company,
		Phone:   req.Phone,
		Address: req.Address,
	}

	if err := s.repo.Create(customer); err != nil {
		log.Error().Err(err).Msg("failed to create customer data")
		return err
	}
	log.Info().Str("name", customer.Name).Msg("customer data successfully created")
	return nil
}

func (s *customerService) Update(id uint, req dto.UpdateCustomerRequest) error {
	customer, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to find customer data by id")
	}

	if customer == nil {
		log.Warn().Uint("id", id).Msg("customer data not found")
		return errors.New("customer data not found")
	}

	customer.Name = req.Name
	customer.Company = req.Company
	customer.Phone = req.Phone
	customer.Address = req.Address

	if err := s.repo.Update(customer); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("update customer data failed")
		return err
	}

	log.Info().Uint("id", id).Msg("successfully updated customer data")
	return nil
}

func (s *customerService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to get customer data by id")
		return err
	}
	log.Info().Uint("id", id).Msg("successfully deleted customer data")
	return nil
}
