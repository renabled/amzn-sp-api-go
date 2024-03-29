// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Address The address.
//
// swagger:model Address
type Address struct {

	// The first line of the address.
	// Required: true
	// Max Length: 60
	// Min Length: 1
	AddressLine1 *string `json:"addressLine1"`

	// Additional address information, if required.
	// Max Length: 60
	// Min Length: 1
	AddressLine2 string `json:"addressLine2,omitempty"`

	// Additional address information, if required.
	// Max Length: 60
	// Min Length: 1
	AddressLine3 string `json:"addressLine3,omitempty"`

	// city
	// Required: true
	City *City `json:"city"`

	// The name of the business or institution associated with the address.
	CompanyName string `json:"companyName,omitempty"`

	// country code
	// Required: true
	CountryCode *CountryCode `json:"countryCode"`

	// The email address of the contact associated with the address.
	// Max Length: 64
	Email string `json:"email,omitempty"`

	// geocode
	Geocode *Geocode `json:"geocode,omitempty"`

	// The name of the person, business or institution at the address.
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// The phone number of the person, business or institution located at that address, including the country calling code.
	// Max Length: 20
	// Min Length: 1
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// postal code
	// Required: true
	PostalCode *PostalCode `json:"postalCode"`

	// state or region
	// Required: true
	StateOrRegion *StateOrRegion `json:"stateOrRegion"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressLine2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressLine3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeocode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
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

	if err := validate.MinLength("addressLine1", "body", *m.AddressLine1, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine1", "body", *m.AddressLine1, 60); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateAddressLine2(formats strfmt.Registry) error {
	if swag.IsZero(m.AddressLine2) { // not required
		return nil
	}

	if err := validate.MinLength("addressLine2", "body", m.AddressLine2, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine2", "body", m.AddressLine2, 60); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateAddressLine3(formats strfmt.Registry) error {
	if swag.IsZero(m.AddressLine3) { // not required
		return nil
	}

	if err := validate.MinLength("addressLine3", "body", m.AddressLine3, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("addressLine3", "body", m.AddressLine3, 60); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	if m.City != nil {
		if err := m.City.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("city")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("city")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if m.CountryCode != nil {
		if err := m.CountryCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("countryCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("countryCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.MaxLength("email", "body", m.Email, 64); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateGeocode(formats strfmt.Registry) error {
	if swag.IsZero(m.Geocode) { // not required
		return nil
	}

	if m.Geocode != nil {
		if err := m.Geocode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geocode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geocode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 50); err != nil {
		return err
	}

	return nil
}

func (m *Address) validatePhoneNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.PhoneNumber) { // not required
		return nil
	}

	if err := validate.MinLength("phoneNumber", "body", m.PhoneNumber, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("phoneNumber", "body", m.PhoneNumber, 20); err != nil {
		return err
	}

	return nil
}

func (m *Address) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("postalCode", "body", m.PostalCode); err != nil {
		return err
	}

	if err := validate.Required("postalCode", "body", m.PostalCode); err != nil {
		return err
	}

	if m.PostalCode != nil {
		if err := m.PostalCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postalCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postalCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateStateOrRegion(formats strfmt.Registry) error {

	if err := validate.Required("stateOrRegion", "body", m.StateOrRegion); err != nil {
		return err
	}

	if err := validate.Required("stateOrRegion", "body", m.StateOrRegion); err != nil {
		return err
	}

	if m.StateOrRegion != nil {
		if err := m.StateOrRegion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateOrRegion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stateOrRegion")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this address based on the context it is used
func (m *Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountryCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGeocode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostalCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateOrRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) contextValidateCity(ctx context.Context, formats strfmt.Registry) error {

	if m.City != nil {
		if err := m.City.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("city")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("city")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateCountryCode(ctx context.Context, formats strfmt.Registry) error {

	if m.CountryCode != nil {
		if err := m.CountryCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("countryCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("countryCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateGeocode(ctx context.Context, formats strfmt.Registry) error {

	if m.Geocode != nil {
		if err := m.Geocode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geocode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("geocode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidatePostalCode(ctx context.Context, formats strfmt.Registry) error {

	if m.PostalCode != nil {
		if err := m.PostalCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postalCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postalCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateStateOrRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.StateOrRegion != nil {
		if err := m.StateOrRegion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stateOrRegion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stateOrRegion")
			}
			return err
		}
	}

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
