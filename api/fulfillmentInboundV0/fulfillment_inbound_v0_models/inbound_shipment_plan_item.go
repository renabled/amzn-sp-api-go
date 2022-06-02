// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InboundShipmentPlanItem Item information used to create an inbound shipment. Returned by the createInboundShipmentPlan operation.
//
// swagger:model InboundShipmentPlanItem
type InboundShipmentPlanItem struct {

	// Amazon's fulfillment network SKU of the item.
	// Required: true
	FulfillmentNetworkSKU *string `json:"FulfillmentNetworkSKU"`

	// prep details list
	PrepDetailsList PrepDetailsList `json:"PrepDetailsList,omitempty"`

	// The item quantity that you are shipping.
	// Required: true
	Quantity *Quantity `json:"Quantity"`

	// The seller SKU of the item.
	// Required: true
	SellerSKU *string `json:"SellerSKU"`
}

// Validate validates this inbound shipment plan item
func (m *InboundShipmentPlanItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFulfillmentNetworkSKU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepDetailsList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellerSKU(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundShipmentPlanItem) validateFulfillmentNetworkSKU(formats strfmt.Registry) error {

	if err := validate.Required("FulfillmentNetworkSKU", "body", m.FulfillmentNetworkSKU); err != nil {
		return err
	}

	return nil
}

func (m *InboundShipmentPlanItem) validatePrepDetailsList(formats strfmt.Registry) error {
	if swag.IsZero(m.PrepDetailsList) { // not required
		return nil
	}

	if err := m.PrepDetailsList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PrepDetailsList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PrepDetailsList")
		}
		return err
	}

	return nil
}

func (m *InboundShipmentPlanItem) validateQuantity(formats strfmt.Registry) error {

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

func (m *InboundShipmentPlanItem) validateSellerSKU(formats strfmt.Registry) error {

	if err := validate.Required("SellerSKU", "body", m.SellerSKU); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this inbound shipment plan item based on the context it is used
func (m *InboundShipmentPlanItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrepDetailsList(ctx, formats); err != nil {
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

func (m *InboundShipmentPlanItem) contextValidatePrepDetailsList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PrepDetailsList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PrepDetailsList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PrepDetailsList")
		}
		return err
	}

	return nil
}

func (m *InboundShipmentPlanItem) contextValidateQuantity(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *InboundShipmentPlanItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundShipmentPlanItem) UnmarshalBinary(b []byte) error {
	var res InboundShipmentPlanItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
