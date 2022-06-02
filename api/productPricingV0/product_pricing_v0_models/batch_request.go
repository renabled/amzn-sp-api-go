// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BatchRequest Common properties of batch requests against individual APIs.
//
// swagger:model BatchRequest
type BatchRequest struct {

	// headers
	Headers HTTPRequestHeaders `json:"headers,omitempty"`

	// method
	// Required: true
	Method *HTTPMethod `json:"method"`

	// The full URI corresponding to the API intended for request, including path parameter substitutions.
	// Required: true
	URI *string `json:"uri"`
}

// Validate validates this batch request
func (m *BatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchRequest) validateHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	if m.Headers != nil {
		if err := m.Headers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("headers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("headers")
			}
			return err
		}
	}

	return nil
}

func (m *BatchRequest) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	if m.Method != nil {
		if err := m.Method.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *BatchRequest) validateURI(formats strfmt.Registry) error {

	if err := validate.Required("uri", "body", m.URI); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this batch request based on the context it is used
func (m *BatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchRequest) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Headers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("headers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("headers")
		}
		return err
	}

	return nil
}

func (m *BatchRequest) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.Method != nil {
		if err := m.Method.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchRequest) UnmarshalBinary(b []byte) error {
	var res BatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
