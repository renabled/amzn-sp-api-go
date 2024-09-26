// Code generated by go-swagger; DO NOT EDIT.

package finances_2024_06_19_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListTransactionsResponse The response schema for the `listTransactions` operation.
//
// swagger:model ListTransactionsResponse
type ListTransactionsResponse struct {

	// The response includes `nextToken` when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextToken` is null. Note that this operation can return empty pages.
	NextToken string `json:"nextToken,omitempty"`

	// transactions
	Transactions Transactions `json:"transactions,omitempty"`
}

// Validate validates this list transactions response
func (m *ListTransactionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTransactionsResponse) validateTransactions(formats strfmt.Registry) error {
	if swag.IsZero(m.Transactions) { // not required
		return nil
	}

	if err := m.Transactions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactions")
		}
		return err
	}

	return nil
}

// ContextValidate validate this list transactions response based on the context it is used
func (m *ListTransactionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransactions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListTransactionsResponse) contextValidateTransactions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Transactions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("transactions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("transactions")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListTransactionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListTransactionsResponse) UnmarshalBinary(b []byte) error {
	var res ListTransactionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
