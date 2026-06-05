package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll() ([]*domain.Customer, error)
	FindByID(id uint) (*domain.Customer, error)
	Create(customer *domain.Customer) error
	Update(customer *domain.Customer) error
	Delete(id uint) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) FindAll() ([]*domain.Customer, error) {
	var customers []*domain.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *customerRepository) FindByID(id uint) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.First(&customer, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) Create(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) Update(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Customer{}, id).Error
}
