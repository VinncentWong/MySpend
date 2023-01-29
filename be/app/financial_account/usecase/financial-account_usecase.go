package usecase

import (
	"module/app/financial_account/repository"
	"module/domain"
	"module/dto/financial_account"
)

type FinancialAccountService struct {
	repository *repository.FinancialAccountRepository
}

func NewFinancialAccountService(r *repository.FinancialAccountRepository) *FinancialAccountService {
	return &FinancialAccountService{
		repository: r,
	}
}

func (s *FinancialAccountService) CreateFinancialAccount(id string, dto financial_account.CreateFinancialAccountDto) (*domain.FinancialAccount, error) {
	result, err := s.repository.CreateFinancialAccount(id, dto)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *FinancialAccountService) GetFinancialAccount(id string) (*[]domain.FinancialAccount, error) {
	result, err := s.repository.GetFinancialAccount(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
