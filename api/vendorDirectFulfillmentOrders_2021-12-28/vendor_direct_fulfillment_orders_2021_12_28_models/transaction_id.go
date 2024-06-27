// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_orders_2021_12_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionID Response containing the transaction ID.
//
// swagger:model TransactionId
type TransactionID struct {

	// GUID assigned by Amazon to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this transaction Id
func (m *TransactionID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this transaction Id based on context it is used
func (m *TransactionID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionID) UnmarshalBinary(b []byte) error {
	var res TransactionID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
