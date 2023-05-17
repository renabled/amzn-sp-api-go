// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderItemBuyerInfo A single order item's buyer information.
//
// swagger:model OrderItemBuyerInfo
type OrderItemBuyerInfo struct {

	// Buyer information for custom orders from the Amazon Custom program.
	BuyerCustomizedInfo *BuyerCustomizedInfoDetail `json:"BuyerCustomizedInfo,omitempty"`

	// A gift message provided by the buyer.
	GiftMessageText string `json:"GiftMessageText,omitempty"`

	// The gift wrap level specified by the buyer.
	GiftWrapLevel string `json:"GiftWrapLevel,omitempty"`

	// The gift wrap price of the item.
	GiftWrapPrice *Money `json:"GiftWrapPrice,omitempty"`

	// The tax on the gift wrap price.
	GiftWrapTax *Money `json:"GiftWrapTax,omitempty"`

	// An Amazon-defined order item identifier.
	// Required: true
	OrderItemID *string `json:"OrderItemId"`
}

// Validate validates this order item buyer info
func (m *OrderItemBuyerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuyerCustomizedInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftWrapPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftWrapTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItemBuyerInfo) validateBuyerCustomizedInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.BuyerCustomizedInfo) { // not required
		return nil
	}

	if m.BuyerCustomizedInfo != nil {
		if err := m.BuyerCustomizedInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerCustomizedInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerCustomizedInfo")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItemBuyerInfo) validateGiftWrapPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.GiftWrapPrice) { // not required
		return nil
	}

	if m.GiftWrapPrice != nil {
		if err := m.GiftWrapPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GiftWrapPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GiftWrapPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItemBuyerInfo) validateGiftWrapTax(formats strfmt.Registry) error {
	if swag.IsZero(m.GiftWrapTax) { // not required
		return nil
	}

	if m.GiftWrapTax != nil {
		if err := m.GiftWrapTax.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GiftWrapTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GiftWrapTax")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItemBuyerInfo) validateOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("OrderItemId", "body", m.OrderItemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this order item buyer info based on the context it is used
func (m *OrderItemBuyerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuyerCustomizedInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGiftWrapPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGiftWrapTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItemBuyerInfo) contextValidateBuyerCustomizedInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.BuyerCustomizedInfo != nil {
		if err := m.BuyerCustomizedInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BuyerCustomizedInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BuyerCustomizedInfo")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItemBuyerInfo) contextValidateGiftWrapPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.GiftWrapPrice != nil {
		if err := m.GiftWrapPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GiftWrapPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GiftWrapPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderItemBuyerInfo) contextValidateGiftWrapTax(ctx context.Context, formats strfmt.Registry) error {

	if m.GiftWrapTax != nil {
		if err := m.GiftWrapTax.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GiftWrapTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GiftWrapTax")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderItemBuyerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderItemBuyerInfo) UnmarshalBinary(b []byte) error {
	var res OrderItemBuyerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
