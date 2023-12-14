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

// PaymentType Payment type of the purchase.
//
// swagger:model PaymentType
type PaymentType string

func NewPaymentType(value PaymentType) *PaymentType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PaymentType.
func (m PaymentType) Pointer() *PaymentType {
	return &m
}

const (

	// PaymentTypePAYTHROUGHAMAZON captures enum value "PAY_THROUGH_AMAZON"
	PaymentTypePAYTHROUGHAMAZON PaymentType = "PAY_THROUGH_AMAZON"

	// PaymentTypePAYDIRECTTOCARRIER captures enum value "PAY_DIRECT_TO_CARRIER"
	PaymentTypePAYDIRECTTOCARRIER PaymentType = "PAY_DIRECT_TO_CARRIER"
)

// for schema
var paymentTypeEnum []interface{}

func init() {
	var res []PaymentType
	if err := json.Unmarshal([]byte(`["PAY_THROUGH_AMAZON","PAY_DIRECT_TO_CARRIER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentTypeEnum = append(paymentTypeEnum, v)
	}
}

func (m PaymentType) validatePaymentTypeEnum(path, location string, value PaymentType) error {
	if err := validate.EnumCase(path, location, value, paymentTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment type
func (m PaymentType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this payment type based on context it is used
func (m PaymentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
