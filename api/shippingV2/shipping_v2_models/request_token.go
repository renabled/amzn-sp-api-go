// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// RequestToken A unique token generated to identify a getRates operation.
//
// swagger:model RequestToken
type RequestToken string

// Validate validates this request token
func (m RequestToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this request token based on context it is used
func (m RequestToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
