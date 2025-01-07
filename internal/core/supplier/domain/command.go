package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	SupplierNameRules          = valid.StringRules().Required().Build()
	SupplierDescriptionRules   = valid.StringRules().Build()
	SupplierPriceRules         = valid.FloatRules[float64]().Required().Build()
	SupplierStockQuantityRules = valid.NumberRules[uint]().Required().Build()
	SupplierCategoryIDRules    = valid.StringRules().Build()
)

// SupplierCreateRequest represents the request to create a Supplier
type SupplierCreateRequest struct {
	ID            uint        `json:"id"`
	Name          string      `json:"name"`
	Description   null.String `json:"description"`
	Price         float64     `json:"price"`
	StockQuantity uint        `json:"stock_quantity"`
	CategoryID    uuid.UUID   `json:"category_i_d"`
	CreatedAt     time.Time   `json:"created_at"`
}

// Validate validates the fields of SupplierCreateRequest
func (c SupplierCreateRequest) Validate() error {
	v := valid.New()

	v.String("name", c.Name, SupplierNameRules...)

	if c.Description.Valid {
		v.String("description", c.Description.String, SupplierDescriptionRules...)
	}

	v.Float64("price", c.Price, SupplierPriceRules...)

	v.Uint("stock_quantity", c.StockQuantity, SupplierStockQuantityRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// SupplierUpdateRequest represents the request to update a Supplier
type SupplierUpdateRequest struct {
	Name          null.String `json:"name"`
	Description   null.String `json:"description"`
	Price         null.Float  `json:"price"`
	StockQuantity null.Int    `json:"stock_quantity"`
	CategoryID    uuid.UUID   `json:"category_i_d"`
	UpdatedAt     null.Time   `json:"updated_at"`
}

// Validate validates the fields of SupplierUpdateRequest
func (u SupplierUpdateRequest) Validate() error {
	v := valid.New()

	if u.Name.Valid {
		v.String("name", u.Name.String, SupplierNameRules...)
	}

	if u.Description.Valid {
		v.String("description", u.Description.String, SupplierDescriptionRules...)
	}

	if u.Price.Valid {
		v.Float64("price", float64(u.Price.Float64), SupplierPriceRules...)
	}

	if u.StockQuantity.Valid {
		v.Uint("stock_quantity", uint(u.StockQuantity.Int64), SupplierStockQuantityRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
