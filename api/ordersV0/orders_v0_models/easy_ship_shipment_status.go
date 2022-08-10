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

// EasyShipShipmentStatus The status of the Amazon Easy-Ship order. This property is included only for Amazon Easy-Ship orders.
//
// swagger:model EasyShipShipmentStatus
type EasyShipShipmentStatus string

func NewEasyShipShipmentStatus(value EasyShipShipmentStatus) *EasyShipShipmentStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated EasyShipShipmentStatus.
func (m EasyShipShipmentStatus) Pointer() *EasyShipShipmentStatus {
	return &m
}

const (

	// EasyShipShipmentStatusPendingSchedule captures enum value "PendingSchedule"
	EasyShipShipmentStatusPendingSchedule EasyShipShipmentStatus = "PendingSchedule"

	// EasyShipShipmentStatusPendingPickUp captures enum value "PendingPickUp"
	EasyShipShipmentStatusPendingPickUp EasyShipShipmentStatus = "PendingPickUp"

	// EasyShipShipmentStatusPendingDropOff captures enum value "PendingDropOff"
	EasyShipShipmentStatusPendingDropOff EasyShipShipmentStatus = "PendingDropOff"

	// EasyShipShipmentStatusLabelCanceled captures enum value "LabelCanceled"
	EasyShipShipmentStatusLabelCanceled EasyShipShipmentStatus = "LabelCanceled"

	// EasyShipShipmentStatusPickedUp captures enum value "PickedUp"
	EasyShipShipmentStatusPickedUp EasyShipShipmentStatus = "PickedUp"

	// EasyShipShipmentStatusDroppedOff captures enum value "DroppedOff"
	EasyShipShipmentStatusDroppedOff EasyShipShipmentStatus = "DroppedOff"

	// EasyShipShipmentStatusAtOriginFC captures enum value "AtOriginFC"
	EasyShipShipmentStatusAtOriginFC EasyShipShipmentStatus = "AtOriginFC"

	// EasyShipShipmentStatusAtDestinationFC captures enum value "AtDestinationFC"
	EasyShipShipmentStatusAtDestinationFC EasyShipShipmentStatus = "AtDestinationFC"

	// EasyShipShipmentStatusDelivered captures enum value "Delivered"
	EasyShipShipmentStatusDelivered EasyShipShipmentStatus = "Delivered"

	// EasyShipShipmentStatusRejectedByBuyer captures enum value "RejectedByBuyer"
	EasyShipShipmentStatusRejectedByBuyer EasyShipShipmentStatus = "RejectedByBuyer"

	// EasyShipShipmentStatusUndeliverable captures enum value "Undeliverable"
	EasyShipShipmentStatusUndeliverable EasyShipShipmentStatus = "Undeliverable"

	// EasyShipShipmentStatusReturningToSeller captures enum value "ReturningToSeller"
	EasyShipShipmentStatusReturningToSeller EasyShipShipmentStatus = "ReturningToSeller"

	// EasyShipShipmentStatusReturnedToSeller captures enum value "ReturnedToSeller"
	EasyShipShipmentStatusReturnedToSeller EasyShipShipmentStatus = "ReturnedToSeller"

	// EasyShipShipmentStatusLost captures enum value "Lost"
	EasyShipShipmentStatusLost EasyShipShipmentStatus = "Lost"

	// EasyShipShipmentStatusOutForDelivery captures enum value "OutForDelivery"
	EasyShipShipmentStatusOutForDelivery EasyShipShipmentStatus = "OutForDelivery"

	// EasyShipShipmentStatusDamaged captures enum value "Damaged"
	EasyShipShipmentStatusDamaged EasyShipShipmentStatus = "Damaged"
)

// for schema
var easyShipShipmentStatusEnum []interface{}

func init() {
	var res []EasyShipShipmentStatus
	if err := json.Unmarshal([]byte(`["PendingSchedule","PendingPickUp","PendingDropOff","LabelCanceled","PickedUp","DroppedOff","AtOriginFC","AtDestinationFC","Delivered","RejectedByBuyer","Undeliverable","ReturningToSeller","ReturnedToSeller","Lost","OutForDelivery","Damaged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		easyShipShipmentStatusEnum = append(easyShipShipmentStatusEnum, v)
	}
}

func (m EasyShipShipmentStatus) validateEasyShipShipmentStatusEnum(path, location string, value EasyShipShipmentStatus) error {
	if err := validate.EnumCase(path, location, value, easyShipShipmentStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this easy ship shipment status
func (m EasyShipShipmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEasyShipShipmentStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this easy ship shipment status based on context it is used
func (m EasyShipShipmentStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
