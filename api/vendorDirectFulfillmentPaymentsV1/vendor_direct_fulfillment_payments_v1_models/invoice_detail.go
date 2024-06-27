// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_payments_v1_models

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

// InvoiceDetail Represents the invoice details, including the invoice number, date, parties involved, payment terms, totals, taxes, charges, and line items.
//
// swagger:model InvoiceDetail
type InvoiceDetail struct {

	// Additional details provided by the selling party, for tax-related or other purposes.
	AdditionalDetails []*AdditionalDetails `json:"additionalDetails"`

	// Name, address and tax details of the party to whom this invoice is issued.
	BillToParty *PartyIdentification `json:"billToParty,omitempty"`

	// Total charge amount details for all line items.
	ChargeDetails []*ChargeDetails `json:"chargeDetails"`

	// Invoice date.
	// Required: true
	// Format: date-time
	InvoiceDate *strfmt.DateTime `json:"invoiceDate"`

	// The unique invoice number.
	// Required: true
	InvoiceNumber *string `json:"invoiceNumber"`

	// Total amount details of the invoice.
	// Required: true
	InvoiceTotal *Money `json:"invoiceTotal"`

	// Provides the details of the items in this invoice.
	// Required: true
	Items []*InvoiceItem `json:"items"`

	// The payment terms for the invoice.
	PaymentTermsCode string `json:"paymentTermsCode,omitempty"`

	// An additional unique reference number used for regulatory or other purposes.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// Name, address and tax details of the party receiving the payment of this invoice.
	// Required: true
	RemitToParty *PartyIdentification `json:"remitToParty"`

	// Warehouse code of the vendor as in the order.
	// Required: true
	ShipFromParty *PartyIdentification `json:"shipFromParty"`

	// Ship-to country code.
	ShipToCountryCode string `json:"shipToCountryCode,omitempty"`

	// Individual tax details per line item.
	TaxTotals []*TaxDetail `json:"taxTotals"`
}

// Validate validates this invoice detail
func (m *InvoiceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillToParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemitToParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipFromParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxTotals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceDetail) validateAdditionalDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalDetails); i++ {
		if swag.IsZero(m.AdditionalDetails[i]) { // not required
			continue
		}

		if m.AdditionalDetails[i] != nil {
			if err := m.AdditionalDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceDetail) validateBillToParty(formats strfmt.Registry) error {
	if swag.IsZero(m.BillToParty) { // not required
		return nil
	}

	if m.BillToParty != nil {
		if err := m.BillToParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billToParty")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) validateChargeDetails(formats strfmt.Registry) error {
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

func (m *InvoiceDetail) validateInvoiceDate(formats strfmt.Registry) error {

	if err := validate.Required("invoiceDate", "body", m.InvoiceDate); err != nil {
		return err
	}

	if err := validate.FormatOf("invoiceDate", "body", "date-time", m.InvoiceDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceDetail) validateInvoiceNumber(formats strfmt.Registry) error {

	if err := validate.Required("invoiceNumber", "body", m.InvoiceNumber); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceDetail) validateInvoiceTotal(formats strfmt.Registry) error {

	if err := validate.Required("invoiceTotal", "body", m.InvoiceTotal); err != nil {
		return err
	}

	if m.InvoiceTotal != nil {
		if err := m.InvoiceTotal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoiceTotal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invoiceTotal")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceDetail) validateRemitToParty(formats strfmt.Registry) error {

	if err := validate.Required("remitToParty", "body", m.RemitToParty); err != nil {
		return err
	}

	if m.RemitToParty != nil {
		if err := m.RemitToParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remitToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remitToParty")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) validateShipFromParty(formats strfmt.Registry) error {

	if err := validate.Required("shipFromParty", "body", m.ShipFromParty); err != nil {
		return err
	}

	if m.ShipFromParty != nil {
		if err := m.ShipFromParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFromParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFromParty")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) validateTaxTotals(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxTotals) { // not required
		return nil
	}

	for i := 0; i < len(m.TaxTotals); i++ {
		if swag.IsZero(m.TaxTotals[i]) { // not required
			continue
		}

		if m.TaxTotals[i] != nil {
			if err := m.TaxTotals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxTotals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxTotals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this invoice detail based on the context it is used
func (m *InvoiceDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBillToParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChargeDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInvoiceTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemitToParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipFromParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxTotals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceDetail) contextValidateAdditionalDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalDetails); i++ {

		if m.AdditionalDetails[i] != nil {
			if err := m.AdditionalDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceDetail) contextValidateBillToParty(ctx context.Context, formats strfmt.Registry) error {

	if m.BillToParty != nil {
		if err := m.BillToParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("billToParty")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) contextValidateChargeDetails(ctx context.Context, formats strfmt.Registry) error {

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

func (m *InvoiceDetail) contextValidateInvoiceTotal(ctx context.Context, formats strfmt.Registry) error {

	if m.InvoiceTotal != nil {
		if err := m.InvoiceTotal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("invoiceTotal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("invoiceTotal")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceDetail) contextValidateRemitToParty(ctx context.Context, formats strfmt.Registry) error {

	if m.RemitToParty != nil {
		if err := m.RemitToParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remitToParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remitToParty")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) contextValidateShipFromParty(ctx context.Context, formats strfmt.Registry) error {

	if m.ShipFromParty != nil {
		if err := m.ShipFromParty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shipFromParty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shipFromParty")
			}
			return err
		}
	}

	return nil
}

func (m *InvoiceDetail) contextValidateTaxTotals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaxTotals); i++ {

		if m.TaxTotals[i] != nil {
			if err := m.TaxTotals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxTotals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxTotals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceDetail) UnmarshalBinary(b []byte) error {
	var res InvoiceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
