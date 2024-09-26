// Code generated by go-swagger; DO NOT EDIT.

package invoices_api_model_2024_06_19_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FileFormat Supported invoice file extensions.
//
// swagger:model FileFormat
type FileFormat string

func NewFileFormat(value FileFormat) *FileFormat {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FileFormat.
func (m FileFormat) Pointer() *FileFormat {
	return &m
}

const (

	// FileFormatXML captures enum value "XML"
	FileFormatXML FileFormat = "XML"
)

// for schema
var fileFormatEnum []interface{}

func init() {
	var res []FileFormat
	if err := json.Unmarshal([]byte(`["XML"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileFormatEnum = append(fileFormatEnum, v)
	}
}

func (m FileFormat) validateFileFormatEnum(path, location string, value FileFormat) error {
	if err := validate.EnumCase(path, location, value, fileFormatEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this file format
func (m FileFormat) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFileFormatEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this file format based on context it is used
func (m FileFormat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
