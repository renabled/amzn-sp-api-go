// Code generated by go-swagger; DO NOT EDIT.

package shipping_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Label The label details of the container.
//
// swagger:model Label
type Label struct {

	// label specification
	LabelSpecification *LabelSpecification `json:"labelSpecification,omitempty"`

	// label stream
	LabelStream LabelStream `json:"labelStream,omitempty"`
}

// Validate validates this label
func (m *Label) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelSpecification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelStream(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Label) validateLabelSpecification(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelSpecification) { // not required
		return nil
	}

	if m.LabelSpecification != nil {
		if err := m.LabelSpecification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labelSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labelSpecification")
			}
			return err
		}
	}

	return nil
}

func (m *Label) validateLabelStream(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelStream) { // not required
		return nil
	}

	if err := m.LabelStream.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labelStream")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labelStream")
		}
		return err
	}

	return nil
}

// ContextValidate validate this label based on the context it is used
func (m *Label) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabelSpecification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelStream(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Label) contextValidateLabelSpecification(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelSpecification != nil {
		if err := m.LabelSpecification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labelSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labelSpecification")
			}
			return err
		}
	}

	return nil
}

func (m *Label) contextValidateLabelStream(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LabelStream.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labelStream")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labelStream")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Label) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Label) UnmarshalBinary(b []byte) error {
	var res Label
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
