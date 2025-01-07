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
	ProductNameRules        = valid.StringRules().Required().Build()
	ProductDescriptionRules = valid.StringRules().Build()
	ProductSupplierRules    = valid.StringRules().Required().Build()

	ProductSkuRules  = valid.StringRules().Required().Build()
	ProductSlugRules = valid.StringRules().Required().Build()
)

// ProductCreateRequest represents the request to create a Product
type ProductCreateRequest struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description null.String     `json:"description"`
	Supplier    string          `json:"supplier"`
	Metadata    json.RawMessage `json:"metadata"`
	Sku         string          `json:"sku"`
	Slug        string          `json:"slug"`
	CreatedAt   time.Time       `json:"created_at"`
}

// Validate validates the fields of ProductCreateRequest
func (c ProductCreateRequest) Validate() error {
	v := valid.New()

	v.String("name", c.Name, ProductNameRules...)

	if c.Description.Valid {
		v.String("description", c.Description.String, ProductDescriptionRules...)
	}

	v.String("supplier", c.Supplier, ProductSupplierRules...)

	v.String("sku", c.Sku, ProductSkuRules...)

	v.String("slug", c.Slug, ProductSlugRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// ProductUpdateRequest represents the request to update a Product
type ProductUpdateRequest struct {
	Name        null.String     `json:"name"`
	Description null.String     `json:"description"`
	Supplier    null.String     `json:"supplier"`
	Metadata    json.RawMessage `json:"metadata"`
	Sku         null.String     `json:"sku"`
	Slug        null.String     `json:"slug"`
	UpdatedAt   null.Time       `json:"updated_at"`
}

// Validate validates the fields of ProductUpdateRequest
func (u ProductUpdateRequest) Validate() error {
	v := valid.New()

	if u.Name.Valid {
		v.String("name", u.Name.String, ProductNameRules...)
	}

	if u.Description.Valid {
		v.String("description", u.Description.String, ProductDescriptionRules...)
	}

	if u.Supplier.Valid {
		v.String("supplier", u.Supplier.String, ProductSupplierRules...)
	}

	if u.Sku.Valid {
		v.String("sku", u.Sku.String, ProductSkuRules...)
	}

	if u.Slug.Valid {
		v.String("slug", u.Slug.String, ProductSlugRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
