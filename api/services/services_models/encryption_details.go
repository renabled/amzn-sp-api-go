// Code generated by go-swagger; DO NOT EDIT.

package services_models

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

// EncryptionDetails Encryption details for required client-side encryption and decryption of document contents.
//
// swagger:model EncryptionDetails
type EncryptionDetails struct {

	// The vector to encrypt or decrypt the document contents using Cipher Block Chaining (CBC).
	// Required: true
	InitializationVector *string `json:"initializationVector"`

	// The encryption key used to encrypt or decrypt the document contents.
	// Required: true
	Key *string `json:"key"`

	// The encryption standard required to encrypt or decrypt the document contents.
	// Required: true
	// Enum: [AES]
	Standard *string `json:"standard"`
}

// Validate validates this encryption details
func (m *EncryptionDetails) Validate(formats strfmt.Registry) error {
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

func (m *EncryptionDetails) validateInitializationVector(formats strfmt.Registry) error {

	if err := validate.Required("initializationVector", "body", m.InitializationVector); err != nil {
		return err
	}

	return nil
}

func (m *EncryptionDetails) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

var encryptionDetailsTypeStandardPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		encryptionDetailsTypeStandardPropEnum = append(encryptionDetailsTypeStandardPropEnum, v)
	}
}

const (

	// EncryptionDetailsStandardAES captures enum value "AES"
	EncryptionDetailsStandardAES string = "AES"
)

// prop value enum
func (m *EncryptionDetails) validateStandardEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, encryptionDetailsTypeStandardPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EncryptionDetails) validateStandard(formats strfmt.Registry) error {

	if err := validate.Required("standard", "body", m.Standard); err != nil {
		return err
	}

	// value enum
	if err := m.validateStandardEnum("standard", "body", *m.Standard); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this encryption details based on context it is used
func (m *EncryptionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionDetails) UnmarshalBinary(b []byte) error {
	var res EncryptionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
