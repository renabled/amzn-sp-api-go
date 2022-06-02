// Code generated by go-swagger; DO NOT EDIT.

package notifications_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// MarketplaceIds A list of marketplace identifiers to subscribe to (e.g. ATVPDKIKX0DER). To receive notifications in every marketplace, do not provide this list.
//
// swagger:model MarketplaceIds
type MarketplaceIds []string

// Validate validates this marketplace ids
func (m MarketplaceIds) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this marketplace ids based on context it is used
func (m MarketplaceIds) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
