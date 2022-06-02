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

// TaxWithholdingEvent A TaxWithholding event on seller's account.
//
// swagger:model TaxWithholdingEvent
type TaxWithholdingEvent struct {

	// The amount which tax was withheld against.
	BaseAmount *Currency `json:"BaseAmount,omitempty"`

	// The date and time when the financial event was posted.
	// Format: date-time
	PostedDate Date `json:"PostedDate,omitempty"`

	// Time period for which tax is withheld.
	TaxWithholdingPeriod *TaxWithholdingPeriod `json:"TaxWithholdingPeriod,omitempty"`

	// The amount of the tax withholding deducted from seller's account.
	WithheldAmount *Currency `json:"WithheldAmount,omitempty"`
}

// Validate validates this tax withholding event
func (m *TaxWithholdingEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxWithholdingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWithheldAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxWithholdingEvent) validateBaseAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseAmount) { // not required
		return nil
	}

	if m.BaseAmount != nil {
		if err := m.BaseAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseAmount")
			}
			return err
		}
	}

	return nil
}

func (m *TaxWithholdingEvent) validatePostedDate(formats strfmt.Registry) error {
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

func (m *TaxWithholdingEvent) validateTaxWithholdingPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxWithholdingPeriod) { // not required
		return nil
	}

	if m.TaxWithholdingPeriod != nil {
		if err := m.TaxWithholdingPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxWithholdingPeriod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxWithholdingPeriod")
			}
			return err
		}
	}

	return nil
}

func (m *TaxWithholdingEvent) validateWithheldAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.WithheldAmount) { // not required
		return nil
	}

	if m.WithheldAmount != nil {
		if err := m.WithheldAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WithheldAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("WithheldAmount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tax withholding event based on the context it is used
func (m *TaxWithholdingEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxWithholdingPeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWithheldAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxWithholdingEvent) contextValidateBaseAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseAmount != nil {
		if err := m.BaseAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseAmount")
			}
			return err
		}
	}

	return nil
}

func (m *TaxWithholdingEvent) contextValidatePostedDate(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TaxWithholdingEvent) contextValidateTaxWithholdingPeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxWithholdingPeriod != nil {
		if err := m.TaxWithholdingPeriod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TaxWithholdingPeriod")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TaxWithholdingPeriod")
			}
			return err
		}
	}

	return nil
}

func (m *TaxWithholdingEvent) contextValidateWithheldAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.WithheldAmount != nil {
		if err := m.WithheldAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WithheldAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("WithheldAmount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxWithholdingEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxWithholdingEvent) UnmarshalBinary(b []byte) error {
	var res TaxWithholdingEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
