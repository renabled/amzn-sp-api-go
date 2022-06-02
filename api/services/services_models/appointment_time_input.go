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

// AppointmentTimeInput The input appointment time details.
//
// swagger:model AppointmentTimeInput
type AppointmentTimeInput struct {

	// The duration of an appointment in minutes.
	DurationInMinutes int64 `json:"durationInMinutes,omitempty"`

	// The date, time in UTC for the start time of an appointment in ISO 8601 format.
	// Required: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"startTime"`
}

// Validate validates this appointment time input
func (m *AppointmentTimeInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppointmentTimeInput) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this appointment time input based on context it is used
func (m *AppointmentTimeInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppointmentTimeInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppointmentTimeInput) UnmarshalBinary(b []byte) error {
	var res AppointmentTimeInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
