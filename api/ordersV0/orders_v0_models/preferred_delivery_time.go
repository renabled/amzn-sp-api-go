// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PreferredDeliveryTime The time window when the delivery is preferred.
//
// swagger:model PreferredDeliveryTime
type PreferredDeliveryTime struct {

	// Business hours when the business is open for deliveries.
	BusinessHours []*BusinessHours `json:"BusinessHours"`

	// Dates when the business is closed in the next 30 days.
	ExceptionDates []*ExceptionDates `json:"ExceptionDates"`
}

// Validate validates this preferred delivery time
func (m *PreferredDeliveryTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessHours(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExceptionDates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreferredDeliveryTime) validateBusinessHours(formats strfmt.Registry) error {
	if swag.IsZero(m.BusinessHours) { // not required
		return nil
	}

	for i := 0; i < len(m.BusinessHours); i++ {
		if swag.IsZero(m.BusinessHours[i]) { // not required
			continue
		}

		if m.BusinessHours[i] != nil {
			if err := m.BusinessHours[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BusinessHours" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BusinessHours" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PreferredDeliveryTime) validateExceptionDates(formats strfmt.Registry) error {
	if swag.IsZero(m.ExceptionDates) { // not required
		return nil
	}

	for i := 0; i < len(m.ExceptionDates); i++ {
		if swag.IsZero(m.ExceptionDates[i]) { // not required
			continue
		}

		if m.ExceptionDates[i] != nil {
			if err := m.ExceptionDates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExceptionDates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExceptionDates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this preferred delivery time based on the context it is used
func (m *PreferredDeliveryTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBusinessHours(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExceptionDates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreferredDeliveryTime) contextValidateBusinessHours(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BusinessHours); i++ {

		if m.BusinessHours[i] != nil {
			if err := m.BusinessHours[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BusinessHours" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BusinessHours" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PreferredDeliveryTime) contextValidateExceptionDates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExceptionDates); i++ {

		if m.ExceptionDates[i] != nil {
			if err := m.ExceptionDates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ExceptionDates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ExceptionDates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreferredDeliveryTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreferredDeliveryTime) UnmarshalBinary(b []byte) error {
	var res PreferredDeliveryTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
