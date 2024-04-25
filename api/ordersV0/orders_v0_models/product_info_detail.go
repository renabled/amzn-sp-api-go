// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProductInfoDetail Product information on the number of items.
//
// swagger:model ProductInfoDetail
type ProductInfoDetail struct {

	// The total number of items that are included in the ASIN.
	NumberOfItems string `json:"NumberOfItems,omitempty"`
}

// Validate validates this product info detail
func (m *ProductInfoDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this product info detail based on context it is used
func (m *ProductInfoDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProductInfoDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductInfoDetail) UnmarshalBinary(b []byte) error {
	var res ProductInfoDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
