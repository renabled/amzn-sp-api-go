// Code generated by go-swagger; DO NOT EDIT.

package awd_2024_05_09_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OutboundOrderCreationData Payload for creating an outbound order.
// Example: {"orderPreferences":[{"orderPreference":"PARTIAL_ORDER","orderPreferenceValue":"SET"}],"packagesToOutbound":[{"count":1,"distributionPackage":{"contents":{"products":[{"quantity":1,"sku":"testPen"}]},"measurements":{"dimensions":{"height":1,"length":1,"unitOfMeasurement":"INCHES","width":1},"volume":{"unitOfMeasurement":"CUIN","volume":1},"weight":{"unitOfMeasurement":"POUNDS","weight":1}},"type":"CASE"}}],"productsToOutbound":[]}
//
// swagger:model OutboundOrderCreationData
type OutboundOrderCreationData struct {

	// Order preferences for the outbound order.
	// Example: [{"orderPreference":"PARTIAL_ORDER","orderPreferenceValue":"SET"}]
	OrderPreferences []*OrderAttribute `json:"orderPreferences"`

	// List of packages to be outbound.
	// Example: [{"count":1,"distributionPackage":{"contents":{"products":[{"quantity":1,"sku":"testPen"}]},"measurements":{"dimensions":{"height":1,"length":1,"unitOfMeasurement":"INCHES","width":1},"volume":{"unitOfMeasurement":"CUIN","volume":1},"weight":{"unitOfMeasurement":"POUNDS","weight":1}},"type":"CASE"}}]
	PackagesToOutbound []*DistributionPackageQuantity `json:"packagesToOutbound"`

	// List of product units to be outbound.
	// Example: [{"quantity":1,"sku":"TestSKU"}]
	ProductsToOutbound []*ProductQuantity `json:"productsToOutbound"`
}

// Validate validates this outbound order creation data
func (m *OutboundOrderCreationData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackagesToOutbound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductsToOutbound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutboundOrderCreationData) validateOrderPreferences(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderPreferences) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderPreferences); i++ {
		if swag.IsZero(m.OrderPreferences[i]) { // not required
			continue
		}

		if m.OrderPreferences[i] != nil {
			if err := m.OrderPreferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderPreferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orderPreferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OutboundOrderCreationData) validatePackagesToOutbound(formats strfmt.Registry) error {
	if swag.IsZero(m.PackagesToOutbound) { // not required
		return nil
	}

	for i := 0; i < len(m.PackagesToOutbound); i++ {
		if swag.IsZero(m.PackagesToOutbound[i]) { // not required
			continue
		}

		if m.PackagesToOutbound[i] != nil {
			if err := m.PackagesToOutbound[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packagesToOutbound" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packagesToOutbound" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OutboundOrderCreationData) validateProductsToOutbound(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductsToOutbound) { // not required
		return nil
	}

	for i := 0; i < len(m.ProductsToOutbound); i++ {
		if swag.IsZero(m.ProductsToOutbound[i]) { // not required
			continue
		}

		if m.ProductsToOutbound[i] != nil {
			if err := m.ProductsToOutbound[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productsToOutbound" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productsToOutbound" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this outbound order creation data based on the context it is used
func (m *OutboundOrderCreationData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderPreferences(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePackagesToOutbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductsToOutbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OutboundOrderCreationData) contextValidateOrderPreferences(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OrderPreferences); i++ {

		if m.OrderPreferences[i] != nil {
			if err := m.OrderPreferences[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("orderPreferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("orderPreferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OutboundOrderCreationData) contextValidatePackagesToOutbound(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PackagesToOutbound); i++ {

		if m.PackagesToOutbound[i] != nil {
			if err := m.PackagesToOutbound[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packagesToOutbound" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("packagesToOutbound" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OutboundOrderCreationData) contextValidateProductsToOutbound(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProductsToOutbound); i++ {

		if m.ProductsToOutbound[i] != nil {
			if err := m.ProductsToOutbound[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productsToOutbound" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productsToOutbound" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OutboundOrderCreationData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutboundOrderCreationData) UnmarshalBinary(b []byte) error {
	var res OutboundOrderCreationData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
