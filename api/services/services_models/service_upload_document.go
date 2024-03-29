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

// ServiceUploadDocument Input for to be uploaded document.
//
// swagger:model ServiceUploadDocument
type ServiceUploadDocument struct {

	// The content length of the to-be-uploaded file
	// Required: true
	// Maximum: 5.24288e+06
	// Minimum: 1
	ContentLength *int64 `json:"contentLength"`

	// An MD5 hash of the content to be submitted to the upload destination. This value is used to determine if the data has been corrupted or tampered with during transit.
	// Pattern: ^[A-Za-z0-9\\+/]{22}={2}$
	ContentMD5 string `json:"contentMD5,omitempty"`

	// The content type of the to-be-uploaded file
	// Required: true
	// Enum: [TIFF JPG PNG JPEG GIF PDF]
	ContentType *string `json:"contentType"`
}

// Validate validates this service upload document
func (m *ServiceUploadDocument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentMD5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceUploadDocument) validateContentLength(formats strfmt.Registry) error {

	if err := validate.Required("contentLength", "body", m.ContentLength); err != nil {
		return err
	}

	if err := validate.MinimumInt("contentLength", "body", *m.ContentLength, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("contentLength", "body", *m.ContentLength, 5.24288e+06, false); err != nil {
		return err
	}

	return nil
}

func (m *ServiceUploadDocument) validateContentMD5(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentMD5) { // not required
		return nil
	}

	if err := validate.Pattern("contentMD5", "body", m.ContentMD5, `^[A-Za-z0-9\\+/]{22}={2}$`); err != nil {
		return err
	}

	return nil
}

var serviceUploadDocumentTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TIFF","JPG","PNG","JPEG","GIF","PDF"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceUploadDocumentTypeContentTypePropEnum = append(serviceUploadDocumentTypeContentTypePropEnum, v)
	}
}

const (

	// ServiceUploadDocumentContentTypeTIFF captures enum value "TIFF"
	ServiceUploadDocumentContentTypeTIFF string = "TIFF"

	// ServiceUploadDocumentContentTypeJPG captures enum value "JPG"
	ServiceUploadDocumentContentTypeJPG string = "JPG"

	// ServiceUploadDocumentContentTypePNG captures enum value "PNG"
	ServiceUploadDocumentContentTypePNG string = "PNG"

	// ServiceUploadDocumentContentTypeJPEG captures enum value "JPEG"
	ServiceUploadDocumentContentTypeJPEG string = "JPEG"

	// ServiceUploadDocumentContentTypeGIF captures enum value "GIF"
	ServiceUploadDocumentContentTypeGIF string = "GIF"

	// ServiceUploadDocumentContentTypePDF captures enum value "PDF"
	ServiceUploadDocumentContentTypePDF string = "PDF"
)

// prop value enum
func (m *ServiceUploadDocument) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceUploadDocumentTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceUploadDocument) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("contentType", "body", m.ContentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", *m.ContentType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service upload document based on context it is used
func (m *ServiceUploadDocument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceUploadDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceUploadDocument) UnmarshalBinary(b []byte) error {
	var res ServiceUploadDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
