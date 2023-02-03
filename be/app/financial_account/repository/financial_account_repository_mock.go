package repository

import (
	"errors"
	"module/domain"
	"module/dto/financial_account"

	"github.com/stretchr/testify/mock"
)

type FinancialAccountRepositoryMock struct {
	Mock mock.Mock
}

func (m *FinancialAccountRepositoryMock) CreateFinancialAccount(id string, dto financial_account.CreateFinancialAccountDto) (*domain.FinancialAccount, error) {
	args := m.Mock.Called(id, dto)
	if args.Get(0) == nil {
		return nil, errors.New("financial account required")
	}
	if args.Get(1) != nil {
		err := args.Get(1).(error)
		return nil, err
	}
	financialAccount := args.Get(0).(*domain.FinancialAccount)
	return financialAccount, nil
}

func (m *FinancialAccountRepositoryMock) GetFinancialAccount(id string) (*[]domain.FinancialAccount, error) {
	args := m.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, errors.New("financial accounts required")
	}
	if args.Get(1) != nil {
		err := args.Get(1).(error)
		return nil, err
	}
	financialAccounts := args.Get(0).(*[]domain.FinancialAccount)
	return financialAccounts, nil
}
