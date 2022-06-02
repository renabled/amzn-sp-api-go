// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScopeOfWork The scope of work for the order.
//
// swagger:model ScopeOfWork
type ScopeOfWork struct {

	// The Amazon Standard Identification Number (ASIN) of the service job.
	Asin string `json:"asin,omitempty"`

	// The number of service jobs.
	Quantity int64 `json:"quantity,omitempty"`

	// A list of skills required to perform the job.
	RequiredSkills []string `json:"requiredSkills"`

	// The title of the service job.
	Title string `json:"title,omitempty"`
}

// Validate validates this scope of work
func (m *ScopeOfWork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this scope of work based on context it is used
func (m *ScopeOfWork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScopeOfWork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScopeOfWork) UnmarshalBinary(b []byte) error {
	var res ScopeOfWork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
