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

// TaxRate Contains the type and rate of tax.
//
// swagger:model TaxRate
type TaxRate struct {

	// Rate of cess tax.
	CessRate float64 `json:"cessRate,omitempty"`

	// Rate of gst tax.
	GstRate float64 `json:"gstRate,omitempty"`

	// Type of tax. Possible values: `CGST`, `SGST`, `IGST`, `TOTAL_TAX`.
	// Max Length: 1024
	// Min Length: 1
	TaxType string `json:"taxType,omitempty"`
}

// Validate validates this tax rate
func (m *TaxRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxRate) validateTaxType(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxType) { // not required
		return nil
	}

	if err := validate.MinLength("taxType", "body", m.TaxType, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("taxType", "body", m.TaxType, 1024); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tax rate based on context it is used
func (m *TaxRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaxRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxRate) UnmarshalBinary(b []byte) error {
	var res TaxRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
