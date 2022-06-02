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

// GetListingOffersBatchResponse The response associated with the getListingOffersBatch API call.
//
// swagger:model GetListingOffersBatchResponse
type GetListingOffersBatchResponse struct {

	// responses
	Responses ListingOffersResponseList `json:"responses,omitempty"`
}

// Validate validates this get listing offers batch response
func (m *GetListingOffersBatchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetListingOffersBatchResponse) validateResponses(formats strfmt.Registry) error {
	if swag.IsZero(m.Responses) { // not required
		return nil
	}

	if err := m.Responses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("responses")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get listing offers batch response based on the context it is used
func (m *GetListingOffersBatchResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResponses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetListingOffersBatchResponse) contextValidateResponses(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Responses.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("responses")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetListingOffersBatchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetListingOffersBatchResponse) UnmarshalBinary(b []byte) error {
	var res GetListingOffersBatchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
