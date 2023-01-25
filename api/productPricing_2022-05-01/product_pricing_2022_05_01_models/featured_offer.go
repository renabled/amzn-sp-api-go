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

// FeaturedOffer featured offer
//
// swagger:model FeaturedOffer
type FeaturedOffer struct {

	// The item condition.
	Condition Condition `json:"condition,omitempty"`

	// An offer identifier used to identify the merchant of the featured offer. Since this may not belong to the requester, the SKU field will be omitted.
	// Required: true
	OfferIdentifier *OfferIdentifier `json:"offerIdentifier"`

	// The current active price of the offer.
	Price *Price `json:"price,omitempty"`
}

// Validate validates this featured offer
func (m *FeaturedOffer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOffer) validateCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	if err := m.Condition.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("condition")
		}
		return err
	}

	return nil
}

func (m *FeaturedOffer) validateOfferIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("offerIdentifier", "body", m.OfferIdentifier); err != nil {
		return err
	}

	if m.OfferIdentifier != nil {
		if err := m.OfferIdentifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("offerIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("offerIdentifier")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOffer) validatePrice(formats strfmt.Registry) error {
	if swag.IsZero(m.Price) { // not required
		return nil
	}

	if m.Price != nil {
		if err := m.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this featured offer based on the context it is used
func (m *FeaturedOffer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOfferIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturedOffer) contextValidateCondition(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Condition.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("condition")
		}
		return err
	}

	return nil
}

func (m *FeaturedOffer) contextValidateOfferIdentifier(ctx context.Context, formats strfmt.Registry) error {

	if m.OfferIdentifier != nil {
		if err := m.OfferIdentifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("offerIdentifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("offerIdentifier")
			}
			return err
		}
	}

	return nil
}

func (m *FeaturedOffer) contextValidatePrice(ctx context.Context, formats strfmt.Registry) error {

	if m.Price != nil {
		if err := m.Price.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeaturedOffer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturedOffer) UnmarshalBinary(b []byte) error {
	var res FeaturedOffer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
