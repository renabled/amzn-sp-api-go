// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PackedItem Represents an item that has been packed into a container for shipping.
//
// swagger:model PackedItem
type PackedItem struct {

	// Buyer's Standard Identification Number (ASIN) of an item. Either `buyerProductIdentifier` or `vendorProductIdentifier` is required.
	BuyerProductIdentifier string `json:"buyerProductIdentifier,omitempty"`

	// Item Sequence Number for the item. This must be the same value as sent in the order for a given item.
	// Required: true
	ItemSequenceNumber *int64 `json:"itemSequenceNumber"`

	// Total item quantity packed in the container.
	// Required: true
	PackedQuantity *ItemQuantity `json:"packedQuantity"`

	// The vendor selected product identification of the item. Should be the same as was sent in the Purchase Order, like SKU Number.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
}

// Validate validates this packed item
func (m *PackedItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemSequenceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackedItem) validateItemSequenceNumber(formats strfmt.Registry) error {

	if err := validate.Required("itemSequenceNumber", "body", m.ItemSequenceNumber); err != nil {
		return err
	}

	return nil
}

func (m *PackedItem) validatePackedQuantity(formats strfmt.Registry) error {

	if err := validate.Required("packedQuantity", "body", m.PackedQuantity); err != nil {
		return err
	}

	if m.PackedQuantity != nil {
		if err := m.PackedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packedQuantity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this packed item based on the context it is used
func (m *PackedItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePackedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackedItem) contextValidatePackedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.PackedQuantity != nil {
		if err := m.PackedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packedQuantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackedItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackedItem) UnmarshalBinary(b []byte) error {
	var res PackedItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
