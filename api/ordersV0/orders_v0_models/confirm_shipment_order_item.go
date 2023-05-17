// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfirmShipmentOrderItem A single order item.
//
// swagger:model ConfirmShipmentOrderItem
type ConfirmShipmentOrderItem struct {

	// The unique identifier of the order item.
	// Required: true
	OrderItemID *string `json:"orderItemId"`

	// The quantity of the item.
	// Required: true
	Quantity *int64 `json:"quantity"`

	// The list of transparency codes.
	TransparencyCodes TransparencyCodeList `json:"transparencyCodes,omitempty"`
}

// Validate validates this confirm shipment order item
func (m *ConfirmShipmentOrderItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransparencyCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfirmShipmentOrderItem) validateOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("orderItemId", "body", m.OrderItemID); err != nil {
		return err
	}

	return nil
}

func (m *ConfirmShipmentOrderItem) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ConfirmShipmentOrderItem) validateTransparencyCodes(formats strfmt.Registry) error {
	if swag.IsZero(m.TransparencyCodes) { // not required
		return nil
	}

	if err := m.TransparencyCodes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transparencyCodes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transparencyCodes")
		}
		return err
	}

	return nil
}

// ContextValidate validate this confirm shipment order item based on the context it is used
func (m *ConfirmShipmentOrderItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransparencyCodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfirmShipmentOrderItem) contextValidateTransparencyCodes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TransparencyCodes.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transparencyCodes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transparencyCodes")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmShipmentOrderItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmShipmentOrderItem) UnmarshalBinary(b []byte) error {
	var res ConfirmShipmentOrderItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}