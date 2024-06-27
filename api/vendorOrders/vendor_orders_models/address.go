// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Address Address of the party.
//
// swagger:model Address
type Address struct {

	// First line of the address.
	// Required: true
	AddressLine1 *string `json:"addressLine1"`

	// Additional address information, if required.
	AddressLine2 string `json:"addressLine2,omitempty"`

	// Additional address information, if required.
	AddressLine3 string `json:"addressLine3,omitempty"`

	// The city where the person, business or institution is located.
	City string `json:"city,omitempty"`

	// The two digit country code. In ISO 3166-1 alpha-2 format.
	// Required: true
	// Max Length: 2
	CountryCode *string `json:"countryCode"`

	// The county where person, business or institution is located.
	County string `json:"county,omitempty"`

	// The district where person, business or institution is located.
	District string `json:"district,omitempty"`

	// The name of the person, business or institution at that address.
	// Required: true
	Name *string `json:"name"`

	// The phone number of the person, business or institution located at that address.
	Phone string `json:"phone,omitempty"`

	// The postal code of that address. It contains a series of letters or digits or both, sometimes including spaces or punctuation.
	PostalCode string `json:"postalCode,omitempty"`

	// The state or region where person, business or institution is located.
	StateOrRegion string `json:"stateOrRegion,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("addressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if err := validate.MaxLength("countryCode", "body", *m.CountryCode, 2); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this address based on context it is used
func (m *Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
