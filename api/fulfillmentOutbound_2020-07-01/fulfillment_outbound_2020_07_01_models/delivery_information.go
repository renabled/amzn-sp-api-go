// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeliveryInformation The delivery information for the package. This information is available after the package is delivered.
//
// swagger:model DeliveryInformation
type DeliveryInformation struct {

	// All of the delivery documents for a package.
	DeliveryDocumentList DeliveryDocumentList `json:"deliveryDocumentList,omitempty"`

	// The drop off location for a package.
	DropOffLocation *DropOffLocation `json:"dropOffLocation,omitempty"`
}

// Validate validates this delivery information
func (m *DeliveryInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeliveryDocumentList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDropOffLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryInformation) validateDeliveryDocumentList(formats strfmt.Registry) error {
	if swag.IsZero(m.DeliveryDocumentList) { // not required
		return nil
	}

	if err := m.DeliveryDocumentList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deliveryDocumentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deliveryDocumentList")
		}
		return err
	}

	return nil
}

func (m *DeliveryInformation) validateDropOffLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.DropOffLocation) { // not required
		return nil
	}

	if m.DropOffLocation != nil {
		if err := m.DropOffLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dropOffLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dropOffLocation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this delivery information based on the context it is used
func (m *DeliveryInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeliveryDocumentList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDropOffLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeliveryInformation) contextValidateDeliveryDocumentList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DeliveryDocumentList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deliveryDocumentList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("deliveryDocumentList")
		}
		return err
	}

	return nil
}

func (m *DeliveryInformation) contextValidateDropOffLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.DropOffLocation != nil {
		if err := m.DropOffLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dropOffLocation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dropOffLocation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeliveryInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeliveryInformation) UnmarshalBinary(b []byte) error {
	var res DeliveryInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
