// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

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

// ConfirmShipmentRequest The request schema for an shipment confirmation.
//
// swagger:model ConfirmShipmentRequest
type ConfirmShipmentRequest struct {

	// The COD collection method (only supported in the JP marketplace).
	// Enum: [DirectPayment]
	CodCollectionMethod string `json:"codCollectionMethod,omitempty"`

	// marketplace Id
	// Required: true
	MarketplaceID *MarketplaceID `json:"marketplaceId"`

	// package detail
	// Required: true
	PackageDetail *PackageDetail `json:"packageDetail"`
}

// Validate validates this confirm shipment request
func (m *ConfirmShipmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodCollectionMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var confirmShipmentRequestTypeCodCollectionMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DirectPayment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		confirmShipmentRequestTypeCodCollectionMethodPropEnum = append(confirmShipmentRequestTypeCodCollectionMethodPropEnum, v)
	}
}

const (

	// ConfirmShipmentRequestCodCollectionMethodDirectPayment captures enum value "DirectPayment"
	ConfirmShipmentRequestCodCollectionMethodDirectPayment string = "DirectPayment"
)

// prop value enum
func (m *ConfirmShipmentRequest) validateCodCollectionMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, confirmShipmentRequestTypeCodCollectionMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConfirmShipmentRequest) validateCodCollectionMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.CodCollectionMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodCollectionMethodEnum("codCollectionMethod", "body", m.CodCollectionMethod); err != nil {
		return err
	}

	return nil
}

func (m *ConfirmShipmentRequest) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketplaceId")
			}
			return err
		}
	}

	return nil
}

func (m *ConfirmShipmentRequest) validatePackageDetail(formats strfmt.Registry) error {

	if err := validate.Required("packageDetail", "body", m.PackageDetail); err != nil {
		return err
	}

	if m.PackageDetail != nil {
		if err := m.PackageDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageDetail")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this confirm shipment request based on the context it is used
func (m *ConfirmShipmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMarketplaceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackageDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfirmShipmentRequest) contextValidateMarketplaceID(ctx context.Context, formats strfmt.Registry) error {

	if m.MarketplaceID != nil {
		if err := m.MarketplaceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketplaceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketplaceId")
			}
			return err
		}
	}

	return nil
}

func (m *ConfirmShipmentRequest) contextValidatePackageDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.PackageDetail != nil {
		if err := m.PackageDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmShipmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmShipmentRequest) UnmarshalBinary(b []byte) error {
	var res ConfirmShipmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
