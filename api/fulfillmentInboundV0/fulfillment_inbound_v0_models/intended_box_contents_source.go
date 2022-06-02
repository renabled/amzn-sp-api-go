// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IntendedBoxContentsSource How the seller intends to provide box contents information for a shipment.
//
// swagger:model IntendedBoxContentsSource
type IntendedBoxContentsSource string

func NewIntendedBoxContentsSource(value IntendedBoxContentsSource) *IntendedBoxContentsSource {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IntendedBoxContentsSource.
func (m IntendedBoxContentsSource) Pointer() *IntendedBoxContentsSource {
	return &m
}

const (

	// IntendedBoxContentsSourceNONE captures enum value "NONE"
	IntendedBoxContentsSourceNONE IntendedBoxContentsSource = "NONE"

	// IntendedBoxContentsSourceFEED captures enum value "FEED"
	IntendedBoxContentsSourceFEED IntendedBoxContentsSource = "FEED"

	// IntendedBoxContentsSourceNr2DBARCODE captures enum value "2D_BARCODE"
	IntendedBoxContentsSourceNr2DBARCODE IntendedBoxContentsSource = "2D_BARCODE"
)

// for schema
var intendedBoxContentsSourceEnum []interface{}

func init() {
	var res []IntendedBoxContentsSource
	if err := json.Unmarshal([]byte(`["NONE","FEED","2D_BARCODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		intendedBoxContentsSourceEnum = append(intendedBoxContentsSourceEnum, v)
	}
}

func (m IntendedBoxContentsSource) validateIntendedBoxContentsSourceEnum(path, location string, value IntendedBoxContentsSource) error {
	if err := validate.EnumCase(path, location, value, intendedBoxContentsSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this intended box contents source
func (m IntendedBoxContentsSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIntendedBoxContentsSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this intended box contents source based on context it is used
func (m IntendedBoxContentsSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
