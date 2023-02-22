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

// PackedQuantity Details of item quantity.
//
// swagger:model packedQuantity
type PackedQuantity struct {

	// Amount of units shipped for a specific item at a shipment level. If the item is present only in certain cartons or pallets within the shipment, please provide this at the appropriate carton or pallet level.
	// Required: true
	Amount *int64 `json:"amount"`

	// Unit of measure for the shipped quantity.
	// Required: true
	// Enum: [Cases Eaches]
	UnitOfMeasure *string `json:"unitOfMeasure"`

	// The case size, in the event that we ordered using cases. Otherwise, 1.
	UnitSize int64 `json:"unitSize,omitempty"`
}

// Validate validates this packed quantity
func (m *PackedQuantity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackedQuantity) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

var packedQuantityTypeUnitOfMeasurePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cases","Eaches"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packedQuantityTypeUnitOfMeasurePropEnum = append(packedQuantityTypeUnitOfMeasurePropEnum, v)
	}
}

const (

	// PackedQuantityUnitOfMeasureCases captures enum value "Cases"
	PackedQuantityUnitOfMeasureCases string = "Cases"

	// PackedQuantityUnitOfMeasureEaches captures enum value "Eaches"
	PackedQuantityUnitOfMeasureEaches string = "Eaches"
)

// prop value enum
func (m *PackedQuantity) validateUnitOfMeasureEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packedQuantityTypeUnitOfMeasurePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PackedQuantity) validateUnitOfMeasure(formats strfmt.Registry) error {

	if err := validate.Required("unitOfMeasure", "body", m.UnitOfMeasure); err != nil {
		return err
	}

	// value enum
	if err := m.validateUnitOfMeasureEnum("unitOfMeasure", "body", *m.UnitOfMeasure); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this packed quantity based on context it is used
func (m *PackedQuantity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackedQuantity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackedQuantity) UnmarshalBinary(b []byte) error {
	var res PackedQuantity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
