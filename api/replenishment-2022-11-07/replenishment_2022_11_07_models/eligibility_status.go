// Code generated by go-swagger; DO NOT EDIT.

package replenishment_2022_11_07_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EligibilityStatus The current eligibility status of an offer.
//
// swagger:model EligibilityStatus
type EligibilityStatus string

func NewEligibilityStatus(value EligibilityStatus) *EligibilityStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated EligibilityStatus.
func (m EligibilityStatus) Pointer() *EligibilityStatus {
	return &m
}

const (

	// EligibilityStatusELIGIBLE captures enum value "ELIGIBLE"
	EligibilityStatusELIGIBLE EligibilityStatus = "ELIGIBLE"

	// EligibilityStatusINELIGIBLE captures enum value "INELIGIBLE"
	EligibilityStatusINELIGIBLE EligibilityStatus = "INELIGIBLE"

	// EligibilityStatusSUSPENDED captures enum value "SUSPENDED"
	EligibilityStatusSUSPENDED EligibilityStatus = "SUSPENDED"

	// EligibilityStatusREPLENISHMENTONLYORDERING captures enum value "REPLENISHMENT_ONLY_ORDERING"
	EligibilityStatusREPLENISHMENTONLYORDERING EligibilityStatus = "REPLENISHMENT_ONLY_ORDERING"
)

// for schema
var eligibilityStatusEnum []interface{}

func init() {
	var res []EligibilityStatus
	if err := json.Unmarshal([]byte(`["ELIGIBLE","INELIGIBLE","SUSPENDED","REPLENISHMENT_ONLY_ORDERING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eligibilityStatusEnum = append(eligibilityStatusEnum, v)
	}
}

func (m EligibilityStatus) validateEligibilityStatusEnum(path, location string, value EligibilityStatus) error {
	if err := validate.EnumCase(path, location, value, eligibilityStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this eligibility status
func (m EligibilityStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEligibilityStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this eligibility status based on context it is used
func (m EligibilityStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}