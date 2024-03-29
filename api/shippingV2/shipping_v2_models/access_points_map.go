// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AccessPointsMap Map of type of access point to list of access points
//
// swagger:model AccessPointsMap
type AccessPointsMap map[string]AccessPointList

// Validate validates this access points map
func (m AccessPointsMap) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := validate.Required(k, "body", m[k]); err != nil {
			return err
		}

		if err := m[k].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(k)
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this access points map based on the context it is used
func (m AccessPointsMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := m[k].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(k)
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName(k)
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
