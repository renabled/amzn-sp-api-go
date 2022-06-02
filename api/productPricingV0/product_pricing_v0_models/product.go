// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Product An item.
//
// swagger:model Product
type Product struct {

	// attribute sets
	AttributeSets AttributeSetList `json:"AttributeSets,omitempty"`

	// competitive pricing
	CompetitivePricing *CompetitivePricingType `json:"CompetitivePricing,omitempty"`

	// identifiers
	// Required: true
	Identifiers *IdentifierType `json:"Identifiers"`

	// offers
	Offers OffersList `json:"Offers,omitempty"`

	// relationships
	Relationships RelationshipList `json:"Relationships,omitempty"`

	// sales rankings
	SalesRankings SalesRankList `json:"SalesRankings,omitempty"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompetitivePricing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSalesRankings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateAttributeSets(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeSets) { // not required
		return nil
	}

	if err := m.AttributeSets.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AttributeSets")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AttributeSets")
		}
		return err
	}

	return nil
}

func (m *Product) validateCompetitivePricing(formats strfmt.Registry) error {
	if swag.IsZero(m.CompetitivePricing) { // not required
		return nil
	}

	if m.CompetitivePricing != nil {
		if err := m.CompetitivePricing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CompetitivePricing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CompetitivePricing")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateIdentifiers(formats strfmt.Registry) error {

	if err := validate.Required("Identifiers", "body", m.Identifiers); err != nil {
		return err
	}

	if m.Identifiers != nil {
		if err := m.Identifiers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Identifiers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Identifiers")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateOffers(formats strfmt.Registry) error {
	if swag.IsZero(m.Offers) { // not required
		return nil
	}

	if err := m.Offers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Offers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Offers")
		}
		return err
	}

	return nil
}

func (m *Product) validateRelationships(formats strfmt.Registry) error {
	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if err := m.Relationships.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Relationships")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Relationships")
		}
		return err
	}

	return nil
}

func (m *Product) validateSalesRankings(formats strfmt.Registry) error {
	if swag.IsZero(m.SalesRankings) { // not required
		return nil
	}

	if err := m.SalesRankings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SalesRankings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SalesRankings")
		}
		return err
	}

	return nil
}

// ContextValidate validate this product based on the context it is used
func (m *Product) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributeSets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCompetitivePricing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOffers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSalesRankings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) contextValidateAttributeSets(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AttributeSets.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AttributeSets")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AttributeSets")
		}
		return err
	}

	return nil
}

func (m *Product) contextValidateCompetitivePricing(ctx context.Context, formats strfmt.Registry) error {

	if m.CompetitivePricing != nil {
		if err := m.CompetitivePricing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CompetitivePricing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("CompetitivePricing")
			}
			return err
		}
	}

	return nil
}

func (m *Product) contextValidateIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	if m.Identifiers != nil {
		if err := m.Identifiers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Identifiers")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Identifiers")
			}
			return err
		}
	}

	return nil
}

func (m *Product) contextValidateOffers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Offers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Offers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Offers")
		}
		return err
	}

	return nil
}

func (m *Product) contextValidateRelationships(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Relationships.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Relationships")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Relationships")
		}
		return err
	}

	return nil
}

func (m *Product) contextValidateSalesRankings(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SalesRankings.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SalesRankings")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("SalesRankings")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
