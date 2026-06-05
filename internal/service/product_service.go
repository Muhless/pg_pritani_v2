package service

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	"pg_pritani/backend/internal/repository"

	"github.com/rs/zerolog/log"
)

type ProductService interface {
	GetAll() ([]*domain.Product, error)
	GetByID(id uint) (*domain.Product, error)
	Create(req dto.CreateProductRequest) error
	Update(id uint, req dto.UpdateProductRequest) error
	Delete(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAll() ([]*domain.Product, error) {
	products, err := s.repo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg("failed to get all products data")
		return nil, err
	}
	return products, nil
}

func (s *productService) GetByID(id uint) (*domain.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to get product by id")
		return nil, err
	}

	if product == nil {
		log.Warn().Uint("id", id).Msg("product not found")
		return nil, errors.New("product not found")
	}
	return product, nil

}

func (s *productService) Create(req dto.CreateProductRequest) error {
	product := &domain.Product{
		Name:  req.Name,
		Type:  req.Type,
		Stock: req.Stock,
		Price: req.Price,
		Photo: req.Photo,
	}

	if err := s.repo.Create(product); err != nil {
		log.Error().Err(err).Msg("failed to create product data")
		return err
	}

	log.Info().Str("name", product.Name).Msg("product data successfully created")
	return nil
}

func (s *productService) Update(id uint, req dto.UpdateProductRequest) error {
	product, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to find product data")
		return err
	}

	if product == nil {
		log.Warn().Uint("id", id).Msg("product not found")
		return errors.New("product not found")
	}

	product.Name = req.Name
	product.Type = req.Type
	product.Stock = req.Stock
	product.Price = req.Price
	product.Photo = req.Photo

	if err := s.repo.Update(product); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to update product data")
		return err
	}
	log.Info().Uint("id", id).Msg("product data successfully updated")
	return nil

}

func (s *productService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		log.Error().Err(err).Uint("id", id).Msg("failed to delete product data")
	}
	log.Info().Uint("id", id).Msg("successfully deleted product data")
	return nil
}
