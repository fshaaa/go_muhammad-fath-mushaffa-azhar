package repository

import (
	"clean-archi-v2/model"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepo() *MockUserRepository {
	return &MockUserRepository{}
}

func (m *MockUserRepository) Create(user model.User) error {
	ret := m.Called(user)
	return ret.Error(0)
}

func (m *MockUserRepository) Find() ([]model.User, error) {
	ret := m.Called()
	return ret.Get(0).([]model.User), ret.Error(1)
}
