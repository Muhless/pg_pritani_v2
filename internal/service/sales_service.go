package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type SalesService interface {
	GetAll() ([]*domain.Sales, error)
	GetByID(id uint) (*domain.Sales, error)
	Create(req dto.CreateSalesRequest) error
	AddPayment(id uint, req dto.UpdateSalesStatusRequest) error
	Delete(id uint) error
}

type salesService struct {
	db          *gorm.DB
	salesRepo   repository.SalesRepository
	productRepo repository.ProductRepository
}

func NewSalesService(
	db *gorm.DB,
	salesRepo repository.SalesRepository,
	productRepo repository.ProductRepository,
) SalesService {
	return &salesService{db, salesRepo, productRepo}
}

func (s *salesService) GetAll() ([]*domain.Sales, error) {
	sales, err := s.salesRepo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to get all sales data")
		return nil, err
	}

	return sales, nil
}

func (s *salesService) GetByID(id uint) (*domain.Sales, error) {
	sales, err := s.salesRepo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to get sales data by id")
		return nil, err
	}
	if sales == nil {
		log.Warn().Uint("id", id).Msg("sales not found")
		return nil, errors.New("sales not found")
	}
	return sales, nil
}

func (s *salesService) Create(req dto.CreateSalesRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			return errors.New("invalid date format, use YYYY-MM-DD")
		}

		var totalPrice float64
		var items []domain.SalesItem

		for _, item := range req.Items {
			product, err := s.productRepo.FindByID(item.ProductID)
			if err != nil {
				return err
			}
			if product == nil {
				errors.New("product not found")
			}

			if product.Stock < item.Quantity {
				log.Warn().Str("product", product.Name).Msg("insufficient stock")
				return errors.New("insufficient stock for product: " + product.Name)
			}

			subtotal := product.Price * float64(item.Quantity)
			totalPrice += subtotal

			items = append(items, domain.SalesItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
				SubTotal:  subtotal,
			})

			product.Stock -= item.Quantity
			if err := tx.Save(product).Error; err != nil {
				return err
			}
		}

		// status
		status := domain.SalesStatusPending
		if req.PaidAmount >= totalPrice {
			status = domain.SalesStatusPaid
		} else if req.PaidAmount > 0 {
			status = domain.SalesStatusPartial
		}

		sales := &domain.Sales{
			CustomerID:      req.CustomerID,
			TotalPrice:      totalPrice,
			PaidAmount:      req.PaidAmount,
			RemainingAmount: totalPrice * req.PaidAmount,
			Status:          status,
			Date:            date,
			Items:           items,
		}

		if err := tx.Create(sales).Error; err != nil {
			log.Error().Err(err).Msg("failed to create sales data")
			return err
		}

		log.Info().Uint("id", sales.ID).Float64("total", totalPrice).Msg("successfully created sales data")
		return nil
	})
}
