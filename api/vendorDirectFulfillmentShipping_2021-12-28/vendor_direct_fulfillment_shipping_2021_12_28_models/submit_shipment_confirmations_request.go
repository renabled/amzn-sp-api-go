// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_2021_12_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubmitShipmentConfirmationsRequest The `submitShipmentConfirmations` request schema.
//
// swagger:model SubmitShipmentConfirmationsRequest
type SubmitShipmentConfirmationsRequest struct {

	// An array of `ShipmentConfirmation` objects, each represents confirmation details for a specific shipment.
	ShipmentConfirmations []*ShipmentConfirmation `json:"shipmentConfirmations"`
}

// Validate validates this submit shipment confirmations request
func (m *SubmitShipmentConfirmationsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipmentConfirmations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitShipmentConfirmationsRequest) validateShipmentConfirmations(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentConfirmations) { // not required
		return nil
	}

	for i := 0; i < len(m.ShipmentConfirmations); i++ {
		if swag.IsZero(m.ShipmentConfirmations[i]) { // not required
			continue
		}

		if m.ShipmentConfirmations[i] != nil {
			if err := m.ShipmentConfirmations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipmentConfirmations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipmentConfirmations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this submit shipment confirmations request based on the context it is used
func (m *SubmitShipmentConfirmationsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShipmentConfirmations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitShipmentConfirmationsRequest) contextValidateShipmentConfirmations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ShipmentConfirmations); i++ {

		if m.ShipmentConfirmations[i] != nil {
			if err := m.ShipmentConfirmations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipmentConfirmations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipmentConfirmations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubmitShipmentConfirmationsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubmitShipmentConfirmationsRequest) UnmarshalBinary(b []byte) error {
	var res SubmitShipmentConfirmationsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
