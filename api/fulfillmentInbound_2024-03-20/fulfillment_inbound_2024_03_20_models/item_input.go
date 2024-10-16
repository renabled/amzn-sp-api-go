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

// ItemInput Defines an item's input parameters.
// Example: {"expiration":"2024-01-01","labelOwner":"AMAZON","manufacturingLotCode":"manufacturingLotCode","msku":"Sunglasses","prepOwner":"AMAZON","quantity":10}
//
// swagger:model ItemInput
type ItemInput struct {

	// The expiration date of the MSKU. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `YYYY-MM-DD`. Items with the same MSKU but different expiration dates cannot go into the same box.
	// Pattern: ^([0-9]{4})-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$
	Expiration string `json:"expiration,omitempty"`

	// label owner
	// Required: true
	LabelOwner *LabelOwner `json:"labelOwner"`

	// The manufacturing lot code.
	// Max Length: 256
	// Min Length: 1
	ManufacturingLotCode string `json:"manufacturingLotCode,omitempty"`

	// The merchant SKU, a merchant-supplied identifier of a specific SKU.
	// Required: true
	// Max Length: 40
	// Min Length: 1
	Msku *string `json:"msku"`

	// prep owner
	// Required: true
	PrepOwner *PrepOwner `json:"prepOwner"`

	// The number of units of the specified MSKU that will be shipped.
	// Required: true
	// Maximum: 10000
	// Minimum: 1
	Quantity *int64 `json:"quantity"`
}

// Validate validates this item input
func (m *ItemInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturingLotCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsku(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemInput) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.Pattern("expiration", "body", m.Expiration, `^([0-9]{4})-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`); err != nil {
		return err
	}

	return nil
}

func (m *ItemInput) validateLabelOwner(formats strfmt.Registry) error {

	if err := validate.Required("labelOwner", "body", m.LabelOwner); err != nil {
		return err
	}

	if err := validate.Required("labelOwner", "body", m.LabelOwner); err != nil {
		return err
	}

	if m.LabelOwner != nil {
		if err := m.LabelOwner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labelOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labelOwner")
			}
			return err
		}
	}

	return nil
}

func (m *ItemInput) validateManufacturingLotCode(formats strfmt.Registry) error {
	if swag.IsZero(m.ManufacturingLotCode) { // not required
		return nil
	}

	if err := validate.MinLength("manufacturingLotCode", "body", m.ManufacturingLotCode, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("manufacturingLotCode", "body", m.ManufacturingLotCode, 256); err != nil {
		return err
	}

	return nil
}

func (m *ItemInput) validateMsku(formats strfmt.Registry) error {

	if err := validate.Required("msku", "body", m.Msku); err != nil {
		return err
	}

	if err := validate.MinLength("msku", "body", *m.Msku, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("msku", "body", *m.Msku, 40); err != nil {
		return err
	}

	return nil
}

func (m *ItemInput) validatePrepOwner(formats strfmt.Registry) error {

	if err := validate.Required("prepOwner", "body", m.PrepOwner); err != nil {
		return err
	}

	if err := validate.Required("prepOwner", "body", m.PrepOwner); err != nil {
		return err
	}

	if m.PrepOwner != nil {
		if err := m.PrepOwner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prepOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prepOwner")
			}
			return err
		}
	}

	return nil
}

func (m *ItemInput) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.MinimumInt("quantity", "body", *m.Quantity, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("quantity", "body", *m.Quantity, 10000, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this item input based on the context it is used
func (m *ItemInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabelOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrepOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemInput) contextValidateLabelOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelOwner != nil {
		if err := m.LabelOwner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labelOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labelOwner")
			}
			return err
		}
	}

	return nil
}

func (m *ItemInput) contextValidatePrepOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.PrepOwner != nil {
		if err := m.PrepOwner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prepOwner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("prepOwner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemInput) UnmarshalBinary(b []byte) error {
	var res ItemInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
