// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaymentExecutionDetailItem Information about a sub-payment method used to pay for a COD order.
//
// swagger:model PaymentExecutionDetailItem
type PaymentExecutionDetailItem struct {

	// payment
	// Required: true
	Payment *Money `json:"Payment"`

	// A sub-payment method for a COD order.
	//
	// Possible values:
	// * `COD`: Cash On Delivery.
	// * `GC`: Gift Card.
	// * `PointsAccount`: Amazon Points.
	// * `Invoice`: Invoice.
	// Required: true
	PaymentMethod *string `json:"PaymentMethod"`
}

// Validate validates this payment execution detail item
func (m *PaymentExecutionDetailItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentExecutionDetailItem) validatePayment(formats strfmt.Registry) error {

	if err := validate.Required("Payment", "body", m.Payment); err != nil {
		return err
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Payment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Payment")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentExecutionDetailItem) validatePaymentMethod(formats strfmt.Registry) error {

	if err := validate.Required("PaymentMethod", "body", m.PaymentMethod); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this payment execution detail item based on the context it is used
func (m *PaymentExecutionDetailItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentExecutionDetailItem) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	if m.Payment != nil {
		if err := m.Payment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Payment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Payment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentExecutionDetailItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentExecutionDetailItem) UnmarshalBinary(b []byte) error {
	var res PaymentExecutionDetailItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
