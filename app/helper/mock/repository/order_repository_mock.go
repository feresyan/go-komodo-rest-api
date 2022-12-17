package repository

import (
	"github.com/stretchr/testify/mock"
	"gokomodo-assesment/app/domain/entity"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) GetOrderBySellerID(in string) ([]*entity.Order, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	response := res.([]*entity.Order)
	return response, call.Error(1)
}

func (m *MockOrderRepository) GetOrderByBuyerID(in string) ([]*entity.Order, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	response := res.([]*entity.Order)
	return response, nil
}

func (m *MockOrderRepository) AcceptOrder(in *entity.Order) error {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil
	}

	return call.Error(0)
}

func (m *MockOrderRepository) AddOrder(in *entity.Order) error {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil
	}

	return call.Error(0)
}
