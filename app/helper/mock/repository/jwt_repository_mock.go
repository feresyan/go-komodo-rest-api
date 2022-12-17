package repository

import (
	"github.com/stretchr/testify/mock"
)

type MockJWTRepository struct {
	mock.Mock
}

func (m *MockJWTRepository) CheckPassword(userPassword string, loginPassword string) error {
	call := m.Called(userPassword, loginPassword)
	res := call.Get(0)
	if res == nil {
		return nil
	}
	return call.Error(0)
}

func (m *MockJWTRepository) GenerateJWT(email, role string) (string, error) {
	call := m.Called(email, role)
	res := call.Get(0)
	if res == "" {
		return "", call.Error(1)
	}

	return "", nil
}
