// Code generated by go-swagger; DO NOT EDIT.

package awd_2024_05_09_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShipmentListing A list of inbound shipment summaries filtered by the attributes specified in the request.
//
// swagger:model ShipmentListing
type ShipmentListing struct {

	// Token to retrieve the next set of paginated results.
	// Example: SampleToken
	NextToken string `json:"nextToken,omitempty"`

	// List of inbound shipment summaries.
	Shipments []*InboundShipmentSummary `json:"shipments"`
}

// Validate validates this shipment listing
func (m *ShipmentListing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentListing) validateShipments(formats strfmt.Registry) error {
	if swag.IsZero(m.Shipments) { // not required
		return nil
	}

	for i := 0; i < len(m.Shipments); i++ {
		if swag.IsZero(m.Shipments[i]) { // not required
			continue
		}

		if m.Shipments[i] != nil {
			if err := m.Shipments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this shipment listing based on the context it is used
func (m *ShipmentListing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShipments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShipmentListing) contextValidateShipments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Shipments); i++ {

		if m.Shipments[i] != nil {
			if err := m.Shipments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShipmentListing) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShipmentListing) UnmarshalBinary(b []byte) error {
	var res ShipmentListing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
