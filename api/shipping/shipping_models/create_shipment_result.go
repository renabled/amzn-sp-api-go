// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateShipmentResult The payload schema for the createShipment operation.
//
// swagger:model CreateShipmentResult
type CreateShipmentResult struct {

	// eligible rates
	// Required: true
	EligibleRates RateList `json:"eligibleRates"`

	// shipment Id
	// Required: true
	ShipmentID *ShipmentID `json:"shipmentId"`
}

// Validate validates this create shipment result
func (m *CreateShipmentResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEligibleRates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShipmentResult) validateEligibleRates(formats strfmt.Registry) error {

	if err := validate.Required("eligibleRates", "body", m.EligibleRates); err != nil {
		return err
	}

	if err := m.EligibleRates.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eligibleRates")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("eligibleRates")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentResult) validateShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("shipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	if err := validate.Required("shipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	if m.ShipmentID != nil {
		if err := m.ShipmentID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create shipment result based on the context it is used
func (m *CreateShipmentResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEligibleRates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShipmentResult) contextValidateEligibleRates(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EligibleRates.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eligibleRates")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("eligibleRates")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentResult) contextValidateShipmentID(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentID != nil {
		if err := m.ShipmentID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateShipmentResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShipmentResult) UnmarshalBinary(b []byte) error {
	var res CreateShipmentResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
