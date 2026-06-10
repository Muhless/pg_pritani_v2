package repository

import (
	"pg_pritani/backend/internal/domain"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	FindAll(page, limit, offset int) ([]*domain.Payment, int64, error)
	FindByID(id uint) (*domain.Payment, error)
	FindBySalesID(salesID uint) ([]*domain.Payment, error)
	Create(payment *domain.Payment) error
	Delete(id uint) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) FindAll(page, limit, offset int) ([]*domain.Payment, int64, error) {
	var payments []*domain.Payment
	var total int64

	if err := r.db.Model(&domain.Payment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.Limit(limit).Offset(offset).Preload("Sales").Find(&payments).Error
	return payments, total, err
}

func (r *paymentRepository) FindByID(id uint) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.Preload("Sales").First(&payment, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *paymentRepository) FindBySalesID(salesID uint) ([]*domain.Payment, error) {
	var payments []*domain.Payment
	err := r.db.Where("sales_id = ?", salesID).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) Create(payment *domain.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Payment{}, id).Error
}
