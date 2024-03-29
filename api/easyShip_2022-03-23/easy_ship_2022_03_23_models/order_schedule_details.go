// Code generated by go-swagger; DO NOT EDIT.

package easy_ship_2022_03_23_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OrderScheduleDetails This object allows users to specify an order to be scheduled. Only the amazonOrderId is required.
//
// swagger:model OrderScheduleDetails
type OrderScheduleDetails struct {

	// amazon order Id
	// Required: true
	AmazonOrderID *AmazonOrderID `json:"amazonOrderId"`

	// package details
	PackageDetails *PackageDetails `json:"packageDetails,omitempty"`
}

// Validate validates this order schedule details
func (m *OrderScheduleDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazonOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderScheduleDetails) validateAmazonOrderID(formats strfmt.Registry) error {

	if err := validate.Required("amazonOrderId", "body", m.AmazonOrderID); err != nil {
		return err
	}

	if err := validate.Required("amazonOrderId", "body", m.AmazonOrderID); err != nil {
		return err
	}

	if m.AmazonOrderID != nil {
		if err := m.AmazonOrderID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amazonOrderId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amazonOrderId")
			}
			return err
		}
	}

	return nil
}

func (m *OrderScheduleDetails) validatePackageDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.PackageDetails) { // not required
		return nil
	}

	if m.PackageDetails != nil {
		if err := m.PackageDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this order schedule details based on the context it is used
func (m *OrderScheduleDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmazonOrderID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackageDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderScheduleDetails) contextValidateAmazonOrderID(ctx context.Context, formats strfmt.Registry) error {

	if m.AmazonOrderID != nil {
		if err := m.AmazonOrderID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amazonOrderId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amazonOrderId")
			}
			return err
		}
	}

	return nil
}

func (m *OrderScheduleDetails) contextValidatePackageDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.PackageDetails != nil {
		if err := m.PackageDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packageDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("packageDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderScheduleDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderScheduleDetails) UnmarshalBinary(b []byte) error {
	var res OrderScheduleDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
