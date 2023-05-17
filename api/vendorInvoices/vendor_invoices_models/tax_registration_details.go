// Code generated by go-swagger; DO NOT EDIT.

package vendor_invoices_models

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

// TaxRegistrationDetails Tax registration details of the entity.
//
// swagger:model TaxRegistrationDetails
type TaxRegistrationDetails struct {

	// The tax registration number for the entity. For example, VAT ID, Consumption Tax ID.
	// Required: true
	TaxRegistrationNumber *string `json:"taxRegistrationNumber"`

	// The tax registration type for the entity.
	// Required: true
	// Enum: [VAT GST]
	TaxRegistrationType *string `json:"taxRegistrationType"`
}

// Validate validates this tax registration details
func (m *TaxRegistrationDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxRegistrationNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxRegistrationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxRegistrationDetails) validateTaxRegistrationNumber(formats strfmt.Registry) error {

	if err := validate.Required("taxRegistrationNumber", "body", m.TaxRegistrationNumber); err != nil {
		return err
	}

	return nil
}

var taxRegistrationDetailsTypeTaxRegistrationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VAT","GST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taxRegistrationDetailsTypeTaxRegistrationTypePropEnum = append(taxRegistrationDetailsTypeTaxRegistrationTypePropEnum, v)
	}
}

const (

	// TaxRegistrationDetailsTaxRegistrationTypeVAT captures enum value "VAT"
	TaxRegistrationDetailsTaxRegistrationTypeVAT string = "VAT"

	// TaxRegistrationDetailsTaxRegistrationTypeGST captures enum value "GST"
	TaxRegistrationDetailsTaxRegistrationTypeGST string = "GST"
)

// prop value enum
func (m *TaxRegistrationDetails) validateTaxRegistrationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taxRegistrationDetailsTypeTaxRegistrationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaxRegistrationDetails) validateTaxRegistrationType(formats strfmt.Registry) error {

	if err := validate.Required("taxRegistrationType", "body", m.TaxRegistrationType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTaxRegistrationTypeEnum("taxRegistrationType", "body", *m.TaxRegistrationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tax registration details based on context it is used
func (m *TaxRegistrationDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaxRegistrationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxRegistrationDetails) UnmarshalBinary(b []byte) error {
	var res TaxRegistrationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
