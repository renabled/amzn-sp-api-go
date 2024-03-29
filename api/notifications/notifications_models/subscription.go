// Code generated by go-swagger; DO NOT EDIT.

package notifications_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subscription Information about the subscription.
//
// swagger:model Subscription
type Subscription struct {

	// The identifier for the destination where notifications will be delivered.
	// Required: true
	DestinationID *string `json:"destinationId"`

	// The version of the payload object to be used in the notification.
	// Required: true
	PayloadVersion *string `json:"payloadVersion"`

	// processing directive
	ProcessingDirective *ProcessingDirective `json:"processingDirective,omitempty"`

	// The subscription identifier generated when the subscription is created.
	// Required: true
	SubscriptionID *string `json:"subscriptionId"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayloadVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDirective(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) validateDestinationID(formats strfmt.Registry) error {

	if err := validate.Required("destinationId", "body", m.DestinationID); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validatePayloadVersion(formats strfmt.Registry) error {

	if err := validate.Required("payloadVersion", "body", m.PayloadVersion); err != nil {
		return err
	}

	return nil
}

func (m *Subscription) validateProcessingDirective(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingDirective) { // not required
		return nil
	}

	if m.ProcessingDirective != nil {
		if err := m.ProcessingDirective.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processingDirective")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processingDirective")
			}
			return err
		}
	}

	return nil
}

func (m *Subscription) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this subscription based on the context it is used
func (m *Subscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProcessingDirective(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subscription) contextValidateProcessingDirective(ctx context.Context, formats strfmt.Registry) error {

	if m.ProcessingDirective != nil {
		if err := m.ProcessingDirective.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("processingDirective")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("processingDirective")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscription) UnmarshalBinary(b []byte) error {
	var res Subscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
