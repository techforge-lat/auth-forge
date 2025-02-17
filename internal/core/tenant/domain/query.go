package domain

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Tenant struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	AppID     uuid.UUID `json:"app_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}
