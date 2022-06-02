// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DimensionType The dimension type attribute of an item.
//
// swagger:model DimensionType
type DimensionType struct {

	// The height attribute of the dimension.
	Height *DecimalWithUnits `json:"Height,omitempty"`

	// The length attribute of the dimension.
	Length *DecimalWithUnits `json:"Length,omitempty"`

	// The weight attribute of the dimension.
	Weight *DecimalWithUnits `json:"Weight,omitempty"`

	// The width attribute of the dimension.
	Width *DecimalWithUnits `json:"Width,omitempty"`
}

// Validate validates this dimension type
func (m *DimensionType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DimensionType) validateHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Height) { // not required
		return nil
	}

	if m.Height != nil {
		if err := m.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Height")
			}
			return err
		}
	}

	return nil
}

func (m *DimensionType) validateLength(formats strfmt.Registry) error {
	if swag.IsZero(m.Length) { // not required
		return nil
	}

	if m.Length != nil {
		if err := m.Length.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Length")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Length")
			}
			return err
		}
	}

	return nil
}

func (m *DimensionType) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if m.Weight != nil {
		if err := m.Weight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Weight")
			}
			return err
		}
	}

	return nil
}

func (m *DimensionType) validateWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.Width) { // not required
		return nil
	}

	if m.Width != nil {
		if err := m.Width.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Width")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dimension type based on the context it is used
func (m *DimensionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLength(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DimensionType) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Height != nil {
		if err := m.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Height")
			}
			return err
		}
	}

	return nil
}

func (m *DimensionType) contextValidateLength(ctx context.Context, formats strfmt.Registry) error {

	if m.Length != nil {
		if err := m.Length.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Length")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Length")
			}
			return err
		}
	}

	return nil
}

func (m *DimensionType) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Weight != nil {
		if err := m.Weight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Weight")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Weight")
			}
			return err
		}
	}

	return nil
}

func (m *DimensionType) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if m.Width != nil {
		if err := m.Width.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Width")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Width")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DimensionType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DimensionType) UnmarshalBinary(b []byte) error {
	var res DimensionType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
