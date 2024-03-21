// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Duration Duration after manufacturing date during which the product is valid for consumption.
//
// swagger:model Duration
type Duration struct {

	// Unit for duration.
	// Required: true
	// Enum: [Days Months]
	DurationUnit *string `json:"durationUnit"`

	// Value for the duration in terms of the durationUnit.
	// Required: true
	DurationValue *int64 `json:"durationValue"`
}

// Validate validates this duration
func (m *Duration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDurationUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDurationValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var durationTypeDurationUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Days","Months"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		durationTypeDurationUnitPropEnum = append(durationTypeDurationUnitPropEnum, v)
	}
}

const (

	// DurationDurationUnitDays captures enum value "Days"
	DurationDurationUnitDays string = "Days"

	// DurationDurationUnitMonths captures enum value "Months"
	DurationDurationUnitMonths string = "Months"
)

// prop value enum
func (m *Duration) validateDurationUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, durationTypeDurationUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Duration) validateDurationUnit(formats strfmt.Registry) error {

	if err := validate.Required("durationUnit", "body", m.DurationUnit); err != nil {
		return err
	}

	// value enum
	if err := m.validateDurationUnitEnum("durationUnit", "body", *m.DurationUnit); err != nil {
		return err
	}

	return nil
}

func (m *Duration) validateDurationValue(formats strfmt.Registry) error {

	if err := validate.Required("durationValue", "body", m.DurationValue); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this duration based on context it is used
func (m *Duration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Duration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Duration) UnmarshalBinary(b []byte) error {
	var res Duration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
