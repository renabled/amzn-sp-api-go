// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// UnitOfMeasurement Indicates the unit of measurement.
//
// swagger:model UnitOfMeasurement
type UnitOfMeasurement string

func NewUnitOfMeasurement(value UnitOfMeasurement) *UnitOfMeasurement {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UnitOfMeasurement.
func (m UnitOfMeasurement) Pointer() *UnitOfMeasurement {
	return &m
}

const (

	// UnitOfMeasurementInches captures enum value "inches"
	UnitOfMeasurementInches UnitOfMeasurement = "inches"

	// UnitOfMeasurementCentimeters captures enum value "centimeters"
	UnitOfMeasurementCentimeters UnitOfMeasurement = "centimeters"
)

// for schema
var unitOfMeasurementEnum []interface{}

func init() {
	var res []UnitOfMeasurement
	if err := json.Unmarshal([]byte(`["inches","centimeters"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		unitOfMeasurementEnum = append(unitOfMeasurementEnum, v)
	}
}

func (m UnitOfMeasurement) validateUnitOfMeasurementEnum(path, location string, value UnitOfMeasurement) error {
	if err := validate.EnumCase(path, location, value, unitOfMeasurementEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this unit of measurement
func (m UnitOfMeasurement) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUnitOfMeasurementEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this unit of measurement based on context it is used
func (m UnitOfMeasurement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
