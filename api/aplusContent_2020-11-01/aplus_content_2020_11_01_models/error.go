// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Error Error response returned when the request is unsuccessful.
//
// swagger:model Error
type Error struct {

	// The code that identifies the type of error condition.
	// Required: true
	// Min Length: 1
	Code *string `json:"code"`

	// Additional information, if available, to clarify the error condition.
	// Min Length: 1
	Details string `json:"details,omitempty"`

	// A human readable description of the error condition.
	// Required: true
	// Min Length: 1
	Message *string `json:"message"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	if err := validate.MinLength("code", "body", *m.Code, 1); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if err := validate.MinLength("details", "body", m.Details, 1); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	if err := validate.MinLength("message", "body", *m.Message, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error based on context it is used
func (m *Error) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
