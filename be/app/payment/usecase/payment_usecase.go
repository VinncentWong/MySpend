package usecase

import (
	"mime/multipart"
	"module/app/payment/repository"
	"module/domain"
	"module/dto/payment"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}

func NewPaymentService(r *repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: r,
	}
}

func (s *PaymentService) CreatePaymentHistory(id string, dto payment.CreatePayment, file *multipart.FileHeader) (*domain.PaymentHistory, error) {
	payment := domain.PaymentHistory{
		BankName: dto.Bankname,
		Price:    dto.TotalPayment,
	}
	result, err := s.repo.CreatePaymentHistory(id, payment, file, dto.Month)
	if err != nil {
		return nil, err
	}
	return result, nil
}
