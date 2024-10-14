// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PrepOwner The owner of the preparations, if special preparations are required.
//
// swagger:model PrepOwner
type PrepOwner string

func NewPrepOwner(value PrepOwner) *PrepOwner {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PrepOwner.
func (m PrepOwner) Pointer() *PrepOwner {
	return &m
}

const (

	// PrepOwnerAMAZON captures enum value "AMAZON"
	PrepOwnerAMAZON PrepOwner = "AMAZON"

	// PrepOwnerSELLER captures enum value "SELLER"
	PrepOwnerSELLER PrepOwner = "SELLER"

	// PrepOwnerNONE captures enum value "NONE"
	PrepOwnerNONE PrepOwner = "NONE"
)

// for schema
var prepOwnerEnum []interface{}

func init() {
	var res []PrepOwner
	if err := json.Unmarshal([]byte(`["AMAZON","SELLER","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		prepOwnerEnum = append(prepOwnerEnum, v)
	}
}

func (m PrepOwner) validatePrepOwnerEnum(path, location string, value PrepOwner) error {
	if err := validate.EnumCase(path, location, value, prepOwnerEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this prep owner
func (m PrepOwner) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePrepOwnerEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this prep owner based on context it is used
func (m PrepOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
