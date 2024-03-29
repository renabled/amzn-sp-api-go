// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaxDetail Indicates the tax specifications associated with the shipment for customs compliance purposes in certain regions.
//
// swagger:model TaxDetail
type TaxDetail struct {

	// The shipper's tax registration number associated with the shipment for customs compliance purposes in certain regions.
	// Required: true
	TaxRegistrationNumber *string `json:"taxRegistrationNumber"`

	// tax type
	// Required: true
	TaxType *TaxType `json:"taxType"`
}

// Validate validates this tax detail
func (m *TaxDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxRegistrationNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxDetail) validateTaxRegistrationNumber(formats strfmt.Registry) error {

	if err := validate.Required("taxRegistrationNumber", "body", m.TaxRegistrationNumber); err != nil {
		return err
	}

	return nil
}

func (m *TaxDetail) validateTaxType(formats strfmt.Registry) error {

	if err := validate.Required("taxType", "body", m.TaxType); err != nil {
		return err
	}

	if err := validate.Required("taxType", "body", m.TaxType); err != nil {
		return err
	}

	if m.TaxType != nil {
		if err := m.TaxType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tax detail based on the context it is used
func (m *TaxDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxDetail) contextValidateTaxType(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxType != nil {
		if err := m.TaxType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxDetail) UnmarshalBinary(b []byte) error {
	var res TaxDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
