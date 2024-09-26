// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdjustmentItem An item in an adjustment to the seller's account.
//
// swagger:model AdjustmentItem
type AdjustmentItem struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN,omitempty"`

	// A unique identifier assigned to products stored in and fulfilled from a fulfillment center.
	FnSKU string `json:"FnSKU,omitempty"`

	// The per-unit value of the item.
	PerUnitAmount *Currency `json:"PerUnitAmount,omitempty"`

	// A short description of the item.
	ProductDescription string `json:"ProductDescription,omitempty"`

	// Represents the number of units in the seller's inventory when the `AdjustmentType` is `FBAInventoryReimbursement`.
	Quantity string `json:"Quantity,omitempty"`

	// The seller SKU of the item. The seller SKU is qualified by the seller's seller ID, which is included with every call to the Selling Partner API.
	SellerSKU string `json:"SellerSKU,omitempty"`

	// The total value of the item.
	TotalAmount *Currency `json:"TotalAmount,omitempty"`
}

// Validate validates this adjustment item
func (m *AdjustmentItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePerUnitAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdjustmentItem) validatePerUnitAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.PerUnitAmount) { // not required
		return nil
	}

	if m.PerUnitAmount != nil {
		if err := m.PerUnitAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PerUnitAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PerUnitAmount")
			}
			return err
		}
	}

	return nil
}

func (m *AdjustmentItem) validateTotalAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalAmount) { // not required
		return nil
	}

	if m.TotalAmount != nil {
		if err := m.TotalAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TotalAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TotalAmount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this adjustment item based on the context it is used
func (m *AdjustmentItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerUnitAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdjustmentItem) contextValidatePerUnitAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.PerUnitAmount != nil {
		if err := m.PerUnitAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PerUnitAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PerUnitAmount")
			}
			return err
		}
	}

	return nil
}

func (m *AdjustmentItem) contextValidateTotalAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalAmount != nil {
		if err := m.TotalAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TotalAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TotalAmount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdjustmentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdjustmentItem) UnmarshalBinary(b []byte) error {
	var res AdjustmentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
