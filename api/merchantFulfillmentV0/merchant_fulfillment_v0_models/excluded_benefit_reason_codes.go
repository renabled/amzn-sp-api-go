// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ExcludedBenefitReasonCodes List of reasons (for example, `LATE_DELIVERY_RISK`) why a benefit is excluded for a shipping offer.
//
// swagger:model ExcludedBenefitReasonCodes
type ExcludedBenefitReasonCodes []string

// Validate validates this excluded benefit reason codes
func (m ExcludedBenefitReasonCodes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this excluded benefit reason codes based on context it is used
func (m ExcludedBenefitReasonCodes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
