package payment

type CreatePayment struct {
	Bankname     string `json:"bankname" validate:"required,oneof=BCA BRI OVO"`
	Month        uint   `json:"month"`
	TotalPayment uint   `json:"totalpayment"`
}
