// Code generated by go-swagger; DO NOT EDIT.

package product_fees_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OptionalFulfillmentProgram An optional enrollment program to return the estimated fees when the offer is fulfilled by Amazon (IsAmazonFulfilled is set to true).
//
// swagger:model OptionalFulfillmentProgram
type OptionalFulfillmentProgram string

func NewOptionalFulfillmentProgram(value OptionalFulfillmentProgram) *OptionalFulfillmentProgram {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OptionalFulfillmentProgram.
func (m OptionalFulfillmentProgram) Pointer() *OptionalFulfillmentProgram {
	return &m
}

const (

	// OptionalFulfillmentProgramFBACORE captures enum value "FBA_CORE"
	OptionalFulfillmentProgramFBACORE OptionalFulfillmentProgram = "FBA_CORE"

	// OptionalFulfillmentProgramFBASNL captures enum value "FBA_SNL"
	OptionalFulfillmentProgramFBASNL OptionalFulfillmentProgram = "FBA_SNL"

	// OptionalFulfillmentProgramFBAEFN captures enum value "FBA_EFN"
	OptionalFulfillmentProgramFBAEFN OptionalFulfillmentProgram = "FBA_EFN"
)

// for schema
var optionalFulfillmentProgramEnum []interface{}

func init() {
	var res []OptionalFulfillmentProgram
	if err := json.Unmarshal([]byte(`["FBA_CORE","FBA_SNL","FBA_EFN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		optionalFulfillmentProgramEnum = append(optionalFulfillmentProgramEnum, v)
	}
}

func (m OptionalFulfillmentProgram) validateOptionalFulfillmentProgramEnum(path, location string, value OptionalFulfillmentProgram) error {
	if err := validate.EnumCase(path, location, value, optionalFulfillmentProgramEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this optional fulfillment program
func (m OptionalFulfillmentProgram) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOptionalFulfillmentProgramEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this optional fulfillment program based on context it is used
func (m OptionalFulfillmentProgram) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
