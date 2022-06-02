// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReasonCodeDetails A return reason code, a description, and an optional description translation.
//
// swagger:model ReasonCodeDetails
type ReasonCodeDetails struct {

	// A human readable description of the return reason code.
	// Required: true
	Description *string `json:"description"`

	// A code that indicates a valid return reason.
	// Required: true
	ReturnReasonCode *string `json:"returnReasonCode"`

	// A translation of the description. The translation is in the language specified in the Language request parameter.
	TranslatedDescription string `json:"translatedDescription,omitempty"`
}

// Validate validates this reason code details
func (m *ReasonCodeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnReasonCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReasonCodeDetails) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ReasonCodeDetails) validateReturnReasonCode(formats strfmt.Registry) error {

	if err := validate.Required("returnReasonCode", "body", m.ReturnReasonCode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this reason code details based on context it is used
func (m *ReasonCodeDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReasonCodeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReasonCodeDetails) UnmarshalBinary(b []byte) error {
	var res ReasonCodeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
