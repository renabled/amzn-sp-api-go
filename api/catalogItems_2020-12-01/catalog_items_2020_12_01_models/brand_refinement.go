// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_2020_12_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BrandRefinement Description of a brand that can be used to get more fine-grained search results.
//
// swagger:model BrandRefinement
type BrandRefinement struct {

	// Brand name. For display and can be used as a search refinement.
	// Required: true
	BrandName *string `json:"brandName"`

	// The estimated number of results that would still be returned if refinement key applied.
	// Required: true
	NumberOfResults *int64 `json:"numberOfResults"`
}

// Validate validates this brand refinement
func (m *BrandRefinement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrandName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrandRefinement) validateBrandName(formats strfmt.Registry) error {

	if err := validate.Required("brandName", "body", m.BrandName); err != nil {
		return err
	}

	return nil
}

func (m *BrandRefinement) validateNumberOfResults(formats strfmt.Registry) error {

	if err := validate.Required("numberOfResults", "body", m.NumberOfResults); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this brand refinement based on context it is used
func (m *BrandRefinement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BrandRefinement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrandRefinement) UnmarshalBinary(b []byte) error {
	var res BrandRefinement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
