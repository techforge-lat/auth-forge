package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Currency struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	ExchangeRate float64   `json:"exchange_rate"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
}
