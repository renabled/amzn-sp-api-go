// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetShipmentsResult Result for the get shipments operation
//
// swagger:model GetShipmentsResult
type GetShipmentsResult struct {

	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken string `json:"NextToken,omitempty"`

	// Information about your inbound shipments.
	ShipmentData InboundShipmentList `json:"ShipmentData,omitempty"`
}

// Validate validates this get shipments result
func (m *GetShipmentsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipmentData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetShipmentsResult) validateShipmentData(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentData) { // not required
		return nil
	}

	if err := m.ShipmentData.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentData")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentData")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get shipments result based on the context it is used
func (m *GetShipmentsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShipmentData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetShipmentsResult) contextValidateShipmentData(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipmentData.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentData")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentData")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetShipmentsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetShipmentsResult) UnmarshalBinary(b []byte) error {
	var res GetShipmentsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
