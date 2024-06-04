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

// ListShipmentContentUpdatePreviewsResponse The `ListShipmentContentUpdatePreviews` response.
//
// swagger:model ListShipmentContentUpdatePreviewsResponse
type ListShipmentContentUpdatePreviewsResponse struct {

	// A list of content update previews in a shipment.
	// Required: true
	ContentUpdatePreviews []*ContentUpdatePreview `json:"contentUpdatePreviews"`

	// pagination
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Validate validates this list shipment content update previews response
func (m *ListShipmentContentUpdatePreviewsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentUpdatePreviews(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListShipmentContentUpdatePreviewsResponse) validateContentUpdatePreviews(formats strfmt.Registry) error {

	if err := validate.Required("contentUpdatePreviews", "body", m.ContentUpdatePreviews); err != nil {
		return err
	}

	for i := 0; i < len(m.ContentUpdatePreviews); i++ {
		if swag.IsZero(m.ContentUpdatePreviews[i]) { // not required
			continue
		}

		if m.ContentUpdatePreviews[i] != nil {
			if err := m.ContentUpdatePreviews[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contentUpdatePreviews" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contentUpdatePreviews" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListShipmentContentUpdatePreviewsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list shipment content update previews response based on the context it is used
func (m *ListShipmentContentUpdatePreviewsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContentUpdatePreviews(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListShipmentContentUpdatePreviewsResponse) contextValidateContentUpdatePreviews(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ContentUpdatePreviews); i++ {

		if m.ContentUpdatePreviews[i] != nil {
			if err := m.ContentUpdatePreviews[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contentUpdatePreviews" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contentUpdatePreviews" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListShipmentContentUpdatePreviewsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListShipmentContentUpdatePreviewsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListShipmentContentUpdatePreviewsResponse) UnmarshalBinary(b []byte) error {
	var res ListShipmentContentUpdatePreviewsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
