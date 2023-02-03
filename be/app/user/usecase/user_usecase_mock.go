package usecase

import (
	"errors"
	"module/domain"
	"module/dto/user"

	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	Mock mock.Mock
}

func (m *UsecaseMock) CreateUser(dto *user.UserRegister) (*domain.User, error) {
	args := m.Mock.Called(dto)
	if args.Get(0) == nil {
		return nil, errors.New("user required")
	}
	if args.Get(1) != nil {
		err, ok := args.Get(1).(error)
		if !ok {
			return nil, errors.New("should be an error")
		}
		return nil, err
	}
	user, ok := args.Get(0).(*domain.User)
	if !ok {
		return nil, errors.New("user required")
	}
	return user, nil
}

func (m *UsecaseMock) Login(dto *user.UserLogin) (*domain.User, error) {
	args := m.Mock.Called(dto)
	if args.Get(0) == nil {
		return nil, errors.New("user required")
	}
	if args.Get(1) != nil {
		err, ok := args.Get(1).(error)
		if !ok {
			return nil, errors.New("should be an error")
		}
		return nil, err
	}
	user, ok := args.Get(0).(*domain.User)
	if !ok {
		return nil, errors.New("user required")
	}
	return user, nil
}

func (m *UsecaseMock) SaveToken(token string, id string) error {
	args := m.Mock.Called(token, id)
	if args.Get(0) == nil {
		return nil
	}
	err := args.Get(0).(error)
	return err
}

func (m *UsecaseMock) GetUserById(id string) (*domain.User, error) {
	args := m.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, errors.New("user required")
	}
	if args.Get(1) != nil {
		err, ok := args.Get(1).(error)
		if !ok {
			return nil, errors.New("should be an error")
		}
		return nil, err
	}
	user, ok := args.Get(0).(*domain.User)
	if !ok {
		return nil, errors.New("should be a user")
	}
	return user, nil
}
