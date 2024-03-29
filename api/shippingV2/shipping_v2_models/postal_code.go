// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// PostalCode The postal code of that address. It contains a series of letters or digits or both, sometimes including spaces or punctuation.
//
// swagger:model PostalCode
type PostalCode string

// Validate validates this postal code
func (m PostalCode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this postal code based on context it is used
func (m PostalCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
