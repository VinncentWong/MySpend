package repository

import (
	"errors"
	"module/domain"
	"module/infrastructure"

	"gorm.io/gorm"
)

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository() *BudgetRepository {
	return &BudgetRepository{
		db: infrastructure.GetDb(),
	}
}

func (r *BudgetRepository) CreateBudget(budget any, userId string) (any, error) {
	fivetyBudget, fivetyOk := budget.(domain.BudgetFivety)
	thirtyBudget, thirtyOk := budget.(domain.BudgetThirty)
	twentyBudget, twentyOk := budget.(domain.BudgetTwenty)
	if !(fivetyOk || thirtyOk || twentyOk) {
		return nil, errors.New("wrong budget type")
	}
	if fivetyOk {
		result := r.db.Model(&domain.BudgetFivety{}).Create(&fivetyBudget)
		if result.Error != nil {
			return nil, result.Error
		}
		return fivetyBudget, nil
	} else if thirtyOk {
		result := r.db.Model(&domain.BudgetThirty{}).Create(&thirtyBudget)
		if result.Error != nil {
			return nil, result.Error
		}
		return fivetyBudget, nil
	} else if twentyOk {
		result := r.db.Model(&domain.BudgetTwenty{}).Create(&twentyBudget)
		if result.Error != nil {
			return nil, result.Error
		}
		return fivetyBudget, nil
	}
	return nil, errors.New("wrong budget type")
}

func (r *BudgetRepository) GetBudget(category string, id string) (any, error) {
	switch category {
	case "50":
		var container domain.BudgetFivety
		result := r.db.Model(&domain.BudgetFivety{}).Where("user_id = ?", id).Find(&container)
		if result.Error != nil {
			return nil, result.Error
		}
		return &container, nil
	case "30":
		var container domain.BudgetThirty
		result := r.db.Model(&domain.BudgetThirty{}).Where("user_id = ?", id).Find(&container)
		if result.Error != nil {
			return nil, result.Error
		}
		return &container, nil
	case "20":
		var container domain.BudgetTwenty
		result := r.db.Model(&domain.BudgetTwenty{}).Where("user_id = ?", id).Find(&container)
		if result.Error != nil {
			return nil, result.Error
		}
		return &container, nil
	}
	return nil, errors.New("wrong category value")
}

func (r *BudgetRepository) UpdateBudget(category string, idBudget string, boolValue bool) (any, error) {
	switch category {
	case "50":
		var container domain.BudgetFivety
		result := r.db.Where("id = ?", idBudget).Take(&container)
		if result.Error != nil {
			return nil, result.Error
		}
		container.IsFinished = boolValue
		result = r.db.Model(&domain.BudgetFivety{}).Where("id = ?", idBudget).Updates(&container)
		if result.Error != nil {
			return nil, result.Error
		}
		return &container, nil
	case "30":
		container := domain.BudgetThirty{
			IsFinished: boolValue,
		}
		result := r.db.Model(&domain.BudgetThirty{}).Where("id = ?", idBudget).Updates(&container)
		if result.Error != nil {
			return nil, result.Error
		}
		return &container, nil
	case "20":
		container := domain.BudgetTwenty{
			IsFinished: boolValue,
		}
		result := r.db.Model(&domain.BudgetTwenty{}).Where("id = ?", idBudget).Updates(&container)
		if result.Error != nil {
			return nil, result.Error
		}
		return &container, nil
	}
	return nil, errors.New("wrong category type")
}
