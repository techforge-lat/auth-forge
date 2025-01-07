package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type InvoiceCalculation struct {
	ID                  uuid.UUID       `json:"id"`
	BillingCustomerID   uuid.UUID       `json:"billing_customer_id"`
	ConsumerCustomerID  uuid.UUID       `json:"consumer_customer_id"`
	SupplierTotalAmount float64         `json:"supplier_total_amount"`
	TotalAmount         float64         `json:"total_amount"`
	PeriodBeginsDate    time.Time       `json:"period_begins_date"`
	PeriodEndsDate      time.Time       `json:"period_ends_date"`
	Hash                string          `json:"hash"`
	Notes               string          `json:"notes"`
	FileURL             string          `json:"file_url"`
	PeriodYear          int             `json:"period_year"`
	PeriodMonth         int             `json:"period_month"`
	Attachments         json.RawMessage `json:"attachments"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           null.Time       `json:"updated_at"`
}
