// Code generated by go-swagger; DO NOT EDIT.

package replenishment_2022_11_07_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ListOffersSortKey The attribute to use to sort the results.
//
// swagger:model ListOffersSortKey
type ListOffersSortKey string

func NewListOffersSortKey(value ListOffersSortKey) *ListOffersSortKey {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ListOffersSortKey.
func (m ListOffersSortKey) Pointer() *ListOffersSortKey {
	return &m
}

const (

	// ListOffersSortKeyASIN captures enum value "ASIN"
	ListOffersSortKeyASIN ListOffersSortKey = "ASIN"

	// ListOffersSortKeySELLINGPARTNERFUNDEDBASEDISCOUNTPERCENTAGE captures enum value "SELLING_PARTNER_FUNDED_BASE_DISCOUNT_PERCENTAGE"
	ListOffersSortKeySELLINGPARTNERFUNDEDBASEDISCOUNTPERCENTAGE ListOffersSortKey = "SELLING_PARTNER_FUNDED_BASE_DISCOUNT_PERCENTAGE"

	// ListOffersSortKeySELLINGPARTNERFUNDEDTIEREDDISCOUNTPERCENTAGE captures enum value "SELLING_PARTNER_FUNDED_TIERED_DISCOUNT_PERCENTAGE"
	ListOffersSortKeySELLINGPARTNERFUNDEDTIEREDDISCOUNTPERCENTAGE ListOffersSortKey = "SELLING_PARTNER_FUNDED_TIERED_DISCOUNT_PERCENTAGE"

	// ListOffersSortKeyAMAZONFUNDEDBASEDISCOUNTPERCENTAGE captures enum value "AMAZON_FUNDED_BASE_DISCOUNT_PERCENTAGE"
	ListOffersSortKeyAMAZONFUNDEDBASEDISCOUNTPERCENTAGE ListOffersSortKey = "AMAZON_FUNDED_BASE_DISCOUNT_PERCENTAGE"

	// ListOffersSortKeyAMAZONFUNDEDTIEREDDISCOUNTPERCENTAGE captures enum value "AMAZON_FUNDED_TIERED_DISCOUNT_PERCENTAGE"
	ListOffersSortKeyAMAZONFUNDEDTIEREDDISCOUNTPERCENTAGE ListOffersSortKey = "AMAZON_FUNDED_TIERED_DISCOUNT_PERCENTAGE"
)

// for schema
var listOffersSortKeyEnum []interface{}

func init() {
	var res []ListOffersSortKey
	if err := json.Unmarshal([]byte(`["ASIN","SELLING_PARTNER_FUNDED_BASE_DISCOUNT_PERCENTAGE","SELLING_PARTNER_FUNDED_TIERED_DISCOUNT_PERCENTAGE","AMAZON_FUNDED_BASE_DISCOUNT_PERCENTAGE","AMAZON_FUNDED_TIERED_DISCOUNT_PERCENTAGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listOffersSortKeyEnum = append(listOffersSortKeyEnum, v)
	}
}

func (m ListOffersSortKey) validateListOffersSortKeyEnum(path, location string, value ListOffersSortKey) error {
	if err := validate.EnumCase(path, location, value, listOffersSortKeyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this list offers sort key
func (m ListOffersSortKey) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateListOffersSortKeyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this list offers sort key based on context it is used
func (m ListOffersSortKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
