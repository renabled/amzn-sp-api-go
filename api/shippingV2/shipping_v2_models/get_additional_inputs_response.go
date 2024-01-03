// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAdditionalInputsResponse The response schema for the getAdditionalInputs operation.
//
// swagger:model GetAdditionalInputsResponse
type GetAdditionalInputsResponse struct {

	// payload
	Payload GetAdditionalInputsResult `json:"payload,omitempty"`
}

// Validate validates this get additional inputs response
func (m *GetAdditionalInputsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get additional inputs response based on context it is used
func (m *GetAdditionalInputsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetAdditionalInputsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAdditionalInputsResponse) UnmarshalBinary(b []byte) error {
	var res GetAdditionalInputsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}