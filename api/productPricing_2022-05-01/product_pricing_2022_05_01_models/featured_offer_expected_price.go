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

// FeaturedOfferExpectedPrice The item price at or below which the target offer may be featured.
//
// swagger:model FeaturedOfferExpectedPrice
type FeaturedOfferExpectedPrice struct {

	// A computed listing price at or below which a seller can expect to become the featured offer (before applicable promotions).
	// Required: true
	ListingPrice *MoneyType `json:"listingPrice"`

	// The number of Amazon Points offered with the purchase of an item, and their monetary value.
	Points *Points `json:"points,omitempty"`
}

// Validate validates this featured offer expected price
func (m *FeaturedOfferExpectedPrice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListingPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOfferExpectedPrice) validateListingPrice(formats strfmt.Registry) error {

	if err := validate.Required("listingPrice", "body", m.ListingPrice); err != nil {
		return err
	}

	if m.ListingPrice != nil {
		if err := m.ListingPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOfferExpectedPrice) validatePoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Points) { // not required
		return nil
	}

	if m.Points != nil {
		if err := m.Points.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("points")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("points")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this featured offer expected price based on the context it is used
func (m *FeaturedOfferExpectedPrice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateListingPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOfferExpectedPrice) contextValidateListingPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.ListingPrice != nil {
		if err := m.ListingPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listingPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listingPrice")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOfferExpectedPrice) contextValidatePoints(ctx context.Context, formats strfmt.Registry) error {

	if m.Points != nil {
		if err := m.Points.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("points")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("points")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeaturedOfferExpectedPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturedOfferExpectedPrice) UnmarshalBinary(b []byte) error {
	var res FeaturedOfferExpectedPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
