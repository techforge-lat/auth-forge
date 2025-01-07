package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type ContractProduct struct {
	ID             uuid.UUID `json:"id"`
	ContractID     uuid.UUID `json:"contract_id"`
	ProductID      uuid.UUID `json:"product_id"`
	ProductPriceID uuid.UUID `json:"product_price_id"`
	SupplierPrice  float64   `json:"supplier_price"`
	Price          float64   `json:"price"`
	PriceType      string    `json:"price_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      null.Time `json:"updated_at"`
}
