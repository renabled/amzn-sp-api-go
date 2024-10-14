// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentUpdatePreview Preview of the changes that will be applied to the shipment.
//
// swagger:model ContentUpdatePreview
type ContentUpdatePreview struct {

	// Identifier of a content update preview.
	// Required: true
	// Max Length: 38
	// Min Length: 38
	// Pattern: ^[a-zA-Z0-9-]*$
	ContentUpdatePreviewID *string `json:"contentUpdatePreviewId"`

	// The time at which the content update expires. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mm:ss.sssZ`.
	// Required: true
	// Format: date-time
	Expiration *strfmt.DateTime `json:"expiration"`

	// requested updates
	// Required: true
	RequestedUpdates *RequestedUpdates `json:"requestedUpdates"`

	// transportation option
	// Required: true
	TransportationOption *TransportationOption `json:"transportationOption"`
}

// Validate validates this content update preview
func (m *ContentUpdatePreview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentUpdatePreviewID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedUpdates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportationOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentUpdatePreview) validateContentUpdatePreviewID(formats strfmt.Registry) error {

	if err := validate.Required("contentUpdatePreviewId", "body", m.ContentUpdatePreviewID); err != nil {
		return err
	}

	if err := validate.MinLength("contentUpdatePreviewId", "body", *m.ContentUpdatePreviewID, 38); err != nil {
		return err
	}

	if err := validate.MaxLength("contentUpdatePreviewId", "body", *m.ContentUpdatePreviewID, 38); err != nil {
		return err
	}

	if err := validate.Pattern("contentUpdatePreviewId", "body", *m.ContentUpdatePreviewID, `^[a-zA-Z0-9-]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ContentUpdatePreview) validateExpiration(formats strfmt.Registry) error {

	if err := validate.Required("expiration", "body", m.Expiration); err != nil {
		return err
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentUpdatePreview) validateRequestedUpdates(formats strfmt.Registry) error {

	if err := validate.Required("requestedUpdates", "body", m.RequestedUpdates); err != nil {
		return err
	}

	if m.RequestedUpdates != nil {
		if err := m.RequestedUpdates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedUpdates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestedUpdates")
			}
			return err
		}
	}

	return nil
}

func (m *ContentUpdatePreview) validateTransportationOption(formats strfmt.Registry) error {

	if err := validate.Required("transportationOption", "body", m.TransportationOption); err != nil {
		return err
	}

	if m.TransportationOption != nil {
		if err := m.TransportationOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transportationOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transportationOption")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this content update preview based on the context it is used
func (m *ContentUpdatePreview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequestedUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransportationOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentUpdatePreview) contextValidateRequestedUpdates(ctx context.Context, formats strfmt.Registry) error {

	if m.RequestedUpdates != nil {
		if err := m.RequestedUpdates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestedUpdates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestedUpdates")
			}
			return err
		}
	}

	return nil
}

func (m *ContentUpdatePreview) contextValidateTransportationOption(ctx context.Context, formats strfmt.Registry) error {

	if m.TransportationOption != nil {
		if err := m.TransportationOption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transportationOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transportationOption")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentUpdatePreview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentUpdatePreview) UnmarshalBinary(b []byte) error {
	var res ContentUpdatePreview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
