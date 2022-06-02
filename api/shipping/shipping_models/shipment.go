// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Shipment The shipment related data.
//
// swagger:model Shipment
type Shipment struct {

	// accepted rate
	AcceptedRate *AcceptedRate `json:"acceptedRate,omitempty"`

	// client reference Id
	// Required: true
	ClientReferenceID *ClientReferenceID `json:"clientReferenceId"`

	// containers
	// Required: true
	Containers ContainerList `json:"containers"`

	// ship from
	// Required: true
	ShipFrom *Address `json:"shipFrom"`

	// ship to
	// Required: true
	ShipTo *Address `json:"shipTo"`

	// shipment Id
	// Required: true
	ShipmentID *ShipmentID `json:"shipmentId"`

	// shipper
	Shipper *Party `json:"shipper,omitempty"`
}

// Validate validates this shipment
func (m *Shipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientReferenceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipper(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Shipment) validateAcceptedRate(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptedRate) { // not required
		return nil
	}

	if m.AcceptedRate != nil {
		if err := m.AcceptedRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptedRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptedRate")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) validateClientReferenceID(formats strfmt.Registry) error {

	if err := validate.Required("clientReferenceId", "body", m.ClientReferenceID); err != nil {
		return err
	}

	if err := validate.Required("clientReferenceId", "body", m.ClientReferenceID); err != nil {
		return err
	}

	if m.ClientReferenceID != nil {
		if err := m.ClientReferenceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clientReferenceId")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("containers", "body", m.Containers); err != nil {
		return err
	}

	if err := m.Containers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("containers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("containers")
		}
		return err
	}

	return nil
}

func (m *Shipment) validateShipFrom(formats strfmt.Registry) error {

	if err := validate.Required("shipFrom", "body", m.ShipFrom); err != nil {
		return err
	}

	if m.ShipFrom != nil {
		if err := m.ShipFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFrom")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) validateShipTo(formats strfmt.Registry) error {

	if err := validate.Required("shipTo", "body", m.ShipTo); err != nil {
		return err
	}

	if m.ShipTo != nil {
		if err := m.ShipTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipTo")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) validateShipmentID(formats strfmt.Registry) error {

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

func (m *Shipment) validateShipper(formats strfmt.Registry) error {
	if swag.IsZero(m.Shipper) { // not required
		return nil
	}

	if m.Shipper != nil {
		if err := m.Shipper.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipper")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipper")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this shipment based on the context it is used
func (m *Shipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptedRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientReferenceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipper(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Shipment) contextValidateAcceptedRate(ctx context.Context, formats strfmt.Registry) error {

	if m.AcceptedRate != nil {
		if err := m.AcceptedRate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptedRate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptedRate")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) contextValidateClientReferenceID(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientReferenceID != nil {
		if err := m.ClientReferenceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientReferenceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clientReferenceId")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) contextValidateContainers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Containers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("containers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("containers")
		}
		return err
	}

	return nil
}

func (m *Shipment) contextValidateShipFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipFrom != nil {
		if err := m.ShipFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFrom")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) contextValidateShipTo(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipTo != nil {
		if err := m.ShipTo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipTo")
			}
			return err
		}
	}

	return nil
}

func (m *Shipment) contextValidateShipmentID(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Shipment) contextValidateShipper(ctx context.Context, formats strfmt.Registry) error {

	if m.Shipper != nil {
		if err := m.Shipper.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipper")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipper")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Shipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Shipment) UnmarshalBinary(b []byte) error {
	var res Shipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
