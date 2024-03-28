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

// TransportationDetailsForShipmentConfirmation Transportation details for this shipment.
//
// swagger:model TransportationDetailsForShipmentConfirmation
type TransportationDetailsForShipmentConfirmation struct {

	// The Bill of Lading (BOL) number is a unique number assigned to each shipment of goods by the vendor or shipper during the creation of the Bill of Lading. This number must be unique for every shipment and cannot be a date/time or single character. The BOL numer is mandatory in Shipment Confirmation message for FTL and LTL shipments, and must match the paper BOL provided with the shipment. Instead of BOL, an alternative reference number (like Delivery Note Number) for the shipment can also be sent in this field.
	BillOfLadingNumber string `json:"billOfLadingNumber,omitempty"`

	// Code that identifies the carrier for the shipment. The Standard Carrier Alpha Code (SCAC) is a unique two to four letter code used to identify a carrier. Carrier SCAC codes are assigned and maintained by the NMFTA (National Motor Freight Association). This field is mandatory for US, CA, MX shipment confirmations.
	CarrierScac string `json:"carrierScac,omitempty"`

	// The field also known as PRO number is a unique number assigned by the carrier. It is used to identify and track the shipment that goes out for delivery. This field is mandatory for UA, CA, MX shipment confirmations.
	CarrierShipmentReferenceNumber string `json:"carrierShipmentReferenceNumber,omitempty"`

	// The mode of transportation for this shipment.
	// Enum: [Road Air Ocean]
	TransportationMode string `json:"transportationMode,omitempty"`
}

// Validate validates this transportation details for shipment confirmation
func (m *TransportationDetailsForShipmentConfirmation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransportationMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var transportationDetailsForShipmentConfirmationTypeTransportationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Road","Air","Ocean"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transportationDetailsForShipmentConfirmationTypeTransportationModePropEnum = append(transportationDetailsForShipmentConfirmationTypeTransportationModePropEnum, v)
	}
}

const (

	// TransportationDetailsForShipmentConfirmationTransportationModeRoad captures enum value "Road"
	TransportationDetailsForShipmentConfirmationTransportationModeRoad string = "Road"

	// TransportationDetailsForShipmentConfirmationTransportationModeAir captures enum value "Air"
	TransportationDetailsForShipmentConfirmationTransportationModeAir string = "Air"

	// TransportationDetailsForShipmentConfirmationTransportationModeOcean captures enum value "Ocean"
	TransportationDetailsForShipmentConfirmationTransportationModeOcean string = "Ocean"
)

// prop value enum
func (m *TransportationDetailsForShipmentConfirmation) validateTransportationModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transportationDetailsForShipmentConfirmationTypeTransportationModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransportationDetailsForShipmentConfirmation) validateTransportationMode(formats strfmt.Registry) error {
	if swag.IsZero(m.TransportationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransportationModeEnum("transportationMode", "body", m.TransportationMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transportation details for shipment confirmation based on context it is used
func (m *TransportationDetailsForShipmentConfirmation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransportationDetailsForShipmentConfirmation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransportationDetailsForShipmentConfirmation) UnmarshalBinary(b []byte) error {
	var res TransportationDetailsForShipmentConfirmation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}