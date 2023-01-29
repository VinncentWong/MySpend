package handler

import (
	"module/app/financial_account/usecase"
	"module/dto/financial_account"
	"module/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FinancialAccountHandler struct {
	usecase *usecase.FinancialAccountService
}

func NewFinancialAccountHandler(u *usecase.FinancialAccountService) *FinancialAccountHandler {
	return &FinancialAccountHandler{
		usecase: u,
	}
}

func (h *FinancialAccountHandler) CreateFinancialAccount(c *gin.Context) {
	id := c.Param("id")
	var container financial_account.CreateFinancialAccountDto
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	err = validator.New().Struct(&container)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, validationError.Error(), false, nil)
		return
	}
	result, err := h.usecase.CreateFinancialAccount(id, container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create financial account", true, result)
}

func (h *FinancialAccountHandler) GetFinancialAccount(c *gin.Context) {
	id := c.Param("id")
	result, err := h.usecase.GetFinancialAccount(id)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get financial account", true, result)
}
