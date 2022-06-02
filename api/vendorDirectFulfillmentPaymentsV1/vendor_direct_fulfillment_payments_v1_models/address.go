// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_payments_v1_models

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

	// Additional street address information, if required.
	AddressLine2 string `json:"addressLine2,omitempty"`

	// Additional street address information, if required.
	AddressLine3 string `json:"addressLine3,omitempty"`

	// The city where the person, business or institution is located.
	// Required: true
	City *string `json:"city"`

	// The two digit country code in ISO 3166-1 alpha-2 format.
	// Required: true
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

	// The postal code of that address. It conatins a series of letters or digits or both, sometimes including spaces or punctuation.
	// Required: true
	PostalCode *string `json:"postalCode"`

	// The state or region where person, business or institution is located.
	// Required: true
	StateOrRegion *string `json:"stateOrRegion"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateOrRegion(formats); err != nil {
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

func (m *Address) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
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

func (m *Address) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("postalCode", "body", m.PostalCode); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateStateOrRegion(formats strfmt.Registry) error {

	if err := validate.Required("stateOrRegion", "body", m.StateOrRegion); err != nil {
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
