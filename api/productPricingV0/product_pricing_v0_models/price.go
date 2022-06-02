// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Price price
//
// swagger:model Price
type Price struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN,omitempty"`

	// product
	Product *Product `json:"Product,omitempty"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU string `json:"SellerSKU,omitempty"`

	// The status of the operation.
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this price
func (m *Price) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Price) validateProduct(formats strfmt.Registry) error {
	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Product")
			}
			return err
		}
	}

	return nil
}

func (m *Price) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this price based on the context it is used
func (m *Price) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProduct(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Price) contextValidateProduct(ctx context.Context, formats strfmt.Registry) error {

	if m.Product != nil {
		if err := m.Product.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Product")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Price) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Price) UnmarshalBinary(b []byte) error {
	var res Price
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
