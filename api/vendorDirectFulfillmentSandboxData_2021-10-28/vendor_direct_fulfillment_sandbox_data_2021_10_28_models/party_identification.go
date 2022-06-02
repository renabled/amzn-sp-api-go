// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_sandbox_data_2021_10_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartyIdentification The identification object for the party information. For example, warehouse code or vendor code. Please refer to specific party for more details.
//
// swagger:model PartyIdentification
type PartyIdentification struct {

	// Assigned identification for the party. For example, warehouse code or vendor code. Please refer to specific party for more details.
	// Required: true
	PartyID *string `json:"partyId"`
}

// Validate validates this party identification
func (m *PartyIdentification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartyIdentification) validatePartyID(formats strfmt.Registry) error {

	if err := validate.Required("partyId", "body", m.PartyID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this party identification based on context it is used
func (m *PartyIdentification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
