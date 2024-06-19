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

// ShipmentSortableField Denotes the field name on which the shipments are to be sorted.
// Example: CREATED_AT
//
// swagger:model ShipmentSortableField
type ShipmentSortableField string

func NewShipmentSortableField(value ShipmentSortableField) *ShipmentSortableField {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ShipmentSortableField.
func (m ShipmentSortableField) Pointer() *ShipmentSortableField {
	return &m
}

const (

	// ShipmentSortableFieldUPDATEDAT captures enum value "UPDATED_AT"
	ShipmentSortableFieldUPDATEDAT ShipmentSortableField = "UPDATED_AT"

	// ShipmentSortableFieldCREATEDAT captures enum value "CREATED_AT"
	ShipmentSortableFieldCREATEDAT ShipmentSortableField = "CREATED_AT"
)

// for schema
var shipmentSortableFieldEnum []interface{}

func init() {
	var res []ShipmentSortableField
	if err := json.Unmarshal([]byte(`["UPDATED_AT","CREATED_AT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shipmentSortableFieldEnum = append(shipmentSortableFieldEnum, v)
	}
}

func (m ShipmentSortableField) validateShipmentSortableFieldEnum(path, location string, value ShipmentSortableField) error {
	if err := validate.EnumCase(path, location, value, shipmentSortableFieldEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this shipment sortable field
func (m ShipmentSortableField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateShipmentSortableFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this shipment sortable field based on context it is used
func (m ShipmentSortableField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
