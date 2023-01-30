package usecase

import (
	"errors"
	"module/app/user/repository"
	"module/domain"
	"module/dto/user"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) CreateUser(dto *user.UserRegister) (*domain.User, error) {
	user, _ := s.repo.GetUserByEmail(dto.Email)
	if user != nil {
		return nil, errors.New("email already exist")
	}
	result, err := s.repo.CreateUser(dto)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) Login(dto *user.UserLogin) (*domain.User, error) {
	result, err := s.repo.Login(dto)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) GetUserById(id string) (*domain.User, error) {
	result, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
