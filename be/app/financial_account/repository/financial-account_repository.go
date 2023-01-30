package repository

import (
	"module/domain"
	"module/dto/financial_account"
	"module/infrastructure"
	"strconv"

	"gorm.io/gorm"
)

type FinancialAccountRepository struct {
	db *gorm.DB
}

func NewFinancialAccountRepository() *FinancialAccountRepository {
	return &FinancialAccountRepository{
		db: infrastructure.GetDb(),
	}
}

func (r *FinancialAccountRepository) CreateFinancialAccount(id string, dto financial_account.CreateFinancialAccountDto) (*domain.FinancialAccount, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	financialAccount := domain.FinancialAccount{
		Bankname:      dto.Bankname,
		Username:      dto.Username,
		Accountnumber: dto.Accountnumber,
	}
	switch financialAccount.Bankname {
	case "BCA":
		financialAccount.TotalBalance = 10000000000
	case "OVO":
		financialAccount.TotalBalance = 750000000
	case "BRI":
		financialAccount.TotalBalance = 500000000
	}
	result := r.db.Model(&domain.User{}).Where("id = ?", id).Take(&domain.User{})
	if result.Error != nil {
		return nil, result.Error
	}
	financialAccount.UserID = uint(userId)
	result = r.db.Create(&financialAccount)
	if result.Error != nil {
		return nil, result.Error
	}
	return &financialAccount, nil
}

func (r *FinancialAccountRepository) GetFinancialAccount(id string) (*[]domain.FinancialAccount, error) {
	var financialAccount []domain.FinancialAccount
	result := r.db.Model(&domain.FinancialAccount{}).Preload("User").Where("user_id = ?", id).Find(&financialAccount)
	if result.Error != nil {
		return nil, result.Error
	}
	return &financialAccount, nil
}
