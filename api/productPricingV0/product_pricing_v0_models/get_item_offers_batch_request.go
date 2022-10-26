// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetItemOffersBatchRequest The request associated with the `getItemOffersBatch` API call.
//
// swagger:model GetItemOffersBatchRequest
type GetItemOffersBatchRequest struct {

	// requests
	Requests ItemOffersRequestList `json:"requests,omitempty"`
}

// Validate validates this get item offers batch request
func (m *GetItemOffersBatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetItemOffersBatchRequest) validateRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.Requests) { // not required
		return nil
	}

	if err := m.Requests.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requests")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requests")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get item offers batch request based on the context it is used
func (m *GetItemOffersBatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetItemOffersBatchRequest) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Requests.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requests")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requests")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetItemOffersBatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetItemOffersBatchRequest) UnmarshalBinary(b []byte) error {
	var res GetItemOffersBatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
