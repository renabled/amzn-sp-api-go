// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateShipmentRequest Request schema.
//
// swagger:model CreateShipmentRequest
type CreateShipmentRequest struct {

	// Hazardous materials options for a package. Consult the terms and conditions for each carrier for more information about hazardous materials.
	HazmatType HazmatType `json:"HazmatType,omitempty"`

	// label format option
	LabelFormatOption *LabelFormatOptionRequest `json:"LabelFormatOption,omitempty"`

	// A list of additional seller inputs required to ship this shipment.
	ShipmentLevelSellerInputsList AdditionalSellerInputsList `json:"ShipmentLevelSellerInputsList,omitempty"`

	// Shipment information required to create a shipment.
	// Required: true
	ShipmentRequestDetails *ShipmentRequestDetails `json:"ShipmentRequestDetails"`

	// shipping service Id
	// Required: true
	ShippingServiceID *ShippingServiceIdentifier `json:"ShippingServiceId"`

	// Identifies a shipping service order made by a carrier.
	ShippingServiceOfferID string `json:"ShippingServiceOfferId,omitempty"`
}

// Validate validates this create shipment request
func (m *CreateShipmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHazmatType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelFormatOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentLevelSellerInputsList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentRequestDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShipmentRequest) validateHazmatType(formats strfmt.Registry) error {
	if swag.IsZero(m.HazmatType) { // not required
		return nil
	}

	if err := m.HazmatType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("HazmatType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("HazmatType")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentRequest) validateLabelFormatOption(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelFormatOption) { // not required
		return nil
	}

	if m.LabelFormatOption != nil {
		if err := m.LabelFormatOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LabelFormatOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LabelFormatOption")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) validateShipmentLevelSellerInputsList(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentLevelSellerInputsList) { // not required
		return nil
	}

	if err := m.ShipmentLevelSellerInputsList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentLevelSellerInputsList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentLevelSellerInputsList")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentRequest) validateShipmentRequestDetails(formats strfmt.Registry) error {

	if err := validate.Required("ShipmentRequestDetails", "body", m.ShipmentRequestDetails); err != nil {
		return err
	}

	if m.ShipmentRequestDetails != nil {
		if err := m.ShipmentRequestDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipmentRequestDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipmentRequestDetails")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) validateShippingServiceID(formats strfmt.Registry) error {

	if err := validate.Required("ShippingServiceId", "body", m.ShippingServiceID); err != nil {
		return err
	}

	if err := validate.Required("ShippingServiceId", "body", m.ShippingServiceID); err != nil {
		return err
	}

	if m.ShippingServiceID != nil {
		if err := m.ShippingServiceID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingServiceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingServiceId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create shipment request based on the context it is used
func (m *CreateShipmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHazmatType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelFormatOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentLevelSellerInputsList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentRequestDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShippingServiceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateShipmentRequest) contextValidateHazmatType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HazmatType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("HazmatType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("HazmatType")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentRequest) contextValidateLabelFormatOption(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelFormatOption != nil {
		if err := m.LabelFormatOption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LabelFormatOption")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LabelFormatOption")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) contextValidateShipmentLevelSellerInputsList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShipmentLevelSellerInputsList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ShipmentLevelSellerInputsList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ShipmentLevelSellerInputsList")
		}
		return err
	}

	return nil
}

func (m *CreateShipmentRequest) contextValidateShipmentRequestDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipmentRequestDetails != nil {
		if err := m.ShipmentRequestDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShipmentRequestDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShipmentRequestDetails")
			}
			return err
		}
	}

	return nil
}

func (m *CreateShipmentRequest) contextValidateShippingServiceID(ctx context.Context, formats strfmt.Registry) error {

	if m.ShippingServiceID != nil {
		if err := m.ShippingServiceID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ShippingServiceId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ShippingServiceId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateShipmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShipmentRequest) UnmarshalBinary(b []byte) error {
	var res CreateShipmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
