// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssociatedItem An item that is associated with an order item. For example, a tire installation service that is purchased with tires.
//
// swagger:model AssociatedItem
type AssociatedItem struct {

	// association type
	AssociationType AssociationType `json:"AssociationType,omitempty"`

	// The order item's order identifier, in 3-7-7 format.
	OrderID string `json:"OrderId,omitempty"`

	// An Amazon-defined item identifier for the associated item.
	OrderItemID string `json:"OrderItemId,omitempty"`
}

// Validate validates this associated item
func (m *AssociatedItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssociatedItem) validateAssociationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AssociationType) { // not required
		return nil
	}

	if err := m.AssociationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AssociationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AssociationType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this associated item based on the context it is used
func (m *AssociatedItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssociationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssociatedItem) contextValidateAssociationType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AssociationType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AssociationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AssociationType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssociatedItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssociatedItem) UnmarshalBinary(b []byte) error {
	var res AssociatedItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
