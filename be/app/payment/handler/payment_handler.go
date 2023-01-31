package handler

import (
	"module/app/payment/usecase"
	"module/dto/payment"
	"module/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
)

type PaymentHandler struct {
	s *usecase.PaymentService
}

func NewPaymentHandler(us *usecase.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		s: us,
	}
}

func (h *PaymentHandler) CreatePaymentHistory(c *gin.Context) {
	id := c.Param("id")
	form, err := c.MultipartForm()
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	formData := form.Value["data"]
	var container payment.CreatePayment
	err = json.Unmarshal([]byte(formData[0]), &container)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	err = validator.New().Struct(&container)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		util.SendResponse(c, http.StatusBadRequest, errs.Error(), false, nil)
		return
	}
	file, err := c.FormFile("photo")
	if err != nil {
		util.SendResponse(c, http.StatusBadRequest, err.Error(), false, nil)
		return
	}
	result, err := h.s.CreatePaymentHistory(id, container, file)
	if err != nil {
		util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
		return
	}
	util.SendResponse(c, http.StatusOK, "success create payment handler", true, result)
}
