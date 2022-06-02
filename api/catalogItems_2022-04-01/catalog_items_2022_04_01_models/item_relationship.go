// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_2022_04_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ItemRelationship Relationship details for an Amazon catalog item.
//
// swagger:model ItemRelationship
type ItemRelationship struct {

	// Identifiers (ASINs) of the related items that are children of this item.
	ChildAsins []string `json:"childAsins"`

	// Identifiers (ASINs) of the related items that are parents of this item.
	ParentAsins []string `json:"parentAsins"`

	// Type of relationship.
	// Example: VARIATION
	// Required: true
	// Enum: [VARIATION PACKAGE_HIERARCHY]
	Type *string `json:"type"`

	// For "VARIATION" relationships, variation theme indicating the combination of Amazon item catalog attributes that define the variation family.
	VariationTheme *ItemVariationTheme `json:"variationTheme,omitempty"`
}

// Validate validates this item relationship
func (m *ItemRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariationTheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var itemRelationshipTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VARIATION","PACKAGE_HIERARCHY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemRelationshipTypeTypePropEnum = append(itemRelationshipTypeTypePropEnum, v)
	}
}

const (

	// ItemRelationshipTypeVARIATION captures enum value "VARIATION"
	ItemRelationshipTypeVARIATION string = "VARIATION"

	// ItemRelationshipTypePACKAGEHIERARCHY captures enum value "PACKAGE_HIERARCHY"
	ItemRelationshipTypePACKAGEHIERARCHY string = "PACKAGE_HIERARCHY"
)

// prop value enum
func (m *ItemRelationship) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemRelationshipTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemRelationship) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ItemRelationship) validateVariationTheme(formats strfmt.Registry) error {
	if swag.IsZero(m.VariationTheme) { // not required
		return nil
	}

	if m.VariationTheme != nil {
		if err := m.VariationTheme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variationTheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variationTheme")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this item relationship based on the context it is used
func (m *ItemRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVariationTheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemRelationship) contextValidateVariationTheme(ctx context.Context, formats strfmt.Registry) error {

	if m.VariationTheme != nil {
		if err := m.VariationTheme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("variationTheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("variationTheme")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemRelationship) UnmarshalBinary(b []byte) error {
	var res ItemRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
