// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Address The postal address information.
//
// swagger:model Address
type Address struct {

	// address line1
	// Required: true
	AddressLine1 *AddressLine1 `json:"AddressLine1"`

	// address line2
	AddressLine2 AddressLine2 `json:"AddressLine2,omitempty"`

	// address line3
	AddressLine3 AddressLine3 `json:"AddressLine3,omitempty"`

	// city
	// Required: true
	City *City `json:"City"`

	// country code
	// Required: true
	CountryCode *CountryCode `json:"CountryCode"`

	// district or county
	DistrictOrCounty DistrictOrCounty `json:"DistrictOrCounty,omitempty"`

	// email
	// Required: true
	Email *EmailAddress `json:"Email"`

	// name
	// Required: true
	Name *AddressName `json:"Name"`

	// phone
	// Required: true
	Phone *PhoneNumber `json:"Phone"`

	// postal code
	// Required: true
	PostalCode *PostalCode `json:"PostalCode"`

	// state or province code
	StateOrProvinceCode StateOrProvinceCode `json:"StateOrProvinceCode,omitempty"`
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

	if err := m.validateDistrictOrCounty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateOrProvinceCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("AddressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	if err := validate.Required("AddressLine1", "body", m.AddressLine1); err != nil {
		return err
	}

	if m.AddressLine1 != nil {
		if err := m.AddressLine1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AddressLine1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AddressLine1")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateAddressLine2(formats strfmt.Registry) error {
	if swag.IsZero(m.AddressLine2) { // not required
		return nil
	}

	if err := m.AddressLine2.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AddressLine2")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AddressLine2")
		}
		return err
	}

	return nil
}

func (m *Address) validateAddressLine3(formats strfmt.Registry) error {
	if swag.IsZero(m.AddressLine3) { // not required
		return nil
	}

	if err := m.AddressLine3.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AddressLine3")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AddressLine3")
		}
		return err
	}

	return nil
}

func (m *Address) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("City", "body", m.City); err != nil {
		return err
	}

	if err := validate.Required("City", "body", m.City); err != nil {
		return err
	}

	if m.City != nil {
		if err := m.City.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("City")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("City")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("CountryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if err := validate.Required("CountryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if m.CountryCode != nil {
		if err := m.CountryCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CountryCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CountryCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateDistrictOrCounty(formats strfmt.Registry) error {
	if swag.IsZero(m.DistrictOrCounty) { // not required
		return nil
	}

	if err := m.DistrictOrCounty.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DistrictOrCounty")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("DistrictOrCounty")
		}
		return err
	}

	return nil
}

func (m *Address) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("Email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.Required("Email", "body", m.Email); err != nil {
		return err
	}

	if m.Email != nil {
		if err := m.Email.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Email")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Email")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Name")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("Phone", "body", m.Phone); err != nil {
		return err
	}

	if err := validate.Required("Phone", "body", m.Phone); err != nil {
		return err
	}

	if m.Phone != nil {
		if err := m.Phone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Phone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Phone")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("PostalCode", "body", m.PostalCode); err != nil {
		return err
	}

	if err := validate.Required("PostalCode", "body", m.PostalCode); err != nil {
		return err
	}

	if m.PostalCode != nil {
		if err := m.PostalCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PostalCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PostalCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateStateOrProvinceCode(formats strfmt.Registry) error {
	if swag.IsZero(m.StateOrProvinceCode) { // not required
		return nil
	}

	if err := m.StateOrProvinceCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StateOrProvinceCode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StateOrProvinceCode")
		}
		return err
	}

	return nil
}

// ContextValidate validate this address based on the context it is used
func (m *Address) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddressLine1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddressLine2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddressLine3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountryCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDistrictOrCounty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostalCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateOrProvinceCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) contextValidateAddressLine1(ctx context.Context, formats strfmt.Registry) error {

	if m.AddressLine1 != nil {
		if err := m.AddressLine1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AddressLine1")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("AddressLine1")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateAddressLine2(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AddressLine2.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AddressLine2")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AddressLine2")
		}
		return err
	}

	return nil
}

func (m *Address) contextValidateAddressLine3(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AddressLine3.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AddressLine3")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AddressLine3")
		}
		return err
	}

	return nil
}

func (m *Address) contextValidateCity(ctx context.Context, formats strfmt.Registry) error {

	if m.City != nil {
		if err := m.City.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("City")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("City")
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
				return ve.ValidateName("CountryCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CountryCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateDistrictOrCounty(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DistrictOrCounty.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DistrictOrCounty")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("DistrictOrCounty")
		}
		return err
	}

	return nil
}

func (m *Address) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	if m.Email != nil {
		if err := m.Email.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Email")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Email")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {
		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Name")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidatePhone(ctx context.Context, formats strfmt.Registry) error {

	if m.Phone != nil {
		if err := m.Phone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Phone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Phone")
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
				return ve.ValidateName("PostalCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("PostalCode")
			}
			return err
		}
	}

	return nil
}

func (m *Address) contextValidateStateOrProvinceCode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StateOrProvinceCode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StateOrProvinceCode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("StateOrProvinceCode")
		}
		return err
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
