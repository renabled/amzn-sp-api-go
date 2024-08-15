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
	"github.com/go-openapi/validate"
)

// InboundShipment Represents an AWD inbound shipment.
//
// swagger:model InboundShipment
type InboundShipment struct {

	// The shipment carrier code.
	CarrierCode *CarrierCode `json:"carrierCode,omitempty"`

	// Timestamp when the shipment was created. The date is returned in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Destination address for this shipment.
	// Required: true
	DestinationAddress *Address `json:"destinationAddress"`

	// Client-provided reference ID that can correlate this shipment to client resources. For example, to map this shipment to an internal bookkeeping order record.
	// Example: TestReferenceId
	ExternalReferenceID string `json:"externalReferenceId,omitempty"`

	// The AWD inbound order ID that this inbound shipment belongs to.
	// Required: true
	// Min Length: 1
	OrderID *string `json:"orderId"`

	// Origin address for this shipment.
	// Required: true
	OriginAddress *Address `json:"originAddress"`

	// Quantity received (at the receiving end) as part of this shipment.
	ReceivedQuantity []*InventoryQuantity `json:"receivedQuantity"`

	// Timestamp when the shipment will be shipped.
	// Format: date-time
	ShipBy strfmt.DateTime `json:"shipBy,omitempty"`

	// Packages that are part of this shipment.
	// Required: true
	ShipmentContainerQuantities []*DistributionPackageQuantity `json:"shipmentContainerQuantities"`

	// Unique shipment ID.
	// Required: true
	// Min Length: 1
	ShipmentID *string `json:"shipmentId"`

	// Quantity details at SKU level for the shipment. This attribute will only appear if the skuQuantities parameter in the request is set to SHOW.
	ShipmentSkuQuantities []*SkuQuantity `json:"shipmentSkuQuantities"`

	// Current status of this shipment.
	// Required: true
	ShipmentStatus *InboundShipmentStatus `json:"shipmentStatus"`

	// Carrier-unique tracking ID for this shipment.
	// Min Length: 1
	TrackingID string `json:"trackingId,omitempty"`

	// Timestamp when the shipment was updated. The date is returned in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// An AWD-provided reference ID that you can use to interact with the warehouse. For example, a carrier appointment booking.
	// Example: TestWarehouseReferenceId
	WarehouseReferenceID string `json:"warehouseReferenceId,omitempty"`
}

// Validate validates this inbound shipment
func (m *InboundShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCarrierCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentContainerQuantities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentSkuQuantities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundShipment) validateCarrierCode(formats strfmt.Registry) error {
	if swag.IsZero(m.CarrierCode) { // not required
		return nil
	}

	if m.CarrierCode != nil {
		if err := m.CarrierCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("carrierCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("carrierCode")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipment) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InboundShipment) validateDestinationAddress(formats strfmt.Registry) error {

	if err := validate.Required("destinationAddress", "body", m.DestinationAddress); err != nil {
		return err
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipment) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("orderId", "body", m.OrderID); err != nil {
		return err
	}

	if err := validate.MinLength("orderId", "body", *m.OrderID, 1); err != nil {
		return err
	}

	return nil
}

func (m *InboundShipment) validateOriginAddress(formats strfmt.Registry) error {

	if err := validate.Required("originAddress", "body", m.OriginAddress); err != nil {
		return err
	}

	if m.OriginAddress != nil {
		if err := m.OriginAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originAddress")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipment) validateReceivedQuantity(formats strfmt.Registry) error {
	if swag.IsZero(m.ReceivedQuantity) { // not required
		return nil
	}

	for i := 0; i < len(m.ReceivedQuantity); i++ {
		if swag.IsZero(m.ReceivedQuantity[i]) { // not required
			continue
		}

		if m.ReceivedQuantity[i] != nil {
			if err := m.ReceivedQuantity[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivedQuantity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivedQuantity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundShipment) validateShipBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipBy) { // not required
		return nil
	}

	if err := validate.FormatOf("shipBy", "body", "date-time", m.ShipBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InboundShipment) validateShipmentContainerQuantities(formats strfmt.Registry) error {

	if err := validate.Required("shipmentContainerQuantities", "body", m.ShipmentContainerQuantities); err != nil {
		return err
	}

	for i := 0; i < len(m.ShipmentContainerQuantities); i++ {
		if swag.IsZero(m.ShipmentContainerQuantities[i]) { // not required
			continue
		}

		if m.ShipmentContainerQuantities[i] != nil {
			if err := m.ShipmentContainerQuantities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipmentContainerQuantities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipmentContainerQuantities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundShipment) validateShipmentID(formats strfmt.Registry) error {

	if err := validate.Required("shipmentId", "body", m.ShipmentID); err != nil {
		return err
	}

	if err := validate.MinLength("shipmentId", "body", *m.ShipmentID, 1); err != nil {
		return err
	}

	return nil
}

func (m *InboundShipment) validateShipmentSkuQuantities(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentSkuQuantities) { // not required
		return nil
	}

	for i := 0; i < len(m.ShipmentSkuQuantities); i++ {
		if swag.IsZero(m.ShipmentSkuQuantities[i]) { // not required
			continue
		}

		if m.ShipmentSkuQuantities[i] != nil {
			if err := m.ShipmentSkuQuantities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipmentSkuQuantities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipmentSkuQuantities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundShipment) validateShipmentStatus(formats strfmt.Registry) error {

	if err := validate.Required("shipmentStatus", "body", m.ShipmentStatus); err != nil {
		return err
	}

	if err := validate.Required("shipmentStatus", "body", m.ShipmentStatus); err != nil {
		return err
	}

	if m.ShipmentStatus != nil {
		if err := m.ShipmentStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentStatus")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipment) validateTrackingID(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackingID) { // not required
		return nil
	}

	if err := validate.MinLength("trackingId", "body", m.TrackingID, 1); err != nil {
		return err
	}

	return nil
}

func (m *InboundShipment) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this inbound shipment based on the context it is used
func (m *InboundShipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCarrierCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestinationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReceivedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentContainerQuantities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentSkuQuantities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InboundShipment) contextValidateCarrierCode(ctx context.Context, formats strfmt.Registry) error {

	if m.CarrierCode != nil {
		if err := m.CarrierCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("carrierCode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("carrierCode")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipment) contextValidateDestinationAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipment) contextValidateOriginAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginAddress != nil {
		if err := m.OriginAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originAddress")
			}
			return err
		}
	}

	return nil
}

func (m *InboundShipment) contextValidateReceivedQuantity(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReceivedQuantity); i++ {

		if m.ReceivedQuantity[i] != nil {
			if err := m.ReceivedQuantity[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("receivedQuantity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("receivedQuantity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundShipment) contextValidateShipmentContainerQuantities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ShipmentContainerQuantities); i++ {

		if m.ShipmentContainerQuantities[i] != nil {
			if err := m.ShipmentContainerQuantities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipmentContainerQuantities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipmentContainerQuantities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundShipment) contextValidateShipmentSkuQuantities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ShipmentSkuQuantities); i++ {

		if m.ShipmentSkuQuantities[i] != nil {
			if err := m.ShipmentSkuQuantities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipmentSkuQuantities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shipmentSkuQuantities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InboundShipment) contextValidateShipmentStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentStatus != nil {
		if err := m.ShipmentStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipmentStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipmentStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InboundShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InboundShipment) UnmarshalBinary(b []byte) error {
	var res InboundShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
