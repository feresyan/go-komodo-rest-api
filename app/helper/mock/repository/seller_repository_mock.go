package repository

import (
	"github.com/stretchr/testify/mock"
	"gokomodo-assesment/app/domain/entity"
)

type MockSellerRepository struct {
	mock.Mock
}

func (m *MockSellerRepository) GetSeller(in *entity.Seller) (*entity.Seller, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	response := res.(*entity.Seller)
	return response, nil
}
