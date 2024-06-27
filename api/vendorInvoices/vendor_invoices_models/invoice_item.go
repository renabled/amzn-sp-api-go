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
	"github.com/go-openapi/validate"
)

// InvoiceItem Details of the item being invoiced.
//
// swagger:model InvoiceItem
type InvoiceItem struct {

	// Individual allowance details per line item.
	AllowanceDetails []*AllowanceDetails `json:"allowanceDetails"`

	// Amazon Standard Identification Number (ASIN) of an item.
	AmazonProductIdentifier string `json:"amazonProductIdentifier,omitempty"`

	// Individual charge details per line item.
	ChargeDetails []*ChargeDetails `json:"chargeDetails"`

	// Details required in order to process a credit note. This information is required only if `invoiceType` is `CreditNote`.
	CreditNoteDetails *CreditNoteDetails `json:"creditNoteDetails,omitempty"`

	// The HSN Tax code. The HSN number cannot contain alphabets.
	HsnCode string `json:"hsnCode,omitempty"`

	// Invoiced quantity of this item. Quantity must be greater than zero.
	// Required: true
	InvoicedQuantity *ItemQuantity `json:"invoicedQuantity"`

	// Unique number related to this line item.
	// Required: true
	ItemSequenceNumber *int64 `json:"itemSequenceNumber"`

	// The item cost to Amazon, which should match the cost on the order. Price information should not be zero or negative. It indicates net unit price. Net cost means VAT is not included in cost.
	// Required: true
	NetCost *Money `json:"netCost"`

	// The Amazon purchase order number for this invoiced line item. Formatting Notes: 8-character alpha-numeric code. This value is mandatory only when `invoiceType` is `Invoice`, and is not required when `invoiceType` is `CreditNote`.
	PurchaseOrderNumber string `json:"purchaseOrderNumber,omitempty"`

	// Individual tax details per line item.
	TaxDetails []*TaxDetails `json:"taxDetails"`

	// The vendor selected product identifier of the item. Should be the same as was provided in the purchase order.
	VendorProductIdentifier string `json:"vendorProductIdentifier,omitempty"`
}

// Validate validates this invoice item
func (m *InvoiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowanceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditNoteDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoicedQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemSequenceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceItem) validateAllowanceDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowanceDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowanceDetails); i++ {
		if swag.IsZero(m.AllowanceDetails[i]) { // not required
			continue
		}

		if m.AllowanceDetails[i] != nil {
			if err := m.AllowanceDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowanceDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowanceDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceItem) validateChargeDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.ChargeDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ChargeDetails); i++ {
		if swag.IsZero(m.ChargeDetails[i]) { // not required
			continue
		}

		if m.ChargeDetails[i] != nil {
			if err := m.ChargeDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chargeDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chargeDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceItem) validateCreditNoteDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.CreditNoteDetails) { // not required
		return nil
	}

	if m.CreditNoteDetails != nil {
		if err := m.CreditNoteDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditNoteDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creditNoteDetails")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceItem) validateInvoicedQuantity(formats strfmt.Registry) error {

	if err := validate.Required("invoicedQuantity", "body", m.InvoicedQuantity); err != nil {
		return err
	}

	if m.InvoicedQuantity != nil {
		if err := m.InvoicedQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoicedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invoicedQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceItem) validateItemSequenceNumber(formats strfmt.Registry) error {

	if err := validate.Required("itemSequenceNumber", "body", m.ItemSequenceNumber); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateNetCost(formats strfmt.Registry) error {

	if err := validate.Required("netCost", "body", m.NetCost); err != nil {
		return err
	}

	if m.NetCost != nil {
		if err := m.NetCost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netCost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netCost")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceItem) validateTaxDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.TaxDetails); i++ {
		if swag.IsZero(m.TaxDetails[i]) { // not required
			continue
		}

		if m.TaxDetails[i] != nil {
			if err := m.TaxDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this invoice item based on the context it is used
func (m *InvoiceItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowanceDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChargeDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreditNoteDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvoicedQuantity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceItem) contextValidateAllowanceDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowanceDetails); i++ {

		if m.AllowanceDetails[i] != nil {
			if err := m.AllowanceDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowanceDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowanceDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceItem) contextValidateChargeDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChargeDetails); i++ {

		if m.ChargeDetails[i] != nil {
			if err := m.ChargeDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chargeDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chargeDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceItem) contextValidateCreditNoteDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.CreditNoteDetails != nil {
		if err := m.CreditNoteDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creditNoteDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creditNoteDetails")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceItem) contextValidateInvoicedQuantity(ctx context.Context, formats strfmt.Registry) error {

	if m.InvoicedQuantity != nil {
		if err := m.InvoicedQuantity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoicedQuantity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invoicedQuantity")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceItem) contextValidateNetCost(ctx context.Context, formats strfmt.Registry) error {

	if m.NetCost != nil {
		if err := m.NetCost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netCost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netCost")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceItem) contextValidateTaxDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaxDetails); i++ {

		if m.TaxDetails[i] != nil {
			if err := m.TaxDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceItem) UnmarshalBinary(b []byte) error {
	var res InvoiceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
