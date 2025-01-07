package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type InvoiceItem struct {
	ID                 uuid.UUID `json:"id"`
	InvoiceID          uuid.UUID `json:"invoice_id"`
	Quantity           int64     `json:"quantity"`
	SupplierUnitPrice  float64   `json:"supplier_unit_price"`
	SupplierTotalPrice float64   `json:"supplier_total_price"`
	UnitPrice          float64   `json:"unit_price"`
	TotalPrice         float64   `json:"total_price"`
	ContractProductID  uuid.UUID `json:"contract_product_id"`
	ProductID          uuid.UUID `json:"product_id"`
	Description        string    `json:"description"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          null.Time `json:"updated_at"`
}
