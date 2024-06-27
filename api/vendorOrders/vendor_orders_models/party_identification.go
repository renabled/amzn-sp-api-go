// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartyIdentification Name, address and tax details of a party.
//
// swagger:model PartyIdentification
type PartyIdentification struct {

	// Identification of the party by address.
	Address *Address `json:"address,omitempty"`

	// Assigned identification for the party. For example, warehouse code or vendor code. Please refer to specific party for more details.
	// Required: true
	PartyID *string `json:"partyId"`

	// Tax registration details of the party.
	TaxInfo *TaxRegistrationDetails `json:"taxInfo,omitempty"`
}

// Validate validates this party identification
func (m *PartyIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyIdentification) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *PartyIdentification) validatePartyID(formats strfmt.Registry) error {

	if err := validate.Required("partyId", "body", m.PartyID); err != nil {
		return err
	}

	return nil
}

func (m *PartyIdentification) validateTaxInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxInfo) { // not required
		return nil
	}

	if m.TaxInfo != nil {
		if err := m.TaxInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this party identification based on the context it is used
func (m *PartyIdentification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyIdentification) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {
		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *PartyIdentification) contextValidateTaxInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxInfo != nil {
		if err := m.TaxInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartyIdentification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartyIdentification) UnmarshalBinary(b []byte) error {
	var res PartyIdentification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
