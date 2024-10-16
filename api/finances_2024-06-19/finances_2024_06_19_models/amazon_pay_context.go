// Code generated by go-swagger; DO NOT EDIT.

package finances_2024_06_19_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AmazonPayContext Additional information related to Amazon Pay.
// Example: {"channel":"MFN","orderType":"Order Type","storeName":"Store 1"}
//
// swagger:model AmazonPayContext
type AmazonPayContext struct {

	// Channel details of related transaction.
	Channel string `json:"channel,omitempty"`

	// The transaction's order type.
	OrderType string `json:"orderType,omitempty"`

	// The name of the store that is related to the transaction.
	StoreName string `json:"storeName,omitempty"`
}

// Validate validates this amazon pay context
func (m *AmazonPayContext) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this amazon pay context based on context it is used
func (m *AmazonPayContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AmazonPayContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmazonPayContext) UnmarshalBinary(b []byte) error {
	var res AmazonPayContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
