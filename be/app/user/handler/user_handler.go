package handler

import (
	"fmt"
	"module/app/user/usecase"
	"module/dto/user"
	"module/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IUserHandler interface {
	CreateUser(c *gin.Context)
	Login(c *gin.Context)
	GetUserById(c *gin.Context)
}

type UserHandler struct {
	usecase usecase.IUserService
}

func NewUserHandler(usecase usecase.IUserService) IUserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var container user.UserRegister
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, "binding exception occured", false, nil)
		return
	}
	err = validator.New().Struct(container)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, validationError.Error(), false, nil)
		return
	}
	result, err := h.usecase.CreateUser(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, fmt.Sprintf("internal server error with message %s", err.Error()), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "sukses create user", true, result)
}

func (h *UserHandler) Login(c *gin.Context) {
	var container user.UserLogin
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, "binding exception occured", false, nil)
		return
	}
	err = validator.New().Struct(container)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, validationError.Error(), false, nil)
		return
	}
	result, err := h.usecase.Login(&container)
	if err != nil {
		util.SendResponse(c, http.StatusUnauthorized, "user unauthorized", false, nil)
		return
	}
	token, err := util.CreateToken(*result)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	err = h.usecase.SaveToken(*token, fmt.Sprintf("%d", result.ID))
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	refreshToken, err := util.CreateToken(*result)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":         result,
		"message":      "user authenticated",
		"success":      true,
		"jwt":          token,
		"refreshtoken": refreshToken,
	})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	result, err := h.usecase.GetUserById(id)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get user", true, result)
}
