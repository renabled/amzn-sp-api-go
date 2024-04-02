// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inboundv2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SelfShipAppointmentDetails Appointment details for carrier pickup or fulfillment center appointments.
//
// swagger:model SelfShipAppointmentDetails
type SelfShipAppointmentDetails struct {

	// Identifier for appointment.
	AppointmentID float64 `json:"appointmentId,omitempty"`

	// appointment slot time
	AppointmentSlotTime *AppointmentSlotTime `json:"appointmentSlotTime,omitempty"`

	// Status of the appointment.
	// Max Length: 1024
	// Min Length: 1
	AppointmentStatus string `json:"appointmentStatus,omitempty"`
}

// Validate validates this self ship appointment details
func (m *SelfShipAppointmentDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointmentSlotTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppointmentStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SelfShipAppointmentDetails) validateAppointmentSlotTime(formats strfmt.Registry) error {
	if swag.IsZero(m.AppointmentSlotTime) { // not required
		return nil
	}

	if m.AppointmentSlotTime != nil {
		if err := m.AppointmentSlotTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appointmentSlotTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appointmentSlotTime")
			}
			return err
		}
	}

	return nil
}

func (m *SelfShipAppointmentDetails) validateAppointmentStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.AppointmentStatus) { // not required
		return nil
	}

	if err := validate.MinLength("appointmentStatus", "body", m.AppointmentStatus, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("appointmentStatus", "body", m.AppointmentStatus, 1024); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this self ship appointment details based on the context it is used
func (m *SelfShipAppointmentDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppointmentSlotTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SelfShipAppointmentDetails) contextValidateAppointmentSlotTime(ctx context.Context, formats strfmt.Registry) error {

	if m.AppointmentSlotTime != nil {
		if err := m.AppointmentSlotTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appointmentSlotTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appointmentSlotTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SelfShipAppointmentDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SelfShipAppointmentDetails) UnmarshalBinary(b []byte) error {
	var res SelfShipAppointmentDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
