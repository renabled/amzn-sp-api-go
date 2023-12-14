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

// DocumentSize The size dimensions of the label.
//
// swagger:model DocumentSize
type DocumentSize struct {

	// The length of the document measured in the units specified.
	// Required: true
	Length *float64 `json:"length"`

	// The unit of measurement.
	// Required: true
	// Enum: [INCH CENTIMETER]
	Unit *string `json:"unit"`

	// The width of the document measured in the units specified.
	// Required: true
	Width *float64 `json:"width"`
}

// Validate validates this document size
func (m *DocumentSize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocumentSize) validateLength(formats strfmt.Registry) error {

	if err := validate.Required("length", "body", m.Length); err != nil {
		return err
	}

	return nil
}

var documentSizeTypeUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INCH","CENTIMETER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		documentSizeTypeUnitPropEnum = append(documentSizeTypeUnitPropEnum, v)
	}
}

const (

	// DocumentSizeUnitINCH captures enum value "INCH"
	DocumentSizeUnitINCH string = "INCH"

	// DocumentSizeUnitCENTIMETER captures enum value "CENTIMETER"
	DocumentSizeUnitCENTIMETER string = "CENTIMETER"
)

// prop value enum
func (m *DocumentSize) validateUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, documentSizeTypeUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DocumentSize) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitEnum("unit", "body", *m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *DocumentSize) validateWidth(formats strfmt.Registry) error {

	if err := validate.Required("width", "body", m.Width); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this document size based on context it is used
func (m *DocumentSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DocumentSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentSize) UnmarshalBinary(b []byte) error {
	var res DocumentSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
