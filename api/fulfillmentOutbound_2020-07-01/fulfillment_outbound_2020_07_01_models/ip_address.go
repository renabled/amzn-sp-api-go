// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_outbound_2020_07_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// IPAddress An IP Address.
//
// swagger:model IpAddress
type IPAddress string

// Validate validates this Ip address
func (m IPAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Ip address based on context it is used
func (m IPAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
