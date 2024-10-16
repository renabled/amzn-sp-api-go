// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Expiry Expiry refers to the collection of dates required  for certain items. These could be either expiryDate or mfgDate and expiryAfterDuration. These are mandatory for perishable items.
//
// swagger:model Expiry
type Expiry struct {

	// Duration after manufacturing date during which the product is valid for consumption.
	ExpiryAfterDuration *Duration `json:"expiryAfterDuration,omitempty"`

	// The date that determines the limit of consumption or use of a product. Its meaning is determined based on the trade item context.
	// Format: date-time
	ExpiryDate strfmt.DateTime `json:"expiryDate,omitempty"`

	// Production, packaging or assembly date determined by the manufacturer. Its meaning is determined based on the trade item context.
	// Format: date-time
	ManufacturerDate strfmt.DateTime `json:"manufacturerDate,omitempty"`
}

// Validate validates this expiry
func (m *Expiry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiryAfterDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturerDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Expiry) validateExpiryAfterDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryAfterDuration) { // not required
		return nil
	}

	if m.ExpiryAfterDuration != nil {
		if err := m.ExpiryAfterDuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiryAfterDuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expiryAfterDuration")
			}
			return err
		}
	}

	return nil
}

func (m *Expiry) validateExpiryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expiryDate", "body", "date-time", m.ExpiryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Expiry) validateManufacturerDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ManufacturerDate) { // not required
		return nil
	}

	if err := validate.FormatOf("manufacturerDate", "body", "date-time", m.ManufacturerDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this expiry based on the context it is used
func (m *Expiry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExpiryAfterDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Expiry) contextValidateExpiryAfterDuration(ctx context.Context, formats strfmt.Registry) error {

	if m.ExpiryAfterDuration != nil {
		if err := m.ExpiryAfterDuration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiryAfterDuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expiryAfterDuration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Expiry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Expiry) UnmarshalBinary(b []byte) error {
	var res Expiry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
