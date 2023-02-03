package usecase

import (
	"module/app/user/repository"
	"module/domain"
	"module/dto/user"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var repositoryMock repository.RepositoryMock = repository.RepositoryMock{
	Mock: mock.Mock{},
}

var userService IUserService = NewUserService(&repositoryMock)

func TestCreateUser(t *testing.T) {
	register := user.UserRegister{
		Name:     "coba",
		Email:    "coba",
		Password: "coba",
	}

	user := domain.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: register.Password,
	}
	// Return specifies the return argument
	repositoryMock.Mock.On("CreateUser", &register).Return(user)
	repositoryMock.Mock.On("GetUserByEmail", register.Email).Return(nil)
	result, err := userService.CreateUser(&register)
	assert.Nil(t, err, "there should no error")
	assert.Equal(t, register.Name, result.Name, "name should be same")
	assert.Equal(t, register.Email, result.Email, "email should be same")
}

func TestLogin(t *testing.T) {
	login := user.UserLogin{
		Email:    "coba",
		Password: "coba",
	}

	user := domain.User{
		Email:    login.Email,
		Password: login.Password,
	}
	repositoryMock.Mock.On("Login", &login).Return(user)
	result, err := userService.Login(&login)
	assert.Nil(t, err, "login should not return an error")
	assert.Equal(t, login.Email, result.Email, "email should same")
	assert.Equal(t, login.Password, result.Password, "password should be same")
}

func TestSaveToken(t *testing.T) {
	token := "token"
	id := "1"
	repositoryMock.Mock.On("SaveToken", token, id).Return(nil)

	err := userService.SaveToken(token, id)
	assert.Nil(t, err, "save token should success")
}

func TestGetUserById(t *testing.T) {
	idString := "1"
	id, _ := strconv.Atoi(idString)
	user := domain.User{
		Model: gorm.Model{
			ID: uint(id),
		},
		Name:     "coba",
		Email:    "coba@gmail.com",
		Password: "cobacobacoba",
	}
	repositoryMock.Mock.On("GetUserById", idString).Return(user)
	result, err := userService.GetUserById(idString)
	assert.Nil(t, err, "this test should not return error")
	assert.Equal(t, user.ID, result.ID, "id should be same")
	assert.Equal(t, user.Name, result.Name, "name should be same")
	assert.Equal(t, user.Email, result.Email, "email should be same")
	assert.Equal(t, user.Password, result.Password, "password should be same")
}
