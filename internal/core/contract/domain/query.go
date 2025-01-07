package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Contract struct {
	ID                uuid.UUID       `json:"id"`
	OwnerCustomerID   uuid.UUID       `json:"owner_customer_id"`
	BillingCustomerID uuid.UUID       `json:"billing_customer_id"`
	BeginsAt          time.Time       `json:"begins_at"`
	EndsAt            time.Time       `json:"ends_at"`
	Commitment        string          `json:"commitment"`
	PaymentFrequency  string          `json:"payment_frequency"`
	RenewalDate       time.Time       `json:"renewal_date"`
	State             string          `json:"state"`
	Attachments       json.RawMessage `json:"attachments"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         null.Time       `json:"updated_at"`
}
