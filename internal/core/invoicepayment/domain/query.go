package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type InvoicePayment struct {
	ID                 uuid.UUID `json:"id"`
	InvoiceID          uuid.UUID `json:"invoice_id"`
	PaymentAccountID   uuid.UUID `json:"payment_account_id"`
	PaymentMethodID    uuid.UUID `json:"payment_method_id"`
	Amount             float64   `json:"amount"`
	Notes              string    `json:"notes"`
	ExchangeRate       float64   `json:"exchange_rate"`
	PaymentDate        time.Time `json:"payment_date"`
	Hash               string    `json:"hash"`
	ReferenceCode      string    `json:"reference_code"`
	IsDetraction       bool      `json:"is_detraction"`
	PenUsdExchangeRate float64   `json:"pen_usd_exchange_rate"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          null.Time `json:"updated_at"`
}
