// Code generated by go-swagger; DO NOT EDIT.

package feeds_2020_09_04_models

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

// FeedDocument feed document
//
// swagger:model FeedDocument
type FeedDocument struct {

	// If present, the feed document contents are compressed using the indicated algorithm.
	// Enum: [GZIP]
	CompressionAlgorithm string `json:"compressionAlgorithm,omitempty"`

	// encryption details
	// Required: true
	EncryptionDetails *FeedDocumentEncryptionDetails `json:"encryptionDetails"`

	// The identifier for the feed document. This identifier is unique only in combination with a seller ID.
	// Required: true
	FeedDocumentID *string `json:"feedDocumentId"`

	// A presigned URL for the feed document. This URL expires after 5 minutes.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this feed document
func (m *FeedDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompressionAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedDocumentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var feedDocumentTypeCompressionAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GZIP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		feedDocumentTypeCompressionAlgorithmPropEnum = append(feedDocumentTypeCompressionAlgorithmPropEnum, v)
	}
}

const (

	// FeedDocumentCompressionAlgorithmGZIP captures enum value "GZIP"
	FeedDocumentCompressionAlgorithmGZIP string = "GZIP"
)

// prop value enum
func (m *FeedDocument) validateCompressionAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, feedDocumentTypeCompressionAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FeedDocument) validateCompressionAlgorithm(formats strfmt.Registry) error {
	if swag.IsZero(m.CompressionAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateCompressionAlgorithmEnum("compressionAlgorithm", "body", m.CompressionAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *FeedDocument) validateEncryptionDetails(formats strfmt.Registry) error {

	if err := validate.Required("encryptionDetails", "body", m.EncryptionDetails); err != nil {
		return err
	}

	if m.EncryptionDetails != nil {
		if err := m.EncryptionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionDetails")
			}
			return err
		}
	}

	return nil
}

func (m *FeedDocument) validateFeedDocumentID(formats strfmt.Registry) error {

	if err := validate.Required("feedDocumentId", "body", m.FeedDocumentID); err != nil {
		return err
	}

	return nil
}

func (m *FeedDocument) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this feed document based on the context it is used
func (m *FeedDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncryptionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeedDocument) contextValidateEncryptionDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionDetails != nil {
		if err := m.EncryptionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeedDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeedDocument) UnmarshalBinary(b []byte) error {
	var res FeedDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
