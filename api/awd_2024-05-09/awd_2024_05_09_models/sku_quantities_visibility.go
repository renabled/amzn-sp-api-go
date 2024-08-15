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

// SkuQuantitiesVisibility Enum to specify if returned shipment should include SKU quantity details
// Example: SHOW
//
// swagger:model SkuQuantitiesVisibility
type SkuQuantitiesVisibility string

func NewSkuQuantitiesVisibility(value SkuQuantitiesVisibility) *SkuQuantitiesVisibility {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SkuQuantitiesVisibility.
func (m SkuQuantitiesVisibility) Pointer() *SkuQuantitiesVisibility {
	return &m
}

const (

	// SkuQuantitiesVisibilitySHOW captures enum value "SHOW"
	SkuQuantitiesVisibilitySHOW SkuQuantitiesVisibility = "SHOW"

	// SkuQuantitiesVisibilityHIDE captures enum value "HIDE"
	SkuQuantitiesVisibilityHIDE SkuQuantitiesVisibility = "HIDE"
)

// for schema
var skuQuantitiesVisibilityEnum []interface{}

func init() {
	var res []SkuQuantitiesVisibility
	if err := json.Unmarshal([]byte(`["SHOW","HIDE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		skuQuantitiesVisibilityEnum = append(skuQuantitiesVisibilityEnum, v)
	}
}

func (m SkuQuantitiesVisibility) validateSkuQuantitiesVisibilityEnum(path, location string, value SkuQuantitiesVisibility) error {
	if err := validate.EnumCase(path, location, value, skuQuantitiesVisibilityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sku quantities visibility
func (m SkuQuantitiesVisibility) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSkuQuantitiesVisibilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sku quantities visibility based on context it is used
func (m SkuQuantitiesVisibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
