package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type InvoiceCalculationItem struct {
	ID                   uuid.UUID `json:"id"`
	InvoiceCalculationID uuid.UUID `json:"invoice_calculation_id"`
	ContractProductID    uuid.UUID `json:"contract_product_id"`
	PeriodBeginsDate     time.Time `json:"period_begins_date"`
	PeriodEndsDate       time.Time `json:"period_ends_date"`
	Quantity             int       `json:"quantity"`
	SupplierUnitAmount   float64   `json:"supplier_unit_amount"`
	SupplierTotalAmount  float64   `json:"supplier_total_amount"`
	UnitAmount           float64   `json:"unit_amount"`
	TotalAmount          float64   `json:"total_amount"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            null.Time `json:"updated_at"`
}
