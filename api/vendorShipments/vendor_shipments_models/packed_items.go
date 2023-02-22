// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackedItems Details of the item being shipped.
//
// swagger:model PackedItems
type PackedItems struct {

	// Buyer Standard Identification Number (ASIN) of an item.
	BuyerProductIdentifier string `json:"buyerProductIdentifier,omitempty"`

	// item details
	ItemDetails *PackageItemDetails `json:"itemDetails,omitempty"`

	// Item sequence number for the item. The first item will be 001, the second 002, and so on. This number is used as a reference to refer to this item from the carton or pallet level.
	ItemSequenceNumber string `json:"itemSequenceNumber,omitempty"`

	// Total item quantity shipped in this shipment.
	PackedQuantity *ItemQuantity `json:"packedQuantity,omitempty"`

	// The vendor selected product identification of the item. Should be the same as was sent in the purchase order.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
}

// Validate validates this packed items
func (m *PackedItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemDetails(formats); err != nil {
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

func (m *PackedItems) validateItemDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemDetails) { // not required
		return nil
	}

	if m.ItemDetails != nil {
		if err := m.ItemDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("itemDetails")
			}
			return err
		}
	}

	return nil
}

func (m *PackedItems) validatePackedQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.PackedQuantity) { // not required
		return nil
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

// ContextValidate validate this packed items based on the context it is used
func (m *PackedItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItemDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackedItems) contextValidateItemDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemDetails != nil {
		if err := m.ItemDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("itemDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("itemDetails")
			}
			return err
		}
	}

	return nil
}

func (m *PackedItems) contextValidatePackedQuantity(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PackedItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackedItems) UnmarshalBinary(b []byte) error {
	var res PackedItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}