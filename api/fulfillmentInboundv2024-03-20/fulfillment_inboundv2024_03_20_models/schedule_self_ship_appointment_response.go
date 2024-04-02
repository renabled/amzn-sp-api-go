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

// ScheduleSelfShipAppointmentResponse `scheduleSelfShipAppointment` response.
// Example: {"selfShipAppointmentDetails":{"appointmentId":1000,"appointmentSlotTime":{"endTime":"2023-03-09T13:15:30Z","startTime":"2023-03-08T13:15:30Z"},"appointmentStatus":"ARRIVAL_SCHEDULED"}}
//
// swagger:model ScheduleSelfShipAppointmentResponse
type ScheduleSelfShipAppointmentResponse struct {

	// self ship appointment details
	// Required: true
	SelfShipAppointmentDetails *SelfShipAppointmentDetails `json:"selfShipAppointmentDetails"`
}

// Validate validates this schedule self ship appointment response
func (m *ScheduleSelfShipAppointmentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfShipAppointmentDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleSelfShipAppointmentResponse) validateSelfShipAppointmentDetails(formats strfmt.Registry) error {

	if err := validate.Required("selfShipAppointmentDetails", "body", m.SelfShipAppointmentDetails); err != nil {
		return err
	}

	if m.SelfShipAppointmentDetails != nil {
		if err := m.SelfShipAppointmentDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selfShipAppointmentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selfShipAppointmentDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this schedule self ship appointment response based on the context it is used
func (m *ScheduleSelfShipAppointmentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelfShipAppointmentDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleSelfShipAppointmentResponse) contextValidateSelfShipAppointmentDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.SelfShipAppointmentDetails != nil {
		if err := m.SelfShipAppointmentDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selfShipAppointmentDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selfShipAppointmentDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleSelfShipAppointmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleSelfShipAppointmentResponse) UnmarshalBinary(b []byte) error {
	var res ScheduleSelfShipAppointmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
