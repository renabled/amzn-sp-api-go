// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PackageDetail Properties of packages
//
// swagger:model PackageDetail
type PackageDetail struct {

	// Identifies the carrier that will deliver the package. This field is required for all marketplaces, see [reference](https://developer-docs.amazon.com/sp-api/changelog/carriercode-value-required-in-shipment-confirmations-for-br-mx-ca-sg-au-in-jp-marketplaces).
	// Required: true
	CarrierCode *string `json:"carrierCode"`

	// Carrier Name that will deliver the package. Required when carrierCode is "Others"
	CarrierName string `json:"carrierName,omitempty"`

	// The list of order items and quantities to be updated.
	// Required: true
	OrderItems ConfirmShipmentOrderItemsList `json:"orderItems"`

	// package reference Id
	// Required: true
	PackageReferenceID *PackageReferenceID `json:"packageReferenceId"`

	// The shipping date for the package. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.
	// Required: true
	// Format: date-time
	ShipDate *strfmt.DateTime `json:"shipDate"`

	// The unique identifier of the supply source.
	ShipFromSupplySourceID string `json:"shipFromSupplySourceId,omitempty"`

	// Ship method to be used for shipping the order.
	ShippingMethod string `json:"shippingMethod,omitempty"`

	// The tracking number used to obtain tracking and delivery information.
	// Required: true
	TrackingNumber *string `json:"trackingNumber"`
}

// Validate validates this package detail
func (m *PackageDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageDetail) validateCarrierCode(formats strfmt.Registry) error {

	if err := validate.Required("carrierCode", "body", m.CarrierCode); err != nil {
		return err
	}

	return nil
}

func (m *PackageDetail) validateOrderItems(formats strfmt.Registry) error {

	if err := validate.Required("orderItems", "body", m.OrderItems); err != nil {
		return err
	}

	if err := m.OrderItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("orderItems")
		}
		return err
	}

	return nil
}

func (m *PackageDetail) validatePackageReferenceID(formats strfmt.Registry) error {

	if err := validate.Required("packageReferenceId", "body", m.PackageReferenceID); err != nil {
		return err
	}

	if err := validate.Required("packageReferenceId", "body", m.PackageReferenceID); err != nil {
		return err
	}

	if m.PackageReferenceID != nil {
		if err := m.PackageReferenceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageReferenceId")
			}
			return err
		}
	}

	return nil
}

func (m *PackageDetail) validateShipDate(formats strfmt.Registry) error {

	if err := validate.Required("shipDate", "body", m.ShipDate); err != nil {
		return err
	}

	if err := validate.FormatOf("shipDate", "body", "date-time", m.ShipDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PackageDetail) validateTrackingNumber(formats strfmt.Registry) error {

	if err := validate.Required("trackingNumber", "body", m.TrackingNumber); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this package detail based on the context it is used
func (m *PackageDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackageReferenceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackageDetail) contextValidateOrderItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.OrderItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("orderItems")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("orderItems")
		}
		return err
	}

	return nil
}

func (m *PackageDetail) contextValidatePackageReferenceID(ctx context.Context, formats strfmt.Registry) error {

	if m.PackageReferenceID != nil {
		if err := m.PackageReferenceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageReferenceId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackageDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackageDetail) UnmarshalBinary(b []byte) error {
	var res PackageDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
