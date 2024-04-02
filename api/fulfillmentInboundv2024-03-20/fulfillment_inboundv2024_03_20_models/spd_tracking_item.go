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

// SpdTrackingItem Contains information used to track and identify a Small Parcel Delivery (SPD) item.
//
// swagger:model SpdTrackingItem
type SpdTrackingItem struct {

	// The ID provided by Amazon that identifies a given box. This ID is comprised of the external shipment ID (which
	//         is generated after transportation has been confirmed) and the index of the box.
	// Max Length: 1024
	// Min Length: 1
	BoxID string `json:"boxId,omitempty"`

	// The tracking ID associated with each box in a non-Amazon partnered Small Parcel Delivery (SPD) shipment.
	// Max Length: 1024
	// Min Length: 1
	TrackingID string `json:"trackingId,omitempty"`

	// Whether or not Amazon has validated the tracking number. If more than 24 hours have passed and the status is
	//         not yet 'VALIDATED', please verify the number and update if necessary. Can be `VALIDATED` or `NOT_VALIDATED`.
	// Max Length: 1024
	// Min Length: 1
	TrackingNumberValidationStatus string `json:"trackingNumberValidationStatus,omitempty"`
}

// Validate validates this spd tracking item
func (m *SpdTrackingItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoxID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingNumberValidationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpdTrackingItem) validateBoxID(formats strfmt.Registry) error {
	if swag.IsZero(m.BoxID) { // not required
		return nil
	}

	if err := validate.MinLength("boxId", "body", m.BoxID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("boxId", "body", m.BoxID, 1024); err != nil {
		return err
	}

	return nil
}

func (m *SpdTrackingItem) validateTrackingID(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackingID) { // not required
		return nil
	}

	if err := validate.MinLength("trackingId", "body", m.TrackingID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("trackingId", "body", m.TrackingID, 1024); err != nil {
		return err
	}

	return nil
}

func (m *SpdTrackingItem) validateTrackingNumberValidationStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackingNumberValidationStatus) { // not required
		return nil
	}

	if err := validate.MinLength("trackingNumberValidationStatus", "body", m.TrackingNumberValidationStatus, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("trackingNumberValidationStatus", "body", m.TrackingNumberValidationStatus, 1024); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this spd tracking item based on context it is used
func (m *SpdTrackingItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpdTrackingItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpdTrackingItem) UnmarshalBinary(b []byte) error {
	var res SpdTrackingItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
