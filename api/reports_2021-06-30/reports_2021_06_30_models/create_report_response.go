// Code generated by go-swagger; DO NOT EDIT.

package reports_2021_06_30_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateReportResponse Response schema.
//
// swagger:model CreateReportResponse
type CreateReportResponse struct {

	// The identifier for the report. This identifier is unique only in combination with a seller ID.
	// Required: true
	ReportID *string `json:"reportId"`
}

// Validate validates this create report response
func (m *CreateReportResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateReportResponse) validateReportID(formats strfmt.Registry) error {

	if err := validate.Required("reportId", "body", m.ReportID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create report response based on context it is used
func (m *CreateReportResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateReportResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateReportResponse) UnmarshalBinary(b []byte) error {
	var res CreateReportResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
