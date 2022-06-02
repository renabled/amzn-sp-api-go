// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ASINPrepInstructions Item preparation instructions to help with item sourcing decisions.
//
// swagger:model ASINPrepInstructions
type ASINPrepInstructions struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN,omitempty"`

	// barcode instruction
	BarcodeInstruction BarcodeInstruction `json:"BarcodeInstruction,omitempty"`

	// prep guidance
	PrepGuidance PrepGuidance `json:"PrepGuidance,omitempty"`

	// prep instruction list
	PrepInstructionList PrepInstructionList `json:"PrepInstructionList,omitempty"`
}

// Validate validates this a s i n prep instructions
func (m *ASINPrepInstructions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBarcodeInstruction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepGuidance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepInstructionList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ASINPrepInstructions) validateBarcodeInstruction(formats strfmt.Registry) error {
	if swag.IsZero(m.BarcodeInstruction) { // not required
		return nil
	}

	if err := m.BarcodeInstruction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BarcodeInstruction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("BarcodeInstruction")
		}
		return err
	}

	return nil
}

func (m *ASINPrepInstructions) validatePrepGuidance(formats strfmt.Registry) error {
	if swag.IsZero(m.PrepGuidance) { // not required
		return nil
	}

	if err := m.PrepGuidance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PrepGuidance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PrepGuidance")
		}
		return err
	}

	return nil
}

func (m *ASINPrepInstructions) validatePrepInstructionList(formats strfmt.Registry) error {
	if swag.IsZero(m.PrepInstructionList) { // not required
		return nil
	}

	if err := m.PrepInstructionList.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PrepInstructionList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PrepInstructionList")
		}
		return err
	}

	return nil
}

// ContextValidate validate this a s i n prep instructions based on the context it is used
func (m *ASINPrepInstructions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBarcodeInstruction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrepGuidance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrepInstructionList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ASINPrepInstructions) contextValidateBarcodeInstruction(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BarcodeInstruction.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BarcodeInstruction")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("BarcodeInstruction")
		}
		return err
	}

	return nil
}

func (m *ASINPrepInstructions) contextValidatePrepGuidance(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PrepGuidance.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PrepGuidance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PrepGuidance")
		}
		return err
	}

	return nil
}

func (m *ASINPrepInstructions) contextValidatePrepInstructionList(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PrepInstructionList.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PrepInstructionList")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("PrepInstructionList")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ASINPrepInstructions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ASINPrepInstructions) UnmarshalBinary(b []byte) error {
	var res ASINPrepInstructions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
