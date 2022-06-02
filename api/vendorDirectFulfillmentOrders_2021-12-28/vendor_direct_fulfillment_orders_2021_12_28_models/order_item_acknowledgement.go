// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_orders_2021_12_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderItemAcknowledgement order item acknowledgement
//
// swagger:model OrderItemAcknowledgement
type OrderItemAcknowledgement struct {

	// Details of quantity acknowledged with the above acknowledgement code.
	// Required: true
	AcknowledgedQuantity *ItemQuantity `json:"acknowledgedQuantity"`

	// Buyer's standard identification number (ASIN) of an item.
	BuyerProductIdentifier string `json:"buyerProductIdentifier,omitempty"`

	// Line item sequence number for the item.
	// Required: true
	ItemSequenceNumber *string `json:"itemSequenceNumber"`

	// The vendor selected product identification of the item. Should be the same as was provided in the purchase order.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
}

// Validate validates this order item acknowledgement
func (m *OrderItemAcknowledgement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcknowledgedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemSequenceNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItemAcknowledgement) validateAcknowledgedQuantity(formats strfmt.Registry) error {

	if err := validate.Required("acknowledgedQuantity", "body", m.AcknowledgedQuantity); err != nil {
		return err
	}

	if m.AcknowledgedQuantity != nil {
		if err := m.AcknowledgedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acknowledgedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acknowledgedQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItemAcknowledgement) validateItemSequenceNumber(formats strfmt.Registry) error {

	if err := validate.Required("itemSequenceNumber", "body", m.ItemSequenceNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order item acknowledgement based on the context it is used
func (m *OrderItemAcknowledgement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcknowledgedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItemAcknowledgement) contextValidateAcknowledgedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.AcknowledgedQuantity != nil {
		if err := m.AcknowledgedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acknowledgedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acknowledgedQuantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderItemAcknowledgement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderItemAcknowledgement) UnmarshalBinary(b []byte) error {
	var res OrderItemAcknowledgement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
