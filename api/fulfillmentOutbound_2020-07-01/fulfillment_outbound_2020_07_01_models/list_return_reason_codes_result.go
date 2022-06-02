// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListReturnReasonCodesResult list return reason codes result
//
// swagger:model ListReturnReasonCodesResult
type ListReturnReasonCodesResult struct {

	// reason code details
	ReasonCodeDetails ReasonCodeDetailsList `json:"reasonCodeDetails,omitempty"`
}

// Validate validates this list return reason codes result
func (m *ListReturnReasonCodesResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReasonCodeDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListReturnReasonCodesResult) validateReasonCodeDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ReasonCodeDetails) { // not required
		return nil
	}

	if err := m.ReasonCodeDetails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reasonCodeDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("reasonCodeDetails")
		}
		return err
	}

	return nil
}

// ContextValidate validate this list return reason codes result based on the context it is used
func (m *ListReturnReasonCodesResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReasonCodeDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListReturnReasonCodesResult) contextValidateReasonCodeDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ReasonCodeDetails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reasonCodeDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("reasonCodeDetails")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListReturnReasonCodesResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListReturnReasonCodesResult) UnmarshalBinary(b []byte) error {
	var res ListReturnReasonCodesResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
