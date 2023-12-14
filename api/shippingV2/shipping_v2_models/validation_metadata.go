// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ValidationMetadata ValidationMetadata Details
//
// swagger:model ValidationMetadata
type ValidationMetadata struct {

	// errorMessage for the error.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// validationStrategy for the error.
	ValidationStrategy string `json:"validationStrategy,omitempty"`

	// Value.
	Value string `json:"value,omitempty"`
}

// Validate validates this validation metadata
func (m *ValidationMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validation metadata based on context it is used
func (m *ValidationMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ValidationMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidationMetadata) UnmarshalBinary(b []byte) error {
	var res ValidationMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
