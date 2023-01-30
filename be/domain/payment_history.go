package domain

import (
	"time"

	"github.com/google/uuid"
)

type PaymentHistory struct {
	ID               uuid.UUID `json:"id"`
	BankName         string    `json:"bankname"`
	Price            uint      `json:"price"`
	LinkPaymentProof string    `json:"linkpaymentproof"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
	User             User      `json:"-"`
	UserID           uint
}
