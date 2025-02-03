package domain

import (
	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	TenantNameRules  = valid.StringRules().Required().Build()
	TenantAppIDRules = valid.StringRules().UUID().Build()
)

// TenantCreateRequest represents the request to create a Tenant
type TenantCreateRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"-"`
	AppID     uuid.UUID `json:"app_id"`
	CreatedAt null.Time `json:"created_at"`
}

// Validate validates the fields of TenantCreateRequest
func (c TenantCreateRequest) Validate() error {
	v := valid.New()

	v.String("id", c.ID.String(), TenantAppIDRules...)
	v.String("name", c.Name, TenantNameRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// TenantUpdateRequest represents the request to update a Tenant
type TenantUpdateRequest struct {
	Name      null.String   `json:"name"`
	Code      null.String   `json:"code"`
	AppID     uuid.NullUUID `json:"app_id"`
	UpdatedAt null.Time     `json:"updated_at"`
}

// Validate validates the fields of TenantUpdateRequest
func (u TenantUpdateRequest) Validate() error {
	v := valid.New()

	if u.Name.Valid {
		v.String("name", u.Name.String, TenantNameRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
