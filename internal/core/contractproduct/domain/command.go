package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	ContractProductContractIDRules     = valid.StringRules().Build()
	ContractProductProductIDRules      = valid.StringRules().Build()
	ContractProductProductPriceIDRules = valid.StringRules().Build()
	ContractProductSupplierPriceRules  = valid.FloatRules[float64]().Required().Build()
	ContractProductPriceRules          = valid.FloatRules[float64]().Required().Build()
	ContractProductPriceTypeRules      = valid.StringRules().Required().Build()
)

// ContractProductCreateRequest represents the request to create a ContractProduct
type ContractProductCreateRequest struct {
	ID             uint      `json:"id"`
	ContractID     uuid.UUID `json:"contract_id"`
	ProductID      uuid.UUID `json:"product_id"`
	ProductPriceID uuid.UUID `json:"product_price_id"`
	SupplierPrice  float64   `json:"supplier_price"`
	Price          float64   `json:"price"`
	PriceType      string    `json:"price_type"`
	CreatedAt      time.Time `json:"created_at"`
}

// Validate validates the fields of ContractProductCreateRequest
func (c ContractProductCreateRequest) Validate() error {
	v := valid.New()

	v.Float64("supplier_price", c.SupplierPrice, ContractProductSupplierPriceRules...)

	v.Float64("price", c.Price, ContractProductPriceRules...)

	v.String("price_type", c.PriceType, ContractProductPriceTypeRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// ContractProductUpdateRequest represents the request to update a ContractProduct
type ContractProductUpdateRequest struct {
	ContractID     uuid.UUID   `json:"contract_id"`
	ProductID      uuid.UUID   `json:"product_id"`
	ProductPriceID uuid.UUID   `json:"product_price_id"`
	SupplierPrice  null.Float  `json:"supplier_price"`
	Price          null.Float  `json:"price"`
	PriceType      null.String `json:"price_type"`
	UpdatedAt      null.Time   `json:"updated_at"`
}

// Validate validates the fields of ContractProductUpdateRequest
func (u ContractProductUpdateRequest) Validate() error {
	v := valid.New()

	if u.SupplierPrice.Valid {
		v.Float64("supplier_price", float64(u.SupplierPrice.Float64), ContractProductSupplierPriceRules...)
	}

	if u.Price.Valid {
		v.Float64("price", float64(u.Price.Float64), ContractProductPriceRules...)
	}

	if u.PriceType.Valid {
		v.String("price_type", u.PriceType.String, ContractProductPriceTypeRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
