package usecase

import (
	"module/app/user/repository"
	"module/dto/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repositoryMock repository.RepositoryMock = repository.RepositoryMock{
	Mock: mock.Mock{},
}

var userService IUserService = NewUserService(&repositoryMock)

func TestCreateUser(t *testing.T) {
	register := user.UserRegister{
		Name:  "coba",
		Email: "coba",
	}

	// Return specifies the return argument
	repositoryMock.Mock.On("CreateUser", &register).Return(&register)
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
	repositoryMock.Mock.On("Login", &login).Return(&login)
	result, err := userService.Login(&login)
	assert.Nil(t, err, "login should not return an error")
	assert.Equal(t, login.Email, result.Email, "email should same")
	assert.Equal(t, login.Password, result.Password, "password should be same")
}

func TestSaveToken(t *testing.T) {
	token := "token"
	id := "1"
	repositoryMock.Mock.On("SaveToken", token, id).Return(token, id)

	err := userService.SaveToken(token, id)
	assert.Nil(t, err, "save token should success")
}

func TestGetUserById(t *testing.T) {
	id := "1"
	repositoryMock.Mock.On("GetUserById", id).Return(id)

	result, err := userService.GetUserById(id)
	assert.Nil(t, err, "this test should not return error")
	assert.Equal(t, "mock", result.Name, "name should be same")
	assert.Equal(t, "mock", result.Email, "email should be same")
	assert.Equal(t, "mock", result.Password, "password should be same")
}
