// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OrderID The Amazon-defined identifier for an order placed by the buyer, in 3-7-7 format.
//
// swagger:model OrderId
type OrderID string

// Validate validates this order Id
func (m OrderID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 20); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this order Id based on context it is used
func (m OrderID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
