package repository

import (
	"github.com/stretchr/testify/mock"
	"gokomodo-assesment/app/domain/entity"
)

type MockBuyerRepository struct {
	mock.Mock
}

func (m *MockBuyerRepository) GetBuyer(in *entity.Buyer) (*entity.Buyer, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	response := res.(*entity.Buyer)
	return response, nil
}
