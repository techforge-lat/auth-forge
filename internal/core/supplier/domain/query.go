package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Supplier struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	StockQuantity uint      `json:"stock_quantity"`
	CategoryID    uuid.UUID `json:"category_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     null.Time `json:"updated_at"`
}
