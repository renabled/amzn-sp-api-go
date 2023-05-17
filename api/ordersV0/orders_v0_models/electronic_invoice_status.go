// Code generated by go-swagger; DO NOT EDIT.

package orders_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ElectronicInvoiceStatus The status of the electronic invoice.
//
// swagger:model ElectronicInvoiceStatus
type ElectronicInvoiceStatus string

func NewElectronicInvoiceStatus(value ElectronicInvoiceStatus) *ElectronicInvoiceStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ElectronicInvoiceStatus.
func (m ElectronicInvoiceStatus) Pointer() *ElectronicInvoiceStatus {
	return &m
}

const (

	// ElectronicInvoiceStatusNotRequired captures enum value "NotRequired"
	ElectronicInvoiceStatusNotRequired ElectronicInvoiceStatus = "NotRequired"

	// ElectronicInvoiceStatusNotFound captures enum value "NotFound"
	ElectronicInvoiceStatusNotFound ElectronicInvoiceStatus = "NotFound"

	// ElectronicInvoiceStatusProcessing captures enum value "Processing"
	ElectronicInvoiceStatusProcessing ElectronicInvoiceStatus = "Processing"

	// ElectronicInvoiceStatusErrored captures enum value "Errored"
	ElectronicInvoiceStatusErrored ElectronicInvoiceStatus = "Errored"

	// ElectronicInvoiceStatusAccepted captures enum value "Accepted"
	ElectronicInvoiceStatusAccepted ElectronicInvoiceStatus = "Accepted"
)

// for schema
var electronicInvoiceStatusEnum []interface{}

func init() {
	var res []ElectronicInvoiceStatus
	if err := json.Unmarshal([]byte(`["NotRequired","NotFound","Processing","Errored","Accepted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		electronicInvoiceStatusEnum = append(electronicInvoiceStatusEnum, v)
	}
}

func (m ElectronicInvoiceStatus) validateElectronicInvoiceStatusEnum(path, location string, value ElectronicInvoiceStatus) error {
	if err := validate.EnumCase(path, location, value, electronicInvoiceStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this electronic invoice status
func (m ElectronicInvoiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateElectronicInvoiceStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this electronic invoice status based on context it is used
func (m ElectronicInvoiceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}