package repository

import (
	"errors"
	"mime/multipart"
	"module/domain"

	"github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	Mock mock.Mock
}

func (m *PaymentRepositoryMock) CreatePaymentHistory(id string, payment domain.PaymentHistory, file *multipart.FileHeader, month uint) (*domain.PaymentHistory, error) {
	args := m.Mock.Called(id, payment, file, month)
	if args.Get(0) == nil {
		return nil, errors.New("payment history required")
	}
	if args.Get(1) != nil {
		err := args.Get(1).(error)
		return nil, err
	}
	paymentHistory, ok := args.Get(0).(*domain.PaymentHistory)
	if !ok {
		return nil, errors.New("payment history required")
	}
	return paymentHistory, nil
}
