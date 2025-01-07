package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Invoice struct {
	ID                    uuid.UUID       `json:"id"`
	BillingCustomerID     uuid.UUID       `json:"billing_customer_id"`
	SupplierTotalPrice    float64         `json:"supplier_total_price"`
	TotalPrice            float64         `json:"total_price"`
	PeriodMonth           int             `json:"period_month"`
	PeriodYear            int             `json:"period_year"`
	OwnerCustomerID       uuid.UUID       `json:"owner_customer_id"`
	Hash                  string          `json:"hash"`
	Notes                 string          `json:"notes"`
	State                 string          `json:"state"`
	InvoiceCalculationID  uuid.UUID       `json:"invoice_calculation_id"`
	HasBeenSentToCustomer bool            `json:"has_been_sent_to_customer"`
	DueDate               time.Time       `json:"due_date"`
	TaxAmount             float64         `json:"tax_amount"`
	DetractionAmount      float64         `json:"detraction_amount"`
	SubTotal              float64         `json:"sub_total"`
	IsDetractionPaid      bool            `json:"is_detraction_paid"`
	PenUsdExchangeRate    float64         `json:"pen_usd_exchange_rate"`
	Attachments           json.RawMessage `json:"attachments"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             null.Time       `json:"updated_at"`
}
