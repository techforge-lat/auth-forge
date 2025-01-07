package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
	ContractOwnerCustomerIDRules   = valid.StringRules().Build()
	ContractBillingCustomerIDRules = valid.StringRules().Build()
	ContractBeginsAtRules          = valid.TimeRules().Required().Build()
	ContractEndsAtRules            = valid.TimeRules().Build()
	ContractCommitmentRules        = valid.StringRules().Required().Build()
	ContractPaymentFrequencyRules  = valid.StringRules().Required().Build()
	ContractRenewalDateRules       = valid.TimeRules().Build()
	ContractStateRules             = valid.StringRules().Required().Build()
)

// ContractCreateRequest represents the request to create a Contract
type ContractCreateRequest struct {
	ID                uint            `json:"id"`
	OwnerCustomerID   uuid.UUID       `json:"owner_customer_id"`
	BillingCustomerID uuid.UUID       `json:"billing_customer_id"`
	BeginsAt          time.Time       `json:"begins_at"`
	EndsAt            null.Time       `json:"ends_at"`
	Commitment        string          `json:"commitment"`
	PaymentFrequency  string          `json:"payment_frequency"`
	RenewalDate       null.Time       `json:"renewal_date"`
	State             string          `json:"state"`
	Attachments       json.RawMessage `json:"attachments"`
	CreatedAt         time.Time       `json:"created_at"`
}

// Validate validates the fields of ContractCreateRequest
func (c ContractCreateRequest) Validate() error {
	v := valid.New()

	v.Time("begins_at", c.BeginsAt, ContractBeginsAtRules...)

	if c.EndsAt.Valid {
		v.Time("ends_at", c.EndsAt.Time, ContractEndsAtRules...)
	}

	v.String("commitment", c.Commitment, ContractCommitmentRules...)

	v.String("payment_frequency", c.PaymentFrequency, ContractPaymentFrequencyRules...)

	if c.RenewalDate.Valid {
		v.Time("renewal_date", c.RenewalDate.Time, ContractRenewalDateRules...)
	}

	v.String("state", c.State, ContractStateRules...)

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}

// ContractUpdateRequest represents the request to update a Contract
type ContractUpdateRequest struct {
	OwnerCustomerID   uuid.UUID       `json:"owner_customer_id"`
	BillingCustomerID uuid.UUID       `json:"billing_customer_id"`
	BeginsAt          null.Time       `json:"begins_at"`
	EndsAt            null.Time       `json:"ends_at"`
	Commitment        null.String     `json:"commitment"`
	PaymentFrequency  null.String     `json:"payment_frequency"`
	RenewalDate       null.Time       `json:"renewal_date"`
	State             null.String     `json:"state"`
	Attachments       json.RawMessage `json:"attachments"`
	UpdatedAt         null.Time       `json:"updated_at"`
}

// Validate validates the fields of ContractUpdateRequest
func (u ContractUpdateRequest) Validate() error {
	v := valid.New()

	if u.BeginsAt.Valid {
		v.Time("begins_at", u.BeginsAt.Time, ContractBeginsAtRules...)
	}

	if u.EndsAt.Valid {
		v.Time("ends_at", u.EndsAt.Time, ContractEndsAtRules...)
	}

	if u.Commitment.Valid {
		v.String("commitment", u.Commitment.String, ContractCommitmentRules...)
	}

	if u.PaymentFrequency.Valid {
		v.String("payment_frequency", u.PaymentFrequency.String, ContractPaymentFrequencyRules...)
	}

	if u.RenewalDate.Valid {
		v.Time("renewal_date", u.RenewalDate.Time, ContractRenewalDateRules...)
	}

	if u.State.Valid {
		v.String("state", u.State.String, ContractStateRules...)
	}

	if v.HasErrors() {
		return errortrace.OnError(v.Errors())
	}

	return nil
}
