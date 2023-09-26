// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubstitutionPreferences substitution preferences
//
// swagger:model SubstitutionPreferences
type SubstitutionPreferences struct {

	// Substitution options for the order item.
	SubstitutionOptions SubstitutionOptionList `json:"SubstitutionOptions,omitempty"`

	// The type of substitution that these preferences represent.
	// Required: true
	// Enum: [CUSTOMER_PREFERENCE AMAZON_RECOMMENDED DO_NOT_SUBSTITUTE]
	SubstitutionType *string `json:"SubstitutionType"`
}

// Validate validates this substitution preferences
func (m *SubstitutionPreferences) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubstitutionOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubstitutionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubstitutionPreferences) validateSubstitutionOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.SubstitutionOptions) { // not required
		return nil
	}

	if err := m.SubstitutionOptions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SubstitutionOptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SubstitutionOptions")
		}
		return err
	}

	return nil
}

var substitutionPreferencesTypeSubstitutionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOMER_PREFERENCE","AMAZON_RECOMMENDED","DO_NOT_SUBSTITUTE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		substitutionPreferencesTypeSubstitutionTypePropEnum = append(substitutionPreferencesTypeSubstitutionTypePropEnum, v)
	}
}

const (

	// SubstitutionPreferencesSubstitutionTypeCUSTOMERPREFERENCE captures enum value "CUSTOMER_PREFERENCE"
	SubstitutionPreferencesSubstitutionTypeCUSTOMERPREFERENCE string = "CUSTOMER_PREFERENCE"

	// SubstitutionPreferencesSubstitutionTypeAMAZONRECOMMENDED captures enum value "AMAZON_RECOMMENDED"
	SubstitutionPreferencesSubstitutionTypeAMAZONRECOMMENDED string = "AMAZON_RECOMMENDED"

	// SubstitutionPreferencesSubstitutionTypeDONOTSUBSTITUTE captures enum value "DO_NOT_SUBSTITUTE"
	SubstitutionPreferencesSubstitutionTypeDONOTSUBSTITUTE string = "DO_NOT_SUBSTITUTE"
)

// prop value enum
func (m *SubstitutionPreferences) validateSubstitutionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, substitutionPreferencesTypeSubstitutionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SubstitutionPreferences) validateSubstitutionType(formats strfmt.Registry) error {

	if err := validate.Required("SubstitutionType", "body", m.SubstitutionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateSubstitutionTypeEnum("SubstitutionType", "body", *m.SubstitutionType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this substitution preferences based on the context it is used
func (m *SubstitutionPreferences) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubstitutionOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubstitutionPreferences) contextValidateSubstitutionOptions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SubstitutionOptions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SubstitutionOptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SubstitutionOptions")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubstitutionPreferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubstitutionPreferences) UnmarshalBinary(b []byte) error {
	var res SubstitutionPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
