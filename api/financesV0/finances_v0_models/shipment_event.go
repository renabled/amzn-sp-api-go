// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShipmentEvent A shipment, refund, guarantee claim, or chargeback.
//
// swagger:model ShipmentEvent
type ShipmentEvent struct {

	// An Amazon-defined identifier for an order.
	AmazonOrderID string `json:"AmazonOrderId,omitempty"`

	// A list of transactions where buyers pay Amazon through one of the credit cards offered by Amazon or where buyers pay a seller directly through COD.
	DirectPaymentList DirectPaymentList `json:"DirectPaymentList,omitempty"`

	// The name of the marketplace where the event occurred.
	MarketplaceName string `json:"MarketplaceName,omitempty"`

	// A list of order-level charge adjustments. These adjustments are applicable to Multi-Channel Fulfillment COD orders.
	OrderChargeAdjustmentList ChargeComponentList `json:"OrderChargeAdjustmentList,omitempty"`

	// A list of order-level charges. These charges are applicable to Multi-Channel Fulfillment COD orders.
	OrderChargeList ChargeComponentList `json:"OrderChargeList,omitempty"`

	// A list of order-level fee adjustments. These adjustments are applicable to Multi-Channel Fulfillment orders.
	OrderFeeAdjustmentList FeeComponentList `json:"OrderFeeAdjustmentList,omitempty"`

	// A list of order-level fees. These charges are applicable to Multi-Channel Fulfillment orders.
	OrderFeeList FeeComponentList `json:"OrderFeeList,omitempty"`

	// The date and time when the financial event was posted.
	// Format: date-time
	PostedDate Date `json:"PostedDate,omitempty"`

	// A seller-defined identifier for an order.
	SellerOrderID string `json:"SellerOrderId,omitempty"`

	// A list of shipment-level fee adjustments.
	ShipmentFeeAdjustmentList FeeComponentList `json:"ShipmentFeeAdjustmentList,omitempty"`

	// A list of shipment-level fees.
	ShipmentFeeList FeeComponentList `json:"ShipmentFeeList,omitempty"`

	// A list of shipment item adjustments.
	ShipmentItemAdjustmentList ShipmentItemList `json:"ShipmentItemAdjustmentList,omitempty"`

	// shipment item list
	ShipmentItemList ShipmentItemList `json:"ShipmentItemList,omitempty"`

	// The name of the store where the event occurred.
	StoreName string `json:"StoreName,omitempty"`
}

// Validate validates this shipment event
func (m *ShipmentEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectPaymentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderChargeAdjustmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderChargeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderFeeAdjustmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderFeeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentFeeAdjustmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentFeeList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentItemAdjustmentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentItemList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentEvent) validateDirectPaymentList(formats strfmt.Registry) error {
	if swag.IsZero(m.DirectPaymentList) { // not required
		return nil
	}

	if err := m.DirectPaymentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DirectPaymentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("DirectPaymentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateOrderChargeAdjustmentList(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderChargeAdjustmentList) { // not required
		return nil
	}

	if err := m.OrderChargeAdjustmentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderChargeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderChargeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateOrderChargeList(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderChargeList) { // not required
		return nil
	}

	if err := m.OrderChargeList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderChargeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderChargeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateOrderFeeAdjustmentList(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderFeeAdjustmentList) { // not required
		return nil
	}

	if err := m.OrderFeeAdjustmentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderFeeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderFeeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateOrderFeeList(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderFeeList) { // not required
		return nil
	}

	if err := m.OrderFeeList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderFeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderFeeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validatePostedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PostedDate) { // not required
		return nil
	}

	if err := m.PostedDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PostedDate")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateShipmentFeeAdjustmentList(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentFeeAdjustmentList) { // not required
		return nil
	}

	if err := m.ShipmentFeeAdjustmentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentFeeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentFeeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateShipmentFeeList(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentFeeList) { // not required
		return nil
	}

	if err := m.ShipmentFeeList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentFeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentFeeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateShipmentItemAdjustmentList(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentItemAdjustmentList) { // not required
		return nil
	}

	if err := m.ShipmentItemAdjustmentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentItemAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentItemAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) validateShipmentItemList(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentItemList) { // not required
		return nil
	}

	if err := m.ShipmentItemList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentItemList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentItemList")
		}
		return err
	}

	return nil
}

// ContextValidate validate this shipment event based on the context it is used
func (m *ShipmentEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirectPaymentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderChargeAdjustmentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderChargeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderFeeAdjustmentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrderFeeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentFeeAdjustmentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentFeeList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentItemAdjustmentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentItemList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentEvent) contextValidateDirectPaymentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DirectPaymentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DirectPaymentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("DirectPaymentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateOrderChargeAdjustmentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderChargeAdjustmentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderChargeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderChargeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateOrderChargeList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderChargeList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderChargeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderChargeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateOrderFeeAdjustmentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderFeeAdjustmentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderFeeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderFeeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateOrderFeeList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderFeeList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("OrderFeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("OrderFeeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidatePostedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PostedDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PostedDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PostedDate")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateShipmentFeeAdjustmentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipmentFeeAdjustmentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentFeeAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentFeeAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateShipmentFeeList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipmentFeeList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentFeeList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentFeeList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateShipmentItemAdjustmentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipmentItemAdjustmentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentItemAdjustmentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentItemAdjustmentList")
		}
		return err
	}

	return nil
}

func (m *ShipmentEvent) contextValidateShipmentItemList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipmentItemList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentItemList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentItemList")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentEvent) UnmarshalBinary(b []byte) error {
	var res ShipmentEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
