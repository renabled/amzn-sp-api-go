// Code generated by go-swagger; DO NOT EDIT.

package data_kiosk_2023_11_15_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateQueryResponse The response for the `createQuery` operation.
//
// swagger:model CreateQueryResponse
type CreateQueryResponse struct {

	// The identifier for the query. This identifier is unique only in combination with a selling partner account ID.
	// Required: true
	QueryID *string `json:"queryId"`
}

// Validate validates this create query response
func (m *CreateQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueryID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateQueryResponse) validateQueryID(formats strfmt.Registry) error {

	if err := validate.Required("queryId", "body", m.QueryID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create query response based on context it is used
func (m *CreateQueryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateQueryResponse) UnmarshalBinary(b []byte) error {
	var res CreateQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
