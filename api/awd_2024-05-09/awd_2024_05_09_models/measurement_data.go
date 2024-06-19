// Code generated by go-swagger; DO NOT EDIT.

package awd_2024_05_09_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MeasurementData Package weight and dimension.
//
// swagger:model MeasurementData
type MeasurementData struct {

	// Dimensions of the package. Dimensions are required when creating an inbound or outbound order.
	Dimensions *PackageDimensions `json:"dimensions,omitempty"`

	// Volume of the package.
	Volume *PackageVolume `json:"volume,omitempty"`

	// Weight of the package.
	// Required: true
	Weight *PackageWeight `json:"weight"`
}

// Validate validates this measurement data
func (m *MeasurementData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MeasurementData) validateDimensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Dimensions) { // not required
		return nil
	}

	if m.Dimensions != nil {
		if err := m.Dimensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *MeasurementData) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

func (m *MeasurementData) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	if m.Weight != nil {
		if err := m.Weight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this measurement data based on the context it is used
func (m *MeasurementData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MeasurementData) contextValidateDimensions(ctx context.Context, formats strfmt.Registry) error {

	if m.Dimensions != nil {
		if err := m.Dimensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dimensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dimensions")
			}
			return err
		}
	}

	return nil
}

func (m *MeasurementData) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

func (m *MeasurementData) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Weight != nil {
		if err := m.Weight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weight")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MeasurementData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeasurementData) UnmarshalBinary(b []byte) error {
	var res MeasurementData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
