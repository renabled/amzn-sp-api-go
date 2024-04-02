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

// GeneratePackingOptionsResponse The `generatePackingOptions` response.
//
// swagger:model GeneratePackingOptionsResponse
type GeneratePackingOptionsResponse struct {

	// UUID for the given operation.
	// Required: true
	// Max Length: 38
	// Min Length: 36
	// Pattern: ^[a-zA-Z0-9-]*$
	OperationID *string `json:"operationId"`
}

// Validate validates this generate packing options response
func (m *GeneratePackingOptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneratePackingOptionsResponse) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	if err := validate.MinLength("operationId", "body", *m.OperationID, 36); err != nil {
		return err
	}

	if err := validate.MaxLength("operationId", "body", *m.OperationID, 38); err != nil {
		return err
	}

	if err := validate.Pattern("operationId", "body", *m.OperationID, `^[a-zA-Z0-9-]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this generate packing options response based on context it is used
func (m *GeneratePackingOptionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GeneratePackingOptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneratePackingOptionsResponse) UnmarshalBinary(b []byte) error {
	var res GeneratePackingOptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
