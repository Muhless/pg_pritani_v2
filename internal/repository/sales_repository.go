package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type SalesRepository interface {
	FindAll() ([]*domain.Sales, error)
	FindByID(id uint) (*domain.Sales, error)
	Create(sales *domain.Sales) error
	Update(sales *domain.Sales) error
	Delete(id uint) error
}

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) SalesRepository {
	return &salesRepository{db}
}

func (r *salesRepository) FindAll() ([]*domain.Sales, error) {
	var sales []*domain.Sales
	err := r.db.Preload("Customer").Preload("Items.Product").Find(&sales).Error
	return sales, err
}

func (r *salesRepository) FindByID(id uint) (*domain.Sales, error) {
	var sales domain.Sales
	err := r.db.Preload("Customer").Preload("Items.Product").First(&sales, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &sales, nil
}

func (r *salesRepository) Create(sales *domain.Sales) error {
	return r.db.Create(sales).Error
}

func (r *salesRepository) Update(sales *domain.Sales) error {
	return r.db.Save(sales).Error
}

func (r *salesRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Sales{}, id).Error
}
