package domain

import (
	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	AppIDRules   = valid.StringRules().UUID().Build()
	AppNameRules = valid.StringRules().Required().MinLength(3).Build()
)

// AppCreateRequest represents the request to create a App
type AppCreateRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"-"`
	CreatedAt null.Time `json:"created_at"`
}

// Validate validates the fields of AppCreateRequest
func (c AppCreateRequest) Validate() error {
	v := valid.New()

	v.String("id", c.ID.String(), AppIDRules...)
	v.String("nombre", c.Name, AppNameRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// AppUpdateRequest represents the request to update a App
type AppUpdateRequest struct {
	Name      null.String `json:"name"`
	UpdatedAt null.Time   `json:"updated_at"`
}

// Validate validates the fields of AppUpdateRequest
func (u AppUpdateRequest) Validate() error {
	v := valid.New()

	if u.Name.Valid {
		v.String("nombre", u.Name.String, AppNameRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
