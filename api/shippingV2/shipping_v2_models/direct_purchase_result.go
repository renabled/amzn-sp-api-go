// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectPurchaseResult The payload for the directPurchaseShipment operation.
//
// swagger:model DirectPurchaseResult
type DirectPurchaseResult struct {

	// package document detail list
	PackageDocumentDetailList PackageDocumentDetailList `json:"packageDocumentDetailList,omitempty"`

	// shipment Id
	// Required: true
	ShipmentID *ShipmentID `json:"shipmentId"`
}

// Validate validates this direct purchase result
func (m *DirectPurchaseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackageDocumentDetailList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectPurchaseResult) validatePackageDocumentDetailList(formats strfmt.Registry) error {
	if swag.IsZero(m.PackageDocumentDetailList) { // not required
		return nil
	}

	if err := m.PackageDocumentDetailList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packageDocumentDetailList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("packageDocumentDetailList")
		}
		return err
	}

	return nil
}

func (m *DirectPurchaseResult) validateShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("shipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	if err := validate.Required("shipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	if m.ShipmentID != nil {
		if err := m.ShipmentID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this direct purchase result based on the context it is used
func (m *DirectPurchaseResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePackageDocumentDetailList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectPurchaseResult) contextValidatePackageDocumentDetailList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PackageDocumentDetailList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packageDocumentDetailList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("packageDocumentDetailList")
		}
		return err
	}

	return nil
}

func (m *DirectPurchaseResult) contextValidateShipmentID(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentID != nil {
		if err := m.ShipmentID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectPurchaseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectPurchaseResult) UnmarshalBinary(b []byte) error {
	var res DirectPurchaseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
