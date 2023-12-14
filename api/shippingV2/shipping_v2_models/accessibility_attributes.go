// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccessibilityAttributes Defines the accessibility details of the access point.
//
// swagger:model AccessibilityAttributes
type AccessibilityAttributes struct {

	// The approximate distance of access point from input postalCode's centroid.
	Distance string `json:"distance,omitempty"`

	// The approximate (static) drive time from input postal code's centroid.
	DriveTime int64 `json:"driveTime,omitempty"`
}

// Validate validates this accessibility attributes
func (m *AccessibilityAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this accessibility attributes based on context it is used
func (m *AccessibilityAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessibilityAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessibilityAttributes) UnmarshalBinary(b []byte) error {
	var res AccessibilityAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
