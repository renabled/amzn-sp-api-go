// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IdentifierType identifier type
//
// swagger:model IdentifierType
type IdentifierType struct {

	// Indicates the item is identified by MarketPlaceId and ASIN.
	MarketplaceASIN *ASINIdentifier `json:"MarketplaceASIN,omitempty"`

	// Indicates the item is identified by MarketPlaceId, SellerId, and SellerSKU.
	SKUIdentifier *SellerSKUIdentifier `json:"SKUIdentifier,omitempty"`
}

// Validate validates this identifier type
func (m *IdentifierType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMarketplaceASIN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSKUIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentifierType) validateMarketplaceASIN(formats strfmt.Registry) error {
	if swag.IsZero(m.MarketplaceASIN) { // not required
		return nil
	}

	if m.MarketplaceASIN != nil {
		if err := m.MarketplaceASIN.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MarketplaceASIN")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MarketplaceASIN")
			}
			return err
		}
	}

	return nil
}

func (m *IdentifierType) validateSKUIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.SKUIdentifier) { // not required
		return nil
	}

	if m.SKUIdentifier != nil {
		if err := m.SKUIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SKUIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SKUIdentifier")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this identifier type based on the context it is used
func (m *IdentifierType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMarketplaceASIN(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSKUIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentifierType) contextValidateMarketplaceASIN(ctx context.Context, formats strfmt.Registry) error {

	if m.MarketplaceASIN != nil {
		if err := m.MarketplaceASIN.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MarketplaceASIN")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MarketplaceASIN")
			}
			return err
		}
	}

	return nil
}

func (m *IdentifierType) contextValidateSKUIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if m.SKUIdentifier != nil {
		if err := m.SKUIdentifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SKUIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("SKUIdentifier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentifierType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentifierType) UnmarshalBinary(b []byte) error {
	var res IdentifierType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
