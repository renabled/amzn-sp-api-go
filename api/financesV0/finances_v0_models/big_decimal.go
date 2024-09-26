// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// BigDecimal A signed decimal number.
//
// swagger:model BigDecimal
type BigDecimal float64

// Validate validates this big decimal
func (m BigDecimal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this big decimal based on context it is used
func (m BigDecimal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
