// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AmazonShipmentDetails Amazon shipment information.
//
// swagger:model AmazonShipmentDetails
type AmazonShipmentDetails struct {

	// This attribute is required only for a Direct Fulfillment shipment. This is the encrypted shipment ID.
	// Required: true
	ShipmentID *string `json:"shipmentId"`
}

// Validate validates this amazon shipment details
func (m *AmazonShipmentDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmazonShipmentDetails) validateShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("shipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this amazon shipment details based on context it is used
func (m *AmazonShipmentDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AmazonShipmentDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmazonShipmentDetails) UnmarshalBinary(b []byte) error {
	var res AmazonShipmentDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
