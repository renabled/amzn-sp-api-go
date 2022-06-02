// Code generated by go-swagger; DO NOT EDIT.

package listings_items_2021_08_01_models

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

// ListingsItemPatchRequest The request body schema for the patchListingsItem operation.
//
// swagger:model ListingsItemPatchRequest
type ListingsItemPatchRequest struct {

	// One or more JSON Patch operations to perform on the listings item.
	// Required: true
	// Min Items: 1
	Patches []*PatchOperation `json:"patches"`

	// The Amazon product type of the listings item.
	// Required: true
	ProductType *string `json:"productType"`
}

// Validate validates this listings item patch request
func (m *ListingsItemPatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListingsItemPatchRequest) validatePatches(formats strfmt.Registry) error {

	if err := validate.Required("patches", "body", m.Patches); err != nil {
		return err
	}

	iPatchesSize := int64(len(m.Patches))

	if err := validate.MinItems("patches", "body", iPatchesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Patches); i++ {
		if swag.IsZero(m.Patches[i]) { // not required
			continue
		}

		if m.Patches[i] != nil {
			if err := m.Patches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListingsItemPatchRequest) validateProductType(formats strfmt.Registry) error {

	if err := validate.Required("productType", "body", m.ProductType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this listings item patch request based on the context it is used
func (m *ListingsItemPatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatches(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListingsItemPatchRequest) contextValidatePatches(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Patches); i++ {

		if m.Patches[i] != nil {
			if err := m.Patches[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patches" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("patches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListingsItemPatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListingsItemPatchRequest) UnmarshalBinary(b []byte) error {
	var res ListingsItemPatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
