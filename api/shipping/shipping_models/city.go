// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// City The city where the person, business or institution is located.
//
// swagger:model City
type City string

// Validate validates this city
func (m City) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 50); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this city based on context it is used
func (m City) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
