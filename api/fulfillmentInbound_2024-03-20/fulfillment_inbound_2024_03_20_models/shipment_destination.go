// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ShipmentDestination The Amazon fulfillment center address and warehouse ID.
//
// swagger:model ShipmentDestination
type ShipmentDestination struct {

	// The address the shipment should be sent to. Empty if the destination type is `AMAZON_OPTIMIZED`.
	Address *Address `json:"address,omitempty"`

	// The type of destination for this shipment. Possible values: `AMAZON_OPTIMIZED`, `AMAZON_WAREHOUSE`.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	DestinationType *string `json:"destinationType"`

	// The warehouse that the shipment should be sent to. Empty if the destination type is `AMAZON_OPTIMIZED`.
	// Max Length: 1024
	// Min Length: 1
	WarehouseID string `json:"warehouseId,omitempty"`
}

// Validate validates this shipment destination
func (m *ShipmentDestination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationType(formats); err != nil {
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

func (m *ShipmentDestination) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *ShipmentDestination) validateDestinationType(formats strfmt.Registry) error {

	if err := validate.Required("destinationType", "body", m.DestinationType); err != nil {
		return err
	}

	if err := validate.MinLength("destinationType", "body", *m.DestinationType, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("destinationType", "body", *m.DestinationType, 1024); err != nil {
		return err
	}

	return nil
}

func (m *ShipmentDestination) validateWarehouseID(formats strfmt.Registry) error {
	if swag.IsZero(m.WarehouseID) { // not required
		return nil
	}

	if err := validate.MinLength("warehouseId", "body", m.WarehouseID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("warehouseId", "body", m.WarehouseID, 1024); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this shipment destination based on the context it is used
func (m *ShipmentDestination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentDestination) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {
		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentDestination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentDestination) UnmarshalBinary(b []byte) error {
	var res ShipmentDestination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
