// Code generated by go-swagger; DO NOT EDIT.

package reports_2020_09_04_models

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

// ReportDocumentEncryptionDetails Encryption details required for decryption of a report document's contents.
//
// swagger:model ReportDocumentEncryptionDetails
type ReportDocumentEncryptionDetails struct {

	// The vector to decrypt the document contents using Cipher Block Chaining (CBC).
	// Required: true
	InitializationVector *string `json:"initializationVector"`

	// The encryption key used to decrypt the document contents.
	// Required: true
	Key *string `json:"key"`

	// The encryption standard required to decrypt the document contents.
	// Required: true
	// Enum: [AES]
	Standard *string `json:"standard"`
}

// Validate validates this report document encryption details
func (m *ReportDocumentEncryptionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInitializationVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportDocumentEncryptionDetails) validateInitializationVector(formats strfmt.Registry) error {

	if err := validate.Required("initializationVector", "body", m.InitializationVector); err != nil {
		return err
	}

	return nil
}

func (m *ReportDocumentEncryptionDetails) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

var reportDocumentEncryptionDetailsTypeStandardPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportDocumentEncryptionDetailsTypeStandardPropEnum = append(reportDocumentEncryptionDetailsTypeStandardPropEnum, v)
	}
}

const (

	// ReportDocumentEncryptionDetailsStandardAES captures enum value "AES"
	ReportDocumentEncryptionDetailsStandardAES string = "AES"
)

// prop value enum
func (m *ReportDocumentEncryptionDetails) validateStandardEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportDocumentEncryptionDetailsTypeStandardPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportDocumentEncryptionDetails) validateStandard(formats strfmt.Registry) error {

	if err := validate.Required("standard", "body", m.Standard); err != nil {
		return err
	}

	// value enum
	if err := m.validateStandardEnum("standard", "body", *m.Standard); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report document encryption details based on context it is used
func (m *ReportDocumentEncryptionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportDocumentEncryptionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportDocumentEncryptionDetails) UnmarshalBinary(b []byte) error {
	var res ReportDocumentEncryptionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
