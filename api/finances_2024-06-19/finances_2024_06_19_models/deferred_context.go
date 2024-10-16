// Code generated by go-swagger; DO NOT EDIT.

package finances_2024_06_19_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeferredContext Additional information related to deferred transactions.
// Example: {"deferralReason":"B2B","deferralStatus":"HOLD","maturityDate":"2024-07-14T00:00:00Z"}
//
// swagger:model DeferredContext
type DeferredContext struct {

	// Deferral policy applied on the transaction.
	//
	// **Examples:** `B2B`,`DD7`
	DeferralReason string `json:"deferralReason,omitempty"`

	// The status of the transaction. For example, `HOLD`,`RELEASE`.
	DeferralStatus string `json:"deferralStatus,omitempty"`

	// The release date of the transaction.
	// Format: date-time
	MaturityDate Date `json:"maturityDate,omitempty"`
}

// Validate validates this deferred context
func (m *DeferredContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaturityDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeferredContext) validateMaturityDate(formats strfmt.Registry) error {
	if swag.IsZero(m.MaturityDate) { // not required
		return nil
	}

	if err := m.MaturityDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("maturityDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("maturityDate")
		}
		return err
	}

	return nil
}

// ContextValidate validate this deferred context based on the context it is used
func (m *DeferredContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaturityDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeferredContext) contextValidateMaturityDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaturityDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("maturityDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("maturityDate")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeferredContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeferredContext) UnmarshalBinary(b []byte) error {
	var res DeferredContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
