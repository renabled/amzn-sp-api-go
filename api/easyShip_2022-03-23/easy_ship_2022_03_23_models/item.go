// Code generated by go-swagger; DO NOT EDIT.

package easy_ship_2022_03_23_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Item Item identifier and serial number information.
//
// swagger:model Item
type Item struct {

	// order item Id
	OrderItemID OrderItemID `json:"orderItemId,omitempty"`

	// order item serial numbers
	OrderItemSerialNumbers OrderItemSerialNumbers `json:"orderItemSerialNumbers,omitempty"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItemSerialNumbers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateOrderItemID(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderItemID) { // not required
		return nil
	}

	if err := m.OrderItemID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderItemId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("orderItemId")
		}
		return err
	}

	return nil
}

func (m *Item) validateOrderItemSerialNumbers(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderItemSerialNumbers) { // not required
		return nil
	}

	if err := m.OrderItemSerialNumbers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderItemSerialNumbers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("orderItemSerialNumbers")
		}
		return err
	}

	return nil
}

// ContextValidate validate this item based on the context it is used
func (m *Item) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderItemID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderItemSerialNumbers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) contextValidateOrderItemID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderItemID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderItemId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("orderItemId")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateOrderItemSerialNumbers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderItemSerialNumbers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderItemSerialNumbers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("orderItemSerialNumbers")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
