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

// LabelPrepType The type of label preparation that is required for the inbound shipment.
//
// swagger:model LabelPrepType
type LabelPrepType string

func NewLabelPrepType(value LabelPrepType) *LabelPrepType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated LabelPrepType.
func (m LabelPrepType) Pointer() *LabelPrepType {
	return &m
}

const (

	// LabelPrepTypeNOLABEL captures enum value "NO_LABEL"
	LabelPrepTypeNOLABEL LabelPrepType = "NO_LABEL"

	// LabelPrepTypeSELLERLABEL captures enum value "SELLER_LABEL"
	LabelPrepTypeSELLERLABEL LabelPrepType = "SELLER_LABEL"

	// LabelPrepTypeAMAZONLABEL captures enum value "AMAZON_LABEL"
	LabelPrepTypeAMAZONLABEL LabelPrepType = "AMAZON_LABEL"
)

// for schema
var labelPrepTypeEnum []interface{}

func init() {
	var res []LabelPrepType
	if err := json.Unmarshal([]byte(`["NO_LABEL","SELLER_LABEL","AMAZON_LABEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		labelPrepTypeEnum = append(labelPrepTypeEnum, v)
	}
}

func (m LabelPrepType) validateLabelPrepTypeEnum(path, location string, value LabelPrepType) error {
	if err := validate.EnumCase(path, location, value, labelPrepTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this label prep type
func (m LabelPrepType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLabelPrepTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this label prep type based on context it is used
func (m LabelPrepType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
