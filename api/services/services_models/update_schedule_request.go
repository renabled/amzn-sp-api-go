// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateScheduleRequest Request schema for the `updateSchedule` operation.
//
// swagger:model UpdateScheduleRequest
type UpdateScheduleRequest struct {

	// List of schedule objects to define the normal working hours of a resource.
	// Required: true
	Schedules AvailabilityRecords `json:"schedules"`
}

// Validate validates this update schedule request
func (m *UpdateScheduleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateScheduleRequest) validateSchedules(formats strfmt.Registry) error {

	if err := validate.Required("schedules", "body", m.Schedules); err != nil {
		return err
	}

	if err := m.Schedules.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("schedules")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("schedules")
		}
		return err
	}

	return nil
}

// ContextValidate validate this update schedule request based on the context it is used
func (m *UpdateScheduleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateScheduleRequest) contextValidateSchedules(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Schedules.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("schedules")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("schedules")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateScheduleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateScheduleRequest) UnmarshalBinary(b []byte) error {
	var res UpdateScheduleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
