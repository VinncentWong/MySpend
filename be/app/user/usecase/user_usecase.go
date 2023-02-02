package usecase

import (
	"errors"
	"module/app/user/repository"
	"module/domain"
	"module/dto/user"
)

type IUserService interface {
	CreateUser(dto *user.UserRegister) (*domain.User, error)
	Login(dto *user.UserLogin) (*domain.User, error)
	SaveToken(token string, id string) error
	GetUserById(id string) (*domain.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
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

func (s *UserService) SaveToken(token string, id string) error {
	err := s.repo.SaveToken(token, id)
	return err
}

func (s *UserService) GetUserById(id string) (*domain.User, error) {
	result, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
