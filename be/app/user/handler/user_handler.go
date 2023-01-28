package handler

import (
	"module/app/user/usecase"
	"module/dto/user"
	"module/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase *usecase.UserService
}

func NewUserHandler(usecase *usecase.UserService) *UserHandler {
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
	result, err := h.usecase.CreateUser(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, "kesalahan internal sistem", false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "sukses membuat user", true, result)
}

func (h *UserHandler) Login(c *gin.Context) {
	var container user.UserLogin
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, "binding exception occured", false, nil)
		return
	}
	result, err := h.usecase.Login(&container)
	if err != nil {
		util.SendResponse(c, http.StatusUnauthorized, "user unauthorized", false, nil)
		return
	}
	token, err := util.CreateToken(result)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, "token creation exception occured", false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "user authenticated", true, gin.H{
		"data": result,
		"jwt":  token,
	})
}
