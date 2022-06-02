// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TemporarilyUnavailableCarrier A carrier who is temporarily unavailable, most likely due to a service outage experienced by the carrier.
//
// swagger:model TemporarilyUnavailableCarrier
type TemporarilyUnavailableCarrier struct {

	// The name of the carrier.
	// Required: true
	CarrierName *string `json:"CarrierName"`
}

// Validate validates this temporarily unavailable carrier
func (m *TemporarilyUnavailableCarrier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemporarilyUnavailableCarrier) validateCarrierName(formats strfmt.Registry) error {

	if err := validate.Required("CarrierName", "body", m.CarrierName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this temporarily unavailable carrier based on context it is used
func (m *TemporarilyUnavailableCarrier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemporarilyUnavailableCarrier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemporarilyUnavailableCarrier) UnmarshalBinary(b []byte) error {
	var res TemporarilyUnavailableCarrier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
