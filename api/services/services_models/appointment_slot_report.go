// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppointmentSlotReport Availability information as per the service context queried.
//
// swagger:model AppointmentSlotReport
type AppointmentSlotReport struct {

	// A list of time windows along with associated capacity in which the service can be performed.
	AppointmentSlots []*AppointmentSlot `json:"appointmentSlots"`

	// End Time up to which the appointment slots are generated in ISO 8601 format.
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// Defines the type of slots.
	// Enum: [REAL_TIME_SCHEDULING NON_REAL_TIME_SCHEDULING]
	SchedulingType string `json:"schedulingType,omitempty"`

	// Start Time from which the appointment slots are generated in ISO 8601 format.
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`
}

// Validate validates this appointment slot report
func (m *AppointmentSlotReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointmentSlots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedulingType(formats); err != nil {
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

func (m *AppointmentSlotReport) validateAppointmentSlots(formats strfmt.Registry) error {
	if swag.IsZero(m.AppointmentSlots) { // not required
		return nil
	}

	for i := 0; i < len(m.AppointmentSlots); i++ {
		if swag.IsZero(m.AppointmentSlots[i]) { // not required
			continue
		}

		if m.AppointmentSlots[i] != nil {
			if err := m.AppointmentSlots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appointmentSlots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appointmentSlots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppointmentSlotReport) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var appointmentSlotReportTypeSchedulingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REAL_TIME_SCHEDULING","NON_REAL_TIME_SCHEDULING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appointmentSlotReportTypeSchedulingTypePropEnum = append(appointmentSlotReportTypeSchedulingTypePropEnum, v)
	}
}

const (

	// AppointmentSlotReportSchedulingTypeREALTIMESCHEDULING captures enum value "REAL_TIME_SCHEDULING"
	AppointmentSlotReportSchedulingTypeREALTIMESCHEDULING string = "REAL_TIME_SCHEDULING"

	// AppointmentSlotReportSchedulingTypeNONREALTIMESCHEDULING captures enum value "NON_REAL_TIME_SCHEDULING"
	AppointmentSlotReportSchedulingTypeNONREALTIMESCHEDULING string = "NON_REAL_TIME_SCHEDULING"
)

// prop value enum
func (m *AppointmentSlotReport) validateSchedulingTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, appointmentSlotReportTypeSchedulingTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AppointmentSlotReport) validateSchedulingType(formats strfmt.Registry) error {
	if swag.IsZero(m.SchedulingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSchedulingTypeEnum("schedulingType", "body", m.SchedulingType); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentSlotReport) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this appointment slot report based on the context it is used
func (m *AppointmentSlotReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppointmentSlots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppointmentSlotReport) contextValidateAppointmentSlots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppointmentSlots); i++ {

		if m.AppointmentSlots[i] != nil {
			if err := m.AppointmentSlots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appointmentSlots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appointmentSlots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppointmentSlotReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppointmentSlotReport) UnmarshalBinary(b []byte) error {
	var res AppointmentSlotReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
