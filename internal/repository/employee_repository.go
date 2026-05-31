package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll() ([]*domain.Employee, error)
	FindByID(id uint) (*domain.Employee, error)
	Update(employee *domain.Employee) error
	Delete(id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) FindAll() ([]*domain.Employee, error) {
	var employees []*domain.Employee
	err := r.db.Find(&employees).Error
	return employees, err
}

func (r *employeeRepository) FindByID(id uint) (*domain.Employee, error) {
	var employee domain.Employee
	err := r.db.First(&employee, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeRepository) Update(employee *domain.Employee) error {
	return r.db.Save(employee).Error
}

func (r *employeeRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Employee{}, id).Error
}
