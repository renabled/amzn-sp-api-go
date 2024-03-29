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

// SupplySourceConfiguration Includes configuration and timezone of a supply source.
//
// swagger:model SupplySourceConfiguration
type SupplySourceConfiguration struct {

	// operational configuration
	OperationalConfiguration *OperationalConfiguration `json:"operationalConfiguration,omitempty"`

	// Please see RFC 6557, should be a canonical time zone ID as listed here: https://www.joda.org/joda-time/timezones.html.
	Timezone string `json:"timezone,omitempty"`
}

// Validate validates this supply source configuration
func (m *SupplySourceConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperationalConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupplySourceConfiguration) validateOperationalConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.OperationalConfiguration) { // not required
		return nil
	}

	if m.OperationalConfiguration != nil {
		if err := m.OperationalConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationalConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationalConfiguration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this supply source configuration based on the context it is used
func (m *SupplySourceConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperationalConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupplySourceConfiguration) contextValidateOperationalConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.OperationalConfiguration != nil {
		if err := m.OperationalConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationalConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationalConfiguration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupplySourceConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupplySourceConfiguration) UnmarshalBinary(b []byte) error {
	var res SupplySourceConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
