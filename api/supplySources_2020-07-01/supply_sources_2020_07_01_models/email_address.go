// Code generated by go-swagger; DO NOT EDIT.

package supply_sources_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EmailAddress The email address to which email messages are delivered.
//
// swagger:model EmailAddress
type EmailAddress string

// Validate validates this email address
func (m EmailAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Pattern("", "body", string(m), `^([a-zA-Z0-9_\-\.]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,5})$`); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this email address based on context it is used
func (m EmailAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
