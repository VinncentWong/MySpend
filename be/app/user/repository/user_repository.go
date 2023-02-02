package repository

import (
	"context"
	"fmt"
	"log"
	"module/domain"
	"module/dto/user"
	"module/infrastructure"
	"module/util"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	CreateUser(dto *user.UserRegister) (*domain.User, error)
	Login(dto *user.UserLogin) (*domain.User, error)
	SaveToken(token string, id string) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
}

type UserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		db:    infrastructure.GetDb(),
		redis: infrastructure.GetRedisClient(),
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

func (ur *UserRepository) SaveToken(token string, id string) error {
	err := ur.redis.Set(context.Background(), fmt.Sprintf("jwt_token/%s", id), token, 15*time.Minute)
	result, _ := err.Result()
	log.Default().Printf("status redis = %v", result)
	return err.Err()
}

func (ur *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var container domain.User
	result := ur.db.Where("email = ?", email).Take(&container)
	if result.Error != nil {
		return nil, result.Error
	}
	return &container, nil
}

func (ur *UserRepository) GetUserById(id string) (*domain.User, error) {
	var container domain.User
	result := ur.db.
		Preload(clause.Associations).
		Where("id = ?", id).Take(&container)
	if result.Error != nil {
		return nil, result.Error
	}
	return &container, nil
}
