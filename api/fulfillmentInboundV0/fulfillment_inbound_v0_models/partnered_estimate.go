// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartneredEstimate The estimated shipping cost for a shipment using an Amazon-partnered carrier.
//
// swagger:model PartneredEstimate
type PartneredEstimate struct {

	// The amount that the Amazon-partnered carrier will charge to ship the inbound shipment.
	// Required: true
	Amount *Amount `json:"Amount"`

	// The date in ISO 8601 date time format by which this estimate must be confirmed. After this date the estimate is no longer valid and cannot be confirmed.
	//
	// Returned only if the TransportStatus value of the inbound shipment is ESTIMATED.
	// Format: date-time
	ConfirmDeadline TimeStampStringType `json:"ConfirmDeadline,omitempty"`

	// The date in ISO 8601 date time format after which a confirmed transportation request can no longer be voided. This date is 24 hours after a Small Parcel shipment transportation request is confirmed or one hour after a Less Than Truckload/Full Truckload (LTL/FTL) shipment transportation request is confirmed. After the void deadline passes the seller's account will be charged for the shipping cost.
	//
	// Returned only if the TransportStatus value of the inbound shipment is CONFIRMED.
	// Format: date-time
	VoidDeadline TimeStampStringType `json:"VoidDeadline,omitempty"`
}

// Validate validates this partnered estimate
func (m *PartneredEstimate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfirmDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoidDeadline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartneredEstimate) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Amount")
			}
			return err
		}
	}

	return nil
}

func (m *PartneredEstimate) validateConfirmDeadline(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfirmDeadline) { // not required
		return nil
	}

	if err := m.ConfirmDeadline.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ConfirmDeadline")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ConfirmDeadline")
		}
		return err
	}

	return nil
}

func (m *PartneredEstimate) validateVoidDeadline(formats strfmt.Registry) error {
	if swag.IsZero(m.VoidDeadline) { // not required
		return nil
	}

	if err := m.VoidDeadline.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("VoidDeadline")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("VoidDeadline")
		}
		return err
	}

	return nil
}

// ContextValidate validate this partnered estimate based on the context it is used
func (m *PartneredEstimate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfirmDeadline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVoidDeadline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartneredEstimate) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {
		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Amount")
			}
			return err
		}
	}

	return nil
}

func (m *PartneredEstimate) contextValidateConfirmDeadline(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConfirmDeadline.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ConfirmDeadline")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ConfirmDeadline")
		}
		return err
	}

	return nil
}

func (m *PartneredEstimate) contextValidateVoidDeadline(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VoidDeadline.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("VoidDeadline")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("VoidDeadline")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartneredEstimate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartneredEstimate) UnmarshalBinary(b []byte) error {
	var res PartneredEstimate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
