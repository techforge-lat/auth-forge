package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	InvoicePaymentInvoiceIDRules        = valid.StringRules().Build()
	InvoicePaymentPaymentAccountIDRules = valid.StringRules().Build()
	InvoicePaymentPaymentMethodIDRules  = valid.StringRules().Build()
	InvoicePaymentAmountRules           = valid.FloatRules[float64]().Required().Build()
	InvoicePaymentNotesRules            = valid.StringRules().Build()
	InvoicePaymentExchangeRateRules     = valid.FloatRules[float64]().Required().Build()
	InvoicePaymentPaymentDateRules      = valid.TimeRules().Build()
	InvoicePaymentHashRules             = valid.StringRules().Required().Build()
	InvoicePaymentReferenceCodeRules    = valid.StringRules().Build()

	InvoicePaymentPenUsdExchangeRateRules = valid.FloatRules[float64]().Build()
)

// InvoicePaymentCreateRequest represents the request to create a InvoicePayment
type InvoicePaymentCreateRequest struct {
	ID                 uuid.UUID   `json:"id"`
	InvoiceID          uuid.UUID   `json:"invoice_id"`
	PaymentAccountID   uuid.UUID   `json:"payment_account_id"`
	PaymentMethodID    uuid.UUID   `json:"payment_method_id"`
	Amount             float64     `json:"amount"`
	Notes              null.String `json:"notes"`
	ExchangeRate       float64     `json:"exchange_rate"`
	PaymentDate        null.Time   `json:"payment_date"`
	Hash               string      `json:"hash"`
	ReferenceCode      null.String `json:"reference_code"`
	IsDetraction       bool        `json:"is_detraction"`
	PenUsdExchangeRate null.Float  `json:"pen_usd_exchange_rate"`
	CreatedAt          time.Time   `json:"created_at"`
}

// Validate validates the fields of InvoicePaymentCreateRequest
func (c InvoicePaymentCreateRequest) Validate() error {
	v := valid.New()

	v.Float64("amount", c.Amount, InvoicePaymentAmountRules...)

	if c.Notes.Valid {
		v.String("notes", c.Notes.String, InvoicePaymentNotesRules...)
	}

	v.Float64("exchange_rate", c.ExchangeRate, InvoicePaymentExchangeRateRules...)

	if c.PaymentDate.Valid {
		v.Time("payment_date", c.PaymentDate.Time, InvoicePaymentPaymentDateRules...)
	}

	v.String("hash", c.Hash, InvoicePaymentHashRules...)

	if c.ReferenceCode.Valid {
		v.String("reference_code", c.ReferenceCode.String, InvoicePaymentReferenceCodeRules...)
	}

	if c.PenUsdExchangeRate.Valid {
		v.Float64("pen_usd_exchange_rate", float64(c.PenUsdExchangeRate.Float64), InvoicePaymentPenUsdExchangeRateRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// InvoicePaymentUpdateRequest represents the request to update a InvoicePayment
type InvoicePaymentUpdateRequest struct {
	InvoiceID          uuid.UUID   `json:"invoice_id"`
	PaymentAccountID   uuid.UUID   `json:"payment_account_id"`
	PaymentMethodID    uuid.UUID   `json:"payment_method_id"`
	Amount             null.Float  `json:"amount"`
	Notes              null.String `json:"notes"`
	ExchangeRate       null.Float  `json:"exchange_rate"`
	PaymentDate        null.Time   `json:"payment_date"`
	Hash               null.String `json:"hash"`
	ReferenceCode      null.String `json:"reference_code"`
	IsDetraction       null.Bool   `json:"is_detraction"`
	PenUsdExchangeRate null.Float  `json:"pen_usd_exchange_rate"`
	UpdatedAt          null.Time   `json:"updated_at"`
}

// Validate validates the fields of InvoicePaymentUpdateRequest
func (u InvoicePaymentUpdateRequest) Validate() error {
	v := valid.New()

	if u.Amount.Valid {
		v.Float64("amount", float64(u.Amount.Float64), InvoicePaymentAmountRules...)
	}

	if u.Notes.Valid {
		v.String("notes", u.Notes.String, InvoicePaymentNotesRules...)
	}

	if u.ExchangeRate.Valid {
		v.Float64("exchange_rate", float64(u.ExchangeRate.Float64), InvoicePaymentExchangeRateRules...)
	}

	if u.PaymentDate.Valid {
		v.Time("payment_date", u.PaymentDate.Time, InvoicePaymentPaymentDateRules...)
	}

	if u.Hash.Valid {
		v.String("hash", u.Hash.String, InvoicePaymentHashRules...)
	}

	if u.ReferenceCode.Valid {
		v.String("reference_code", u.ReferenceCode.String, InvoicePaymentReferenceCodeRules...)
	}

	if u.PenUsdExchangeRate.Valid {
		v.Float64("pen_usd_exchange_rate", float64(u.PenUsdExchangeRate.Float64), InvoicePaymentPenUsdExchangeRateRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
