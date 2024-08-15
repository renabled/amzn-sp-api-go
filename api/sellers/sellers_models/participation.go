// Code generated by go-swagger; DO NOT EDIT.

package sellers_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Participation Information that is specific to a seller in a marketplace.
//
// swagger:model Participation
type Participation struct {

	// Specifies if the seller has suspended listings. `true` if the seller Listing Status is set to Inactive, otherwise `false`.
	// Required: true
	HasSuspendedListings *bool `json:"hasSuspendedListings"`

	// If `true`, the seller participates in the marketplace. Otherwise `false`.
	// Required: true
	IsParticipating *bool `json:"isParticipating"`
}

// Validate validates this participation
func (m *Participation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasSuspendedListings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsParticipating(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Participation) validateHasSuspendedListings(formats strfmt.Registry) error {

	if err := validate.Required("hasSuspendedListings", "body", m.HasSuspendedListings); err != nil {
		return err
	}

	return nil
}

func (m *Participation) validateIsParticipating(formats strfmt.Registry) error {

	if err := validate.Required("isParticipating", "body", m.IsParticipating); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this participation based on context it is used
func (m *Participation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Participation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Participation) UnmarshalBinary(b []byte) error {
	var res Participation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
