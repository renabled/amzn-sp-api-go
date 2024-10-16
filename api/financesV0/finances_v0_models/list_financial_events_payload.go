// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListFinancialEventsPayload The payload for the `listFinancialEvents` operation.
//
// swagger:model ListFinancialEventsPayload
type ListFinancialEventsPayload struct {

	// financial events
	FinancialEvents *FinancialEvents `json:"FinancialEvents,omitempty"`

	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken string `json:"NextToken,omitempty"`
}

// Validate validates this list financial events payload
func (m *ListFinancialEventsPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinancialEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListFinancialEventsPayload) validateFinancialEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.FinancialEvents) { // not required
		return nil
	}

	if m.FinancialEvents != nil {
		if err := m.FinancialEvents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FinancialEvents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FinancialEvents")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list financial events payload based on the context it is used
func (m *ListFinancialEventsPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFinancialEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListFinancialEventsPayload) contextValidateFinancialEvents(ctx context.Context, formats strfmt.Registry) error {

	if m.FinancialEvents != nil {
		if err := m.FinancialEvents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FinancialEvents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("FinancialEvents")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListFinancialEventsPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListFinancialEventsPayload) UnmarshalBinary(b []byte) error {
	var res ListFinancialEventsPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
