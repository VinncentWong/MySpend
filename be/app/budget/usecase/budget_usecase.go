package usecase

import (
	"errors"
	"module/app/budget/cache"
	"module/app/budget/repository"
	"module/domain"
	"module/dto/budget"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type BudgetService struct {
	c    *cache.BudgetCache
	repo *repository.BudgetRepository
}

func NewBudgetService(r *repository.BudgetRepository, c *cache.BudgetCache) *BudgetService {
	return &BudgetService{
		c:    c,
		repo: r,
	}
}

func (s *BudgetService) CreateBudget(category string, dto budget.CreateBudget, userId string) (any, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	from, err := time.Parse("2006-01-02 15:04:05", dto.From)
	if err != nil {
		return nil, err
	}
	until, err := time.Parse("2006-01-02 15:04:05", dto.Until)
	if err != nil {
		return nil, err
	}
	switch category {
	case "50":
		budgetFivety := domain.BudgetFivety{
			ID:         uuid.New(),
			Budgetname: dto.BudgetName,
			Bankname:   dto.BankName,
			Amount:     dto.Amount,
			TypeBudget: dto.Type,
			From:       from,
			Until:      until,
			UserID:     uint(id),
		}
		result, err := s.repo.CreateBudget(budgetFivety, userId)
		if err != nil {
			return nil, err
		}
		return &result, nil
	case "30":
		budgetThirty := domain.BudgetThirty{
			ID:         uuid.New(),
			Budgetname: dto.BudgetName,
			Bankname:   dto.BankName,
			Amount:     dto.Amount,
			TypeBudget: dto.Type,
			From:       from,
			Until:      until,
			UserID:     uint(id),
		}
		result, err := s.repo.CreateBudget(budgetThirty, userId)
		if err != nil {
			return nil, err
		}
		return &result, nil
	case "20":
		budgetTwenty := domain.BudgetTwenty{
			ID:         uuid.New(),
			Budgetname: dto.BudgetName,
			Bankname:   dto.BankName,
			Amount:     dto.Amount,
			TypeBudget: dto.Type,
			From:       from,
			Until:      until,
			UserID:     uint(id),
		}
		result, err := s.repo.CreateBudget(budgetTwenty, userId)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New("wrong category type")
}

func (s *BudgetService) GetBudget(category string, userId string) (any, error) {
	result, err := s.c.GetBudget(category, userId)
	if err == nil {
		return result, nil
	}
	result, err = s.repo.GetBudget(category, userId)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (s *BudgetService) UpdateBudget(category string, idBudget string, dto budget.UpdateBudget) (any, error) {
	result, err := s.repo.UpdateBudget(category, idBudget, dto.IsFinished)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
