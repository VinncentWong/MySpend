package repository

import (
	"module/domain"
	"module/dto/user"
	"module/infrastructure"
	"module/util"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: infrastructure.GetDb(),
	}
}

func (ur *UserRepository) CreateUser(dto *user.UserRegister) (*domain.User, error) {
	bytePassword, err := util.Hash(dto.Password)
	if err != nil {
		return nil, err
	}
	user := domain.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: string(*bytePassword),
	}
	result := ur.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) Login(dto *user.UserLogin) (*domain.User, error) {
	var container domain.User
	result := ur.db.Where("email = ?", dto.Email).Take(&container)
	if result.Error != nil {
		return nil, result.Error
	}
	err := util.CompareHash(dto.Password, container.Password)
	if err == nil {
		return &container, nil
	} else {
		return nil, err
	}
}
