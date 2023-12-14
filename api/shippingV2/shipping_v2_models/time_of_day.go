// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeOfDay Denotes time of the day, used for defining opening or closing time of access points
//
// swagger:model TimeOfDay
type TimeOfDay struct {

	// hour of day
	HourOfDay int64 `json:"hourOfDay,omitempty"`

	// minute of hour
	MinuteOfHour int64 `json:"minuteOfHour,omitempty"`

	// second of minute
	SecondOfMinute int64 `json:"secondOfMinute,omitempty"`
}

// Validate validates this time of day
func (m *TimeOfDay) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this time of day based on context it is used
func (m *TimeOfDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeOfDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeOfDay) UnmarshalBinary(b []byte) error {
	var res TimeOfDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
