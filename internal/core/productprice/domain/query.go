package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type ProductPrice struct {
	ID                 uuid.UUID `json:"id"`
	ProductID          uuid.UUID `json:"product_id"`
	BeginsAt           time.Time `json:"begins_at"`
	EndsAt             time.Time `json:"ends_at"`
	IsActive           bool      `json:"is_active"`
	SupplierPrice      float64   `json:"supplier_price"`
	Price              float64   `json:"price"`
	ContractCommitment string    `json:"contract_commitment"`
	PaymentFrequency   string    `json:"payment_frequency"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          null.Time `json:"updated_at"`
}
