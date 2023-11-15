// Code generated by go-swagger; DO NOT EDIT.

package supply_sources_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SupplySourceStatus The `SupplySource` status
//
// swagger:model SupplySourceStatus
type SupplySourceStatus string

func NewSupplySourceStatus(value SupplySourceStatus) *SupplySourceStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SupplySourceStatus.
func (m SupplySourceStatus) Pointer() *SupplySourceStatus {
	return &m
}

const (

	// SupplySourceStatusActive captures enum value "Active"
	SupplySourceStatusActive SupplySourceStatus = "Active"

	// SupplySourceStatusInactive captures enum value "Inactive"
	SupplySourceStatusInactive SupplySourceStatus = "Inactive"
)

// for schema
var supplySourceStatusEnum []interface{}

func init() {
	var res []SupplySourceStatus
	if err := json.Unmarshal([]byte(`["Active","Inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supplySourceStatusEnum = append(supplySourceStatusEnum, v)
	}
}

func (m SupplySourceStatus) validateSupplySourceStatusEnum(path, location string, value SupplySourceStatus) error {
	if err := validate.EnumCase(path, location, value, supplySourceStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this supply source status
func (m SupplySourceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSupplySourceStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this supply source status based on context it is used
func (m SupplySourceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}