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

// InboundShipmentResult Result of an inbound shipment operation
//
// swagger:model InboundShipmentResult
type InboundShipmentResult struct {

	// The shipment identifier submitted in the request.
	// Required: true
	ShipmentID *string `json:"ShipmentId"`
}

// Validate validates this inbound shipment result
func (m *InboundShipmentResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundShipmentResult) validateShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("ShipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this inbound shipment result based on context it is used
func (m *InboundShipmentResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InboundShipmentResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundShipmentResult) UnmarshalBinary(b []byte) error {
	var res InboundShipmentResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
