package financial_account

type CreateFinancialAccountDto struct {
	Bankname      string `validate:"required, oneof=BCA BRI OVO"`
	Accountnumber string `validate:"required"`
	Username      string `validate:"required"`
}
