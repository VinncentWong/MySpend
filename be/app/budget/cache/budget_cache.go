package cache

import (
	"context"
	"errors"
	"fmt"
	"module/domain"
	"module/infrastructure"

	"github.com/go-redis/cache/v8"
)

type BudgetCache struct {
	c *cache.Cache
}

func NewBudgetCache() *BudgetCache {
	return &BudgetCache{
		c: infrastructure.GetCache(),
	}
}

func (c *BudgetCache) GetBudget(category, id string) (any, error) {
	switch category {
	case "50":
		budgetFivety := domain.BudgetFivety{}
		err := c.c.Get(context.Background(), fmt.Sprintf("getbudget/%v/%v", category, id), &budgetFivety)
		if err != nil {
			return nil, err
		}
		return &budgetFivety, nil
	case "30":
		budgetThirty := domain.BudgetThirty{}
		err := c.c.Get(context.Background(), fmt.Sprintf("getbudget/%v/%v", category, id), &budgetThirty)
		if err != nil {
			return nil, err
		}
		return &budgetThirty, nil
	case "20":
		budgetTwenty := domain.BudgetTwenty{}
		err := c.c.Get(context.Background(), fmt.Sprintf("getbudget/%v/%v", category, id), &budgetTwenty)
		if err != nil {
			return nil, err
		}
		return &budgetTwenty, nil
	}
	return nil, errors.New("wrong category type")
}
