// Code generated by go-swagger; DO NOT EDIT.

package awd_2024_05_09_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InventorySummary Summary of inventory per SKU.
//
// swagger:model InventorySummary
type InventorySummary struct {

	// inventory details
	InventoryDetails *InventoryDetails `json:"inventoryDetails,omitempty"`

	// The seller or merchant SKU.
	// Required: true
	Sku *string `json:"sku"`

	// Total quantity that is in-transit from the seller and has not yet been received at an AWD Distribution Center
	TotalInboundQuantity int64 `json:"totalInboundQuantity,omitempty"`

	// Total quantity that is present in AWD distribution centers.
	TotalOnhandQuantity int64 `json:"totalOnhandQuantity,omitempty"`
}

// Validate validates this inventory summary
func (m *InventorySummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInventoryDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventorySummary) validateInventoryDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.InventoryDetails) { // not required
		return nil
	}

	if m.InventoryDetails != nil {
		if err := m.InventoryDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventoryDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventoryDetails")
			}
			return err
		}
	}

	return nil
}

func (m *InventorySummary) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this inventory summary based on the context it is used
func (m *InventorySummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInventoryDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventorySummary) contextValidateInventoryDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.InventoryDetails != nil {
		if err := m.InventoryDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventoryDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventoryDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventorySummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventorySummary) UnmarshalBinary(b []byte) error {
	var res InventorySummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
