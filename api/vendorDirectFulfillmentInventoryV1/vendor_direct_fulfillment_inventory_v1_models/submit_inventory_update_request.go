// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_inventory_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubmitInventoryUpdateRequest The request body for the `submitInventoryUpdate` operation.
//
// swagger:model SubmitInventoryUpdateRequest
type SubmitInventoryUpdateRequest struct {

	// Inventory details required to update some or all items for the requested warehouse.
	Inventory *InventoryUpdate `json:"inventory,omitempty"`
}

// Validate validates this submit inventory update request
func (m *SubmitInventoryUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInventory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitInventoryUpdateRequest) validateInventory(formats strfmt.Registry) error {
	if swag.IsZero(m.Inventory) { // not required
		return nil
	}

	if m.Inventory != nil {
		if err := m.Inventory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this submit inventory update request based on the context it is used
func (m *SubmitInventoryUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInventory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitInventoryUpdateRequest) contextValidateInventory(ctx context.Context, formats strfmt.Registry) error {

	if m.Inventory != nil {
		if err := m.Inventory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inventory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inventory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubmitInventoryUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubmitInventoryUpdateRequest) UnmarshalBinary(b []byte) error {
	var res SubmitInventoryUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
