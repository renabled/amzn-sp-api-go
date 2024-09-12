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

// ShippingInstructions The shipping-related information of a delivery.
//
// swagger:model ShippingInstructions
type ShippingInstructions struct {

	// The name of the carrier that delivers the package.
	CarrierCode string `json:"CarrierCode,omitempty"`

	// Models constraints around shipping
	ShippingConstraints *FulfillmentPlanShippingConstraints `json:"ShippingConstraints,omitempty"`

	// The ship method that is used for the order.
	ShippingMethod string `json:"ShippingMethod,omitempty"`
}

// Validate validates this shipping instructions
func (m *ShippingInstructions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShippingConstraints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingInstructions) validateShippingConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingConstraints) { // not required
		return nil
	}

	if m.ShippingConstraints != nil {
		if err := m.ShippingConstraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingConstraints")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this shipping instructions based on the context it is used
func (m *ShippingInstructions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShippingConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingInstructions) contextValidateShippingConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingConstraints != nil {
		if err := m.ShippingConstraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingConstraints")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShippingInstructions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShippingInstructions) UnmarshalBinary(b []byte) error {
	var res ShippingInstructions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
