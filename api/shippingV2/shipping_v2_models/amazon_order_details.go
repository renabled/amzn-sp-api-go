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

// AmazonOrderDetails Amazon order information. This is required if the shipment source channel is Amazon.
//
// swagger:model AmazonOrderDetails
type AmazonOrderDetails struct {

	// The Amazon order ID associated with the Amazon order fulfilled by this shipment.
	// Required: true
	OrderID *string `json:"orderId"`
}

// Validate validates this amazon order details
func (m *AmazonOrderDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmazonOrderDetails) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("orderId", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this amazon order details based on context it is used
func (m *AmazonOrderDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AmazonOrderDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmazonOrderDetails) UnmarshalBinary(b []byte) error {
	var res AmazonOrderDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}