// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inboundv2024_03_20_models

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

// CustomPlacementInput Provide units going to the warehouse.
// Example: {"items":[{"expiration":"2024-01-01","labelOwner":"AMAZON","manufacturingLotCode":"manufacturingLotCode","msku":"Sunglasses","prepOwner":"AMAZON","quantity":10}],"warehouseId":"YYZ14"}
//
// swagger:model CustomPlacementInput
type CustomPlacementInput struct {

	// Items included while creating Inbound Plan.
	// Required: true
	// Max Items: 2000
	// Min Items: 1
	Items []*ItemInput `json:"items"`

	// Warehouse Id.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	WarehouseID *string `json:"warehouseId"`
}

// Validate validates this custom placement input
func (m *CustomPlacementInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarehouseID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomPlacementInput) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	iItemsSize := int64(len(m.Items))

	if err := validate.MinItems("items", "body", iItemsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("items", "body", iItemsSize, 2000); err != nil {
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

func (m *CustomPlacementInput) validateWarehouseID(formats strfmt.Registry) error {

	if err := validate.Required("warehouseId", "body", m.WarehouseID); err != nil {
		return err
	}

	if err := validate.MinLength("warehouseId", "body", *m.WarehouseID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("warehouseId", "body", *m.WarehouseID, 1024); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this custom placement input based on the context it is used
func (m *CustomPlacementInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomPlacementInput) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CustomPlacementInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomPlacementInput) UnmarshalBinary(b []byte) error {
	var res CustomPlacementInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
