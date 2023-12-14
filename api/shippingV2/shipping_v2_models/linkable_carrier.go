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

// LinkableCarrier Info About Linkable Carrier
//
// swagger:model LinkableCarrier
type LinkableCarrier struct {

	// carrier Id
	CarrierID CarrierID `json:"carrierId,omitempty"`

	// linkable account types
	LinkableAccountTypes LinkableAccountTypeList `json:"linkableAccountTypes,omitempty"`
}

// Validate validates this linkable carrier
func (m *LinkableCarrier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkableAccountTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkableCarrier) validateCarrierID(formats strfmt.Registry) error {
	if swag.IsZero(m.CarrierID) { // not required
		return nil
	}

	if err := m.CarrierID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("carrierId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("carrierId")
		}
		return err
	}

	return nil
}

func (m *LinkableCarrier) validateLinkableAccountTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkableAccountTypes) { // not required
		return nil
	}

	if err := m.LinkableAccountTypes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("linkableAccountTypes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("linkableAccountTypes")
		}
		return err
	}

	return nil
}

// ContextValidate validate this linkable carrier based on the context it is used
func (m *LinkableCarrier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCarrierID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkableAccountTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkableCarrier) contextValidateCarrierID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CarrierID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("carrierId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("carrierId")
		}
		return err
	}

	return nil
}

func (m *LinkableCarrier) contextValidateLinkableAccountTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LinkableAccountTypes.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("linkableAccountTypes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("linkableAccountTypes")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinkableCarrier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkableCarrier) UnmarshalBinary(b []byte) error {
	var res LinkableCarrier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
