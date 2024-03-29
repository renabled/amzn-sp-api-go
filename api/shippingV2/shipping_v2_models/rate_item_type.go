// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RateItemType Type of the rateItem.
//
// swagger:model RateItemType
type RateItemType string

func NewRateItemType(value RateItemType) *RateItemType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RateItemType.
func (m RateItemType) Pointer() *RateItemType {
	return &m
}

const (

	// RateItemTypeMANDATORY captures enum value "MANDATORY"
	RateItemTypeMANDATORY RateItemType = "MANDATORY"

	// RateItemTypeOPTIONAL captures enum value "OPTIONAL"
	RateItemTypeOPTIONAL RateItemType = "OPTIONAL"

	// RateItemTypeINCLUDED captures enum value "INCLUDED"
	RateItemTypeINCLUDED RateItemType = "INCLUDED"
)

// for schema
var rateItemTypeEnum []interface{}

func init() {
	var res []RateItemType
	if err := json.Unmarshal([]byte(`["MANDATORY","OPTIONAL","INCLUDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rateItemTypeEnum = append(rateItemTypeEnum, v)
	}
}

func (m RateItemType) validateRateItemTypeEnum(path, location string, value RateItemType) error {
	if err := validate.EnumCase(path, location, value, rateItemTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this rate item type
func (m RateItemType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRateItemTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this rate item type based on context it is used
func (m RateItemType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
