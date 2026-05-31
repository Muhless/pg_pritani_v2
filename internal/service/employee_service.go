package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"

	"github.com/rs/zerolog/log"
)

type EmployeeService interface {
	GetAll() ([]*domain.Employee, error)
	GetByID(id uint) (*domain.Employee, error)
	Update(id uint, req dto.UpdateEmployeeRequest) error
	Delete(id uint) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}

func (s *employeeService) GetAll() ([]*domain.Employee, error) {
	employees, err := s.repo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to get all employees data")
		return nil, err
	}
	return employees, nil
}

func (s *employeeService) GetByID(id uint) (*domain.Employee, error) {
	employee, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("id not found")
		return nil, err
	}
	if employee == nil {
		log.Warn().Uint("id", id).Msg("employee data not found")
		return nil, err
	}
	return employee, nil
}

func (s *employeeService) Update(id uint, req dto.UpdateEmployeeRequest) error {
	employee, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("id not found")
		return err
	}
	if employee == nil {
		log.Warn().Uint("id", id).Msg("employee not found")
		return errors.New("employee not found")
	}

	employee.Name = req.Name
	employee.Email = req.Email
	employee.Phone = req.Phone
	employee.Address = req.Address
	employee.Photo = req.Photo
	employee.IsActive = req.IsActive

	if err := s.repo.Update(employee); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to update employee data")
		return err
	}
	log.Info().Uint("id", id).Msg("employee data successfully updated")
	return nil

}

func (s *employeeService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("employee data not found")
		return err
	}
	log.Info().Uint("id", id).Msg("employee data successfully deleted")
	return nil
}
