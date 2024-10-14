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

// Quote The estimated shipping cost associated with the transportation option.
//
// swagger:model Quote
type Quote struct {

	// cost
	// Required: true
	Cost *Currency `json:"cost"`

	// The time at which this transportation option quote expires. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime with pattern `yyyy-MM-ddTHH:mm:ss.sssZ`.
	// Format: date-time
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// Voidable until timestamp.
	// Format: date-time
	VoidableUntil strfmt.DateTime `json:"voidableUntil,omitempty"`
}

// Validate validates this quote
func (m *Quote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoidableUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quote) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *Quote) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quote) validateVoidableUntil(formats strfmt.Registry) error {
	if swag.IsZero(m.VoidableUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("voidableUntil", "body", "date-time", m.VoidableUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this quote based on the context it is used
func (m *Quote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quote) contextValidateCost(ctx context.Context, formats strfmt.Registry) error {

	if m.Cost != nil {
		if err := m.Cost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Quote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quote) UnmarshalBinary(b []byte) error {
	var res Quote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
