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

// CurrencyCode The currency code.
//
// swagger:model CurrencyCode
type CurrencyCode string

func NewCurrencyCode(value CurrencyCode) *CurrencyCode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CurrencyCode.
func (m CurrencyCode) Pointer() *CurrencyCode {
	return &m
}

const (

	// CurrencyCodeUSD captures enum value "USD"
	CurrencyCodeUSD CurrencyCode = "USD"

	// CurrencyCodeGBP captures enum value "GBP"
	CurrencyCodeGBP CurrencyCode = "GBP"
)

// for schema
var currencyCodeEnum []interface{}

func init() {
	var res []CurrencyCode
	if err := json.Unmarshal([]byte(`["USD","GBP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyCodeEnum = append(currencyCodeEnum, v)
	}
}

func (m CurrencyCode) validateCurrencyCodeEnum(path, location string, value CurrencyCode) error {
	if err := validate.EnumCase(path, location, value, currencyCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this currency code
func (m CurrencyCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCurrencyCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this currency code based on context it is used
func (m CurrencyCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
