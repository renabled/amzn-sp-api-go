// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ShipmentStatus The shipment status to apply.
//
// swagger:model ShipmentStatus
type ShipmentStatus string

func NewShipmentStatus(value ShipmentStatus) *ShipmentStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ShipmentStatus.
func (m ShipmentStatus) Pointer() *ShipmentStatus {
	return &m
}

const (

	// ShipmentStatusReadyForPickup captures enum value "ReadyForPickup"
	ShipmentStatusReadyForPickup ShipmentStatus = "ReadyForPickup"

	// ShipmentStatusPickedUp captures enum value "PickedUp"
	ShipmentStatusPickedUp ShipmentStatus = "PickedUp"

	// ShipmentStatusRefusedPickup captures enum value "RefusedPickup"
	ShipmentStatusRefusedPickup ShipmentStatus = "RefusedPickup"
)

// for schema
var shipmentStatusEnum []interface{}

func init() {
	var res []ShipmentStatus
	if err := json.Unmarshal([]byte(`["ReadyForPickup","PickedUp","RefusedPickup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shipmentStatusEnum = append(shipmentStatusEnum, v)
	}
}

func (m ShipmentStatus) validateShipmentStatusEnum(path, location string, value ShipmentStatus) error {
	if err := validate.EnumCase(path, location, value, shipmentStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this shipment status
func (m ShipmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateShipmentStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this shipment status based on context it is used
func (m ShipmentStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
