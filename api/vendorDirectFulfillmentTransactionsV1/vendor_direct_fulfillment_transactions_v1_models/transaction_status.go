// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_transactions_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionStatus The payload for the getTransactionStatus operation.
//
// swagger:model TransactionStatus
type TransactionStatus struct {

	// transaction status
	TransactionStatus *Transaction `json:"transactionStatus,omitempty"`
}

// Validate validates this transaction status
func (m *TransactionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransactionStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionStatus) validateTransactionStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.TransactionStatus) { // not required
		return nil
	}

	if m.TransactionStatus != nil {
		if err := m.TransactionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactionStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transactionStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transaction status based on the context it is used
func (m *TransactionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransactionStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionStatus) contextValidateTransactionStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.TransactionStatus != nil {
		if err := m.TransactionStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transactionStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transactionStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionStatus) UnmarshalBinary(b []byte) error {
	var res TransactionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
