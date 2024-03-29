// Code generated by go-swagger; DO NOT EDIT.

package shipping_v2_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ServiceIds A list of ServiceId.
//
// swagger:model ServiceIds
type ServiceIds []string

// Validate validates this service ids
func (m ServiceIds) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service ids based on context it is used
func (m ServiceIds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
