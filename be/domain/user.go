package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name             string             `json:"name"`
	Email            string             `json:"email"`
	Password         string             `json:"password"`
	Activepoint      uint               `json:"activepoint"`
	PaymentHistory   []PaymentHistory   `gorm:"foreignKey:UserID"`
	FinancialAccount []FinancialAccount `gorm:"foreignKey:UserID"`
	BudgetFivety     []BudgetFivety     `gorm:"foreignKey:UserID"`
	BudgetThirty     []BudgetThirty     `gorm:"foreignKey:UserID"`
	BudgetTwenty     []BudgetTwenty     `gorm:"foreignKey:UserID"`
}
