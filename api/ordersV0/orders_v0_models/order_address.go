// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderAddress The shipping address for the order.
//
// swagger:model OrderAddress
type OrderAddress struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	// Required: true
	AmazonOrderID *string `json:"AmazonOrderId"`

	// The company name of the contact buyer. For IBA orders, the buyer company must be Amazon entities.
	BuyerCompanyName string `json:"BuyerCompanyName,omitempty"`

	// delivery preferences
	DeliveryPreferences *DeliveryPreferences `json:"DeliveryPreferences,omitempty"`

	// The shipping address for the order.
	//
	// **Note**: `ShippingAddress` is only available for orders with the following status values: `Unshipped`, `PartiallyShipped`, `Shipped`, and `InvoiceUnconfirmed`.
	ShippingAddress *Address `json:"ShippingAddress,omitempty"`
}

// Validate validates this order address
func (m *OrderAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazonOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliveryPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderAddress) validateAmazonOrderID(formats strfmt.Registry) error {

	if err := validate.Required("AmazonOrderId", "body", m.AmazonOrderID); err != nil {
		return err
	}

	return nil
}

func (m *OrderAddress) validateDeliveryPreferences(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryPreferences) { // not required
		return nil
	}

	if m.DeliveryPreferences != nil {
		if err := m.DeliveryPreferences.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryPreferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryPreferences")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAddress) validateShippingAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingAddress) { // not required
		return nil
	}

	if m.ShippingAddress != nil {
		if err := m.ShippingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingAddress")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order address based on the context it is used
func (m *OrderAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeliveryPreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderAddress) contextValidateDeliveryPreferences(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliveryPreferences != nil {
		if err := m.DeliveryPreferences.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeliveryPreferences")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeliveryPreferences")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAddress) contextValidateShippingAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingAddress != nil {
		if err := m.ShippingAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderAddress) UnmarshalBinary(b []byte) error {
	var res OrderAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
