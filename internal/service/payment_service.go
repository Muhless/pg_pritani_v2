package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type PaymentService interface {
	GetAll() ([]*domain.Payment, error)
	GetByID(id uint) (*domain.Payment, error)
	GetBySalesID(salesID uint) ([]*domain.Payment, error)
	Create(req dto.CreatePaymentRequest) error
	Delete(id uint) error
}

type paymentService struct {
	db          *gorm.DB
	paymentRepo repository.PaymentRepository
	salesRepo   repository.SalesRepository
}

func NewPaymentService(db *gorm.DB, paymentRepo repository.PaymentRepository, salesRepo repository.SalesRepository) PaymentService {
	return &paymentService{db, paymentRepo, salesRepo}
}

func (s *paymentService) GetAll() ([]*domain.Payment, error) {
	payments, err := s.paymentRepo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to get all payments data")
		return nil, err
	}

	return payments, err
}

func (s *paymentService) GetByID(id uint) (*domain.Payment, error) {
	payment, err := s.paymentRepo.FindByID(id)
	if err != nil {
		log.Info().Uint("id", id).Msg("failed to get payment data by id")
		return nil, err
	}

	if payment == nil {
		log.Warn().Uint("id", id).Msg("payment data not found")
		return nil, errors.New("payment data not found")
	}
	return payment, nil
}

func (s *paymentService) GetBySalesID(salesID uint) ([]*domain.Payment, error) {
	payments, err := s.paymentRepo.FindBySalesID(salesID)
	if err != nil {
		log.Info().Uint("sales_id", salesID).Msg("failed to get all payments data by sales id")
		return nil, err
	}
	return payments, nil

}

func (s *paymentService) Create(req dto.CreatePaymentRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		sales, err := s.salesRepo.FindByID(req.SalesID)
		if err != nil {
			return err
		}

		if sales == nil {
			return errors.New("sales data not found")
		}

		if sales.Status == domain.SalesStatusPaid {
			return errors.New("sales already paid")
		}
		if sales.Status == domain.SalesStatusCancelled {
			return errors.New("sales is cancelled")
		}

		payment := &domain.Payment{
			SalesID: req.SalesID,
			Amount:  req.Amount,
			Method:  domain.PaymentMethod(req.Method),
			Date:    req.Date,
			Note:    req.Note,
		}

		if err := tx.Create(payment).Error; err != nil {
			log.Error().Err(err).Msg("failed to create payment data")
			return err
		}

		sales.PaidAmount += req.Amount
		sales.RemainingAmount = sales.TotalPrice - sales.PaidAmount

		if sales.PaidAmount >= sales.TotalPrice {
			sales.Status = domain.SalesStatusPaid
			sales.RemainingAmount = 0
		} else {
			sales.Status = domain.SalesStatusPartial
		}

		if err := tx.Save(sales).Error; err != nil {
			log.Error().Err(err).Msg("failed to update sales data after payment")
			return err
		}

		log.Info().Uint("sales_id", req.SalesID).Float64("amount", req.Amount).Msg("successfully created payment data")
		return nil
	})
}

func (s *paymentService) Delete(id uint) error {
	if err := s.paymentRepo.Delete(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to delete payment data by id")
		return err
	}
	log.Info().Uint("id", id).Msg("successfully deleted payment data")
	return nil
}
