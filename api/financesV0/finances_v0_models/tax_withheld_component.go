// Code generated by go-swagger; DO NOT EDIT.

package finances_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaxWithheldComponent Information about the taxes withheld.
//
// swagger:model TaxWithheldComponent
type TaxWithheldComponent struct {

	// The tax collection model applied to the item.
	//
	// Possible values:
	//
	// * `MarketplaceFacilitator`: Tax is withheld and remitted to the taxing authority by Amazon on behalf of the seller.
	// * `Standard`: Tax is paid to the seller and not remitted to the taxing authority by Amazon.
	TaxCollectionModel string `json:"TaxCollectionModel,omitempty"`

	// A list of charges that represent the types and amounts of taxes withheld.
	TaxesWithheld ChargeComponentList `json:"TaxesWithheld,omitempty"`
}

// Validate validates this tax withheld component
func (m *TaxWithheldComponent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxesWithheld(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxWithheldComponent) validateTaxesWithheld(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxesWithheld) { // not required
		return nil
	}

	if err := m.TaxesWithheld.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TaxesWithheld")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("TaxesWithheld")
		}
		return err
	}

	return nil
}

// ContextValidate validate this tax withheld component based on the context it is used
func (m *TaxWithheldComponent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxesWithheld(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxWithheldComponent) contextValidateTaxesWithheld(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TaxesWithheld.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("TaxesWithheld")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("TaxesWithheld")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxWithheldComponent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxWithheldComponent) UnmarshalBinary(b []byte) error {
	var res TaxWithheldComponent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
