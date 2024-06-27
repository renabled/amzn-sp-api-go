// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_orders_v1_models

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

// TaxDetails The order's tax details.
//
// swagger:model TaxDetails
type TaxDetails struct {

	// tax amount
	// Required: true
	TaxAmount *Money `json:"taxAmount"`

	// tax rate
	TaxRate Decimal `json:"taxRate,omitempty"`

	// taxable amount
	TaxableAmount *Money `json:"taxableAmount,omitempty"`

	// Tax type.
	// Enum: [CONSUMPTION GST MwSt. PST TOTAL TVA VAT]
	Type string `json:"type,omitempty"`
}

// Validate validates this tax details
func (m *TaxDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxableAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxDetails) validateTaxAmount(formats strfmt.Registry) error {

	if err := validate.Required("taxAmount", "body", m.TaxAmount); err != nil {
		return err
	}

	if m.TaxAmount != nil {
		if err := m.TaxAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxAmount")
			}
			return err
		}
	}

	return nil
}

func (m *TaxDetails) validateTaxRate(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxRate) { // not required
		return nil
	}

	if err := m.TaxRate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("taxRate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("taxRate")
		}
		return err
	}

	return nil
}

func (m *TaxDetails) validateTaxableAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxableAmount) { // not required
		return nil
	}

	if m.TaxableAmount != nil {
		if err := m.TaxableAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxableAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxableAmount")
			}
			return err
		}
	}

	return nil
}

var taxDetailsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONSUMPTION","GST","MwSt.","PST","TOTAL","TVA","VAT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taxDetailsTypeTypePropEnum = append(taxDetailsTypeTypePropEnum, v)
	}
}

const (

	// TaxDetailsTypeCONSUMPTION captures enum value "CONSUMPTION"
	TaxDetailsTypeCONSUMPTION string = "CONSUMPTION"

	// TaxDetailsTypeGST captures enum value "GST"
	TaxDetailsTypeGST string = "GST"

	// TaxDetailsTypeMwStDot captures enum value "MwSt."
	TaxDetailsTypeMwStDot string = "MwSt."

	// TaxDetailsTypePST captures enum value "PST"
	TaxDetailsTypePST string = "PST"

	// TaxDetailsTypeTOTAL captures enum value "TOTAL"
	TaxDetailsTypeTOTAL string = "TOTAL"

	// TaxDetailsTypeTVA captures enum value "TVA"
	TaxDetailsTypeTVA string = "TVA"

	// TaxDetailsTypeVAT captures enum value "VAT"
	TaxDetailsTypeVAT string = "VAT"
)

// prop value enum
func (m *TaxDetails) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taxDetailsTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaxDetails) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this tax details based on the context it is used
func (m *TaxDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxableAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxDetails) contextValidateTaxAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxAmount != nil {
		if err := m.TaxAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxAmount")
			}
			return err
		}
	}

	return nil
}

func (m *TaxDetails) contextValidateTaxRate(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TaxRate.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("taxRate")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("taxRate")
		}
		return err
	}

	return nil
}

func (m *TaxDetails) contextValidateTaxableAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.TaxableAmount != nil {
		if err := m.TaxableAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taxableAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taxableAmount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxDetails) UnmarshalBinary(b []byte) error {
	var res TaxDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
