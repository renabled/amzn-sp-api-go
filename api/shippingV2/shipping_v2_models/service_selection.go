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

// ServiceSelection Service Selection Criteria.
//
// swagger:model ServiceSelection
type ServiceSelection struct {

	// service Id
	// Required: true
	ServiceID ServiceIds `json:"serviceId"`
}

// Validate validates this service selection
func (m *ServiceSelection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceSelection) validateServiceID(formats strfmt.Registry) error {

	if err := validate.Required("serviceId", "body", m.ServiceID); err != nil {
		return err
	}

	if err := m.ServiceID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceId")
		}
		return err
	}

	return nil
}

// ContextValidate validate this service selection based on the context it is used
func (m *ServiceSelection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServiceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceSelection) contextValidateServiceID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServiceID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("serviceId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("serviceId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSelection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSelection) UnmarshalBinary(b []byte) error {
	var res ServiceSelection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
