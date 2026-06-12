package service_test

import (
	"errors"
	"pg_pritani/backend/internal/domain"
	"pg_pritani/backend/internal/dto"
	repoMock "pg_pritani/backend/internal/repository/mock"
	"pg_pritani/backend/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllProducts(t *testing.T) {
	mockRepo := new(repoMock.ProductRepositoryMock)
	svc := service.NewProductService(mockRepo)

	expected := []*domain.Product{
		{Name: "Produt A", Price: 20000, Stock: 10},
		{Name: "Produt B", Price: 40000, Stock: 20},
	}

	mockRepo.On("FindAll", 1, 10, 0).Return(expected, int64(2), nil)

	products, total, err := svc.GetAll(1, 10)
	assert.Error(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, products, 2)
	mockRepo.AssertExpectations(t)

}

func TestGetProductByID(t *testing.T) {
	mockRepo := new(repoMock.ProductRepositoryMock)
	svc := service.NewProductService(mockRepo)

	expected := &domain.Product{Name: "Product A", Price: 30000}

	mockRepo.On("FindByID", uint(1)).Return(expected, nil)

	product, err := svc.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Product A", product.Name)
	mockRepo.AssertExpectations(t)

}

func TestCreateProduct(t *testing.T) {
	mockRepo := new(repoMock.ProductRepositoryMock)
	svc := service.NewProductService(mockRepo)

	req := dto.CreateProductRequest{
		Name:  "Product A",
		Type:  "Type A",
		Stock: 10,
		Price: 30000,
	}

	mockRepo.On("Create", mock.AnythingOfType("*domain.Product")).Return(nil)
	err := svc.Create(req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}

func TestCreateProduct_Failed(t *testing.T) {
	mockRepo := new(repoMock.ProductRepositoryMock)
	svc := service.NewProductService(mockRepo)

	req := dto.CreateProductRequest{
		Name:  "Product A",
		Type:  "Type A",
		Stock: 10,
		Price: 10000,
	}

	mockRepo.On("Create", mock.AnythingOfType("*domain.Product")).Return(errors.New("database error"))

	err := svc.Create(req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(repoMock.ProductRepositoryMock)
	svc := service.NewProductService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)
	err := svc.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}
