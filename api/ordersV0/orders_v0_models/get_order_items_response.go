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

// GetOrderItemsResponse The response schema for the getOrderItems operation.
//
// swagger:model GetOrderItemsResponse
type GetOrderItemsResponse struct {

	// One or more unexpected errors occurred during the getOrderItems operation.
	Errors ErrorList `json:"errors,omitempty"`

	// The payload for the getOrderItems operation.
	Payload *OrderItemsList `json:"payload,omitempty"`
}

// Validate validates this get order items response
func (m *GetOrderItemsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetOrderItemsResponse) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if err := m.Errors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *GetOrderItemsResponse) validatePayload(formats strfmt.Registry) error {
	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get order items response based on the context it is used
func (m *GetOrderItemsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetOrderItemsResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Errors.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *GetOrderItemsResponse) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	if m.Payload != nil {
		if err := m.Payload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetOrderItemsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOrderItemsResponse) UnmarshalBinary(b []byte) error {
	var res GetOrderItemsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}