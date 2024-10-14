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

// SetPrepDetailsRequest The `setPrepDetails` request.
// Example: {"marketplaceId":"A2EUQ1WTGCTBG2","mskuPrepDetails":[{"msku":"msku","prepCategory":"NONE","prepTypes":["ITEM_NO_PREP"]}]}
//
// swagger:model SetPrepDetailsRequest
type SetPrepDetailsRequest struct {

	// The marketplace ID. For a list of possible values, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	// Required: true
	// Max Length: 20
	// Min Length: 1
	MarketplaceID *string `json:"marketplaceId"`

	// A list of MSKUs and related prep details.
	// Required: true
	// Max Items: 100
	// Min Items: 1
	MskuPrepDetails []*MskuPrepDetailInput `json:"mskuPrepDetails"`
}

// Validate validates this set prep details request
func (m *SetPrepDetailsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMarketplaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMskuPrepDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetPrepDetailsRequest) validateMarketplaceID(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceId", "body", m.MarketplaceID); err != nil {
		return err
	}

	if err := validate.MinLength("marketplaceId", "body", *m.MarketplaceID, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("marketplaceId", "body", *m.MarketplaceID, 20); err != nil {
		return err
	}

	return nil
}

func (m *SetPrepDetailsRequest) validateMskuPrepDetails(formats strfmt.Registry) error {

	if err := validate.Required("mskuPrepDetails", "body", m.MskuPrepDetails); err != nil {
		return err
	}

	iMskuPrepDetailsSize := int64(len(m.MskuPrepDetails))

	if err := validate.MinItems("mskuPrepDetails", "body", iMskuPrepDetailsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("mskuPrepDetails", "body", iMskuPrepDetailsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(m.MskuPrepDetails); i++ {
		if swag.IsZero(m.MskuPrepDetails[i]) { // not required
			continue
		}

		if m.MskuPrepDetails[i] != nil {
			if err := m.MskuPrepDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mskuPrepDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mskuPrepDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this set prep details request based on the context it is used
func (m *SetPrepDetailsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMskuPrepDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetPrepDetailsRequest) contextValidateMskuPrepDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MskuPrepDetails); i++ {

		if m.MskuPrepDetails[i] != nil {
			if err := m.MskuPrepDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mskuPrepDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mskuPrepDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetPrepDetailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetPrepDetailsRequest) UnmarshalBinary(b []byte) error {
	var res SetPrepDetailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
