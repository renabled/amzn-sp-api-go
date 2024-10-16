// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeliveryPreferences The delivery preferences applied to the destination address. These preferences will be applied when possible and are best effort.
// This feature is currently supported only in the JP marketplace and not applicable for other marketplaces.
//
// swagger:model DeliveryPreferences
type DeliveryPreferences struct {

	// Additional delivery instructions. For example, this could be instructions on how to enter a building, nearby landmark or navigation instructions, 'Beware of dogs', etc.
	// Max Length: 250
	DeliveryInstructions string `json:"deliveryInstructions,omitempty"`

	// The preferred location to leave packages at the destination address.
	DropOffLocation *DropOffLocation `json:"dropOffLocation,omitempty"`
}

// Validate validates this delivery preferences
func (m *DeliveryPreferences) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeliveryInstructions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDropOffLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryPreferences) validateDeliveryInstructions(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryInstructions) { // not required
		return nil
	}

	if err := validate.MaxLength("deliveryInstructions", "body", m.DeliveryInstructions, 250); err != nil {
		return err
	}

	return nil
}

func (m *DeliveryPreferences) validateDropOffLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.DropOffLocation) { // not required
		return nil
	}

	if m.DropOffLocation != nil {
		if err := m.DropOffLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dropOffLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dropOffLocation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delivery preferences based on the context it is used
func (m *DeliveryPreferences) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDropOffLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryPreferences) contextValidateDropOffLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.DropOffLocation != nil {
		if err := m.DropOffLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dropOffLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dropOffLocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeliveryPreferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeliveryPreferences) UnmarshalBinary(b []byte) error {
	var res DeliveryPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
