// Code generated by go-swagger; DO NOT EDIT.

package awd_2024_05_09_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OutboundShipmentStatus Possible shipment statuses for outbound shipments.
// Example: CREATED
//
// swagger:model OutboundShipmentStatus
type OutboundShipmentStatus string

func NewOutboundShipmentStatus(value OutboundShipmentStatus) *OutboundShipmentStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OutboundShipmentStatus.
func (m OutboundShipmentStatus) Pointer() *OutboundShipmentStatus {
	return &m
}

const (

	// OutboundShipmentStatusCREATED captures enum value "CREATED"
	OutboundShipmentStatusCREATED OutboundShipmentStatus = "CREATED"

	// OutboundShipmentStatusINTRANSIT captures enum value "IN_TRANSIT"
	OutboundShipmentStatusINTRANSIT OutboundShipmentStatus = "IN_TRANSIT"

	// OutboundShipmentStatusDELIVERED captures enum value "DELIVERED"
	OutboundShipmentStatusDELIVERED OutboundShipmentStatus = "DELIVERED"

	// OutboundShipmentStatusRECEIVING captures enum value "RECEIVING"
	OutboundShipmentStatusRECEIVING OutboundShipmentStatus = "RECEIVING"

	// OutboundShipmentStatusRECEIVED captures enum value "RECEIVED"
	OutboundShipmentStatusRECEIVED OutboundShipmentStatus = "RECEIVED"

	// OutboundShipmentStatusCLOSED captures enum value "CLOSED"
	OutboundShipmentStatusCLOSED OutboundShipmentStatus = "CLOSED"

	// OutboundShipmentStatusCANCELLED captures enum value "CANCELLED"
	OutboundShipmentStatusCANCELLED OutboundShipmentStatus = "CANCELLED"

	// OutboundShipmentStatusFAILED captures enum value "FAILED"
	OutboundShipmentStatusFAILED OutboundShipmentStatus = "FAILED"
)

// for schema
var outboundShipmentStatusEnum []interface{}

func init() {
	var res []OutboundShipmentStatus
	if err := json.Unmarshal([]byte(`["CREATED","IN_TRANSIT","DELIVERED","RECEIVING","RECEIVED","CLOSED","CANCELLED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		outboundShipmentStatusEnum = append(outboundShipmentStatusEnum, v)
	}
}

func (m OutboundShipmentStatus) validateOutboundShipmentStatusEnum(path, location string, value OutboundShipmentStatus) error {
	if err := validate.EnumCase(path, location, value, outboundShipmentStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this outbound shipment status
func (m OutboundShipmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOutboundShipmentStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this outbound shipment status based on context it is used
func (m OutboundShipmentStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
