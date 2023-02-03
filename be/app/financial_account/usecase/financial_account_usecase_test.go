package usecase

import (
	"module/app/financial_account/repository"
	"module/domain"
	"module/dto/financial_account"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var financialAccountRepositoryMock repository.FinancialAccountRepositoryMock = repository.FinancialAccountRepositoryMock{
	Mock: mock.Mock{},
}

var financialAccountService FinancialAccountService = *NewFinancialAccountService(&financialAccountRepositoryMock)

func TestCreateFinancialAccount(t *testing.T) {
	data := []struct {
		Id               string
		Dto              financial_account.CreateFinancialAccountDto
		FinancialAccount domain.FinancialAccount
	}{
		{
			Id: "1",
			Dto: financial_account.CreateFinancialAccountDto{
				Username:      "coba1",
				Bankname:      "bank1",
				Accountnumber: "111111",
			},
			FinancialAccount: domain.FinancialAccount{
				Username:      "coba1",
				Bankname:      "bank1",
				Accountnumber: "111111",
			},
		},
		{
			Id: "2",
			Dto: financial_account.CreateFinancialAccountDto{
				Username:      "coba2",
				Bankname:      "bank2",
				Accountnumber: "222222",
			},
			FinancialAccount: domain.FinancialAccount{
				Username:      "coba2",
				Bankname:      "bank2",
				Accountnumber: "222222",
			},
		},
		{
			Id: "3",
			Dto: financial_account.CreateFinancialAccountDto{
				Username:      "coba3",
				Bankname:      "bank3",
				Accountnumber: "333333",
			},
			FinancialAccount: domain.FinancialAccount{
				Username:      "coba3",
				Bankname:      "bank3",
				Accountnumber: "333333",
			},
		},
		{
			Id: "4",
			Dto: financial_account.CreateFinancialAccountDto{
				Username:      "coba4",
				Bankname:      "bank4",
				Accountnumber: "444444",
			},
			FinancialAccount: domain.FinancialAccount{
				Username:      "coba4",
				Bankname:      "bank4",
				Accountnumber: "444444",
			},
		},
		{
			Id: "5",
			Dto: financial_account.CreateFinancialAccountDto{
				Username:      "coba5",
				Bankname:      "bank5",
				Accountnumber: "555555",
			},
			FinancialAccount: domain.FinancialAccount{
				Username:      "coba5",
				Bankname:      "bank5",
				Accountnumber: "555555",
			},
		},
	}
	for _, d := range data {
		t.Run("create financial account test", func(t *testing.T) {
			financialAccountRepositoryMock.Mock.On("CreateFinancialAccount", d.Id, d.Dto).Return(&d.FinancialAccount, nil)

			result, err := financialAccountService.CreateFinancialAccount(d.Id, d.Dto)
			assert.Nil(t, err, "should not return an error")
			assert.Equal(t, d.FinancialAccount.Username, result.Username, "username should be equal")
			assert.Equal(t, d.FinancialAccount.Bankname, result.Bankname, "bankname should be equal")
			assert.Equal(t, d.FinancialAccount.Accountnumber, result.Accountnumber, "accountnumber should be equal")
		})
	}
}

func TestGetFinancialAccount(t *testing.T) {
	data := []struct {
		Id      string
		ArrData *[]domain.FinancialAccount
	}{
		{
			Id: "1",
			ArrData: &[]domain.FinancialAccount{
				{
					Bankname:     "bank1",
					Username:     "user1",
					TotalBalance: 1000,
				},
				{
					Bankname:     "bank2",
					Username:     "user2",
					TotalBalance: 2000,
				},
				{
					Bankname:     "bank3",
					Username:     "user3",
					TotalBalance: 3000,
				},
			},
		},
		{
			Id: "2",
			ArrData: &[]domain.FinancialAccount{
				{
					Bankname:     "bank4",
					Username:     "user4",
					TotalBalance: 4000,
				},
				{
					Bankname:     "bank5",
					Username:     "user5",
					TotalBalance: 5000,
				},
				{
					Bankname:     "bank6",
					Username:     "user6",
					TotalBalance: 6000,
				},
			},
		},
		{
			Id: "3",
			ArrData: &[]domain.FinancialAccount{
				{
					Bankname:     "bank7",
					Username:     "user7",
					TotalBalance: 7000,
				},
				{
					Bankname:     "bank8",
					Username:     "user8",
					TotalBalance: 8000,
				},
				{
					Bankname:     "bank9",
					Username:     "user9",
					TotalBalance: 9000,
				},
			},
		},
	}

	for _, d := range data {
		t.Run("get financial account test", func(t *testing.T) {
			financialAccountRepositoryMock.Mock.On("GetFinancialAccount", d.Id).Return(d.ArrData, nil)
			result, err := financialAccountService.GetFinancialAccount(d.Id)
			assert.Nil(t, err, "should not return an error")
			for i, v := range *d.ArrData {
				resultAccount := (*result)[i]
				assert.Equal(t, v.Bankname, resultAccount.Bankname, "bankname should be same")
				assert.Equal(t, v.Username, resultAccount.Username, "username should be same")
				assert.Equal(t, v.TotalBalance, resultAccount.TotalBalance, "total balance should be same")
			}
		})
	}
}
