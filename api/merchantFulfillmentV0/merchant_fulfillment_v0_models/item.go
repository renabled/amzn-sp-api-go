// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Item An Amazon order item identifier and a quantity.
//
// swagger:model Item
type Item struct {

	// item description
	ItemDescription ItemDescription `json:"ItemDescription,omitempty"`

	// A list of additional seller inputs required to ship this item using the chosen shipping service.
	ItemLevelSellerInputsList AdditionalSellerInputsList `json:"ItemLevelSellerInputsList,omitempty"`

	// item weight
	ItemWeight *Weight `json:"ItemWeight,omitempty"`

	// order item Id
	// Required: true
	OrderItemID *OrderItemID `json:"OrderItemId"`

	// quantity
	// Required: true
	Quantity *ItemQuantity `json:"Quantity"`

	// transparency code list
	TransparencyCodeList TransparencyCodeList `json:"TransparencyCodeList,omitempty"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemLevelSellerInputsList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransparencyCodeList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateItemDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemDescription) { // not required
		return nil
	}

	if err := m.ItemDescription.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemDescription")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemDescription")
		}
		return err
	}

	return nil
}

func (m *Item) validateItemLevelSellerInputsList(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemLevelSellerInputsList) { // not required
		return nil
	}

	if err := m.ItemLevelSellerInputsList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemLevelSellerInputsList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemLevelSellerInputsList")
		}
		return err
	}

	return nil
}

func (m *Item) validateItemWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemWeight) { // not required
		return nil
	}

	if m.ItemWeight != nil {
		if err := m.ItemWeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemWeight")
			}
			return err
		}
	}

	return nil
}

func (m *Item) validateOrderItemID(formats strfmt.Registry) error {

	if err := validate.Required("OrderItemId", "body", m.OrderItemID); err != nil {
		return err
	}

	if err := validate.Required("OrderItemId", "body", m.OrderItemID); err != nil {
		return err
	}

	if m.OrderItemID != nil {
		if err := m.OrderItemID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OrderItemId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OrderItemId")
			}
			return err
		}
	}

	return nil
}

func (m *Item) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("Quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.Required("Quantity", "body", m.Quantity); err != nil {
		return err
	}

	if m.Quantity != nil {
		if err := m.Quantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Quantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Quantity")
			}
			return err
		}
	}

	return nil
}

func (m *Item) validateTransparencyCodeList(formats strfmt.Registry) error {
	if swag.IsZero(m.TransparencyCodeList) { // not required
		return nil
	}

	if err := m.TransparencyCodeList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TransparencyCodeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("TransparencyCodeList")
		}
		return err
	}

	return nil
}

// ContextValidate validate this item based on the context it is used
func (m *Item) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItemDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemLevelSellerInputsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderItemID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransparencyCodeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) contextValidateItemDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemDescription.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemDescription")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemDescription")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateItemLevelSellerInputsList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemLevelSellerInputsList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ItemLevelSellerInputsList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ItemLevelSellerInputsList")
		}
		return err
	}

	return nil
}

func (m *Item) contextValidateItemWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.ItemWeight != nil {
		if err := m.ItemWeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ItemWeight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ItemWeight")
			}
			return err
		}
	}

	return nil
}

func (m *Item) contextValidateOrderItemID(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderItemID != nil {
		if err := m.OrderItemID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OrderItemId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OrderItemId")
			}
			return err
		}
	}

	return nil
}

func (m *Item) contextValidateQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.Quantity != nil {
		if err := m.Quantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Quantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Quantity")
			}
			return err
		}
	}

	return nil
}

func (m *Item) contextValidateTransparencyCodeList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TransparencyCodeList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TransparencyCodeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("TransparencyCodeList")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Item) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Item) UnmarshalBinary(b []byte) error {
	var res Item
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
