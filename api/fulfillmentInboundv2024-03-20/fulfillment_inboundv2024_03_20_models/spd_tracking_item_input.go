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

// SpdTrackingItemInput Small Parcel Delivery (SPD) tracking items input information.
// Example: {"boxId":"FBA10ABC0YY100001","trackingId":"FBA10002000"}
//
// swagger:model SpdTrackingItemInput
type SpdTrackingItemInput struct {

	// The ID provided by Amazon that identifies a given box. This ID is comprised of the external shipment ID (which
	//         is generated after transportation has been confirmed) and the index of the box.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	BoxID *string `json:"boxId"`

	// The tracking Id associated with each box in a non-Amazon partnered Small Parcel Delivery (SPD) shipment. The seller must provide this information.
	// Required: true
	// Max Length: 1024
	// Min Length: 1
	TrackingID *string `json:"trackingId"`
}

// Validate validates this spd tracking item input
func (m *SpdTrackingItemInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoxID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpdTrackingItemInput) validateBoxID(formats strfmt.Registry) error {

	if err := validate.Required("boxId", "body", m.BoxID); err != nil {
		return err
	}

	if err := validate.MinLength("boxId", "body", *m.BoxID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("boxId", "body", *m.BoxID, 1024); err != nil {
		return err
	}

	return nil
}

func (m *SpdTrackingItemInput) validateTrackingID(formats strfmt.Registry) error {

	if err := validate.Required("trackingId", "body", m.TrackingID); err != nil {
		return err
	}

	if err := validate.MinLength("trackingId", "body", *m.TrackingID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("trackingId", "body", *m.TrackingID, 1024); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this spd tracking item input based on context it is used
func (m *SpdTrackingItemInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpdTrackingItemInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpdTrackingItemInput) UnmarshalBinary(b []byte) error {
	var res SpdTrackingItemInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
