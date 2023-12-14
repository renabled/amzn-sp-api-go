// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CollectionsFormDocument Collection Form Document Details
//
// swagger:model CollectionsFormDocument
type CollectionsFormDocument struct {

	// Base64 document Value of Collection.
	Base64EncodedContent string `json:"base64EncodedContent,omitempty"`

	// Collection Document format is PDF.
	DocumentFormat string `json:"documentFormat,omitempty"`
}

// Validate validates this collections form document
func (m *CollectionsFormDocument) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this collections form document based on context it is used
func (m *CollectionsFormDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CollectionsFormDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectionsFormDocument) UnmarshalBinary(b []byte) error {
	var res CollectionsFormDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
