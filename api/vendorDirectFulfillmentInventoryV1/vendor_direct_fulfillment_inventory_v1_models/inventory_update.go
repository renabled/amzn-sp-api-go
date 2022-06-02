// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_inventory_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InventoryUpdate inventory update
//
// swagger:model InventoryUpdate
type InventoryUpdate struct {

	// When true, this request contains a full feed. Otherwise, this request contains a partial feed. When sending a full feed, you must send information about all items in the warehouse. Any items not in the full feed are updated as not available. When sending a partial feed, only include the items that need an update to inventory. The status of other items will remain unchanged.
	// Required: true
	IsFullUpdate *bool `json:"isFullUpdate"`

	// A list of inventory items with updated details, including quantity available.
	// Required: true
	Items []*ItemDetails `json:"items"`

	// ID of the selling party or vendor.
	// Required: true
	SellingParty *PartyIdentification `json:"sellingParty"`
}

// Validate validates this inventory update
func (m *InventoryUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsFullUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellingParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryUpdate) validateIsFullUpdate(formats strfmt.Registry) error {

	if err := validate.Required("isFullUpdate", "body", m.IsFullUpdate); err != nil {
		return err
	}

	return nil
}

func (m *InventoryUpdate) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InventoryUpdate) validateSellingParty(formats strfmt.Registry) error {

	if err := validate.Required("sellingParty", "body", m.SellingParty); err != nil {
		return err
	}

	if m.SellingParty != nil {
		if err := m.SellingParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this inventory update based on the context it is used
func (m *InventoryUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellingParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryUpdate) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InventoryUpdate) contextValidateSellingParty(ctx context.Context, formats strfmt.Registry) error {

	if m.SellingParty != nil {
		if err := m.SellingParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryUpdate) UnmarshalBinary(b []byte) error {
	var res InventoryUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
