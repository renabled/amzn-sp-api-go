// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Amount The monetary value.
//
// swagger:model Amount
type Amount struct {

	// currency code
	// Required: true
	CurrencyCode *CurrencyCode `json:"CurrencyCode"`

	// The amount.
	// Required: true
	Value *BigDecimalType `json:"Value"`
}

// Validate validates this amount
func (m *Amount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrencyCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Amount) validateCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("CurrencyCode", "body", m.CurrencyCode); err != nil {
		return err
	}

	if err := validate.Required("CurrencyCode", "body", m.CurrencyCode); err != nil {
		return err
	}

	if m.CurrencyCode != nil {
		if err := m.CurrencyCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrencyCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CurrencyCode")
			}
			return err
		}
	}

	return nil
}

func (m *Amount) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("Value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.Required("Value", "body", m.Value); err != nil {
		return err
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this amount based on the context it is used
func (m *Amount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrencyCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Amount) contextValidateCurrencyCode(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrencyCode != nil {
		if err := m.CurrencyCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CurrencyCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CurrencyCode")
			}
			return err
		}
	}

	return nil
}

func (m *Amount) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Amount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Amount) UnmarshalBinary(b []byte) error {
	var res Amount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
