package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	InvoiceBillingCustomerIDRules    = valid.StringRules().Build()
	InvoiceSupplierTotalPriceRules   = valid.FloatRules[float64]().Required().Build()
	InvoiceTotalPriceRules           = valid.FloatRules[float64]().Required().Build()
	InvoicePeriodMonthRules          = valid.NumberRules[int64]().Required().Build()
	InvoicePeriodYearRules           = valid.NumberRules[int64]().Required().Build()
	InvoiceOwnerCustomerIDRules      = valid.StringRules().Build()
	InvoiceHashRules                 = valid.StringRules().Required().Build()
	InvoiceNotesRules                = valid.StringRules().Build()
	InvoiceStateRules                = valid.StringRules().Build()
	InvoiceInvoiceCalculationIDRules = valid.StringRules().Build()

	InvoiceDueDateRules          = valid.TimeRules().Build()
	InvoiceTaxAmountRules        = valid.FloatRules[float64]().Required().Build()
	InvoiceDetractionAmountRules = valid.FloatRules[float64]().Required().Build()
	InvoiceSubTotalRules         = valid.FloatRules[float64]().Required().Build()

	InvoicePenUsdExchangeRateRules = valid.FloatRules[float64]().Build()
)

// InvoiceCreateRequest represents the request to create a Invoice
type InvoiceCreateRequest struct {
	ID                    uuid.UUID       `json:"id"`
	BillingCustomerID     uuid.UUID       `json:"billing_customer_id"`
	SupplierTotalPrice    float64         `json:"supplier_total_price"`
	TotalPrice            float64         `json:"total_price"`
	PeriodMonth           int64           `json:"period_month"`
	PeriodYear            int64           `json:"period_year"`
	OwnerCustomerID       uuid.UUID       `json:"owner_customer_id"`
	Hash                  string          `json:"hash"`
	Notes                 null.String     `json:"notes"`
	State                 null.String     `json:"state"`
	InvoiceCalculationID  uuid.UUID       `json:"invoice_calculation_id"`
	HasBeenSentToCustomer bool            `json:"has_been_sent_to_customer"`
	DueDate               null.Time       `json:"due_date"`
	TaxAmount             float64         `json:"tax_amount"`
	DetractionAmount      float64         `json:"detraction_amount"`
	SubTotal              float64         `json:"sub_total"`
	IsDetractionPaid      bool            `json:"is_detraction_paid"`
	PenUsdExchangeRate    null.Float      `json:"pen_usd_exchange_rate"`
	Attachments           json.RawMessage `json:"attachments"`
	CreatedAt             time.Time       `json:"created_at"`
}

// Validate validates the fields of InvoiceCreateRequest
func (c InvoiceCreateRequest) Validate() error {
	v := valid.New()

	v.Float64("supplier_total_price", c.SupplierTotalPrice, InvoiceSupplierTotalPriceRules...)

	v.Float64("total_price", c.TotalPrice, InvoiceTotalPriceRules...)

	v.Int("period_month", c.PeriodMonth, InvoicePeriodMonthRules...)

	v.Int("period_year", c.PeriodYear, InvoicePeriodYearRules...)

	v.String("hash", c.Hash, InvoiceHashRules...)

	if c.Notes.Valid {
		v.String("notes", c.Notes.String, InvoiceNotesRules...)
	}

	if c.State.Valid {
		v.String("state", c.State.String, InvoiceStateRules...)
	}

	if c.DueDate.Valid {
		v.Time("due_date", c.DueDate.Time, InvoiceDueDateRules...)
	}

	v.Float64("tax_amount", c.TaxAmount, InvoiceTaxAmountRules...)

	v.Float64("detraction_amount", c.DetractionAmount, InvoiceDetractionAmountRules...)

	v.Float64("sub_total", c.SubTotal, InvoiceSubTotalRules...)

	if c.PenUsdExchangeRate.Valid {
		v.Float64("pen_usd_exchange_rate", float64(c.PenUsdExchangeRate.Float64), InvoicePenUsdExchangeRateRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// InvoiceUpdateRequest represents the request to update a Invoice
type InvoiceUpdateRequest struct {
	BillingCustomerID     uuid.UUID       `json:"billing_customer_id"`
	SupplierTotalPrice    null.Float      `json:"supplier_total_price"`
	TotalPrice            null.Float      `json:"total_price"`
	PeriodMonth           null.Int        `json:"period_month"`
	PeriodYear            null.Int        `json:"period_year"`
	OwnerCustomerID       uuid.UUID       `json:"owner_customer_id"`
	Hash                  null.String     `json:"hash"`
	Notes                 null.String     `json:"notes"`
	State                 null.String     `json:"state"`
	InvoiceCalculationID  uuid.UUID       `json:"invoice_calculation_id"`
	HasBeenSentToCustomer null.Bool       `json:"has_been_sent_to_customer"`
	DueDate               null.Time       `json:"due_date"`
	TaxAmount             null.Float      `json:"tax_amount"`
	DetractionAmount      null.Float      `json:"detraction_amount"`
	SubTotal              null.Float      `json:"sub_total"`
	IsDetractionPaid      null.Bool       `json:"is_detraction_paid"`
	PenUsdExchangeRate    null.Float      `json:"pen_usd_exchange_rate"`
	Attachments           json.RawMessage `json:"attachments"`
	UpdatedAt             null.Time       `json:"updated_at"`
}

// Validate validates the fields of InvoiceUpdateRequest
func (u InvoiceUpdateRequest) Validate() error {
	v := valid.New()

	if u.SupplierTotalPrice.Valid {
		v.Float64("supplier_total_price", float64(u.SupplierTotalPrice.Float64), InvoiceSupplierTotalPriceRules...)
	}

	if u.TotalPrice.Valid {
		v.Float64("total_price", float64(u.TotalPrice.Float64), InvoiceTotalPriceRules...)
	}

	if u.PeriodMonth.Valid {
		v.Int("period_month", u.PeriodMonth.Int64, InvoicePeriodMonthRules...)
	}

	if u.PeriodYear.Valid {
		v.Int("period_year", u.PeriodYear.Int64, InvoicePeriodYearRules...)
	}

	if u.Hash.Valid {
		v.String("hash", u.Hash.String, InvoiceHashRules...)
	}

	if u.Notes.Valid {
		v.String("notes", u.Notes.String, InvoiceNotesRules...)
	}

	if u.State.Valid {
		v.String("state", u.State.String, InvoiceStateRules...)
	}

	if u.DueDate.Valid {
		v.Time("due_date", u.DueDate.Time, InvoiceDueDateRules...)
	}

	if u.TaxAmount.Valid {
		v.Float64("tax_amount", float64(u.TaxAmount.Float64), InvoiceTaxAmountRules...)
	}

	if u.DetractionAmount.Valid {
		v.Float64("detraction_amount", float64(u.DetractionAmount.Float64), InvoiceDetractionAmountRules...)
	}

	if u.SubTotal.Valid {
		v.Float64("sub_total", float64(u.SubTotal.Float64), InvoiceSubTotalRules...)
	}

	if u.PenUsdExchangeRate.Valid {
		v.Float64("pen_usd_exchange_rate", float64(u.PenUsdExchangeRate.Float64), InvoicePenUsdExchangeRateRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
