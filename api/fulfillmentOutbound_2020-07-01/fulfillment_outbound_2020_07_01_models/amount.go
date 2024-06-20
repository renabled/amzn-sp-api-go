// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

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

// Amount A quantity based on unit of measure.
//
// swagger:model Amount
type Amount struct {

	// The unit of measure for the amount.
	// Required: true
	// Enum: [Eaches]
	UnitOfMeasure *string `json:"unitOfMeasure"`

	// The amount of a product in the associated unit of measure.
	// Required: true
	Value *Decimal `json:"value"`
}

// Validate validates this amount
func (m *Amount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnitOfMeasure(formats); err != nil {
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

var amountTypeUnitOfMeasurePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Eaches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		amountTypeUnitOfMeasurePropEnum = append(amountTypeUnitOfMeasurePropEnum, v)
	}
}

const (

	// AmountUnitOfMeasureEaches captures enum value "Eaches"
	AmountUnitOfMeasureEaches string = "Eaches"
)

// prop value enum
func (m *Amount) validateUnitOfMeasureEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, amountTypeUnitOfMeasurePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Amount) validateUnitOfMeasure(formats strfmt.Registry) error {

	if err := validate.Required("unitOfMeasure", "body", m.UnitOfMeasure); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitOfMeasureEnum("unitOfMeasure", "body", *m.UnitOfMeasure); err != nil {
		return err
	}

	return nil
}

func (m *Amount) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this amount based on the context it is used
func (m *Amount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Amount) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Amount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Amount) UnmarshalBinary(b []byte) error {
	var res Amount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
