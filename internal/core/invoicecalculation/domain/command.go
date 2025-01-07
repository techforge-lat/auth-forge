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
	InvoiceCalculationBillingCustomerIDRules   = valid.StringRules().Build()
	InvoiceCalculationConsumerCustomerIDRules  = valid.StringRules().Build()
	InvoiceCalculationSupplierTotalAmountRules = valid.FloatRules[float64]().Required().Build()
	InvoiceCalculationTotalAmountRules         = valid.FloatRules[float64]().Required().Build()
	InvoiceCalculationPeriodBeginsDateRules    = valid.TimeRules().Required().Build()
	InvoiceCalculationPeriodEndsDateRules      = valid.TimeRules().Required().Build()
	InvoiceCalculationHashRules                = valid.StringRules().Required().Build()
	InvoiceCalculationNotesRules               = valid.StringRules().Build()
	InvoiceCalculationFileURLRules             = valid.StringRules().Build()
	InvoiceCalculationPeriodYearRules          = valid.NumberRules[int64]().Required().Build()
	InvoiceCalculationPeriodMonthRules         = valid.NumberRules[int64]().Required().Build()
)

// InvoiceCalculationCreateRequest represents the request to create a InvoiceCalculation
type InvoiceCalculationCreateRequest struct {
	ID                  uuid.UUID       `json:"id"`
	BillingCustomerID   uuid.UUID       `json:"billing_customer_id"`
	ConsumerCustomerID  uuid.UUID       `json:"consumer_customer_id"`
	SupplierTotalAmount float64         `json:"supplier_total_amount"`
	TotalAmount         float64         `json:"total_amount"`
	PeriodBeginsDate    time.Time       `json:"period_begins_date"`
	PeriodEndsDate      time.Time       `json:"period_ends_date"`
	Hash                string          `json:"hash"`
	Notes               null.String     `json:"notes"`
	FileURL             null.String     `json:"file_url"`
	PeriodYear          int64           `json:"period_year"`
	PeriodMonth         int64           `json:"period_month"`
	Attachments         json.RawMessage `json:"attachments"`
	CreatedAt           time.Time       `json:"created_at"`
}

// Validate validates the fields of InvoiceCalculationCreateRequest
func (c InvoiceCalculationCreateRequest) Validate() error {
	v := valid.New()

	v.Float64("supplier_total_amount", c.SupplierTotalAmount, InvoiceCalculationSupplierTotalAmountRules...)

	v.Float64("total_amount", c.TotalAmount, InvoiceCalculationTotalAmountRules...)

	v.Time("period_begins_date", c.PeriodBeginsDate, InvoiceCalculationPeriodBeginsDateRules...)

	v.Time("period_ends_date", c.PeriodEndsDate, InvoiceCalculationPeriodEndsDateRules...)

	v.String("hash", c.Hash, InvoiceCalculationHashRules...)

	if c.Notes.Valid {
		v.String("notes", c.Notes.String, InvoiceCalculationNotesRules...)
	}

	if c.FileURL.Valid {
		v.String("file_url", c.FileURL.String, InvoiceCalculationFileURLRules...)
	}

	v.Int("period_year", c.PeriodYear, InvoiceCalculationPeriodYearRules...)

	v.Int("period_month", c.PeriodMonth, InvoiceCalculationPeriodMonthRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// InvoiceCalculationUpdateRequest represents the request to update a InvoiceCalculation
type InvoiceCalculationUpdateRequest struct {
	BillingCustomerID   uuid.UUID       `json:"billing_customer_id"`
	ConsumerCustomerID  uuid.UUID       `json:"consumer_customer_id"`
	SupplierTotalAmount null.Float      `json:"supplier_total_amount"`
	TotalAmount         null.Float      `json:"total_amount"`
	PeriodBeginsDate    null.Time       `json:"period_begins_date"`
	PeriodEndsDate      null.Time       `json:"period_ends_date"`
	Hash                null.String     `json:"hash"`
	Notes               null.String     `json:"notes"`
	FileURL             null.String     `json:"file_url"`
	PeriodYear          null.Int        `json:"period_year"`
	PeriodMonth         null.Int        `json:"period_month"`
	Attachments         json.RawMessage `json:"attachments"`
	UpdatedAt           null.Time       `json:"updated_at"`
}

// Validate validates the fields of InvoiceCalculationUpdateRequest
func (u InvoiceCalculationUpdateRequest) Validate() error {
	v := valid.New()

	if u.SupplierTotalAmount.Valid {
		v.Float64("supplier_total_amount", float64(u.SupplierTotalAmount.Float64), InvoiceCalculationSupplierTotalAmountRules...)
	}

	if u.TotalAmount.Valid {
		v.Float64("total_amount", float64(u.TotalAmount.Float64), InvoiceCalculationTotalAmountRules...)
	}

	if u.PeriodBeginsDate.Valid {
		v.Time("period_begins_date", u.PeriodBeginsDate.Time, InvoiceCalculationPeriodBeginsDateRules...)
	}

	if u.PeriodEndsDate.Valid {
		v.Time("period_ends_date", u.PeriodEndsDate.Time, InvoiceCalculationPeriodEndsDateRules...)
	}

	if u.Hash.Valid {
		v.String("hash", u.Hash.String, InvoiceCalculationHashRules...)
	}

	if u.Notes.Valid {
		v.String("notes", u.Notes.String, InvoiceCalculationNotesRules...)
	}

	if u.FileURL.Valid {
		v.String("file_url", u.FileURL.String, InvoiceCalculationFileURLRules...)
	}

	if u.PeriodYear.Valid {
		v.Int("period_year", u.PeriodYear.Int64, InvoiceCalculationPeriodYearRules...)
	}

	if u.PeriodMonth.Valid {
		v.Int("period_month", u.PeriodMonth.Int64, InvoiceCalculationPeriodMonthRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
