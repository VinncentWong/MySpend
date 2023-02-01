package handler

import (
	"module/app/budget/usecase"
	"module/dto/budget"
	"module/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BudgetHandler struct {
	s *usecase.BudgetService
}

func NewBudgetHandler(s *usecase.BudgetService) *BudgetHandler {
	return &BudgetHandler{
		s: s,
	}
}

func (h *BudgetHandler) CreateBudget(c *gin.Context) {
	category := c.Query("category")
	userId := c.Query("userId")
	var container budget.CreateBudget
	err := c.ShouldBindJSON(&container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	errs := validator.New().Struct(&container)
	if errs != nil {
		arrErr := errs.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, arrErr.Error(), false, nil)
		return
	}
	result, err := h.s.CreateBudget(category, container, userId)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusCreated, "success create budget", true, result)
}

func (h *BudgetHandler) GetBudget(c *gin.Context) {
	category := c.Query("category")
	userId := c.Query("userId")
	result, err := h.s.GetBudget(category, userId)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success get budget", true, result)
}

func (h *BudgetHandler) UpdateBudget(c *gin.Context) {
	category := c.Query("category")
	idBudget := c.Query("idBudget")
	dto := budget.UpdateBudget{}
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	err = validator.New().Struct(&dto)
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	result, err := h.s.UpdateBudget(category, idBudget, dto)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success update budget", true, result)
}
