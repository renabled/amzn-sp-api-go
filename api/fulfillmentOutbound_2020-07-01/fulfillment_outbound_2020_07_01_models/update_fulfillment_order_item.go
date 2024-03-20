// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateFulfillmentOrderItem Item information for updating a fulfillment order.
//
// swagger:model UpdateFulfillmentOrderItem
type UpdateFulfillmentOrderItem struct {

	// Item-specific text that displays in recipient-facing materials such as the outbound shipment packing slip.
	// Max Length: 250
	DisplayableComment string `json:"displayableComment,omitempty"`

	// Amazon's fulfillment network SKU of the item.
	FulfillmentNetworkSku string `json:"fulfillmentNetworkSku,omitempty"`

	// A message to the gift recipient, if applicable.
	// Max Length: 512
	GiftMessage string `json:"giftMessage,omitempty"`

	// Indicates whether the item is sellable or unsellable.
	OrderItemDisposition string `json:"orderItemDisposition,omitempty"`

	// The monetary value assigned by the seller to this item. This is a required field for India MCF orders.
	PerUnitDeclaredValue *Money `json:"perUnitDeclaredValue,omitempty"`

	// The amount to be collected from the recipient for this item in a COD (Cash On Delivery) order.
	PerUnitPrice *Money `json:"perUnitPrice,omitempty"`

	// The tax on the amount to be collected from the recipient for this item in a COD (Cash On Delivery) order.
	PerUnitTax *Money `json:"perUnitTax,omitempty"`

	// quantity
	// Required: true
	Quantity *Quantity `json:"quantity"`

	// Identifies the fulfillment order item to update. Created with a previous call to the `createFulfillmentOrder` operation.
	// Required: true
	// Max Length: 50
	SellerFulfillmentOrderItemID *string `json:"sellerFulfillmentOrderItemId"`

	// The seller SKU of the item.
	SellerSku string `json:"sellerSku,omitempty"`
}

// Validate validates this update fulfillment order item
func (m *UpdateFulfillmentOrderItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayableComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerUnitDeclaredValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerUnitPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerUnitTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerFulfillmentOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateFulfillmentOrderItem) validateDisplayableComment(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayableComment) { // not required
		return nil
	}

	if err := validate.MaxLength("displayableComment", "body", m.DisplayableComment, 250); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) validateGiftMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.GiftMessage) { // not required
		return nil
	}

	if err := validate.MaxLength("giftMessage", "body", m.GiftMessage, 512); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) validatePerUnitDeclaredValue(formats strfmt.Registry) error {
	if swag.IsZero(m.PerUnitDeclaredValue) { // not required
		return nil
	}

	if m.PerUnitDeclaredValue != nil {
		if err := m.PerUnitDeclaredValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perUnitDeclaredValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perUnitDeclaredValue")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) validatePerUnitPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.PerUnitPrice) { // not required
		return nil
	}

	if m.PerUnitPrice != nil {
		if err := m.PerUnitPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perUnitPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perUnitPrice")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) validatePerUnitTax(formats strfmt.Registry) error {
	if swag.IsZero(m.PerUnitTax) { // not required
		return nil
	}

	if m.PerUnitTax != nil {
		if err := m.PerUnitTax.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perUnitTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perUnitTax")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if m.Quantity != nil {
		if err := m.Quantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quantity")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) validateSellerFulfillmentOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("sellerFulfillmentOrderItemId", "body", m.SellerFulfillmentOrderItemID); err != nil {
		return err
	}

	if err := validate.MaxLength("sellerFulfillmentOrderItemId", "body", *m.SellerFulfillmentOrderItemID, 50); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update fulfillment order item based on the context it is used
func (m *UpdateFulfillmentOrderItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerUnitDeclaredValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerUnitPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerUnitTax(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateFulfillmentOrderItem) contextValidatePerUnitDeclaredValue(ctx context.Context, formats strfmt.Registry) error {

	if m.PerUnitDeclaredValue != nil {
		if err := m.PerUnitDeclaredValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perUnitDeclaredValue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perUnitDeclaredValue")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) contextValidatePerUnitPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.PerUnitPrice != nil {
		if err := m.PerUnitPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perUnitPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perUnitPrice")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) contextValidatePerUnitTax(ctx context.Context, formats strfmt.Registry) error {

	if m.PerUnitTax != nil {
		if err := m.PerUnitTax.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perUnitTax")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perUnitTax")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFulfillmentOrderItem) contextValidateQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.Quantity != nil {
		if err := m.Quantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("quantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateFulfillmentOrderItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateFulfillmentOrderItem) UnmarshalBinary(b []byte) error {
	var res UpdateFulfillmentOrderItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
