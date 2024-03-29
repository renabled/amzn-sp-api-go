// Code generated by go-swagger; DO NOT EDIT.

package definitions_product_types_2020_09_01_models

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

// ProductTypeList A list of Amazon product types with definitions available.
//
// swagger:model ProductTypeList
type ProductTypeList struct {

	// Amazon product type version identifier.
	// Required: true
	ProductTypeVersion *string `json:"productTypeVersion"`

	// product types
	// Required: true
	ProductTypes []*ProductType `json:"productTypes"`
}

// Validate validates this product type list
func (m *ProductTypeList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProductTypeVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductTypeList) validateProductTypeVersion(formats strfmt.Registry) error {

	if err := validate.Required("productTypeVersion", "body", m.ProductTypeVersion); err != nil {
		return err
	}

	return nil
}

func (m *ProductTypeList) validateProductTypes(formats strfmt.Registry) error {

	if err := validate.Required("productTypes", "body", m.ProductTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.ProductTypes); i++ {
		if swag.IsZero(m.ProductTypes[i]) { // not required
			continue
		}

		if m.ProductTypes[i] != nil {
			if err := m.ProductTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this product type list based on the context it is used
func (m *ProductTypeList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProductTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductTypeList) contextValidateProductTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProductTypes); i++ {

		if m.ProductTypes[i] != nil {
			if err := m.ProductTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductTypeList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductTypeList) UnmarshalBinary(b []byte) error {
	var res ProductTypeList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
