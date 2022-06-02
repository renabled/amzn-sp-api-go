// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_transactions_v1_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Transaction The transaction status details.
//
// swagger:model Transaction
type Transaction struct {

	// Error code and message for the failed transaction. Only available when transaction status is 'Failure'.
	Errors ErrorList `json:"errors,omitempty"`

	// Current processing status of the transaction.
	// Required: true
	// Enum: [Failure Processing Success]
	Status *string `json:"status"`

	// The unique identifier sent in the 'transactionId' field in response to the post request of a specific transaction.
	// Required: true
	TransactionID *string `json:"transactionId"`
}

// Validate validates this transaction
func (m *Transaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if err := m.Errors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

var transactionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Failure","Processing","Success"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionTypeStatusPropEnum = append(transactionTypeStatusPropEnum, v)
	}
}

const (

	// TransactionStatusFailure captures enum value "Failure"
	TransactionStatusFailure string = "Failure"

	// TransactionStatusProcessing captures enum value "Processing"
	TransactionStatusProcessing string = "Processing"

	// TransactionStatusSuccess captures enum value "Success"
	TransactionStatusSuccess string = "Success"
)

// prop value enum
func (m *Transaction) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transactionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Transaction) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Transaction) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transactionId", "body", m.TransactionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transaction based on the context it is used
func (m *Transaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Transaction) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Errors.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Transaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transaction) UnmarshalBinary(b []byte) error {
	var res Transaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
