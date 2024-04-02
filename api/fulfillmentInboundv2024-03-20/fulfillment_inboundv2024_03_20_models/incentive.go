// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inboundv2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Incentive Contains details about cost related modifications to the placement cost.
//
// swagger:model Incentive
type Incentive struct {

	// Description of the incentive.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	Description *string `json:"description"`

	// Target of the incentive. Can be 'Placement Services' or 'Fulfillment Fee Discount'.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	Target *string `json:"target"`

	// Type of incentive. Can be `FEE` or `DISCOUNT`.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	Type *string `json:"type"`

	// value
	// Required: true
	Value *Currency `json:"value"`
}

// Validate validates this incentive
func (m *Incentive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Incentive) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 1024); err != nil {
		return err
	}

	return nil
}

func (m *Incentive) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	if err := validate.MinLength("target", "body", *m.Target, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("target", "body", *m.Target, 1024); err != nil {
		return err
	}

	return nil
}

func (m *Incentive) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", *m.Type, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", *m.Type, 1024); err != nil {
		return err
	}

	return nil
}

func (m *Incentive) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this incentive based on the context it is used
func (m *Incentive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Incentive) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Incentive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Incentive) UnmarshalBinary(b []byte) error {
	var res Incentive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
