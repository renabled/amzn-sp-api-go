// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinkableAccountType Info About Linkable Account Type
//
// swagger:model LinkableAccountType
type LinkableAccountType struct {

	// account type
	AccountType AccountType `json:"accountType,omitempty"`

	// carrier account inputs
	CarrierAccountInputs CarrierAccountInputsList `json:"carrierAccountInputs,omitempty"`
}

// Validate validates this linkable account type
func (m *LinkableAccountType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCarrierAccountInputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkableAccountType) validateAccountType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := m.AccountType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *LinkableAccountType) validateCarrierAccountInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.CarrierAccountInputs) { // not required
		return nil
	}

	if err := m.CarrierAccountInputs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("carrierAccountInputs")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("carrierAccountInputs")
		}
		return err
	}

	return nil
}

// ContextValidate validate this linkable account type based on the context it is used
func (m *LinkableAccountType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCarrierAccountInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkableAccountType) contextValidateAccountType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccountType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accountType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accountType")
		}
		return err
	}

	return nil
}

func (m *LinkableAccountType) contextValidateCarrierAccountInputs(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CarrierAccountInputs.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("carrierAccountInputs")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("carrierAccountInputs")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkableAccountType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkableAccountType) UnmarshalBinary(b []byte) error {
	var res LinkableAccountType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
