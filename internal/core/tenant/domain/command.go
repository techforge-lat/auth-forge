package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	nameRules   = valid.StringRules().Required().MaxLength(255).Build()
	domainRules = valid.StringRules().Required().MaxLength(255).Build()
)

// TenantCreateRequest represents the request to create a Cashbox
type TenantCreateRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate validates the fields of CashboxCreateRequest
func (c TenantCreateRequest) Validate() error {
	v := valid.New()

	v.String("name", c.Name, nameRules...)
	v.String("domain", c.Domain, domainRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// TenantUpdateRequest represents the request to update a Cashbox
type TenantUpdateRequest struct {
	Name      null.String `json:"name"`
	Domain    null.String `json:"domain"`
	UpdatedAt null.Time   `json:"updated_at"`
}

// Validate validates the fields of CashboxUpdateRequest
func (u TenantUpdateRequest) Validate() error {
	v := valid.New()

	if u.Name.Valid {
		v.String("name", u.Name.String, nameRules...)
	}

	if u.Domain.Valid {
		v.String("domain", u.Domain.String, domainRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
