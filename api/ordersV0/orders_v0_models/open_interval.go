// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenInterval The time interval for which the business is open.
//
// swagger:model OpenInterval
type OpenInterval struct {

	// The time when the business closes.
	EndTime *OpenTimeInterval `json:"EndTime,omitempty"`

	// The time when the business opens.
	StartTime *OpenTimeInterval `json:"StartTime,omitempty"`
}

// Validate validates this open interval
func (m *OpenInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenInterval) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if m.EndTime != nil {
		if err := m.EndTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EndTime")
			}
			return err
		}
	}

	return nil
}

func (m *OpenInterval) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if m.StartTime != nil {
		if err := m.StartTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StartTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StartTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this open interval based on the context it is used
func (m *OpenInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenInterval) contextValidateEndTime(ctx context.Context, formats strfmt.Registry) error {

	if m.EndTime != nil {
		if err := m.EndTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EndTime")
			}
			return err
		}
	}

	return nil
}

func (m *OpenInterval) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if m.StartTime != nil {
		if err := m.StartTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StartTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StartTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenInterval) UnmarshalBinary(b []byte) error {
	var res OpenInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}