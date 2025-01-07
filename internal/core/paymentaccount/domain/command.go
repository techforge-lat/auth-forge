package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	PaymentAccountCurrencyIDRules        = valid.StringRules().Build()
	PaymentAccountNameRules              = valid.StringRules().Required().Build()
	PaymentAccountAccountHolderNameRules = valid.StringRules().Required().Build()
	PaymentAccountAccountNumberRules     = valid.StringRules().Build()
)

// PaymentAccountCreateRequest represents the request to create a PaymentAccount
type PaymentAccountCreateRequest struct {
	ID                uuid.UUID   `json:"id"`
	CurrencyID        uuid.UUID   `json:"currency_id"`
	Name              string      `json:"name"`
	AccountHolderName string      `json:"account_holder_name"`
	AccountNumber     null.String `json:"account_number"`
	CreatedAt         time.Time   `json:"created_at"`
}

// Validate validates the fields of PaymentAccountCreateRequest
func (c PaymentAccountCreateRequest) Validate() error {
	v := valid.New()

	v.String("name", c.Name, PaymentAccountNameRules...)

	v.String("account_holder_name", c.AccountHolderName, PaymentAccountAccountHolderNameRules...)

	if c.AccountNumber.Valid {
		v.String("account_number", c.AccountNumber.String, PaymentAccountAccountNumberRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// PaymentAccountUpdateRequest represents the request to update a PaymentAccount
type PaymentAccountUpdateRequest struct {
	CurrencyID        uuid.UUID   `json:"currency_id"`
	Name              null.String `json:"name"`
	AccountHolderName null.String `json:"account_holder_name"`
	AccountNumber     null.String `json:"account_number"`
	UpdatedAt         null.Time   `json:"updated_at"`
}

// Validate validates the fields of PaymentAccountUpdateRequest
func (u PaymentAccountUpdateRequest) Validate() error {
	v := valid.New()

	if u.Name.Valid {
		v.String("name", u.Name.String, PaymentAccountNameRules...)
	}

	if u.AccountHolderName.Valid {
		v.String("account_holder_name", u.AccountHolderName.String, PaymentAccountAccountHolderNameRules...)
	}

	if u.AccountNumber.Valid {
		v.String("account_number", u.AccountNumber.String, PaymentAccountAccountNumberRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
