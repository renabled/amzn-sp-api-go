// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// PackageDimension Number representing the given package dimension.
//
// swagger:model PackageDimension
type PackageDimension float64

// Validate validates this package dimension
func (m PackageDimension) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this package dimension based on context it is used
func (m PackageDimension) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
