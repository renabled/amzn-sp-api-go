// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_2022_05_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetFeaturedOfferExpectedPriceBatchRequest The request body for the getFeaturedOfferExpectedPriceBatch operation.
//
// swagger:model GetFeaturedOfferExpectedPriceBatchRequest
type GetFeaturedOfferExpectedPriceBatchRequest struct {

	// requests
	Requests FeaturedOfferExpectedPriceRequestList `json:"requests,omitempty"`
}

// Validate validates this get featured offer expected price batch request
func (m *GetFeaturedOfferExpectedPriceBatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFeaturedOfferExpectedPriceBatchRequest) validateRequests(formats strfmt.Registry) error {
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

// ContextValidate validate this get featured offer expected price batch request based on the context it is used
func (m *GetFeaturedOfferExpectedPriceBatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFeaturedOfferExpectedPriceBatchRequest) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetFeaturedOfferExpectedPriceBatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFeaturedOfferExpectedPriceBatchRequest) UnmarshalBinary(b []byte) error {
	var res GetFeaturedOfferExpectedPriceBatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
