// Code generated by go-swagger; DO NOT EDIT.

package vendor_invoices_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubmitInvoicesRequest The request schema for the submitInvoices operation.
//
// swagger:model SubmitInvoicesRequest
type SubmitInvoicesRequest struct {

	// invoices
	Invoices []*Invoice `json:"invoices"`
}

// Validate validates this submit invoices request
func (m *SubmitInvoicesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvoices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitInvoicesRequest) validateInvoices(formats strfmt.Registry) error {
	if swag.IsZero(m.Invoices) { // not required
		return nil
	}

	for i := 0; i < len(m.Invoices); i++ {
		if swag.IsZero(m.Invoices[i]) { // not required
			continue
		}

		if m.Invoices[i] != nil {
			if err := m.Invoices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invoices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("invoices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this submit invoices request based on the context it is used
func (m *SubmitInvoicesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvoices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubmitInvoicesRequest) contextValidateInvoices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Invoices); i++ {

		if m.Invoices[i] != nil {
			if err := m.Invoices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invoices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("invoices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubmitInvoicesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubmitInvoicesRequest) UnmarshalBinary(b []byte) error {
	var res SubmitInvoicesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
