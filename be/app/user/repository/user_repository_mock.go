package repository

import (
	"errors"
	"module/domain"
	"module/dto/user"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

// Called make sure that mock call the method with exact parameter
func (m *RepositoryMock) CreateUser(dto *user.UserRegister) (*domain.User, error) {
	args := m.Mock.Called(dto)
	if args.Get(0) == nil {
		return nil, errors.New("user required")
	}
	user, ok := args.Get(0).(domain.User)
	if !ok {
		return nil, errors.New("user required")
	}
	return &user, nil
}

func (m *RepositoryMock) Login(dto *user.UserLogin) (*domain.User, error) {
	args := m.Mock.Called(dto)
	if args.Get(0) == nil {
		return nil, errors.New("user required")
	}
	user, ok := args.Get(0).(domain.User)
	if !ok {
		return nil, errors.New("user required")
	}
	return &user, nil
}

func (m *RepositoryMock) SaveToken(token string, id string) error {
	args := m.Mock.Called(token, id)
	if args.Get(0) == nil {
		return nil
	}
	err := args.Get(0).(error)
	return errors.New(err.Error())
}

func (m *RepositoryMock) GetUserByEmail(email string) (*domain.User, error) {
	args := m.Mock.Called(email)
	if args.Get(0) == nil {
		return nil, errors.New("email argument required")
	}
	return &domain.User{
		Name:     "mock",
		Email:    "mock",
		Password: "mock",
	}, nil
}

func (m *RepositoryMock) GetUserById(id string) (*domain.User, error) {
	args := m.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, errors.New("argument required")
	}
	user, ok := args.Get(0).(domain.User)
	if !ok {
		return nil, errors.New("user required")
	}
	return &user, nil
}
