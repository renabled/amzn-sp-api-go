// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Weight The weight of a package.
//
// swagger:model Weight
type Weight struct {

	// unit
	// Required: true
	Unit *UnitOfWeight `json:"unit"`

	// Value of a weight.
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	Value *float64 `json:"value"`
}

// Validate validates this weight
func (m *Weight) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
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

func (m *Weight) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

func (m *Weight) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.Minimum("value", "body", *m.Value, 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("value", "body", *m.Value, 100000, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this weight based on the context it is used
func (m *Weight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Weight) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.Unit != nil {
		if err := m.Unit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Weight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Weight) UnmarshalBinary(b []byte) error {
	var res Weight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
