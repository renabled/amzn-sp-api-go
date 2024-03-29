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

// ServicesCapability The services capability of a supply source.
//
// swagger:model ServicesCapability
type ServicesCapability struct {

	// When true, `SupplySource` supports the Service capability.
	IsSupported bool `json:"isSupported,omitempty"`

	// operational configuration
	OperationalConfiguration *OperationalConfiguration `json:"operationalConfiguration,omitempty"`
}

// Validate validates this services capability
func (m *ServicesCapability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperationalConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicesCapability) validateOperationalConfiguration(formats strfmt.Registry) error {
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

// ContextValidate validate this services capability based on the context it is used
func (m *ServicesCapability) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperationalConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicesCapability) contextValidateOperationalConfiguration(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ServicesCapability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicesCapability) UnmarshalBinary(b []byte) error {
	var res ServicesCapability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
