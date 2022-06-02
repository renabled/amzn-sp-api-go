// Code generated by go-swagger; DO NOT EDIT.

package feeds_2020_09_04_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateFeedSpecification create feed specification
//
// swagger:model CreateFeedSpecification
type CreateFeedSpecification struct {

	// feed options
	FeedOptions FeedOptions `json:"feedOptions,omitempty"`

	// The feed type.
	// Required: true
	FeedType *string `json:"feedType"`

	// The document identifier returned by the createFeedDocument operation. Encrypt and upload the feed document contents before calling the createFeed operation.
	// Required: true
	InputFeedDocumentID *string `json:"inputFeedDocumentId"`

	// A list of identifiers for marketplaces that you want the feed to be applied to.
	// Required: true
	// Max Items: 25
	// Min Items: 1
	MarketplaceIds []string `json:"marketplaceIds"`
}

// Validate validates this create feed specification
func (m *CreateFeedSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeedOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputFeedDocumentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketplaceIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFeedSpecification) validateFeedOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.FeedOptions) { // not required
		return nil
	}

	if m.FeedOptions != nil {
		if err := m.FeedOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feedOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("feedOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateFeedSpecification) validateFeedType(formats strfmt.Registry) error {

	if err := validate.Required("feedType", "body", m.FeedType); err != nil {
		return err
	}

	return nil
}

func (m *CreateFeedSpecification) validateInputFeedDocumentID(formats strfmt.Registry) error {

	if err := validate.Required("inputFeedDocumentId", "body", m.InputFeedDocumentID); err != nil {
		return err
	}

	return nil
}

func (m *CreateFeedSpecification) validateMarketplaceIds(formats strfmt.Registry) error {

	if err := validate.Required("marketplaceIds", "body", m.MarketplaceIds); err != nil {
		return err
	}

	iMarketplaceIdsSize := int64(len(m.MarketplaceIds))

	if err := validate.MinItems("marketplaceIds", "body", iMarketplaceIdsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("marketplaceIds", "body", iMarketplaceIdsSize, 25); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create feed specification based on the context it is used
func (m *CreateFeedSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeedOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateFeedSpecification) contextValidateFeedOptions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FeedOptions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("feedOptions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("feedOptions")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateFeedSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFeedSpecification) UnmarshalBinary(b []byte) error {
	var res CreateFeedSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
