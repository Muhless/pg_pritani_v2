package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]*domain.Product, error)
	FindByID(id uint) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll() ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Find(&product, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &product, err

}

func (r *productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}
