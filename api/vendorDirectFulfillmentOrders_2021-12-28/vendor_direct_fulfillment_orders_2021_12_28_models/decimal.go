// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_orders_2021_12_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Decimal A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation.
//
// swagger:model Decimal
type Decimal string

// Validate validates this decimal
func (m Decimal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this decimal based on context it is used
func (m Decimal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
