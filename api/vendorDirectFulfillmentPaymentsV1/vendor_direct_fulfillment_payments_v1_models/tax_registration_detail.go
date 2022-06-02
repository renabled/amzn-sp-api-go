// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_payments_v1_models

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

// TaxRegistrationDetail Tax registration details of the entity.
//
// swagger:model TaxRegistrationDetail
type TaxRegistrationDetail struct {

	// Address associated with the tax registration number.
	TaxRegistrationAddress *Address `json:"taxRegistrationAddress,omitempty"`

	// Tax registration message that can be used for additional tax related details.
	TaxRegistrationMessage string `json:"taxRegistrationMessage,omitempty"`

	// Tax registration number for the party. For example, VAT ID.
	// Required: true
	TaxRegistrationNumber *string `json:"taxRegistrationNumber"`

	// Tax registration type for the entity.
	// Enum: [VAT GST]
	TaxRegistrationType string `json:"taxRegistrationType,omitempty"`
}

// Validate validates this tax registration detail
func (m *TaxRegistrationDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxRegistrationAddress(formats); err != nil {
		res = append(res, err)
	}

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

func (m *TaxRegistrationDetail) validateTaxRegistrationAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxRegistrationAddress) { // not required
		return nil
	}

	if m.TaxRegistrationAddress != nil {
		if err := m.TaxRegistrationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxRegistrationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxRegistrationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *TaxRegistrationDetail) validateTaxRegistrationNumber(formats strfmt.Registry) error {

	if err := validate.Required("taxRegistrationNumber", "body", m.TaxRegistrationNumber); err != nil {
		return err
	}

	return nil
}

var taxRegistrationDetailTypeTaxRegistrationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VAT","GST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taxRegistrationDetailTypeTaxRegistrationTypePropEnum = append(taxRegistrationDetailTypeTaxRegistrationTypePropEnum, v)
	}
}

const (

	// TaxRegistrationDetailTaxRegistrationTypeVAT captures enum value "VAT"
	TaxRegistrationDetailTaxRegistrationTypeVAT string = "VAT"

	// TaxRegistrationDetailTaxRegistrationTypeGST captures enum value "GST"
	TaxRegistrationDetailTaxRegistrationTypeGST string = "GST"
)

// prop value enum
func (m *TaxRegistrationDetail) validateTaxRegistrationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taxRegistrationDetailTypeTaxRegistrationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaxRegistrationDetail) validateTaxRegistrationType(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxRegistrationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTaxRegistrationTypeEnum("taxRegistrationType", "body", m.TaxRegistrationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this tax registration detail based on the context it is used
func (m *TaxRegistrationDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxRegistrationAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxRegistrationDetail) contextValidateTaxRegistrationAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxRegistrationAddress != nil {
		if err := m.TaxRegistrationAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxRegistrationAddress")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxRegistrationAddress")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxRegistrationDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxRegistrationDetail) UnmarshalBinary(b []byte) error {
	var res TaxRegistrationDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
