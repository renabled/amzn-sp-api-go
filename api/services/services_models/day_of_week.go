// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DayOfWeek The day of the week.
//
// swagger:model DayOfWeek
type DayOfWeek string

func NewDayOfWeek(value DayOfWeek) *DayOfWeek {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DayOfWeek.
func (m DayOfWeek) Pointer() *DayOfWeek {
	return &m
}

const (

	// DayOfWeekMONDAY captures enum value "MONDAY"
	DayOfWeekMONDAY DayOfWeek = "MONDAY"

	// DayOfWeekTUESDAY captures enum value "TUESDAY"
	DayOfWeekTUESDAY DayOfWeek = "TUESDAY"

	// DayOfWeekWEDNESDAY captures enum value "WEDNESDAY"
	DayOfWeekWEDNESDAY DayOfWeek = "WEDNESDAY"

	// DayOfWeekTHURSDAY captures enum value "THURSDAY"
	DayOfWeekTHURSDAY DayOfWeek = "THURSDAY"

	// DayOfWeekFRIDAY captures enum value "FRIDAY"
	DayOfWeekFRIDAY DayOfWeek = "FRIDAY"

	// DayOfWeekSATURDAY captures enum value "SATURDAY"
	DayOfWeekSATURDAY DayOfWeek = "SATURDAY"

	// DayOfWeekSUNDAY captures enum value "SUNDAY"
	DayOfWeekSUNDAY DayOfWeek = "SUNDAY"
)

// for schema
var dayOfWeekEnum []interface{}

func init() {
	var res []DayOfWeek
	if err := json.Unmarshal([]byte(`["MONDAY","TUESDAY","WEDNESDAY","THURSDAY","FRIDAY","SATURDAY","SUNDAY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dayOfWeekEnum = append(dayOfWeekEnum, v)
	}
}

func (m DayOfWeek) validateDayOfWeekEnum(path, location string, value DayOfWeek) error {
	if err := validate.EnumCase(path, location, value, dayOfWeekEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this day of week
func (m DayOfWeek) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDayOfWeekEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this day of week based on context it is used
func (m DayOfWeek) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
