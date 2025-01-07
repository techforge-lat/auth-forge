package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	CurrencyNameRules         = valid.StringRules().Required().Build()
	CurrencySymbolRules       = valid.StringRules().Required().Build()
	CurrencyExchangeRateRules = valid.FloatRules[float64]().Required().Build()
)

// CurrencyCreateRequest represents the request to create a Currency
type CurrencyCreateRequest struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	ExchangeRate float64   `json:"exchange_rate"`
	CreatedAt    time.Time `json:"created_at"`
}

// Validate validates the fields of CurrencyCreateRequest
func (c CurrencyCreateRequest) Validate() error {
	v := valid.New()

	v.String("name", c.Name, CurrencyNameRules...)

	v.String("symbol", c.Symbol, CurrencySymbolRules...)

	v.Float64("exchange_rate", c.ExchangeRate, CurrencyExchangeRateRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// CurrencyUpdateRequest represents the request to update a Currency
type CurrencyUpdateRequest struct {
	Name         null.String `json:"name"`
	Symbol       null.String `json:"symbol"`
	ExchangeRate null.Float  `json:"exchange_rate"`
	UpdatedAt    null.Time   `json:"updated_at"`
}

// Validate validates the fields of CurrencyUpdateRequest
func (u CurrencyUpdateRequest) Validate() error {
	v := valid.New()

	if u.Name.Valid {
		v.String("name", u.Name.String, CurrencyNameRules...)
	}

	if u.Symbol.Valid {
		v.String("symbol", u.Symbol.String, CurrencySymbolRules...)
	}

	if u.ExchangeRate.Valid {
		v.Float64("exchange_rate", float64(u.ExchangeRate.Float64), CurrencyExchangeRateRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
