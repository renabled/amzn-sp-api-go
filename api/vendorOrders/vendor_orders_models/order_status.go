// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderStatus Current status of a purchase order.
//
// swagger:model OrderStatus
type OrderStatus struct {

	// Detailed order status.
	// Required: true
	ItemStatus ItemStatus `json:"itemStatus"`

	// The date when the purchase order was last updated. Must be in ISO-8601 date/time format.
	// Format: date-time
	LastUpdatedDate strfmt.DateTime `json:"lastUpdatedDate,omitempty"`

	// The date the purchase order was placed. Must be in ISO-8601 date/time format.
	// Required: true
	// Format: date-time
	PurchaseOrderDate *strfmt.DateTime `json:"purchaseOrderDate"`

	// The buyer's purchase order number for this order. Formatting Notes: 8-character alpha-numeric code.
	// Required: true
	PurchaseOrderNumber *string `json:"purchaseOrderNumber"`

	// The status of the buyer's purchase order for this order.
	// Required: true
	// Enum: [OPEN CLOSED]
	PurchaseOrderStatus *string `json:"purchaseOrderStatus"`

	// Name/Address and tax details of the selling party.
	// Required: true
	SellingParty *PartyIdentification `json:"sellingParty"`

	// Name/Address and tax details of the ship to party. Find a list of fulfillment center addresses for a region on the [Resources page of Amazon Vendor Central](https://vendorcentral.amazon.com/hz/vendor/members/support/help/node/GPZ88XH8HQM97ZV6).
	// Required: true
	ShipToParty *PartyIdentification `json:"shipToParty"`
}

// Validate validates this order status
func (m *OrderStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellingParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipToParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderStatus) validateItemStatus(formats strfmt.Registry) error {

	if err := validate.Required("itemStatus", "body", m.ItemStatus); err != nil {
		return err
	}

	if err := m.ItemStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("itemStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("itemStatus")
		}
		return err
	}

	return nil
}

func (m *OrderStatus) validateLastUpdatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedDate", "body", "date-time", m.LastUpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrderStatus) validatePurchaseOrderDate(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderDate", "body", m.PurchaseOrderDate); err != nil {
		return err
	}

	if err := validate.FormatOf("purchaseOrderDate", "body", "date-time", m.PurchaseOrderDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OrderStatus) validatePurchaseOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderNumber", "body", m.PurchaseOrderNumber); err != nil {
		return err
	}

	return nil
}

var orderStatusTypePurchaseOrderStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN","CLOSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderStatusTypePurchaseOrderStatusPropEnum = append(orderStatusTypePurchaseOrderStatusPropEnum, v)
	}
}

const (

	// OrderStatusPurchaseOrderStatusOPEN captures enum value "OPEN"
	OrderStatusPurchaseOrderStatusOPEN string = "OPEN"

	// OrderStatusPurchaseOrderStatusCLOSED captures enum value "CLOSED"
	OrderStatusPurchaseOrderStatusCLOSED string = "CLOSED"
)

// prop value enum
func (m *OrderStatus) validatePurchaseOrderStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, orderStatusTypePurchaseOrderStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OrderStatus) validatePurchaseOrderStatus(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderStatus", "body", m.PurchaseOrderStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validatePurchaseOrderStatusEnum("purchaseOrderStatus", "body", *m.PurchaseOrderStatus); err != nil {
		return err
	}

	return nil
}

func (m *OrderStatus) validateSellingParty(formats strfmt.Registry) error {

	if err := validate.Required("sellingParty", "body", m.SellingParty); err != nil {
		return err
	}

	if m.SellingParty != nil {
		if err := m.SellingParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderStatus) validateShipToParty(formats strfmt.Registry) error {

	if err := validate.Required("shipToParty", "body", m.ShipToParty); err != nil {
		return err
	}

	if m.ShipToParty != nil {
		if err := m.ShipToParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipToParty")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order status based on the context it is used
func (m *OrderStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItemStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellingParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipToParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderStatus) contextValidateItemStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ItemStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("itemStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("itemStatus")
		}
		return err
	}

	return nil
}

func (m *OrderStatus) contextValidateSellingParty(ctx context.Context, formats strfmt.Registry) error {

	if m.SellingParty != nil {
		if err := m.SellingParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sellingParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sellingParty")
			}
			return err
		}
	}

	return nil
}

func (m *OrderStatus) contextValidateShipToParty(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipToParty != nil {
		if err := m.ShipToParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipToParty")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderStatus) UnmarshalBinary(b []byte) error {
	var res OrderStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
