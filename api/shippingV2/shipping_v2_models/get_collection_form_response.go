// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetCollectionFormResponse The Response  for the GetCollectionFormResponse operation.
//
// swagger:model GetCollectionFormResponse
type GetCollectionFormResponse struct {

	// collections form document
	CollectionsFormDocument *CollectionsFormDocument `json:"collectionsFormDocument,omitempty"`
}

// Validate validates this get collection form response
func (m *GetCollectionFormResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectionsFormDocument(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCollectionFormResponse) validateCollectionsFormDocument(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectionsFormDocument) { // not required
		return nil
	}

	if m.CollectionsFormDocument != nil {
		if err := m.CollectionsFormDocument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionsFormDocument")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionsFormDocument")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get collection form response based on the context it is used
func (m *GetCollectionFormResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectionsFormDocument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCollectionFormResponse) contextValidateCollectionsFormDocument(ctx context.Context, formats strfmt.Registry) error {

	if m.CollectionsFormDocument != nil {
		if err := m.CollectionsFormDocument.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionsFormDocument")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionsFormDocument")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCollectionFormResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCollectionFormResponse) UnmarshalBinary(b []byte) error {
	var res GetCollectionFormResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}