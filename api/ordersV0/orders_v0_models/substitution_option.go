// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubstitutionOption substitution option
//
// swagger:model SubstitutionOption
type SubstitutionOption struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN,omitempty"`

	// Measurement information for the substitution option.
	Measurement *Measurement `json:"Measurement,omitempty"`

	// The number of items to be picked for this substitution option.
	QuantityOrdered int64 `json:"QuantityOrdered,omitempty"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU string `json:"SellerSKU,omitempty"`

	// The title of the item.
	Title string `json:"Title,omitempty"`
}

// Validate validates this substitution option
func (m *SubstitutionOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeasurement(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubstitutionOption) validateMeasurement(formats strfmt.Registry) error {
	if swag.IsZero(m.Measurement) { // not required
		return nil
	}

	if m.Measurement != nil {
		if err := m.Measurement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Measurement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Measurement")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this substitution option based on the context it is used
func (m *SubstitutionOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMeasurement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubstitutionOption) contextValidateMeasurement(ctx context.Context, formats strfmt.Registry) error {

	if m.Measurement != nil {
		if err := m.Measurement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Measurement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Measurement")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubstitutionOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubstitutionOption) UnmarshalBinary(b []byte) error {
	var res SubstitutionOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}