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

// FeaturedOfferExpectedPriceRequestParams The parameters for an individual request.
//
// swagger:model FeaturedOfferExpectedPriceRequestParams
type FeaturedOfferExpectedPriceRequestParams struct {

	// marketplace Id
	// Required: true
	MarketplaceID *MarketplaceID `json:"marketplaceId"`

	// sku
	// Required: true
	Sku *Sku `json:"sku"`
}

// Validate validates this featured offer expected price request params
func (m *FeaturedOfferExpectedPriceRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSku(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOfferExpectedPriceRequestParams) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketplaceId")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOfferExpectedPriceRequestParams) validateSku(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	if err := validate.Required("sku", "body", m.Sku); err != nil {
		return err
	}

	if m.Sku != nil {
		if err := m.Sku.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sku")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sku")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this featured offer expected price request params based on the context it is used
func (m *FeaturedOfferExpectedPriceRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMarketplaceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSku(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOfferExpectedPriceRequestParams) contextValidateMarketplaceID(ctx context.Context, formats strfmt.Registry) error {

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketplaceId")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOfferExpectedPriceRequestParams) contextValidateSku(ctx context.Context, formats strfmt.Registry) error {

	if m.Sku != nil {
		if err := m.Sku.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sku")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sku")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeaturedOfferExpectedPriceRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturedOfferExpectedPriceRequestParams) UnmarshalBinary(b []byte) error {
	var res FeaturedOfferExpectedPriceRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
