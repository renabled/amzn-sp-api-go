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

// TDSReimbursementEvent An event related to a Tax-Deducted-at-Source (TDS) reimbursement.
//
// swagger:model TDSReimbursementEvent
type TDSReimbursementEvent struct {

	// The date and time when the financial event was posted.
	// Format: date-time
	PostedDate Date `json:"PostedDate,omitempty"`

	// The amount reimbursed.
	ReimbursedAmount *Currency `json:"ReimbursedAmount,omitempty"`

	// The Tax-Deducted-at-Source (TDS) identifier.
	TDSOrderID string `json:"TDSOrderId,omitempty"`
}

// Validate validates this t d s reimbursement event
func (m *TDSReimbursementEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePostedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReimbursedAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TDSReimbursementEvent) validatePostedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PostedDate) { // not required
		return nil
	}

	if err := m.PostedDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PostedDate")
		}
		return err
	}

	return nil
}

func (m *TDSReimbursementEvent) validateReimbursedAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.ReimbursedAmount) { // not required
		return nil
	}

	if m.ReimbursedAmount != nil {
		if err := m.ReimbursedAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReimbursedAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ReimbursedAmount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this t d s reimbursement event based on the context it is used
func (m *TDSReimbursementEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePostedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReimbursedAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TDSReimbursementEvent) contextValidatePostedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PostedDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PostedDate")
		}
		return err
	}

	return nil
}

func (m *TDSReimbursementEvent) contextValidateReimbursedAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.ReimbursedAmount != nil {
		if err := m.ReimbursedAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ReimbursedAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ReimbursedAmount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TDSReimbursementEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TDSReimbursementEvent) UnmarshalBinary(b []byte) error {
	var res TDSReimbursementEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
