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

// GetRatesRequest The request schema for the getRates operation. When the channelType is Amazon, the shipTo address is not required and will be ignored.
//
// swagger:model GetRatesRequest
type GetRatesRequest struct {

	// channel details
	// Required: true
	ChannelDetails *ChannelDetails `json:"channelDetails"`

	// client reference details
	ClientReferenceDetails ClientReferenceDetails `json:"clientReferenceDetails,omitempty"`

	// destination access point details
	DestinationAccessPointDetails *AccessPointDetails `json:"destinationAccessPointDetails,omitempty"`

	// packages
	// Required: true
	Packages PackageList `json:"packages"`

	// The return to address.
	ReturnTo *Address `json:"returnTo,omitempty"`

	// The ship date and time (the requested pickup). This defaults to the current date and time.
	// Format: date-time
	ShipDate strfmt.DateTime `json:"shipDate,omitempty"`

	// The ship from address.
	// Required: true
	ShipFrom *Address `json:"shipFrom"`

	// The ship to address.
	ShipTo *Address `json:"shipTo,omitempty"`

	// shipment type
	ShipmentType ShipmentType `json:"shipmentType,omitempty"`

	// This field describe shipper instruction.
	ShipperInstruction *ShipperInstruction `json:"shipperInstruction,omitempty"`

	// tax details
	TaxDetails TaxDetailList `json:"taxDetails,omitempty"`

	// value added services
	ValueAddedServices *ValueAddedServiceDetails `json:"valueAddedServices,omitempty"`
}

// Validate validates this get rates request
func (m *GetRatesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientReferenceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAccessPointDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipperInstruction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueAddedServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetRatesRequest) validateChannelDetails(formats strfmt.Registry) error {

	if err := validate.Required("channelDetails", "body", m.ChannelDetails); err != nil {
		return err
	}

	if m.ChannelDetails != nil {
		if err := m.ChannelDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channelDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channelDetails")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) validateClientReferenceDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientReferenceDetails) { // not required
		return nil
	}

	if err := m.ClientReferenceDetails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("clientReferenceDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("clientReferenceDetails")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) validateDestinationAccessPointDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationAccessPointDetails) { // not required
		return nil
	}

	if m.DestinationAccessPointDetails != nil {
		if err := m.DestinationAccessPointDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAccessPointDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAccessPointDetails")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) validatePackages(formats strfmt.Registry) error {

	if err := validate.Required("packages", "body", m.Packages); err != nil {
		return err
	}

	if err := m.Packages.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packages")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("packages")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) validateReturnTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ReturnTo) { // not required
		return nil
	}

	if m.ReturnTo != nil {
		if err := m.ReturnTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("returnTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("returnTo")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) validateShipDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipDate) { // not required
		return nil
	}

	if err := validate.FormatOf("shipDate", "body", "date-time", m.ShipDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetRatesRequest) validateShipFrom(formats strfmt.Registry) error {

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

func (m *GetRatesRequest) validateShipTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipTo) { // not required
		return nil
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

func (m *GetRatesRequest) validateShipmentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentType) { // not required
		return nil
	}

	if err := m.ShipmentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) validateShipperInstruction(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipperInstruction) { // not required
		return nil
	}

	if m.ShipperInstruction != nil {
		if err := m.ShipperInstruction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipperInstruction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipperInstruction")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) validateTaxDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxDetails) { // not required
		return nil
	}

	if err := m.TaxDetails.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("taxDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("taxDetails")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) validateValueAddedServices(formats strfmt.Registry) error {
	if swag.IsZero(m.ValueAddedServices) { // not required
		return nil
	}

	if m.ValueAddedServices != nil {
		if err := m.ValueAddedServices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueAddedServices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueAddedServices")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get rates request based on the context it is used
func (m *GetRatesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannelDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientReferenceDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDestinationAccessPointDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReturnTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipperInstruction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValueAddedServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetRatesRequest) contextValidateChannelDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ChannelDetails != nil {
		if err := m.ChannelDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channelDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channelDetails")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) contextValidateClientReferenceDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ClientReferenceDetails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("clientReferenceDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("clientReferenceDetails")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) contextValidateDestinationAccessPointDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.DestinationAccessPointDetails != nil {
		if err := m.DestinationAccessPointDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAccessPointDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("destinationAccessPointDetails")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) contextValidatePackages(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Packages.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("packages")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("packages")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) contextValidateReturnTo(ctx context.Context, formats strfmt.Registry) error {

	if m.ReturnTo != nil {
		if err := m.ReturnTo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("returnTo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("returnTo")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) contextValidateShipFrom(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetRatesRequest) contextValidateShipTo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetRatesRequest) contextValidateShipmentType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipmentType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shipmentType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shipmentType")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) contextValidateShipperInstruction(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipperInstruction != nil {
		if err := m.ShipperInstruction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipperInstruction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipperInstruction")
			}
			return err
		}
	}

	return nil
}

func (m *GetRatesRequest) contextValidateTaxDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TaxDetails.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("taxDetails")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("taxDetails")
		}
		return err
	}

	return nil
}

func (m *GetRatesRequest) contextValidateValueAddedServices(ctx context.Context, formats strfmt.Registry) error {

	if m.ValueAddedServices != nil {
		if err := m.ValueAddedServices.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueAddedServices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueAddedServices")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetRatesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRatesRequest) UnmarshalBinary(b []byte) error {
	var res GetRatesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}