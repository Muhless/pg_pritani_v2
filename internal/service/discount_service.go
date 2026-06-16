package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"
	"time"

	"github.com/rs/zerolog/log"
)

type DiscountService interface {
	GetAll(page, limit, offset int) ([]*domain.Discount, error)
	GetByID(id uint) (*domain.Discount, error)
	Create(req dto.CreateDiscountRequest) error
	Update(id uint, req dto.UpdateDiscountRequest) error
	Delete(id uint) error
}

type discountService struct {
	repo repository.DiscountRepository
}

func NewDiscountService(repo repository.DiscountRepository) DiscountService {
	return &discountService{repo}
}

func (s *discountService) GetAll(page, limit, offset int) ([]*domain.Discount, int64, error) {
	offset = (page - 1) * limit
	discounts, total, err := s.repo.FindAll(page, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("failed to get all discounts data")
		return nil, 0, err
	}
	return discounts, total, nil
}

func (s *discountService) GetByID(id uint) (*domain.Discount, error) {
	discount, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to get discount data by id")
		return nil, err
	}

	if discount == nil {
		log.Warn().Uint("id", id).Msg("discount data not found")
		return nil, errors.New("discount data not found")
	}

	return discount, nil
}

func (s *discountService) Create(req dto.CreateDiscountRequest) error {
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return errors.New("invalid start_date format, use YYYY-MM-DD")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return errors.New("invalid end_date format,use YYYY-MM-DD")
	}

	if endDate.Before(startDate) {
		return errors.New("end_date must be after start_date")
	}

	discount := &domain.Discount{
		ProductID: req.ProductID,
		Name:      req.Name,
		Type:      domain.DiscountType(req.Type),
		Value:     req.Value,
		StartDate: startDate,
		EndDate:   endDate,
		IsActive:  req.IsActive,
	}

	if err := s.repo.Create(discount); err != nil {
		log.Error().Err(err).Msg("failed to create discount data")
		return err
	}

	log.Info().Str("name", discount.Name).Msg("successfully created discount data")
	return nil
}

func (s *discountService) Update(id uint, req dto.UpdateDiscountRequest) error {
	discount, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to get discount data by id")
		return err
	}

	if discount == nil {
		return errors.New("discount data not found")
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return errors.New("invalid start_date format, use YYYY-MM-DD")
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return errors.New("invalid end_date format, use YYYY-MM-DD")
	}

	if endDate.Before(startDate) {
		return errors.New("end_date must be after start_date")
	}

	discount.Name = req.Name
	discount.Type = domain.DiscountType(req.Type)
	discount.Value = req.Value
	discount.StartDate = startDate
	discount.EndDate = endDate
	discount.IsActive = req.IsActive

	if err := s.repo.Update(discount); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to update discount data")
		return err
	}

	log.Info().Uint("id", id).Msg("successfully updated discount data")
	return nil
}

func (s *discountService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to delete discount data")
		return err
	}
	log.Info().Uint("id", id).Msg("successfully deleted discount data")
	return nil
}
