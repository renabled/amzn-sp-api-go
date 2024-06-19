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

// VolumeUnitOfMeasurement Unit of measurement for the package volume.
//
// swagger:model VolumeUnitOfMeasurement
type VolumeUnitOfMeasurement string

func NewVolumeUnitOfMeasurement(value VolumeUnitOfMeasurement) *VolumeUnitOfMeasurement {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VolumeUnitOfMeasurement.
func (m VolumeUnitOfMeasurement) Pointer() *VolumeUnitOfMeasurement {
	return &m
}

const (

	// VolumeUnitOfMeasurementCUIN captures enum value "CU_IN"
	VolumeUnitOfMeasurementCUIN VolumeUnitOfMeasurement = "CU_IN"

	// VolumeUnitOfMeasurementCBM captures enum value "CBM"
	VolumeUnitOfMeasurementCBM VolumeUnitOfMeasurement = "CBM"

	// VolumeUnitOfMeasurementCC captures enum value "CC"
	VolumeUnitOfMeasurementCC VolumeUnitOfMeasurement = "CC"
)

// for schema
var volumeUnitOfMeasurementEnum []interface{}

func init() {
	var res []VolumeUnitOfMeasurement
	if err := json.Unmarshal([]byte(`["CU_IN","CBM","CC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeUnitOfMeasurementEnum = append(volumeUnitOfMeasurementEnum, v)
	}
}

func (m VolumeUnitOfMeasurement) validateVolumeUnitOfMeasurementEnum(path, location string, value VolumeUnitOfMeasurement) error {
	if err := validate.EnumCase(path, location, value, volumeUnitOfMeasurementEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this volume unit of measurement
func (m VolumeUnitOfMeasurement) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVolumeUnitOfMeasurementEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this volume unit of measurement based on context it is used
func (m VolumeUnitOfMeasurement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
