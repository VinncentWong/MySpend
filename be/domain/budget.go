package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	Transportation uint = iota
	FoodAndDrink
	Groceries
	Insurance
	Utilities
	Technology
	Savings
	Investments
	Emergency
	Staycation
	SelfReward
	Gym
	Others
)

type BudgetTwenty struct {
	ID         uuid.UUID
	Budgetname string    `json:"budgetname"`
	Bankname   string    `json:"bankname"`
	TypeBudget uint      `json:"typebudget"`
	IsFinished bool      `json:"isfinished"`
	From       time.Time `json:"from"`
	Until      time.Time `json:"until"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt  time.Time `json:"deletedAt"`
	User       User      `json:"-"`
	UserID     uint
}

type BudgetThirty struct {
	ID         uuid.UUID
	Budgetname string    `json:"budgetname"`
	Bankname   string    `json:"bankname"`
	TypeBudget uint      `json:"typebudget"`
	IsFinished bool      `json:"isfinished"`
	From       time.Time `json:"from"`
	Until      time.Time `json:"until"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt  time.Time `json:"deletedAt"`
	User       User      `json:"-"`
	UserID     uint
}

type BudgetFivety struct {
	ID         uuid.UUID
	Budgetname string    `json:"budgetname"`
	Bankname   string    `json:"bankname"`
	TypeBudget uint      `json:"typebudget"`
	IsFinished bool      `json:"isfinished"`
	From       time.Time `json:"from"`
	Until      time.Time `json:"until"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt  time.Time `json:"deletedAt"`
	User       User      `json:"-"`
	UserID     uint
}
