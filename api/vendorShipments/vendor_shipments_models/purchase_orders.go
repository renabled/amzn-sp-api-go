// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

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

// PurchaseOrders Transport Request pickup date
//
// swagger:model purchaseOrders
type PurchaseOrders struct {

	// A list of the items that are associated to the PO in this transport and their associated details.
	Items []*PurchaseOrderItems `json:"items"`

	// Purchase order numbers involved in this shipment, list all the PO that are involved as part of this shipment.
	// Format: date-time
	PurchaseOrderDate strfmt.DateTime `json:"purchaseOrderDate,omitempty"`

	// Purchase order numbers involved in this shipment, list all the PO that are involved as part of this shipment.
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`

	// Date range in which shipment is expected for these purchase orders.
	ShipWindow string `json:"shipWindow,omitempty"`
}

// Validate validates this purchase orders
func (m *PurchaseOrders) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchaseOrders) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PurchaseOrders) validatePurchaseOrderDate(formats strfmt.Registry) error {
	if swag.IsZero(m.PurchaseOrderDate) { // not required
		return nil
	}

	if err := validate.FormatOf("purchaseOrderDate", "body", "date-time", m.PurchaseOrderDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this purchase orders based on the context it is used
func (m *PurchaseOrders) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurchaseOrders) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PurchaseOrders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurchaseOrders) UnmarshalBinary(b []byte) error {
	var res PurchaseOrders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}