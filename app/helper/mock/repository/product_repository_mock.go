package repository

import (
	"github.com/stretchr/testify/mock"
	"gokomodo-assesment/app/domain/entity"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetListProductBySellerID(in string) ([]*entity.Product, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	response := res.([]*entity.Product)
	return response, nil
}

func (m *MockProductRepository) GetListProduct() ([]*entity.Product, error) {
	call := m.Called()
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	response := res.([]*entity.Product)
	return response, nil
}

func (m *MockProductRepository) AddProduct(in *entity.Product) error {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil
	}

	return call.Error(0)
}
