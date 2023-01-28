package usecase

import (
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
