// Code generated by go-swagger; DO NOT EDIT.

package notifications_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateDestinationRequest The request schema for the `createDestination` operation.
//
// swagger:model CreateDestinationRequest
type CreateDestinationRequest struct {

	// A developer-defined name to help identify this destination.
	// Required: true
	Name *string `json:"name"`

	// The information required to create a destination resource. Applications should use one resource type (sqs or eventBridge) per destination.
	// Required: true
	ResourceSpecification *DestinationResourceSpecification `json:"resourceSpecification"`
}

// Validate validates this create destination request
func (m *CreateDestinationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceSpecification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDestinationRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateDestinationRequest) validateResourceSpecification(formats strfmt.Registry) error {

	if err := validate.Required("resourceSpecification", "body", m.ResourceSpecification); err != nil {
		return err
	}

	if m.ResourceSpecification != nil {
		if err := m.ResourceSpecification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceSpecification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create destination request based on the context it is used
func (m *CreateDestinationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceSpecification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDestinationRequest) contextValidateResourceSpecification(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceSpecification != nil {
		if err := m.ResourceSpecification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceSpecification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateDestinationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDestinationRequest) UnmarshalBinary(b []byte) error {
	var res CreateDestinationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
