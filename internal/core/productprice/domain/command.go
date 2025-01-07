package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	ProductPriceProductIDRules = valid.StringRules().Build()
	ProductPriceBeginsAtRules  = valid.TimeRules().Required().Build()
	ProductPriceEndsAtRules    = valid.TimeRules().Build()

	ProductPriceSupplierPriceRules      = valid.FloatRules[float64]().Required().Build()
	ProductPricePriceRules              = valid.FloatRules[float64]().Required().Build()
	ProductPriceContractCommitmentRules = valid.StringRules().Required().Build()
	ProductPricePaymentFrequencyRules   = valid.StringRules().Required().Build()
)

// ProductPriceCreateRequest represents the request to create a ProductPrice
type ProductPriceCreateRequest struct {
	ID                 uuid.UUID `json:"id"`
	ProductID          uuid.UUID `json:"product_id"`
	BeginsAt           time.Time `json:"begins_at"`
	EndsAt             null.Time `json:"ends_at"`
	IsActive           bool      `json:"is_active"`
	SupplierPrice      float64   `json:"supplier_price"`
	Price              float64   `json:"price"`
	ContractCommitment string    `json:"contract_commitment"`
	PaymentFrequency   string    `json:"payment_frequency"`
	CreatedAt          time.Time `json:"created_at"`
}

// Validate validates the fields of ProductPriceCreateRequest
func (c ProductPriceCreateRequest) Validate() error {
	v := valid.New()

	v.Time("begins_at", c.BeginsAt, ProductPriceBeginsAtRules...)

	if c.EndsAt.Valid {
		v.Time("ends_at", c.EndsAt.Time, ProductPriceEndsAtRules...)
	}

	v.Float64("supplier_price", c.SupplierPrice, ProductPriceSupplierPriceRules...)

	v.Float64("price", c.Price, ProductPricePriceRules...)

	v.String("contract_commitment", c.ContractCommitment, ProductPriceContractCommitmentRules...)

	v.String("payment_frequency", c.PaymentFrequency, ProductPricePaymentFrequencyRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// ProductPriceUpdateRequest represents the request to update a ProductPrice
type ProductPriceUpdateRequest struct {
	ProductID          uuid.UUID   `json:"product_id"`
	BeginsAt           null.Time   `json:"begins_at"`
	EndsAt             null.Time   `json:"ends_at"`
	IsActive           null.Bool   `json:"is_active"`
	SupplierPrice      null.Float  `json:"supplier_price"`
	Price              null.Float  `json:"price"`
	ContractCommitment null.String `json:"contract_commitment"`
	PaymentFrequency   null.String `json:"payment_frequency"`
	UpdatedAt          null.Time   `json:"updated_at"`
}

// Validate validates the fields of ProductPriceUpdateRequest
func (u ProductPriceUpdateRequest) Validate() error {
	v := valid.New()

	if u.BeginsAt.Valid {
		v.Time("begins_at", u.BeginsAt.Time, ProductPriceBeginsAtRules...)
	}

	if u.EndsAt.Valid {
		v.Time("ends_at", u.EndsAt.Time, ProductPriceEndsAtRules...)
	}

	if u.SupplierPrice.Valid {
		v.Float64("supplier_price", float64(u.SupplierPrice.Float64), ProductPriceSupplierPriceRules...)
	}

	if u.Price.Valid {
		v.Float64("price", float64(u.Price.Float64), ProductPricePriceRules...)
	}

	if u.ContractCommitment.Valid {
		v.String("contract_commitment", u.ContractCommitment.String, ProductPriceContractCommitmentRules...)
	}

	if u.PaymentFrequency.Valid {
		v.String("payment_frequency", u.PaymentFrequency.String, ProductPricePaymentFrequencyRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
