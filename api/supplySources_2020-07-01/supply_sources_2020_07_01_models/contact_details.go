// Code generated by go-swagger; DO NOT EDIT.

package supply_sources_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactDetails The contact details
//
// swagger:model ContactDetails
type ContactDetails struct {

	// primary
	Primary *ContactDetailsPrimary `json:"primary,omitempty"`
}

// Validate validates this contact details
func (m *ContactDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactDetails) validatePrimary(formats strfmt.Registry) error {
	if swag.IsZero(m.Primary) { // not required
		return nil
	}

	if m.Primary != nil {
		if err := m.Primary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this contact details based on the context it is used
func (m *ContactDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrimary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactDetails) contextValidatePrimary(ctx context.Context, formats strfmt.Registry) error {

	if m.Primary != nil {
		if err := m.Primary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactDetails) UnmarshalBinary(b []byte) error {
	var res ContactDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ContactDetailsPrimary contact details primary
//
// swagger:model ContactDetailsPrimary
type ContactDetailsPrimary struct {

	// email
	Email EmailAddress `json:"email,omitempty"`

	// The phone number of the person, business or institution.
	Phone string `json:"phone,omitempty"`
}

// Validate validates this contact details primary
func (m *ContactDetailsPrimary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactDetailsPrimary) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := m.Email.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("primary" + "." + "email")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("primary" + "." + "email")
		}
		return err
	}

	return nil
}

// ContextValidate validate this contact details primary based on the context it is used
func (m *ContactDetailsPrimary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactDetailsPrimary) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Email.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("primary" + "." + "email")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("primary" + "." + "email")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactDetailsPrimary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactDetailsPrimary) UnmarshalBinary(b []byte) error {
	var res ContactDetailsPrimary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
