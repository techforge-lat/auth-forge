package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Product struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Supplier    string          `json:"supplier"`
	Metadata    json.RawMessage `json:"metadata"`
	Sku         string          `json:"sku"`
	Slug        string          `json:"slug"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   null.Time       `json:"updated_at"`
}
