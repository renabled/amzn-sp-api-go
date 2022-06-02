// Code generated by go-swagger; DO NOT EDIT.

package notifications_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MarketplaceFilter Use this event filter to customize your subscription to send notifications for only the specified marketplaceId's.
//
// swagger:model MarketplaceFilter
type MarketplaceFilter struct {

	// marketplace ids
	MarketplaceIds MarketplaceIds `json:"marketplaceIds,omitempty"`
}

// Validate validates this marketplace filter
func (m *MarketplaceFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMarketplaceIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceFilter) validateMarketplaceIds(formats strfmt.Registry) error {
	if swag.IsZero(m.MarketplaceIds) { // not required
		return nil
	}

	if err := m.MarketplaceIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("marketplaceIds")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("marketplaceIds")
		}
		return err
	}

	return nil
}

// ContextValidate validate this marketplace filter based on the context it is used
func (m *MarketplaceFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMarketplaceIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceFilter) contextValidateMarketplaceIds(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MarketplaceIds.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("marketplaceIds")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("marketplaceIds")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceFilter) UnmarshalBinary(b []byte) error {
	var res MarketplaceFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
