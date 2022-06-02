// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostContentDocumentAsinRelationsRequest post content document asin relations request
//
// swagger:model PostContentDocumentAsinRelationsRequest
type PostContentDocumentAsinRelationsRequest struct {

	// asin set
	// Required: true
	AsinSet AsinSet `json:"asinSet"`
}

// Validate validates this post content document asin relations request
func (m *PostContentDocumentAsinRelationsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsinSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostContentDocumentAsinRelationsRequest) validateAsinSet(formats strfmt.Registry) error {

	if err := validate.Required("asinSet", "body", m.AsinSet); err != nil {
		return err
	}

	if err := m.AsinSet.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("asinSet")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("asinSet")
		}
		return err
	}

	return nil
}

// ContextValidate validate this post content document asin relations request based on the context it is used
func (m *PostContentDocumentAsinRelationsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAsinSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostContentDocumentAsinRelationsRequest) contextValidateAsinSet(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AsinSet.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("asinSet")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("asinSet")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostContentDocumentAsinRelationsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostContentDocumentAsinRelationsRequest) UnmarshalBinary(b []byte) error {
	var res PostContentDocumentAsinRelationsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
