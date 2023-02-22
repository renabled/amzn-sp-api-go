// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipments_models

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

// TransportationDetails transportation details
//
// swagger:model TransportationDetails
type TransportationDetails struct {

	// Bill Of Lading (BOL) number is the unique number assigned by the vendor. The BOL present in the Shipment Confirmation message ideally matches the paper BOL provided with the shipment, but that is no must. Instead of BOL, an alternative reference number (like Delivery Note Number) for the shipment can also be sent in this field.
	BillOfLadingNumber string `json:"billOfLadingNumber,omitempty"`

	// Indicates the carrier details and their contact informations
	CarrierDetails *CarrierDetails `json:"carrierDetails,omitempty"`

	// Estimated Date on which shipment will be delivered from Vendor to Buyer
	// Format: date-time
	EstimatedDeliveryDate strfmt.DateTime `json:"estimatedDeliveryDate,omitempty"`

	// The type of shipment.
	// Enum: [TruckLoad LessThanTruckLoad SmallParcel]
	ShipMode string `json:"shipMode,omitempty"`

	// Date on which shipment will be delivered from Vendor to Buyer
	// Format: date-time
	ShipmentDeliveryDate strfmt.DateTime `json:"shipmentDeliveryDate,omitempty"`

	// Date when shipment is performed by the Vendor to Buyer
	// Format: date-time
	ShippedDate strfmt.DateTime `json:"shippedDate,omitempty"`

	// The mode of transportation for this shipment.
	// Enum: [Road Air Ocean]
	TransportationMode string `json:"transportationMode,omitempty"`
}

// Validate validates this transportation details
func (m *TransportationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimatedDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentDeliveryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransportationMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportationDetails) validateCarrierDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.CarrierDetails) { // not required
		return nil
	}

	if m.CarrierDetails != nil {
		if err := m.CarrierDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("carrierDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("carrierDetails")
			}
			return err
		}
	}

	return nil
}

func (m *TransportationDetails) validateEstimatedDeliveryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EstimatedDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("estimatedDeliveryDate", "body", "date-time", m.EstimatedDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var transportationDetailsTypeShipModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TruckLoad","LessThanTruckLoad","SmallParcel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transportationDetailsTypeShipModePropEnum = append(transportationDetailsTypeShipModePropEnum, v)
	}
}

const (

	// TransportationDetailsShipModeTruckLoad captures enum value "TruckLoad"
	TransportationDetailsShipModeTruckLoad string = "TruckLoad"

	// TransportationDetailsShipModeLessThanTruckLoad captures enum value "LessThanTruckLoad"
	TransportationDetailsShipModeLessThanTruckLoad string = "LessThanTruckLoad"

	// TransportationDetailsShipModeSmallParcel captures enum value "SmallParcel"
	TransportationDetailsShipModeSmallParcel string = "SmallParcel"
)

// prop value enum
func (m *TransportationDetails) validateShipModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transportationDetailsTypeShipModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransportationDetails) validateShipMode(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateShipModeEnum("shipMode", "body", m.ShipMode); err != nil {
		return err
	}

	return nil
}

func (m *TransportationDetails) validateShipmentDeliveryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentDeliveryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("shipmentDeliveryDate", "body", "date-time", m.ShipmentDeliveryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransportationDetails) validateShippedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ShippedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("shippedDate", "body", "date-time", m.ShippedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var transportationDetailsTypeTransportationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Road","Air","Ocean"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transportationDetailsTypeTransportationModePropEnum = append(transportationDetailsTypeTransportationModePropEnum, v)
	}
}

const (

	// TransportationDetailsTransportationModeRoad captures enum value "Road"
	TransportationDetailsTransportationModeRoad string = "Road"

	// TransportationDetailsTransportationModeAir captures enum value "Air"
	TransportationDetailsTransportationModeAir string = "Air"

	// TransportationDetailsTransportationModeOcean captures enum value "Ocean"
	TransportationDetailsTransportationModeOcean string = "Ocean"
)

// prop value enum
func (m *TransportationDetails) validateTransportationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transportationDetailsTypeTransportationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransportationDetails) validateTransportationMode(formats strfmt.Registry) error {
	if swag.IsZero(m.TransportationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransportationModeEnum("transportationMode", "body", m.TransportationMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this transportation details based on the context it is used
func (m *TransportationDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCarrierDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransportationDetails) contextValidateCarrierDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.CarrierDetails != nil {
		if err := m.CarrierDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("carrierDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("carrierDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransportationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportationDetails) UnmarshalBinary(b []byte) error {
	var res TransportationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
