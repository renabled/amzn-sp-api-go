// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrackingDetailsInput Tracking information input for Less-Than-Truckload (LTL) and Small Parcel Delivery (SPD) shipments.
// Example: {"ltlTrackingDetail":{"billOfLadingNumber":"billOfLadingNumber","freightBillNumber":["freightBillNumber1"]}}
//
// swagger:model TrackingDetailsInput
type TrackingDetailsInput struct {

	// ltl tracking detail
	LtlTrackingDetail *LtlTrackingDetailInput `json:"ltlTrackingDetail,omitempty"`

	// spd tracking detail
	SpdTrackingDetail *SpdTrackingDetailInput `json:"spdTrackingDetail,omitempty"`
}

// Validate validates this tracking details input
func (m *TrackingDetailsInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLtlTrackingDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpdTrackingDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrackingDetailsInput) validateLtlTrackingDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.LtlTrackingDetail) { // not required
		return nil
	}

	if m.LtlTrackingDetail != nil {
		if err := m.LtlTrackingDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ltlTrackingDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ltlTrackingDetail")
			}
			return err
		}
	}

	return nil
}

func (m *TrackingDetailsInput) validateSpdTrackingDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.SpdTrackingDetail) { // not required
		return nil
	}

	if m.SpdTrackingDetail != nil {
		if err := m.SpdTrackingDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spdTrackingDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spdTrackingDetail")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tracking details input based on the context it is used
func (m *TrackingDetailsInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLtlTrackingDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpdTrackingDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrackingDetailsInput) contextValidateLtlTrackingDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.LtlTrackingDetail != nil {
		if err := m.LtlTrackingDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ltlTrackingDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ltlTrackingDetail")
			}
			return err
		}
	}

	return nil
}

func (m *TrackingDetailsInput) contextValidateSpdTrackingDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.SpdTrackingDetail != nil {
		if err := m.SpdTrackingDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spdTrackingDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spdTrackingDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrackingDetailsInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrackingDetailsInput) UnmarshalBinary(b []byte) error {
	var res TrackingDetailsInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
