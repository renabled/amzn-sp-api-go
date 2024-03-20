// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateFulfillmentReturnRequest The `createFulfillmentReturn` operation creates a fulfillment return for items that were fulfilled using the `createFulfillmentOrder` operation. For calls to `createFulfillmentReturn`, you must include `ReturnReasonCode` values returned by a previous call to the `listReturnReasonCodes` operation.
//
// swagger:model CreateFulfillmentReturnRequest
type CreateFulfillmentReturnRequest struct {

	// items
	// Required: true
	Items CreateReturnItemList `json:"items"`
}

// Validate validates this create fulfillment return request
func (m *CreateFulfillmentReturnRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFulfillmentReturnRequest) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	if err := m.Items.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("items")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("items")
		}
		return err
	}

	return nil
}

// ContextValidate validate this create fulfillment return request based on the context it is used
func (m *CreateFulfillmentReturnRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFulfillmentReturnRequest) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Items.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("items")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("items")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateFulfillmentReturnRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFulfillmentReturnRequest) UnmarshalBinary(b []byte) error {
	var res CreateFulfillmentReturnRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
