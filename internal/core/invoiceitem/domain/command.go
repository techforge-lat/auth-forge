package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	InvoiceItemInvoiceIDRules          = valid.StringRules().Build()
	InvoiceItemQuantityRules           = valid.NumberRules[int64]().Required().Build()
	InvoiceItemSupplierUnitPriceRules  = valid.FloatRules[float64]().Required().Build()
	InvoiceItemSupplierTotalPriceRules = valid.FloatRules[float64]().Required().Build()
	InvoiceItemUnitPriceRules          = valid.FloatRules[float64]().Required().Build()
	InvoiceItemTotalPriceRules         = valid.FloatRules[float64]().Required().Build()
	InvoiceItemContractProductIDRules  = valid.StringRules().Build()
	InvoiceItemProductIDRules          = valid.StringRules().Build()
	InvoiceItemDescriptionRules        = valid.StringRules().Build()
)

// InvoiceItemCreateRequest represents the request to create a InvoiceItem
type InvoiceItemCreateRequest struct {
	ID                 uuid.UUID   `json:"id"`
	InvoiceID          uuid.UUID   `json:"invoice_id"`
	Quantity           int64       `json:"quantity"`
	SupplierUnitPrice  float64     `json:"supplier_unit_price"`
	SupplierTotalPrice float64     `json:"supplier_total_price"`
	UnitPrice          float64     `json:"unit_price"`
	TotalPrice         float64     `json:"total_price"`
	ContractProductID  uuid.UUID   `json:"contract_product_id"`
	ProductID          uuid.UUID   `json:"product_id"`
	Description        null.String `json:"description"`
	CreatedAt          time.Time   `json:"created_at"`
}

// Validate validates the fields of InvoiceItemCreateRequest
func (c InvoiceItemCreateRequest) Validate() error {
	v := valid.New()

	v.Int("quantity", c.Quantity, InvoiceItemQuantityRules...)

	v.Float64("supplier_unit_price", c.SupplierUnitPrice, InvoiceItemSupplierUnitPriceRules...)

	v.Float64("supplier_total_price", c.SupplierTotalPrice, InvoiceItemSupplierTotalPriceRules...)

	v.Float64("unit_price", c.UnitPrice, InvoiceItemUnitPriceRules...)

	v.Float64("total_price", c.TotalPrice, InvoiceItemTotalPriceRules...)

	if c.Description.Valid {
		v.String("description", c.Description.String, InvoiceItemDescriptionRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// InvoiceItemUpdateRequest represents the request to update a InvoiceItem
type InvoiceItemUpdateRequest struct {
	InvoiceID          uuid.UUID   `json:"invoice_id"`
	Quantity           null.Int    `json:"quantity"`
	SupplierUnitPrice  null.Float  `json:"supplier_unit_price"`
	SupplierTotalPrice null.Float  `json:"supplier_total_price"`
	UnitPrice          null.Float  `json:"unit_price"`
	TotalPrice         null.Float  `json:"total_price"`
	ContractProductID  uuid.UUID   `json:"contract_product_id"`
	ProductID          uuid.UUID   `json:"product_id"`
	Description        null.String `json:"description"`
	UpdatedAt          null.Time   `json:"updated_at"`
}

// Validate validates the fields of InvoiceItemUpdateRequest
func (u InvoiceItemUpdateRequest) Validate() error {
	v := valid.New()

	if u.Quantity.Valid {
		v.Int("quantity", u.Quantity.Int64, InvoiceItemQuantityRules...)
	}

	if u.SupplierUnitPrice.Valid {
		v.Float64("supplier_unit_price", float64(u.SupplierUnitPrice.Float64), InvoiceItemSupplierUnitPriceRules...)
	}

	if u.SupplierTotalPrice.Valid {
		v.Float64("supplier_total_price", float64(u.SupplierTotalPrice.Float64), InvoiceItemSupplierTotalPriceRules...)
	}

	if u.UnitPrice.Valid {
		v.Float64("unit_price", float64(u.UnitPrice.Float64), InvoiceItemUnitPriceRules...)
	}

	if u.TotalPrice.Valid {
		v.Float64("total_price", float64(u.TotalPrice.Float64), InvoiceItemTotalPriceRules...)
	}

	if u.Description.Valid {
		v.String("description", u.Description.String, InvoiceItemDescriptionRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
