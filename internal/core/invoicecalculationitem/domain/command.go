package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	InvoiceCalculationItemInvoiceCalculationIDRules = valid.StringRules().Build()
	InvoiceCalculationItemContractProductIDRules    = valid.StringRules().Build()
	InvoiceCalculationItemPeriodBeginsDateRules     = valid.TimeRules().Required().Build()
	InvoiceCalculationItemPeriodEndsDateRules       = valid.TimeRules().Required().Build()
	InvoiceCalculationItemQuantityRules             = valid.NumberRules[int64]().Required().Build()
	InvoiceCalculationItemSupplierUnitAmountRules   = valid.FloatRules[float64]().Required().Build()
	InvoiceCalculationItemSupplierTotalAmountRules  = valid.FloatRules[float64]().Required().Build()
	InvoiceCalculationItemUnitAmountRules           = valid.FloatRules[float64]().Required().Build()
	InvoiceCalculationItemTotalAmountRules          = valid.FloatRules[float64]().Required().Build()
)

// InvoiceCalculationItemCreateRequest represents the request to create a InvoiceCalculationItem
type InvoiceCalculationItemCreateRequest struct {
	ID                   uuid.UUID `json:"id"`
	InvoiceCalculationID uuid.UUID `json:"invoice_calculation_id"`
	ContractProductID    uuid.UUID `json:"contract_product_id"`
	PeriodBeginsDate     time.Time `json:"period_begins_date"`
	PeriodEndsDate       time.Time `json:"period_ends_date"`
	Quantity             int64     `json:"quantity"`
	SupplierUnitAmount   float64   `json:"supplier_unit_amount"`
	SupplierTotalAmount  float64   `json:"supplier_total_amount"`
	UnitAmount           float64   `json:"unit_amount"`
	TotalAmount          float64   `json:"total_amount"`
	CreatedAt            time.Time `json:"created_at"`
}

// Validate validates the fields of InvoiceCalculationItemCreateRequest
func (c InvoiceCalculationItemCreateRequest) Validate() error {
	v := valid.New()

	v.Time("period_begins_date", c.PeriodBeginsDate, InvoiceCalculationItemPeriodBeginsDateRules...)

	v.Time("period_ends_date", c.PeriodEndsDate, InvoiceCalculationItemPeriodEndsDateRules...)

	v.Int("quantity", c.Quantity, InvoiceCalculationItemQuantityRules...)

	v.Float64("supplier_unit_amount", c.SupplierUnitAmount, InvoiceCalculationItemSupplierUnitAmountRules...)

	v.Float64("supplier_total_amount", c.SupplierTotalAmount, InvoiceCalculationItemSupplierTotalAmountRules...)

	v.Float64("unit_amount", c.UnitAmount, InvoiceCalculationItemUnitAmountRules...)

	v.Float64("total_amount", c.TotalAmount, InvoiceCalculationItemTotalAmountRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// InvoiceCalculationItemUpdateRequest represents the request to update a InvoiceCalculationItem
type InvoiceCalculationItemUpdateRequest struct {
	InvoiceCalculationID uuid.UUID  `json:"invoice_calculation_id"`
	ContractProductID    uuid.UUID  `json:"contract_product_id"`
	PeriodBeginsDate     null.Time  `json:"period_begins_date"`
	PeriodEndsDate       null.Time  `json:"period_ends_date"`
	Quantity             null.Int   `json:"quantity"`
	SupplierUnitAmount   null.Float `json:"supplier_unit_amount"`
	SupplierTotalAmount  null.Float `json:"supplier_total_amount"`
	UnitAmount           null.Float `json:"unit_amount"`
	TotalAmount          null.Float `json:"total_amount"`
	UpdatedAt            null.Time  `json:"updated_at"`
}

// Validate validates the fields of InvoiceCalculationItemUpdateRequest
func (u InvoiceCalculationItemUpdateRequest) Validate() error {
	v := valid.New()

	if u.PeriodBeginsDate.Valid {
		v.Time("period_begins_date", u.PeriodBeginsDate.Time, InvoiceCalculationItemPeriodBeginsDateRules...)
	}

	if u.PeriodEndsDate.Valid {
		v.Time("period_ends_date", u.PeriodEndsDate.Time, InvoiceCalculationItemPeriodEndsDateRules...)
	}

	if u.Quantity.Valid {
		v.Int("quantity", u.Quantity.Int64, InvoiceCalculationItemQuantityRules...)
	}

	if u.SupplierUnitAmount.Valid {
		v.Float64("supplier_unit_amount", float64(u.SupplierUnitAmount.Float64), InvoiceCalculationItemSupplierUnitAmountRules...)
	}

	if u.SupplierTotalAmount.Valid {
		v.Float64("supplier_total_amount", float64(u.SupplierTotalAmount.Float64), InvoiceCalculationItemSupplierTotalAmountRules...)
	}

	if u.UnitAmount.Valid {
		v.Float64("unit_amount", float64(u.UnitAmount.Float64), InvoiceCalculationItemUnitAmountRules...)
	}

	if u.TotalAmount.Valid {
		v.Float64("total_amount", float64(u.TotalAmount.Float64), InvoiceCalculationItemTotalAmountRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
