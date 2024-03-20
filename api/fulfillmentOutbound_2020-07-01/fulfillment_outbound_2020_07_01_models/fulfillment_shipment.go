// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

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

// FulfillmentShipment Delivery and item information for a shipment in a fulfillment order.
//
// swagger:model FulfillmentShipment
type FulfillmentShipment struct {

	// A shipment identifier assigned by Amazon.
	// Required: true
	AmazonShipmentID *string `json:"amazonShipmentId"`

	// The estimated arrival date and time of the shipment. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format. Note that this value can change over time. If a shipment includes more than one package, `estimatedArrivalDate` applies to all of the packages in the shipment. If the shipment has been cancelled, `estimatedArrivalDate` is not returned.
	// Format: date-time
	EstimatedArrivalDate Timestamp `json:"estimatedArrivalDate,omitempty"`

	// An identifier for the fulfillment center that the shipment will be sent from.
	// Required: true
	FulfillmentCenterID *string `json:"fulfillmentCenterId"`

	// fulfillment shipment item
	// Required: true
	FulfillmentShipmentItem FulfillmentShipmentItemList `json:"fulfillmentShipmentItem"`

	// fulfillment shipment package
	FulfillmentShipmentPackage FulfillmentShipmentPackageList `json:"fulfillmentShipmentPackage,omitempty"`

	// The current status of the shipment.
	// Required: true
	// Enum: [PENDING SHIPPED CANCELLED_BY_FULFILLER CANCELLED_BY_SELLER]
	FulfillmentShipmentStatus *string `json:"fulfillmentShipmentStatus"`

	// The meaning of the `shippingDate` value depends on the current status of the shipment. If the current value of `FulfillmentShipmentStatus` is:
	//
	// * Pending - `shippingDate` represents the estimated time that the shipment will leave the Amazon fulfillment center.
	//
	// * Shipped - `shippingDate` represents the date that the shipment left the Amazon fulfillment center.
	// If a shipment includes more than one package, `shippingDate` applies to all of the packages in the shipment. If the value of `FulfillmentShipmentStatus` is `CancelledByFulfiller` or `CancelledBySeller`, `shippingDate` is not returned. The value must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
	// Format: date-time
	ShippingDate Timestamp `json:"shippingDate,omitempty"`

	// Provides additional insight into shipment timeline. Primairly used to communicate that actual delivery dates aren't available.
	ShippingNotes []string `json:"shippingNotes"`
}

// Validate validates this fulfillment shipment
func (m *FulfillmentShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazonShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimatedArrivalDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentCenterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentShipmentItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentShipmentPackage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFulfillmentShipmentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentShipment) validateAmazonShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("amazonShipmentId", "body", m.AmazonShipmentID); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentShipment) validateEstimatedArrivalDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EstimatedArrivalDate) { // not required
		return nil
	}

	if err := m.EstimatedArrivalDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("estimatedArrivalDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("estimatedArrivalDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentShipment) validateFulfillmentCenterID(formats strfmt.Registry) error {

	if err := validate.Required("fulfillmentCenterId", "body", m.FulfillmentCenterID); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentShipment) validateFulfillmentShipmentItem(formats strfmt.Registry) error {

	if err := validate.Required("fulfillmentShipmentItem", "body", m.FulfillmentShipmentItem); err != nil {
		return err
	}

	if err := m.FulfillmentShipmentItem.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentShipmentItem")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentShipmentItem")
		}
		return err
	}

	return nil
}

func (m *FulfillmentShipment) validateFulfillmentShipmentPackage(formats strfmt.Registry) error {
	if swag.IsZero(m.FulfillmentShipmentPackage) { // not required
		return nil
	}

	if err := m.FulfillmentShipmentPackage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentShipmentPackage")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentShipmentPackage")
		}
		return err
	}

	return nil
}

var fulfillmentShipmentTypeFulfillmentShipmentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","SHIPPED","CANCELLED_BY_FULFILLER","CANCELLED_BY_SELLER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fulfillmentShipmentTypeFulfillmentShipmentStatusPropEnum = append(fulfillmentShipmentTypeFulfillmentShipmentStatusPropEnum, v)
	}
}

const (

	// FulfillmentShipmentFulfillmentShipmentStatusPENDING captures enum value "PENDING"
	FulfillmentShipmentFulfillmentShipmentStatusPENDING string = "PENDING"

	// FulfillmentShipmentFulfillmentShipmentStatusSHIPPED captures enum value "SHIPPED"
	FulfillmentShipmentFulfillmentShipmentStatusSHIPPED string = "SHIPPED"

	// FulfillmentShipmentFulfillmentShipmentStatusCANCELLEDBYFULFILLER captures enum value "CANCELLED_BY_FULFILLER"
	FulfillmentShipmentFulfillmentShipmentStatusCANCELLEDBYFULFILLER string = "CANCELLED_BY_FULFILLER"

	// FulfillmentShipmentFulfillmentShipmentStatusCANCELLEDBYSELLER captures enum value "CANCELLED_BY_SELLER"
	FulfillmentShipmentFulfillmentShipmentStatusCANCELLEDBYSELLER string = "CANCELLED_BY_SELLER"
)

// prop value enum
func (m *FulfillmentShipment) validateFulfillmentShipmentStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fulfillmentShipmentTypeFulfillmentShipmentStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FulfillmentShipment) validateFulfillmentShipmentStatus(formats strfmt.Registry) error {

	if err := validate.Required("fulfillmentShipmentStatus", "body", m.FulfillmentShipmentStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateFulfillmentShipmentStatusEnum("fulfillmentShipmentStatus", "body", *m.FulfillmentShipmentStatus); err != nil {
		return err
	}

	return nil
}

func (m *FulfillmentShipment) validateShippingDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippingDate) { // not required
		return nil
	}

	if err := m.ShippingDate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shippingDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shippingDate")
		}
		return err
	}

	return nil
}

// ContextValidate validate this fulfillment shipment based on the context it is used
func (m *FulfillmentShipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEstimatedArrivalDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentShipmentItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFulfillmentShipmentPackage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FulfillmentShipment) contextValidateEstimatedArrivalDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EstimatedArrivalDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("estimatedArrivalDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("estimatedArrivalDate")
		}
		return err
	}

	return nil
}

func (m *FulfillmentShipment) contextValidateFulfillmentShipmentItem(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentShipmentItem.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentShipmentItem")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentShipmentItem")
		}
		return err
	}

	return nil
}

func (m *FulfillmentShipment) contextValidateFulfillmentShipmentPackage(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FulfillmentShipmentPackage.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fulfillmentShipmentPackage")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fulfillmentShipmentPackage")
		}
		return err
	}

	return nil
}

func (m *FulfillmentShipment) contextValidateShippingDate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShippingDate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shippingDate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shippingDate")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FulfillmentShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FulfillmentShipment) UnmarshalBinary(b []byte) error {
	var res FulfillmentShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
