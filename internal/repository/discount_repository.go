package repository

import (
	"pg_pritani/backend/internal/domain"
	"time"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	FindAll(page, limit, offset int) ([]*domain.Discount, int64, error)
	FindByID(id uint) (*domain.Discount, error)
	FindByActiveProductID(ProductID uint) (*domain.Discount, error)
	Create(discount *domain.Discount) error
	Update(discount *domain.Discount) error
	Delete(id uint) error
}

type discountRepository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &discountRepository{db}
}

func (r *discountRepository) FindAll(page, limit, offset int) ([]*domain.Discount, int64, error) {
	var discounts []*domain.Discount
	var total int64

	if err := r.db.Model(&domain.Discount{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.Preload("Product").Limit(limit).Offset(offset).Find(&discounts).Error
	return discounts, total, err
}

func (r *discountRepository) FindByID(id uint) (*domain.Discount, error) {
	var discount domain.Discount
	err := r.db.Preload("Product").First(&discount).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &discount, nil
}

func (r *discountRepository) FindByActiveProductID(productID uint) (*domain.Discount, error) {
	var discount domain.Discount
	now := time.Now()
	err := r.db.Where(
		"product_id = ? AND is_active = true AND start_date<=? AND end_date >= ?",
		productID, now, now,
	).First(&discount).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &discount, nil
}

func (r *discountRepository) Create(discount *domain.Discount) error {
	return r.db.Create(discount).Error
}

func (r *discountRepository) Update(discount *domain.Discount) error {
	return r.db.Save(discount).Error
}

func (r *discountRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Discount{}, id).Error
}
