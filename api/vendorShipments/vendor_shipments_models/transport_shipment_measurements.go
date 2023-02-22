// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransportShipmentMeasurements Shipment measurement details.
//
// swagger:model TransportShipmentMeasurements
type TransportShipmentMeasurements struct {

	// Total Volume of the shipment.
	ShipmentVolume *Volume `json:"shipmentVolume,omitempty"`

	// Total Weight of the shipment.
	ShipmentWeight *Weight `json:"shipmentWeight,omitempty"`

	// Total number of cartons present in the shipment. Provide the cartonCount only for non-palletized shipments.
	TotalCartonCount int64 `json:"totalCartonCount,omitempty"`

	// Total number of Non Stackable Pallets present in the shipment.
	TotalPalletNonStackable int64 `json:"totalPalletNonStackable,omitempty"`

	// Total number of Stackable Pallets present in the shipment.
	TotalPalletStackable int64 `json:"totalPalletStackable,omitempty"`
}

// Validate validates this transport shipment measurements
func (m *TransportShipmentMeasurements) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipmentVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportShipmentMeasurements) validateShipmentVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentVolume) { // not required
		return nil
	}

	if m.ShipmentVolume != nil {
		if err := m.ShipmentVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentVolume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentVolume")
			}
			return err
		}
	}

	return nil
}

func (m *TransportShipmentMeasurements) validateShipmentWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentWeight) { // not required
		return nil
	}

	if m.ShipmentWeight != nil {
		if err := m.ShipmentWeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentWeight")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transport shipment measurements based on the context it is used
func (m *TransportShipmentMeasurements) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShipmentVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportShipmentMeasurements) contextValidateShipmentVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentVolume != nil {
		if err := m.ShipmentVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentVolume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentVolume")
			}
			return err
		}
	}

	return nil
}

func (m *TransportShipmentMeasurements) contextValidateShipmentWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentWeight != nil {
		if err := m.ShipmentWeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentWeight")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransportShipmentMeasurements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportShipmentMeasurements) UnmarshalBinary(b []byte) error {
	var res TransportShipmentMeasurements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}