// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_2022_05_01_models

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

// ShippingOption The shipping option available for the offer.
//
// swagger:model ShippingOption
type ShippingOption struct {

	// Shipping price for the offer.
	// Required: true
	Price *MoneyType `json:"price"`

	// The type of the shipping option.
	// Required: true
	// Enum: [DEFAULT]
	ShippingOptionType interface{} `json:"shippingOptionType"`
}

// Validate validates this shipping option
func (m *ShippingOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingOptionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingOption) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	if m.Price != nil {
		if err := m.Price.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

var shippingOptionTypeShippingOptionTypePropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`["DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shippingOptionTypeShippingOptionTypePropEnum = append(shippingOptionTypeShippingOptionTypePropEnum, v)
	}
}

// prop value enum
func (m *ShippingOption) validateShippingOptionTypeEnum(path, location string, value interface{}) error {
	if err := validate.EnumCase(path, location, value, shippingOptionTypeShippingOptionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ShippingOption) validateShippingOptionType(formats strfmt.Registry) error {

	if m.ShippingOptionType == nil {
		return errors.Required("shippingOptionType", "body", nil)
	}

	return nil
}

// ContextValidate validate this shipping option based on the context it is used
func (m *ShippingOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShippingOption) contextValidatePrice(ctx context.Context, formats strfmt.Registry) error {

	if m.Price != nil {
		if err := m.Price.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("price")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("price")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShippingOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShippingOption) UnmarshalBinary(b []byte) error {
	var res ShippingOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
