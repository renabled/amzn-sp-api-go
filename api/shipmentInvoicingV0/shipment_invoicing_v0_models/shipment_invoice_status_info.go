// Code generated by go-swagger; DO NOT EDIT.

package shipment_invoicing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShipmentInvoiceStatusInfo The shipment invoice status information.
//
// swagger:model ShipmentInvoiceStatusInfo
type ShipmentInvoiceStatusInfo struct {

	// The Amazon-defined shipment identifier.
	AmazonShipmentID string `json:"AmazonShipmentId,omitempty"`

	// invoice status
	InvoiceStatus ShipmentInvoiceStatus `json:"InvoiceStatus,omitempty"`
}

// Validate validates this shipment invoice status info
func (m *ShipmentInvoiceStatusInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvoiceStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentInvoiceStatusInfo) validateInvoiceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.InvoiceStatus) { // not required
		return nil
	}

	if err := m.InvoiceStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InvoiceStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InvoiceStatus")
		}
		return err
	}

	return nil
}

// ContextValidate validate this shipment invoice status info based on the context it is used
func (m *ShipmentInvoiceStatusInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvoiceStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentInvoiceStatusInfo) contextValidateInvoiceStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.InvoiceStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("InvoiceStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("InvoiceStatus")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentInvoiceStatusInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentInvoiceStatusInfo) UnmarshalBinary(b []byte) error {
	var res ShipmentInvoiceStatusInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
