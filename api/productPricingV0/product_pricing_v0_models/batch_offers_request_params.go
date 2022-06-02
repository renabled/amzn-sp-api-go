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

// BatchOffersRequestParams batch offers request params
//
// swagger:model BatchOffersRequestParams
type BatchOffersRequestParams struct {

	// customer type
	CustomerType CustomerType `json:"CustomerType,omitempty"`

	// item condition
	// Required: true
	ItemCondition *ItemCondition `json:"ItemCondition"`

	// marketplace Id
	// Required: true
	MarketplaceID *MarketplaceID `json:"MarketplaceId"`
}

// Validate validates this batch offers request params
func (m *BatchOffersRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchOffersRequestParams) validateCustomerType(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerType) { // not required
		return nil
	}

	if err := m.CustomerType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CustomerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CustomerType")
		}
		return err
	}

	return nil
}

func (m *BatchOffersRequestParams) validateItemCondition(formats strfmt.Registry) error {

	if err := validate.Required("ItemCondition", "body", m.ItemCondition); err != nil {
		return err
	}

	if err := validate.Required("ItemCondition", "body", m.ItemCondition); err != nil {
		return err
	}

	if m.ItemCondition != nil {
		if err := m.ItemCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemCondition")
			}
			return err
		}
	}

	return nil
}

func (m *BatchOffersRequestParams) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("MarketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if err := validate.Required("MarketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MarketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MarketplaceId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this batch offers request params based on the context it is used
func (m *BatchOffersRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarketplaceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchOffersRequestParams) contextValidateCustomerType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CustomerType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CustomerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("CustomerType")
		}
		return err
	}

	return nil
}

func (m *BatchOffersRequestParams) contextValidateItemCondition(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemCondition != nil {
		if err := m.ItemCondition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemCondition")
			}
			return err
		}
	}

	return nil
}

func (m *BatchOffersRequestParams) contextValidateMarketplaceID(ctx context.Context, formats strfmt.Registry) error {

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MarketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MarketplaceId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchOffersRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchOffersRequestParams) UnmarshalBinary(b []byte) error {
	var res BatchOffersRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
