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

// SupportedDocumentSpecification Document specification that is supported for a service offering.
//
// swagger:model SupportedDocumentSpecification
type SupportedDocumentSpecification struct {

	// format
	// Required: true
	Format *DocumentFormat `json:"format"`

	// print options
	// Required: true
	PrintOptions PrintOptionList `json:"printOptions"`

	// size
	// Required: true
	Size *DocumentSize `json:"size"`
}

// Validate validates this supported document specification
func (m *SupportedDocumentSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrintOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportedDocumentSpecification) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	if m.Format != nil {
		if err := m.Format.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("format")
			}
			return err
		}
	}

	return nil
}

func (m *SupportedDocumentSpecification) validatePrintOptions(formats strfmt.Registry) error {

	if err := validate.Required("printOptions", "body", m.PrintOptions); err != nil {
		return err
	}

	if err := m.PrintOptions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("printOptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("printOptions")
		}
		return err
	}

	return nil
}

func (m *SupportedDocumentSpecification) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this supported document specification based on the context it is used
func (m *SupportedDocumentSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrintOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportedDocumentSpecification) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.Format != nil {
		if err := m.Format.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("format")
			}
			return err
		}
	}

	return nil
}

func (m *SupportedDocumentSpecification) contextValidatePrintOptions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PrintOptions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("printOptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("printOptions")
		}
		return err
	}

	return nil
}

func (m *SupportedDocumentSpecification) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportedDocumentSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedDocumentSpecification) UnmarshalBinary(b []byte) error {
	var res SupportedDocumentSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
