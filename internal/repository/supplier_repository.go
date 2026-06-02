package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type SupplierRepository interface {
	FindAll() ([]*domain.Supplier, error)
	FindByID(id uint) (*domain.Supplier, error)
	Create(supplier *domain.Supplier) error
	Update(supplier *domain.Supplier) error
	Delete(id uint) error
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) FindAll() ([]*domain.Supplier, error) {
	var suppliers []*domain.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *supplierRepository) FindByID(id uint) (*domain.Supplier, error) {
	var supplier domain.Supplier
	err := r.db.Find(&supplier, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &supplier, err

}

func (r *supplierRepository) Create(supplier *domain.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) Update(supplier *domain.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *supplierRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Supplier{}, id).Error
}
