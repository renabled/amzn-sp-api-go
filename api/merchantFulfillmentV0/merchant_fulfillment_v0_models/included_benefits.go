// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment_v0_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IncludedBenefits A list of included benefits.
//
// swagger:model IncludedBenefits
type IncludedBenefits []string

// Validate validates this included benefits
func (m IncludedBenefits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this included benefits based on context it is used
func (m IncludedBenefits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
