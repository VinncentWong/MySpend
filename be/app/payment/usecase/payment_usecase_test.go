package usecase

import (
	"mime/multipart"
	"module/app/payment/repository"
	"module/domain"
	"module/dto/payment"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var paymentRepositoryMock repository.PaymentRepositoryMock = repository.PaymentRepositoryMock{
	Mock: mock.Mock{},
}

var paymentService IPaymentService = NewPaymentService(&paymentRepositoryMock)

func TestCreatePaymentHistory(t *testing.T) {

	data := []struct {
		Id             string
		Payment        payment.CreatePayment
		FileHeader     multipart.FileHeader
		PaymentHistory domain.PaymentHistory
	}{
		{
			Id: "1",
			Payment: payment.CreatePayment{
				Bankname:     "coba",
				TotalPayment: 1000,
				Month:        3,
			},
			FileHeader: multipart.FileHeader{
				Filename: "coba.txt",
				Size:     1000,
			},
			PaymentHistory: domain.PaymentHistory{
				BankName: "coba",
				Price:    1000,
			},
		},
		{
			Id: "2",
			Payment: payment.CreatePayment{
				Bankname:     "coba2",
				TotalPayment: 2000,
				Month:        6,
			},
			FileHeader: multipart.FileHeader{
				Filename: "coba2.txt",
				Size:     3000,
			},
			PaymentHistory: domain.PaymentHistory{
				BankName: "coba2",
				Price:    2000,
			},
		},
		{
			Id: "3",
			Payment: payment.CreatePayment{
				Bankname:     "coba3",
				TotalPayment: 3000,
				Month:        12,
			},
			FileHeader: multipart.FileHeader{
				Filename: "coba3.txt",
				Size:     3000,
			},
			PaymentHistory: domain.PaymentHistory{
				BankName: "coba3",
				Price:    3000,
			},
		},
	}

	// ValueSource
	for _, d := range data {
		t.Run("payment_usecase", func(t *testing.T) {
			paymentRepositoryMock.Mock.On("CreatePaymentHistory", d.Id, d.PaymentHistory, &d.FileHeader, d.Payment.Month).Return(&d.PaymentHistory, nil)
			result, err := paymentService.CreatePaymentHistory(d.Id, d.Payment, &d.FileHeader)
			assert.Nil(t, err, "should not return an error")
			assert.Equal(t, d.PaymentHistory.BankName, result.BankName)
			assert.Equal(t, d.PaymentHistory.Price, result.Price, "total payment should be same")

		})
	}
}
