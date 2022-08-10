// Code generated by go-swagger; DO NOT EDIT.

package services_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RangeSlotCapacityQuery Request schema for the `getRangeSlotCapacity` operation. This schema is used to define the time range and capacity types that are being queried.
//
// swagger:model RangeSlotCapacityQuery
type RangeSlotCapacityQuery struct {

	// An array of capacity types which are being requested. Default value is `[SCHEDULED_CAPACITY]`.
	CapacityTypes []CapacityType `json:"capacityTypes"`

	// End date time up to which the capacity slots are being requested in ISO 8601 format.
	// Required: true
	// Format: date-time
	EndDateTime *strfmt.DateTime `json:"endDateTime"`

	// Start date time from which the capacity slots are being requested in ISO 8601 format.
	// Required: true
	// Format: date-time
	StartDateTime *strfmt.DateTime `json:"startDateTime"`
}

// Validate validates this range slot capacity query
func (m *RangeSlotCapacityQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RangeSlotCapacityQuery) validateCapacityTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.CapacityTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.CapacityTypes); i++ {

		if err := m.CapacityTypes[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacityTypes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacityTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *RangeSlotCapacityQuery) validateEndDateTime(formats strfmt.Registry) error {

	if err := validate.Required("endDateTime", "body", m.EndDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("endDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RangeSlotCapacityQuery) validateStartDateTime(formats strfmt.Registry) error {

	if err := validate.Required("startDateTime", "body", m.StartDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("startDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this range slot capacity query based on the context it is used
func (m *RangeSlotCapacityQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacityTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RangeSlotCapacityQuery) contextValidateCapacityTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CapacityTypes); i++ {

		if err := m.CapacityTypes[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacityTypes" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacityTypes" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RangeSlotCapacityQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RangeSlotCapacityQuery) UnmarshalBinary(b []byte) error {
	var res RangeSlotCapacityQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
