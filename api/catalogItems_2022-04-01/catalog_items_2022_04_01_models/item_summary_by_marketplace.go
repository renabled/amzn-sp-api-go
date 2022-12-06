// Code generated by go-swagger; DO NOT EDIT.

package catalog_items_2022_04_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ItemSummaryByMarketplace Summary details of an Amazon catalog item for the indicated Amazon marketplace.
//
// swagger:model ItemSummaryByMarketplace
type ItemSummaryByMarketplace struct {

	// Identifies an Amazon catalog item is intended for an adult audience or is sexual in nature.
	AdultProduct bool `json:"adultProduct,omitempty"`

	// Identifies an Amazon catalog item is autographed by a player or celebrity.
	Autographed bool `json:"autographed,omitempty"`

	// Name of the brand associated with an Amazon catalog item.
	Brand string `json:"brand,omitempty"`

	// Classification (browse node) associated with an Amazon catalog item.
	BrowseClassification *ItemBrowseClassification `json:"browseClassification,omitempty"`

	// Name of the color associated with an Amazon catalog item.
	Color string `json:"color,omitempty"`

	// Individual contributors to the creation of an item, such as the authors or actors.
	Contributors []*ItemContributor `json:"contributors"`

	// Classification type associated with the Amazon catalog item.
	// Enum: [BASE_PRODUCT OTHER PRODUCT_BUNDLE VARIATION_PARENT]
	ItemClassification string `json:"itemClassification,omitempty"`

	// Name, or title, associated with an Amazon catalog item.
	ItemName string `json:"itemName,omitempty"`

	// Name of the manufacturer associated with an Amazon catalog item.
	Manufacturer string `json:"manufacturer,omitempty"`

	// Amazon marketplace identifier.
	// Required: true
	MarketplaceID *string `json:"marketplaceId"`

	// Identifies an Amazon catalog item is memorabilia valued for its connection with historical events, culture, or entertainment.
	Memorabilia bool `json:"memorabilia,omitempty"`

	// Model number associated with an Amazon catalog item.
	ModelNumber string `json:"modelNumber,omitempty"`

	// Quantity of an Amazon catalog item in one package.
	PackageQuantity int64 `json:"packageQuantity,omitempty"`

	// Part number associated with an Amazon catalog item.
	PartNumber string `json:"partNumber,omitempty"`

	// First date on which an Amazon catalog item is shippable to customers.
	// Format: date
	ReleaseDate strfmt.Date `json:"releaseDate,omitempty"`

	// Name of the size associated with an Amazon catalog item.
	Size string `json:"size,omitempty"`

	// Name of the style associated with an Amazon catalog item.
	Style string `json:"style,omitempty"`

	// Identifies an Amazon catalog item is eligible for trade-in.
	TradeInEligible bool `json:"tradeInEligible,omitempty"`

	// Identifier of the website display group associated with an Amazon catalog item.
	WebsiteDisplayGroup string `json:"websiteDisplayGroup,omitempty"`

	// Display name of the website display group associated with an Amazon catalog item.
	WebsiteDisplayGroupName string `json:"websiteDisplayGroupName,omitempty"`
}

// Validate validates this item summary by marketplace
func (m *ItemSummaryByMarketplace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrowseClassification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContributors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemClassification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSummaryByMarketplace) validateBrowseClassification(formats strfmt.Registry) error {
	if swag.IsZero(m.BrowseClassification) { // not required
		return nil
	}

	if m.BrowseClassification != nil {
		if err := m.BrowseClassification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("browseClassification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("browseClassification")
			}
			return err
		}
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateContributors(formats strfmt.Registry) error {
	if swag.IsZero(m.Contributors) { // not required
		return nil
	}

	for i := 0; i < len(m.Contributors); i++ {
		if swag.IsZero(m.Contributors[i]) { // not required
			continue
		}

		if m.Contributors[i] != nil {
			if err := m.Contributors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contributors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contributors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var itemSummaryByMarketplaceTypeItemClassificationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BASE_PRODUCT","OTHER","PRODUCT_BUNDLE","VARIATION_PARENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemSummaryByMarketplaceTypeItemClassificationPropEnum = append(itemSummaryByMarketplaceTypeItemClassificationPropEnum, v)
	}
}

const (

	// ItemSummaryByMarketplaceItemClassificationBASEPRODUCT captures enum value "BASE_PRODUCT"
	ItemSummaryByMarketplaceItemClassificationBASEPRODUCT string = "BASE_PRODUCT"

	// ItemSummaryByMarketplaceItemClassificationOTHER captures enum value "OTHER"
	ItemSummaryByMarketplaceItemClassificationOTHER string = "OTHER"

	// ItemSummaryByMarketplaceItemClassificationPRODUCTBUNDLE captures enum value "PRODUCT_BUNDLE"
	ItemSummaryByMarketplaceItemClassificationPRODUCTBUNDLE string = "PRODUCT_BUNDLE"

	// ItemSummaryByMarketplaceItemClassificationVARIATIONPARENT captures enum value "VARIATION_PARENT"
	ItemSummaryByMarketplaceItemClassificationVARIATIONPARENT string = "VARIATION_PARENT"
)

// prop value enum
func (m *ItemSummaryByMarketplace) validateItemClassificationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemSummaryByMarketplaceTypeItemClassificationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemSummaryByMarketplace) validateItemClassification(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemClassification) { // not required
		return nil
	}

	// value enum
	if err := m.validateItemClassificationEnum("itemClassification", "body", m.ItemClassification); err != nil {
		return err
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	return nil
}

func (m *ItemSummaryByMarketplace) validateReleaseDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("releaseDate", "body", "date", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this item summary by marketplace based on the context it is used
func (m *ItemSummaryByMarketplace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBrowseClassification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContributors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSummaryByMarketplace) contextValidateBrowseClassification(ctx context.Context, formats strfmt.Registry) error {

	if m.BrowseClassification != nil {
		if err := m.BrowseClassification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("browseClassification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("browseClassification")
			}
			return err
		}
	}

	return nil
}

func (m *ItemSummaryByMarketplace) contextValidateContributors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Contributors); i++ {

		if m.Contributors[i] != nil {
			if err := m.Contributors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contributors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contributors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemSummaryByMarketplace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemSummaryByMarketplace) UnmarshalBinary(b []byte) error {
	var res ItemSummaryByMarketplace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
