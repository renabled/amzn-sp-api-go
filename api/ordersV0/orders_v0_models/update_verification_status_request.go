// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateVerificationStatusRequest The request body for the `updateVerificationStatus` operation.
//
// swagger:model UpdateVerificationStatusRequest
type UpdateVerificationStatusRequest struct {

	// The updated values of the `VerificationStatus` field.
	// Required: true
	RegulatedOrderVerificationStatus *UpdateVerificationStatusRequestBody `json:"regulatedOrderVerificationStatus"`
}

// Validate validates this update verification status request
func (m *UpdateVerificationStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegulatedOrderVerificationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateVerificationStatusRequest) validateRegulatedOrderVerificationStatus(formats strfmt.Registry) error {

	if err := validate.Required("regulatedOrderVerificationStatus", "body", m.RegulatedOrderVerificationStatus); err != nil {
		return err
	}

	if m.RegulatedOrderVerificationStatus != nil {
		if err := m.RegulatedOrderVerificationStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regulatedOrderVerificationStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("regulatedOrderVerificationStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update verification status request based on the context it is used
func (m *UpdateVerificationStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegulatedOrderVerificationStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateVerificationStatusRequest) contextValidateRegulatedOrderVerificationStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.RegulatedOrderVerificationStatus != nil {
		if err := m.RegulatedOrderVerificationStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("regulatedOrderVerificationStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("regulatedOrderVerificationStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateVerificationStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateVerificationStatusRequest) UnmarshalBinary(b []byte) error {
	var res UpdateVerificationStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
