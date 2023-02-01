package budget

import (
	"module/domain"
)

type CreateBudget struct {
	BudgetName string      `json:"budgetname" validate:"required"`
	BankName   string      `json:"bankname" validate:"required,oneof=BCA BRI OVO"`
	Type       domain.Type `json:"type" validate:"min=0,max=12"`
	Amount     uint        `json:"amount"`
	From       string      `json:"from"`
	Until      string      `json:"until"`
}

type UpdateBudget struct {
	IsFinished bool `json:"isfinished" validate:"required"`
}
