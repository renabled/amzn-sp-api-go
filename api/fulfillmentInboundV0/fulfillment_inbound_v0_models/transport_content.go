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

// TransportContent Inbound shipment information, including carrier details, shipment status, and the workflow status for a request for shipment with an Amazon-partnered carrier.
//
// swagger:model TransportContent
type TransportContent struct {

	// transport details
	// Required: true
	TransportDetails *TransportDetailOutput `json:"TransportDetails"`

	// transport header
	// Required: true
	TransportHeader *TransportHeader `json:"TransportHeader"`

	// transport result
	// Required: true
	TransportResult *TransportResult `json:"TransportResult"`
}

// Validate validates this transport content
func (m *TransportContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransportDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportContent) validateTransportDetails(formats strfmt.Registry) error {

	if err := validate.Required("TransportDetails", "body", m.TransportDetails); err != nil {
		return err
	}

	if m.TransportDetails != nil {
		if err := m.TransportDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportDetails")
			}
			return err
		}
	}

	return nil
}

func (m *TransportContent) validateTransportHeader(formats strfmt.Registry) error {

	if err := validate.Required("TransportHeader", "body", m.TransportHeader); err != nil {
		return err
	}

	if m.TransportHeader != nil {
		if err := m.TransportHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportHeader")
			}
			return err
		}
	}

	return nil
}

func (m *TransportContent) validateTransportResult(formats strfmt.Registry) error {

	if err := validate.Required("TransportResult", "body", m.TransportResult); err != nil {
		return err
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

// ContextValidate validate this transport content based on the context it is used
func (m *TransportContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransportDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransportHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransportResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportContent) contextValidateTransportDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.TransportDetails != nil {
		if err := m.TransportDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportDetails")
			}
			return err
		}
	}

	return nil
}

func (m *TransportContent) contextValidateTransportHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.TransportHeader != nil {
		if err := m.TransportHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TransportHeader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TransportHeader")
			}
			return err
		}
	}

	return nil
}

func (m *TransportContent) contextValidateTransportResult(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TransportContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportContent) UnmarshalBinary(b []byte) error {
	var res TransportContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
