// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GenerateSelfShipAppointmentSlotsRequest The `generateSelfShipAppointmentSlots` request.
// Example: {"desiredEndDate":"2024-01-06T14:48:00.000Z","desiredStartDate":"2024-01-05T14:48:00.000Z"}
//
// swagger:model GenerateSelfShipAppointmentSlotsRequest
type GenerateSelfShipAppointmentSlotsRequest struct {

	// The desired end date. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
	// Format: date-time
	DesiredEndDate strfmt.DateTime `json:"desiredEndDate,omitempty"`

	// The desired start date. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format.
	// Format: date-time
	DesiredStartDate strfmt.DateTime `json:"desiredStartDate,omitempty"`
}

// Validate validates this generate self ship appointment slots request
func (m *GenerateSelfShipAppointmentSlotsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesiredEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateSelfShipAppointmentSlotsRequest) validateDesiredEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DesiredEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("desiredEndDate", "body", "date-time", m.DesiredEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GenerateSelfShipAppointmentSlotsRequest) validateDesiredStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DesiredStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("desiredStartDate", "body", "date-time", m.DesiredStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this generate self ship appointment slots request based on context it is used
func (m *GenerateSelfShipAppointmentSlotsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenerateSelfShipAppointmentSlotsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateSelfShipAppointmentSlotsRequest) UnmarshalBinary(b []byte) error {
	var res GenerateSelfShipAppointmentSlotsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
