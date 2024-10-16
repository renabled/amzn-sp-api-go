// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonTransportResult Common container for transport result
//
// swagger:model CommonTransportResult
type CommonTransportResult struct {

	// transport result
	TransportResult *TransportResult `json:"TransportResult,omitempty"`
}

// Validate validates this common transport result
func (m *CommonTransportResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransportResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonTransportResult) validateTransportResult(formats strfmt.Registry) error {
	if swag.IsZero(m.TransportResult) { // not required
		return nil
	}

	if m.TransportResult != nil {
		if err := m.TransportResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportResult")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportResult")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this common transport result based on the context it is used
func (m *CommonTransportResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransportResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonTransportResult) contextValidateTransportResult(ctx context.Context, formats strfmt.Registry) error {

	if m.TransportResult != nil {
		if err := m.TransportResult.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportResult")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportResult")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonTransportResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonTransportResult) UnmarshalBinary(b []byte) error {
	var res CommonTransportResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
