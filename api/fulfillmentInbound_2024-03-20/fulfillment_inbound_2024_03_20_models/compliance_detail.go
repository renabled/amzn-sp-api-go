// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ComplianceDetail Contains item identifiers and related tax information.
//
// swagger:model ComplianceDetail
type ComplianceDetail struct {

	// The Amazon Standard Identification Number, which identifies the detail page identifier.
	// Max Length: 10
	// Min Length: 1
	Asin string `json:"asin,omitempty"`

	// The Fulfillment Network SKU, which identifies a real fulfillable item with catalog data and condition.
	// Max Length: 10
	// Min Length: 1
	Fnsku string `json:"fnsku,omitempty"`

	// The merchant SKU, a merchant-supplied identifier for a specific SKU.
	// Max Length: 40
	// Min Length: 1
	Msku string `json:"msku,omitempty"`

	// tax details
	TaxDetails *TaxDetails `json:"taxDetails,omitempty"`
}

// Validate validates this compliance detail
func (m *ComplianceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFnsku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComplianceDetail) validateAsin(formats strfmt.Registry) error {
	if swag.IsZero(m.Asin) { // not required
		return nil
	}

	if err := validate.MinLength("asin", "body", m.Asin, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("asin", "body", m.Asin, 10); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceDetail) validateFnsku(formats strfmt.Registry) error {
	if swag.IsZero(m.Fnsku) { // not required
		return nil
	}

	if err := validate.MinLength("fnsku", "body", m.Fnsku, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("fnsku", "body", m.Fnsku, 10); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceDetail) validateMsku(formats strfmt.Registry) error {
	if swag.IsZero(m.Msku) { // not required
		return nil
	}

	if err := validate.MinLength("msku", "body", m.Msku, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("msku", "body", m.Msku, 40); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceDetail) validateTaxDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxDetails) { // not required
		return nil
	}

	if m.TaxDetails != nil {
		if err := m.TaxDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this compliance detail based on the context it is used
func (m *ComplianceDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComplianceDetail) contextValidateTaxDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxDetails != nil {
		if err := m.TaxDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceDetail) UnmarshalBinary(b []byte) error {
	var res ComplianceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
