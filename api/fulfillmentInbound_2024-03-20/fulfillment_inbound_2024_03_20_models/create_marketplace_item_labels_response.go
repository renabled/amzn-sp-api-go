// Code generated by go-swagger; DO NOT EDIT.

package fulfillment_inbound_2024_03_20_models

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

// CreateMarketplaceItemLabelsResponse The `createMarketplaceItemLabels` response.
//
// swagger:model CreateMarketplaceItemLabelsResponse
type CreateMarketplaceItemLabelsResponse struct {

	// Resources to download the requested document.
	// Required: true
	DocumentDownloads []*DocumentDownload `json:"documentDownloads"`
}

// Validate validates this create marketplace item labels response
func (m *CreateMarketplaceItemLabelsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentDownloads(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMarketplaceItemLabelsResponse) validateDocumentDownloads(formats strfmt.Registry) error {

	if err := validate.Required("documentDownloads", "body", m.DocumentDownloads); err != nil {
		return err
	}

	for i := 0; i < len(m.DocumentDownloads); i++ {
		if swag.IsZero(m.DocumentDownloads[i]) { // not required
			continue
		}

		if m.DocumentDownloads[i] != nil {
			if err := m.DocumentDownloads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documentDownloads" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("documentDownloads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create marketplace item labels response based on the context it is used
func (m *CreateMarketplaceItemLabelsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocumentDownloads(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMarketplaceItemLabelsResponse) contextValidateDocumentDownloads(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DocumentDownloads); i++ {

		if m.DocumentDownloads[i] != nil {
			if err := m.DocumentDownloads[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documentDownloads" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("documentDownloads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMarketplaceItemLabelsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMarketplaceItemLabelsResponse) UnmarshalBinary(b []byte) error {
	var res CreateMarketplaceItemLabelsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
