// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

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

// Weight The weight in the units indicated.
//
// swagger:model Weight
type Weight struct {

	// The unit of measurement.
	// Required: true
	// Enum: [GRAM KILOGRAM OUNCE POUND]
	Unit *string `json:"unit"`

	// The measurement value.
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this weight
func (m *Weight) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var weightTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GRAM","KILOGRAM","OUNCE","POUND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		weightTypeUnitPropEnum = append(weightTypeUnitPropEnum, v)
	}
}

const (

	// WeightUnitGRAM captures enum value "GRAM"
	WeightUnitGRAM string = "GRAM"

	// WeightUnitKILOGRAM captures enum value "KILOGRAM"
	WeightUnitKILOGRAM string = "KILOGRAM"

	// WeightUnitOUNCE captures enum value "OUNCE"
	WeightUnitOUNCE string = "OUNCE"

	// WeightUnitPOUND captures enum value "POUND"
	WeightUnitPOUND string = "POUND"
)

// prop value enum
func (m *Weight) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, weightTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Weight) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *Weight) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this weight based on context it is used
func (m *Weight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Weight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Weight) UnmarshalBinary(b []byte) error {
	var res Weight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
