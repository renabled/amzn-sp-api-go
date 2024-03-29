// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Reservation Reservation object reduces the capacity of a resource.
//
// swagger:model Reservation
type Reservation struct {

	// `AvailabilityRecord` to represent the capacity of a resource over a time range.
	// Required: true
	Availability *AvailabilityRecord `json:"availability"`

	// Unique identifier for a reservation. If present, it is treated as an update reservation request and will update the corresponding reservation. Otherwise, it is treated as a new create reservation request.
	ReservationID string `json:"reservationId,omitempty"`

	// Type of reservation.
	// Required: true
	// Enum: [APPOINTMENT TRAVEL VACATION BREAK TRAINING]
	Type *string `json:"type"`
}

// Validate validates this reservation
func (m *Reservation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Reservation) validateAvailability(formats strfmt.Registry) error {

	if err := validate.Required("availability", "body", m.Availability); err != nil {
		return err
	}

	if m.Availability != nil {
		if err := m.Availability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("availability")
			}
			return err
		}
	}

	return nil
}

var reservationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APPOINTMENT","TRAVEL","VACATION","BREAK","TRAINING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reservationTypeTypePropEnum = append(reservationTypeTypePropEnum, v)
	}
}

const (

	// ReservationTypeAPPOINTMENT captures enum value "APPOINTMENT"
	ReservationTypeAPPOINTMENT string = "APPOINTMENT"

	// ReservationTypeTRAVEL captures enum value "TRAVEL"
	ReservationTypeTRAVEL string = "TRAVEL"

	// ReservationTypeVACATION captures enum value "VACATION"
	ReservationTypeVACATION string = "VACATION"

	// ReservationTypeBREAK captures enum value "BREAK"
	ReservationTypeBREAK string = "BREAK"

	// ReservationTypeTRAINING captures enum value "TRAINING"
	ReservationTypeTRAINING string = "TRAINING"
)

// prop value enum
func (m *Reservation) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reservationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Reservation) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this reservation based on the context it is used
func (m *Reservation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Reservation) contextValidateAvailability(ctx context.Context, formats strfmt.Registry) error {

	if m.Availability != nil {
		if err := m.Availability.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availability")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("availability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Reservation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reservation) UnmarshalBinary(b []byte) error {
	var res Reservation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
