// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateFulfillmentReturnResult The result for the createFulfillmentReturn operation.
//
// swagger:model CreateFulfillmentReturnResult
type CreateFulfillmentReturnResult struct {

	// invalid return items
	InvalidReturnItems InvalidReturnItemList `json:"invalidReturnItems,omitempty"`

	// return authorizations
	ReturnAuthorizations ReturnAuthorizationList `json:"returnAuthorizations,omitempty"`

	// return items
	ReturnItems ReturnItemList `json:"returnItems,omitempty"`
}

// Validate validates this create fulfillment return result
func (m *CreateFulfillmentReturnResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvalidReturnItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFulfillmentReturnResult) validateInvalidReturnItems(formats strfmt.Registry) error {
	if swag.IsZero(m.InvalidReturnItems) { // not required
		return nil
	}

	if err := m.InvalidReturnItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("invalidReturnItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("invalidReturnItems")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentReturnResult) validateReturnAuthorizations(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnAuthorizations) { // not required
		return nil
	}

	if err := m.ReturnAuthorizations.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("returnAuthorizations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("returnAuthorizations")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentReturnResult) validateReturnItems(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnItems) { // not required
		return nil
	}

	if err := m.ReturnItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("returnItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("returnItems")
		}
		return err
	}

	return nil
}

// ContextValidate validate this create fulfillment return result based on the context it is used
func (m *CreateFulfillmentReturnResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvalidReturnItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReturnAuthorizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReturnItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFulfillmentReturnResult) contextValidateInvalidReturnItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InvalidReturnItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("invalidReturnItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("invalidReturnItems")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentReturnResult) contextValidateReturnAuthorizations(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ReturnAuthorizations.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("returnAuthorizations")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("returnAuthorizations")
		}
		return err
	}

	return nil
}

func (m *CreateFulfillmentReturnResult) contextValidateReturnItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ReturnItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("returnItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("returnItems")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateFulfillmentReturnResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFulfillmentReturnResult) UnmarshalBinary(b []byte) error {
	var res CreateFulfillmentReturnResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
