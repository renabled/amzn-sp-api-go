// Code generated by go-swagger; DO NOT EDIT.

package vendor_direct_fulfillment_shipping_v1_models

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

// PackingSlip Packing slip information.
//
// swagger:model PackingSlip
type PackingSlip struct {

	// A Base64encoded string of the packing slip PDF.
	// Required: true
	Content *string `json:"content"`

	// The format of the file such as PDF, JPEG etc.
	// Enum: [application/pdf]
	ContentType string `json:"contentType,omitempty"`

	// Purchase order number of the shipment that the packing slip is for.
	// Required: true
	// Pattern: ^[a-zA-Z0-9]+$
	PurchaseOrderNumber *string `json:"purchaseOrderNumber"`
}

// Validate validates this packing slip
func (m *PackingSlip) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseOrderNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackingSlip) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

var packingSlipTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["application/pdf"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		packingSlipTypeContentTypePropEnum = append(packingSlipTypeContentTypePropEnum, v)
	}
}

const (

	// PackingSlipContentTypeApplicationPdf captures enum value "application/pdf"
	PackingSlipContentTypeApplicationPdf string = "application/pdf"
)

// prop value enum
func (m *PackingSlip) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, packingSlipTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PackingSlip) validateContentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *PackingSlip) validatePurchaseOrderNumber(formats strfmt.Registry) error {

	if err := validate.Required("purchaseOrderNumber", "body", m.PurchaseOrderNumber); err != nil {
		return err
	}

	if err := validate.Pattern("purchaseOrderNumber", "body", *m.PurchaseOrderNumber, `^[a-zA-Z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this packing slip based on context it is used
func (m *PackingSlip) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackingSlip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackingSlip) UnmarshalBinary(b []byte) error {
	var res PackingSlip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
