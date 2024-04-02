// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inboundv2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WindowInput Contains only a starting DateTime.
//
// swagger:model WindowInput
type WindowInput struct {

	// The start date of the window. The time component must be zero.
	// Required: true
	// Format: date-time
	Start *strfmt.DateTime `json:"start"`
}

// Validate validates this window input
func (m *WindowInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WindowInput) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this window input based on context it is used
func (m *WindowInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WindowInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WindowInput) UnmarshalBinary(b []byte) error {
	var res WindowInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
