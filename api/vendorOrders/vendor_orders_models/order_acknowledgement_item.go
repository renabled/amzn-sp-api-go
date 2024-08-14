// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderAcknowledgementItem Details of the item being acknowledged.
//
// swagger:model OrderAcknowledgementItem
type OrderAcknowledgementItem struct {

	// Amazon Standard Identification Number (ASIN) of an item.
	AmazonProductIdentifier string `json:"amazonProductIdentifier,omitempty"`

	// The discount multiplier that should be applied to the price if a vendor sells books with a list price. This is a multiplier factor to arrive at a final discounted price. A multiplier of .90 would be the factor if a 10% discount is given.
	DiscountMultiplier string `json:"discountMultiplier,omitempty"`

	// This is used to indicate acknowledged quantity.
	// Required: true
	ItemAcknowledgements []*OrderItemAcknowledgement `json:"itemAcknowledgements"`

	// Line item sequence number for the item.
	ItemSequenceNumber string `json:"itemSequenceNumber,omitempty"`

	// The list price of an item per each or weight unit. Required only if a vendor sells books at list price.
	ListPrice *Money `json:"listPrice,omitempty"`

	// The net cost of an item per each or weight unit that must match the cost on the invoice. This is a required field. If left blank, Amazon systems will reject the file. Price information must not be zero or negative.
	NetCost *Money `json:"netCost,omitempty"`

	// The quantity of this item ordered.
	// Required: true
	OrderedQuantity *ItemQuantity `json:"orderedQuantity"`

	// The vendor selected product identification of the item. Should be the same as was sent in the purchase order.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
}

// Validate validates this order acknowledgement item
func (m *OrderAcknowledgementItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemAcknowledgements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderAcknowledgementItem) validateItemAcknowledgements(formats strfmt.Registry) error {

	if err := validate.Required("itemAcknowledgements", "body", m.ItemAcknowledgements); err != nil {
		return err
	}

	for i := 0; i < len(m.ItemAcknowledgements); i++ {
		if swag.IsZero(m.ItemAcknowledgements[i]) { // not required
			continue
		}

		if m.ItemAcknowledgements[i] != nil {
			if err := m.ItemAcknowledgements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderAcknowledgementItem) validateListPrice(formats strfmt.Registry) error {
	if swag.IsZero(m.ListPrice) { // not required
		return nil
	}

	if m.ListPrice != nil {
		if err := m.ListPrice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAcknowledgementItem) validateNetCost(formats strfmt.Registry) error {
	if swag.IsZero(m.NetCost) { // not required
		return nil
	}

	if m.NetCost != nil {
		if err := m.NetCost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netCost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netCost")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAcknowledgementItem) validateOrderedQuantity(formats strfmt.Registry) error {

	if err := validate.Required("orderedQuantity", "body", m.OrderedQuantity); err != nil {
		return err
	}

	if m.OrderedQuantity != nil {
		if err := m.OrderedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderedQuantity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order acknowledgement item based on the context it is used
func (m *OrderAcknowledgementItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItemAcknowledgements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateListPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderAcknowledgementItem) contextValidateItemAcknowledgements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ItemAcknowledgements); i++ {

		if m.ItemAcknowledgements[i] != nil {
			if err := m.ItemAcknowledgements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("itemAcknowledgements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderAcknowledgementItem) contextValidateListPrice(ctx context.Context, formats strfmt.Registry) error {

	if m.ListPrice != nil {
		if err := m.ListPrice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listPrice")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("listPrice")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAcknowledgementItem) contextValidateNetCost(ctx context.Context, formats strfmt.Registry) error {

	if m.NetCost != nil {
		if err := m.NetCost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netCost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netCost")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAcknowledgementItem) contextValidateOrderedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderedQuantity != nil {
		if err := m.OrderedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderedQuantity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderAcknowledgementItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderAcknowledgementItem) UnmarshalBinary(b []byte) error {
	var res OrderAcknowledgementItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
