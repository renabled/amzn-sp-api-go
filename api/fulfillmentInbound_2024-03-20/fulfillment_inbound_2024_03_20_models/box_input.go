// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

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

// BoxInput Input information for a given box.
// Example: {"contentInformationSource":"BOX_CONTENT_PROVIDED","dimensions":{"height":5,"length":3,"unitOfMeasurement":"CM","width":4},"items":[{"expiration":"2024-01-01","labelOwner":"AMAZON","manufacturingLotCode":"manufacturingLotCode","msku":"Sunglasses","prepOwner":"AMAZON","quantity":10}],"quantity":2,"weight":{"unit":"KG","value":5.5}}
//
// swagger:model BoxInput
type BoxInput struct {

	// content information source
	// Required: true
	ContentInformationSource *BoxContentInformationSource `json:"contentInformationSource"`

	// dimensions
	// Required: true
	Dimensions *Dimensions `json:"dimensions"`

	// The items and their quantity in the box. This must be empty if the box `contentInformationSource` is `BARCODE_2D` or `MANUAL_PROCESS`.
	Items []*ItemInput `json:"items"`

	// The number of containers where all other properties like weight or dimensions are identical.
	// Required: true
	// Maximum: 10000
	// Minimum: 1
	Quantity *int64 `json:"quantity"`

	// weight
	// Required: true
	Weight *Weight `json:"weight"`
}

// Validate validates this box input
func (m *BoxInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentInformationSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
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

func (m *BoxInput) validateContentInformationSource(formats strfmt.Registry) error {

	if err := validate.Required("contentInformationSource", "body", m.ContentInformationSource); err != nil {
		return err
	}

	if err := validate.Required("contentInformationSource", "body", m.ContentInformationSource); err != nil {
		return err
	}

	if m.ContentInformationSource != nil {
		if err := m.ContentInformationSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentInformationSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentInformationSource")
			}
			return err
		}
	}

	return nil
}

func (m *BoxInput) validateDimensions(formats strfmt.Registry) error {

	if err := validate.Required("dimensions", "body", m.Dimensions); err != nil {
		return err
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

func (m *BoxInput) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BoxInput) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.MinimumInt("quantity", "body", *m.Quantity, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("quantity", "body", *m.Quantity, 10000, false); err != nil {
		return err
	}

	return nil
}

func (m *BoxInput) validateWeight(formats strfmt.Registry) error {

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

// ContextValidate validate this box input based on the context it is used
func (m *BoxInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContentInformationSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
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

func (m *BoxInput) contextValidateContentInformationSource(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentInformationSource != nil {
		if err := m.ContentInformationSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contentInformationSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("contentInformationSource")
			}
			return err
		}
	}

	return nil
}

func (m *BoxInput) contextValidateDimensions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BoxInput) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BoxInput) contextValidateWeight(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BoxInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BoxInput) UnmarshalBinary(b []byte) error {
	var res BoxInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
