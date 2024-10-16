// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_2021_12_28_models

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

// ShippingLabelRequest Represents the request payload to create a shipping label. Contains the purchase order number, selling party, ship from party, and a list of containers or packages in the shipment.
//
// swagger:model ShippingLabelRequest
type ShippingLabelRequest struct {

	// A list of the packages in this shipment.
	Containers []*Container `json:"containers"`

	// Purchase order number of the order for which to create a shipping label.
	// Required: true
	// Pattern: ^[a-zA-Z0-9]+$
	PurchaseOrderNumber *string `json:"purchaseOrderNumber"`

	// ID of the selling party or vendor.
	// Required: true
	SellingParty *PartyIdentification `json:"sellingParty"`

	// Warehouse code of vendor.
	// Required: true
	ShipFromParty *PartyIdentification `json:"shipFromParty"`
}

// Validate validates this shipping label request
func (m *ShippingLabelRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSellingParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFromParty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingLabelRequest) validateContainers(formats strfmt.Registry) error {
	if swag.IsZero(m.Containers) { // not required
		return nil
	}

	for i := 0; i < len(m.Containers); i++ {
		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {
			if err := m.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ShippingLabelRequest) validatePurchaseOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderNumber", "body", m.PurchaseOrderNumber); err != nil {
		return err
	}

	if err := validate.Pattern("purchaseOrderNumber", "body", *m.PurchaseOrderNumber, `^[a-zA-Z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ShippingLabelRequest) validateSellingParty(formats strfmt.Registry) error {

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

func (m *ShippingLabelRequest) validateShipFromParty(formats strfmt.Registry) error {

	if err := validate.Required("shipFromParty", "body", m.ShipFromParty); err != nil {
		return err
	}

	if m.ShipFromParty != nil {
		if err := m.ShipFromParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFromParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFromParty")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this shipping label request based on the context it is used
func (m *ShippingLabelRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSellingParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFromParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingLabelRequest) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Containers); i++ {

		if m.Containers[i] != nil {
			if err := m.Containers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ShippingLabelRequest) contextValidateSellingParty(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ShippingLabelRequest) contextValidateShipFromParty(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipFromParty != nil {
		if err := m.ShipFromParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFromParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFromParty")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShippingLabelRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShippingLabelRequest) UnmarshalBinary(b []byte) error {
	var res ShippingLabelRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
