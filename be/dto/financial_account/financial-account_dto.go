package financial_account

type CreateFinancialAccountDto struct {
	Bankname      string `validate:"required,oneof=BCA BRI OVO" json:"bankname"`
	Accountnumber string `validate:"required" json:"account_number"`
	Username      string `validate:"required" json:"username"`
}
