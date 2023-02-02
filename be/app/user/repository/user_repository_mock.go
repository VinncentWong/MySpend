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

func (m *RepositoryMock) CreateUser(dto *user.UserRegister) (*domain.User, error) {
	args := m.Mock.Called(dto)
	if args.Get(0) == nil {
		return nil, errors.New("no user register dto")
	}
	register, ok := args.Get(0).(*user.UserRegister)
	if !ok {
		return nil, errors.New("wrong argument type")
	}
	user := domain.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: register.Password,
	}
	return &user, nil
}

func (m *RepositoryMock) Login(dto *user.UserLogin) (*domain.User, error) {
	args := m.Mock.Called(dto)
	if args.Get(0) == nil {
		return nil, errors.New("no user register dto")
	}
	login, ok := args.Get(0).(*user.UserLogin)
	if !ok {
		return nil, errors.New("wrong argument type")
	}
	user := domain.User{
		Email:    login.Email,
		Password: login.Password,
	}
	return &user, nil
}

func (m *RepositoryMock) SaveToken(token string, id string) error {
	args := m.Mock.Called(token, id)
	if args.Get(0) == nil || args.Get(1) == nil {
		return errors.New("argument required")
	}
	return nil
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
	return &domain.User{
		Name:     "mock",
		Email:    "mock",
		Password: "mock",
	}, nil
}
