package mock

import (
	"pg_pritani/backend/internal/domain"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) FindAll(page, limit, offset int) ([]*domain.Product, int64, error) {
	args := m.Called(page, limit, offset)
	return args.Get(0).([]*domain.Product), args.Get(1).(int64), args.Error(2)
}

func (m *ProductRepositoryMock) FindByID(id uint) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Create(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) Update(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *ProductRepositoryMock) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
