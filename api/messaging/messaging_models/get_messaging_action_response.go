// Code generated by go-swagger; DO NOT EDIT.

package messaging_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetMessagingActionResponse Describes a messaging action that can be taken for an order. Provides a JSON Hypertext Application Language (HAL) link to the JSON schema document that describes the expected input.
//
// swagger:model GetMessagingActionResponse
type GetMessagingActionResponse struct {

	// embedded
	Embedded *GetMessagingActionResponseEmbedded `json:"_embedded,omitempty"`

	// links
	Links *GetMessagingActionResponseLinks `json:"_links,omitempty"`

	// errors
	Errors ErrorList `json:"errors,omitempty"`

	// payload
	Payload *MessagingAction `json:"payload,omitempty"`
}

// Validate validates this get messaging action response
func (m *GetMessagingActionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmbedded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMessagingActionResponse) validateEmbedded(formats strfmt.Registry) error {
	if swag.IsZero(m.Embedded) { // not required
		return nil
	}

	if m.Embedded != nil {
		if err := m.Embedded.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *GetMessagingActionResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *GetMessagingActionResponse) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if err := m.Errors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *GetMessagingActionResponse) validatePayload(formats strfmt.Registry) error {
	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get messaging action response based on the context it is used
func (m *GetMessagingActionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmbedded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayload(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMessagingActionResponse) contextValidateEmbedded(ctx context.Context, formats strfmt.Registry) error {

	if m.Embedded != nil {
		if err := m.Embedded.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded")
			}
			return err
		}
	}

	return nil
}

func (m *GetMessagingActionResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *GetMessagingActionResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Errors.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errors")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errors")
		}
		return err
	}

	return nil
}

func (m *GetMessagingActionResponse) contextValidatePayload(ctx context.Context, formats strfmt.Registry) error {

	if m.Payload != nil {
		if err := m.Payload.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payload")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetMessagingActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMessagingActionResponse) UnmarshalBinary(b []byte) error {
	var res GetMessagingActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetMessagingActionResponseEmbedded get messaging action response embedded
//
// swagger:model GetMessagingActionResponseEmbedded
type GetMessagingActionResponseEmbedded struct {

	// schema
	Schema *GetSchemaResponse `json:"schema,omitempty"`
}

// Validate validates this get messaging action response embedded
func (m *GetMessagingActionResponseEmbedded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMessagingActionResponseEmbedded) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get messaging action response embedded based on the context it is used
func (m *GetMessagingActionResponseEmbedded) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMessagingActionResponseEmbedded) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_embedded" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_embedded" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetMessagingActionResponseEmbedded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMessagingActionResponseEmbedded) UnmarshalBinary(b []byte) error {
	var res GetMessagingActionResponseEmbedded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetMessagingActionResponseLinks get messaging action response links
//
// swagger:model GetMessagingActionResponseLinks
type GetMessagingActionResponseLinks struct {

	// schema
	// Required: true
	Schema *LinkObject `json:"schema"`

	// self
	// Required: true
	Self *LinkObject `json:"self"`
}

// Validate validates this get messaging action response links
func (m *GetMessagingActionResponseLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMessagingActionResponseLinks) validateSchema(formats strfmt.Registry) error {

	if err := validate.Required("_links"+"."+"schema", "body", m.Schema); err != nil {
		return err
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

func (m *GetMessagingActionResponseLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("_links"+"."+"self", "body", m.Self); err != nil {
		return err
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get messaging action response links based on the context it is used
func (m *GetMessagingActionResponseLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMessagingActionResponseLinks) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "schema")
			}
			return err
		}
	}

	return nil
}

func (m *GetMessagingActionResponseLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetMessagingActionResponseLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMessagingActionResponseLinks) UnmarshalBinary(b []byte) error {
	var res GetMessagingActionResponseLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
