// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Status The status of the package being shipped.
//
// swagger:model Status
type Status string

func NewStatus(value Status) *Status {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Status.
func (m Status) Pointer() *Status {
	return &m
}

const (

	// StatusPreTransit captures enum value "PreTransit"
	StatusPreTransit Status = "PreTransit"

	// StatusInTransit captures enum value "InTransit"
	StatusInTransit Status = "InTransit"

	// StatusDelivered captures enum value "Delivered"
	StatusDelivered Status = "Delivered"

	// StatusLost captures enum value "Lost"
	StatusLost Status = "Lost"

	// StatusOutForDelivery captures enum value "OutForDelivery"
	StatusOutForDelivery Status = "OutForDelivery"

	// StatusRejected captures enum value "Rejected"
	StatusRejected Status = "Rejected"

	// StatusUndeliverable captures enum value "Undeliverable"
	StatusUndeliverable Status = "Undeliverable"

	// StatusDeliveryAttempted captures enum value "DeliveryAttempted"
	StatusDeliveryAttempted Status = "DeliveryAttempted"

	// StatusPickupCancelled captures enum value "PickupCancelled"
	StatusPickupCancelled Status = "PickupCancelled"
)

// for schema
var statusEnum []interface{}

func init() {
	var res []Status
	if err := json.Unmarshal([]byte(`["PreTransit","InTransit","Delivered","Lost","OutForDelivery","Rejected","Undeliverable","DeliveryAttempted","PickupCancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusEnum = append(statusEnum, v)
	}
}

func (m Status) validateStatusEnum(path, location string, value Status) error {
	if err := validate.EnumCase(path, location, value, statusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this status
func (m Status) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this status based on context it is used
func (m Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
