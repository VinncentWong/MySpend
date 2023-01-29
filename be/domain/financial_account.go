package domain

import "gorm.io/gorm"

type FinancialAccount struct {
	gorm.Model
	Bankname      string `json:"bankname"`
	Username      string `json:"username"`
	Accountnumber string `json:"accountnumber"`
	TotalBalance  uint   `json:"totalbalance"`
	UserID        uint
}
