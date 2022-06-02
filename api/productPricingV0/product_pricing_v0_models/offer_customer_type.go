// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OfferCustomerType offer customer type
//
// swagger:model OfferCustomerType
type OfferCustomerType string

func NewOfferCustomerType(value OfferCustomerType) *OfferCustomerType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OfferCustomerType.
func (m OfferCustomerType) Pointer() *OfferCustomerType {
	return &m
}

const (

	// OfferCustomerTypeB2C captures enum value "B2C"
	OfferCustomerTypeB2C OfferCustomerType = "B2C"

	// OfferCustomerTypeB2B captures enum value "B2B"
	OfferCustomerTypeB2B OfferCustomerType = "B2B"
)

// for schema
var offerCustomerTypeEnum []interface{}

func init() {
	var res []OfferCustomerType
	if err := json.Unmarshal([]byte(`["B2C","B2B"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		offerCustomerTypeEnum = append(offerCustomerTypeEnum, v)
	}
}

func (m OfferCustomerType) validateOfferCustomerTypeEnum(path, location string, value OfferCustomerType) error {
	if err := validate.EnumCase(path, location, value, offerCustomerTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this offer customer type
func (m OfferCustomerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOfferCustomerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this offer customer type based on context it is used
func (m OfferCustomerType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
