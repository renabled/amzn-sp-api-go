// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_2021_12_28_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustomerInvoiceList Represents a list of customer invoices, potentially paginated.
//
// swagger:model CustomerInvoiceList
type CustomerInvoiceList struct {

	// Represents a customer invoice within the `CustomerInvoiceList`.
	CustomerInvoices []*CustomerInvoice `json:"customerInvoices"`

	// The pagination elements required to retrieve the remaining data.
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Validate validates this customer invoice list
func (m *CustomerInvoiceList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerInvoices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerInvoiceList) validateCustomerInvoices(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerInvoices) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomerInvoices); i++ {
		if swag.IsZero(m.CustomerInvoices[i]) { // not required
			continue
		}

		if m.CustomerInvoices[i] != nil {
			if err := m.CustomerInvoices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customerInvoices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customerInvoices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerInvoiceList) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this customer invoice list based on the context it is used
func (m *CustomerInvoiceList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomerInvoices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerInvoiceList) contextValidateCustomerInvoices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomerInvoices); i++ {

		if m.CustomerInvoices[i] != nil {
			if err := m.CustomerInvoices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customerInvoices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customerInvoices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustomerInvoiceList) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerInvoiceList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerInvoiceList) UnmarshalBinary(b []byte) error {
	var res CustomerInvoiceList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
