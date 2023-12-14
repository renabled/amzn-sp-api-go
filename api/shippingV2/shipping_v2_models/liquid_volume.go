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

// LiquidVolume Liquid Volume.
//
// swagger:model LiquidVolume
type LiquidVolume struct {

	// The unit of measurement.
	// Required: true
	// Enum: [ML L FL_OZ GAL PT QT C]
	Unit *string `json:"unit"`

	// The measurement value.
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this liquid volume
func (m *LiquidVolume) Validate(formats strfmt.Registry) error {
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

var liquidVolumeTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ML","L","FL_OZ","GAL","PT","QT","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		liquidVolumeTypeUnitPropEnum = append(liquidVolumeTypeUnitPropEnum, v)
	}
}

const (

	// LiquidVolumeUnitML captures enum value "ML"
	LiquidVolumeUnitML string = "ML"

	// LiquidVolumeUnitL captures enum value "L"
	LiquidVolumeUnitL string = "L"

	// LiquidVolumeUnitFLOZ captures enum value "FL_OZ"
	LiquidVolumeUnitFLOZ string = "FL_OZ"

	// LiquidVolumeUnitGAL captures enum value "GAL"
	LiquidVolumeUnitGAL string = "GAL"

	// LiquidVolumeUnitPT captures enum value "PT"
	LiquidVolumeUnitPT string = "PT"

	// LiquidVolumeUnitQT captures enum value "QT"
	LiquidVolumeUnitQT string = "QT"

	// LiquidVolumeUnitC captures enum value "C"
	LiquidVolumeUnitC string = "C"
)

// prop value enum
func (m *LiquidVolume) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, liquidVolumeTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LiquidVolume) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *LiquidVolume) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this liquid volume based on context it is used
func (m *LiquidVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LiquidVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LiquidVolume) UnmarshalBinary(b []byte) error {
	var res LiquidVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
