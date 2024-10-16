// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductAdsPaymentEvent A Sponsored Products payment event.
//
// swagger:model ProductAdsPaymentEvent
type ProductAdsPaymentEvent struct {

	// Base amount of the transaction, before tax.
	BaseValue *Currency `json:"baseValue,omitempty"`

	// The identifier for the invoice that includes the transaction.
	InvoiceID string `json:"invoiceId,omitempty"`

	// The date and time when the financial event was posted.
	// Format: date-time
	PostedDate Date `json:"postedDate,omitempty"`

	// Tax amount of the transaction.
	TaxValue *Currency `json:"taxValue,omitempty"`

	// Indicates if the transaction is for a charge or a refund.
	//
	// Possible values:
	//
	// * `charge`
	//
	// * `refund`
	TransactionType string `json:"transactionType,omitempty"`

	// The total amount of the transaction. Equal to `baseValue` + `taxValue`.
	TransactionValue *Currency `json:"transactionValue,omitempty"`
}

// Validate validates this product ads payment event
func (m *ProductAdsPaymentEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductAdsPaymentEvent) validateBaseValue(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseValue) { // not required
		return nil
	}

	if m.BaseValue != nil {
		if err := m.BaseValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseValue")
			}
			return err
		}
	}

	return nil
}

func (m *ProductAdsPaymentEvent) validatePostedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PostedDate) { // not required
		return nil
	}

	if err := m.PostedDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("postedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("postedDate")
		}
		return err
	}

	return nil
}

func (m *ProductAdsPaymentEvent) validateTaxValue(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxValue) { // not required
		return nil
	}

	if m.TaxValue != nil {
		if err := m.TaxValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxValue")
			}
			return err
		}
	}

	return nil
}

func (m *ProductAdsPaymentEvent) validateTransactionValue(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionValue) { // not required
		return nil
	}

	if m.TransactionValue != nil {
		if err := m.TransactionValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactionValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transactionValue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this product ads payment event based on the context it is used
func (m *ProductAdsPaymentEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransactionValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductAdsPaymentEvent) contextValidateBaseValue(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseValue != nil {
		if err := m.BaseValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseValue")
			}
			return err
		}
	}

	return nil
}

func (m *ProductAdsPaymentEvent) contextValidatePostedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PostedDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("postedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("postedDate")
		}
		return err
	}

	return nil
}

func (m *ProductAdsPaymentEvent) contextValidateTaxValue(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxValue != nil {
		if err := m.TaxValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxValue")
			}
			return err
		}
	}

	return nil
}

func (m *ProductAdsPaymentEvent) contextValidateTransactionValue(ctx context.Context, formats strfmt.Registry) error {

	if m.TransactionValue != nil {
		if err := m.TransactionValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactionValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transactionValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductAdsPaymentEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductAdsPaymentEvent) UnmarshalBinary(b []byte) error {
	var res ProductAdsPaymentEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
