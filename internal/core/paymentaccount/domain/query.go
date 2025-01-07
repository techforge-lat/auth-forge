package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type PaymentAccount struct {
	ID                uuid.UUID `json:"id"`
	CurrencyID        uuid.UUID `json:"currency_id"`
	Name              string    `json:"name"`
	AccountHolderName string    `json:"account_holder_name"`
	AccountNumber     string    `json:"account_number"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         null.Time `json:"updated_at"`
}
