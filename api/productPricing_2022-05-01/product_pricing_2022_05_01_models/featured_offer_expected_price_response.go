// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_2022_05_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FeaturedOfferExpectedPriceResponse featured offer expected price response
//
// swagger:model FeaturedOfferExpectedPriceResponse
type FeaturedOfferExpectedPriceResponse struct {
	BatchResponse

	// body
	Body *FeaturedOfferExpectedPriceResponseBody `json:"body,omitempty"`

	// Use these request parameters to map the response back to the request.
	// Required: true
	Request *FeaturedOfferExpectedPriceRequestParams `json:"request"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *FeaturedOfferExpectedPriceResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BatchResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BatchResponse = aO0

	// AO1
	var dataAO1 struct {
		Body *FeaturedOfferExpectedPriceResponseBody `json:"body,omitempty"`

		Request *FeaturedOfferExpectedPriceRequestParams `json:"request"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Body = dataAO1.Body

	m.Request = dataAO1.Request

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m FeaturedOfferExpectedPriceResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BatchResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Body *FeaturedOfferExpectedPriceResponseBody `json:"body,omitempty"`

		Request *FeaturedOfferExpectedPriceRequestParams `json:"request"`
	}

	dataAO1.Body = m.Body

	dataAO1.Request = m.Request

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this featured offer expected price response
func (m *FeaturedOfferExpectedPriceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BatchResponse
	if err := m.BatchResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOfferExpectedPriceResponse) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOfferExpectedPriceResponse) validateRequest(formats strfmt.Registry) error {

	if err := validate.Required("request", "body", m.Request); err != nil {
		return err
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this featured offer expected price response based on the context it is used
func (m *FeaturedOfferExpectedPriceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BatchResponse
	if err := m.BatchResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBody(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOfferExpectedPriceResponse) contextValidateBody(ctx context.Context, formats strfmt.Registry) error {

	if m.Body != nil {
		if err := m.Body.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOfferExpectedPriceResponse) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {
		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeaturedOfferExpectedPriceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturedOfferExpectedPriceResponse) UnmarshalBinary(b []byte) error {
	var res FeaturedOfferExpectedPriceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}