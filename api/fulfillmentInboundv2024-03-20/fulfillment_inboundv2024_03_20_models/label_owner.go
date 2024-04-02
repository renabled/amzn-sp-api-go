// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inboundv2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// LabelOwner Specifies who will label the items. Options include `AMAZON` and `SELLER`.
//
// swagger:model LabelOwner
type LabelOwner string

func NewLabelOwner(value LabelOwner) *LabelOwner {
	return &value
}

// Pointer returns a pointer to a freshly-allocated LabelOwner.
func (m LabelOwner) Pointer() *LabelOwner {
	return &m
}

const (

	// LabelOwnerAMAZON captures enum value "AMAZON"
	LabelOwnerAMAZON LabelOwner = "AMAZON"

	// LabelOwnerSELLER captures enum value "SELLER"
	LabelOwnerSELLER LabelOwner = "SELLER"
)

// for schema
var labelOwnerEnum []interface{}

func init() {
	var res []LabelOwner
	if err := json.Unmarshal([]byte(`["AMAZON","SELLER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		labelOwnerEnum = append(labelOwnerEnum, v)
	}
}

func (m LabelOwner) validateLabelOwnerEnum(path, location string, value LabelOwner) error {
	if err := validate.EnumCase(path, location, value, labelOwnerEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this label owner
func (m LabelOwner) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLabelOwnerEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this label owner based on context it is used
func (m LabelOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
