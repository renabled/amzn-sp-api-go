// Code generated by go-swagger; DO NOT EDIT.

package awd_2024_05_09_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OutboundStatus Statuses supported for an outbound order.
//
// swagger:model OutboundStatus
type OutboundStatus string

func NewOutboundStatus(value OutboundStatus) *OutboundStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OutboundStatus.
func (m OutboundStatus) Pointer() *OutboundStatus {
	return &m
}

const (

	// OutboundStatusCONFIRMED captures enum value "CONFIRMED"
	OutboundStatusCONFIRMED OutboundStatus = "CONFIRMED"

	// OutboundStatusDRAFT captures enum value "DRAFT"
	OutboundStatusDRAFT OutboundStatus = "DRAFT"

	// OutboundStatusELIGIBLE captures enum value "ELIGIBLE"
	OutboundStatusELIGIBLE OutboundStatus = "ELIGIBLE"

	// OutboundStatusEXECUTING captures enum value "EXECUTING"
	OutboundStatusEXECUTING OutboundStatus = "EXECUTING"

	// OutboundStatusFAILURE captures enum value "FAILURE"
	OutboundStatusFAILURE OutboundStatus = "FAILURE"

	// OutboundStatusINELIGIBLE captures enum value "INELIGIBLE"
	OutboundStatusINELIGIBLE OutboundStatus = "INELIGIBLE"

	// OutboundStatusINVENTORYOUTBOUND captures enum value "INVENTORY_OUTBOUND"
	OutboundStatusINVENTORYOUTBOUND OutboundStatus = "INVENTORY_OUTBOUND"

	// OutboundStatusSUCCESS captures enum value "SUCCESS"
	OutboundStatusSUCCESS OutboundStatus = "SUCCESS"

	// OutboundStatusVALIDATING captures enum value "VALIDATING"
	OutboundStatusVALIDATING OutboundStatus = "VALIDATING"
)

// for schema
var outboundStatusEnum []interface{}

func init() {
	var res []OutboundStatus
	if err := json.Unmarshal([]byte(`["CONFIRMED","DRAFT","ELIGIBLE","EXECUTING","FAILURE","INELIGIBLE","INVENTORY_OUTBOUND","SUCCESS","VALIDATING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		outboundStatusEnum = append(outboundStatusEnum, v)
	}
}

func (m OutboundStatus) validateOutboundStatusEnum(path, location string, value OutboundStatus) error {
	if err := validate.EnumCase(path, location, value, outboundStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this outbound status
func (m OutboundStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOutboundStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this outbound status based on context it is used
func (m OutboundStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
