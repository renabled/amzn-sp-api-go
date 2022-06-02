// Code generated by go-swagger; DO NOT EDIT.

package listings_items_2020_09_01_models

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

// Issue An issue with a listings item.
//
// swagger:model Issue
type Issue struct {

	// Name of the attribute associated with the issue, if applicable.
	AttributeName string `json:"attributeName,omitempty"`

	// An issue code that identifies the type of issue.
	// Required: true
	Code *string `json:"code"`

	// A message that describes the issue.
	// Required: true
	Message *string `json:"message"`

	// The severity of the issue.
	// Required: true
	// Enum: [ERROR WARNING INFO]
	Severity *string `json:"severity"`
}

// Validate validates this issue
func (m *Issue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Issue) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Issue) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

var issueTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ERROR","WARNING","INFO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		issueTypeSeverityPropEnum = append(issueTypeSeverityPropEnum, v)
	}
}

const (

	// IssueSeverityERROR captures enum value "ERROR"
	IssueSeverityERROR string = "ERROR"

	// IssueSeverityWARNING captures enum value "WARNING"
	IssueSeverityWARNING string = "WARNING"

	// IssueSeverityINFO captures enum value "INFO"
	IssueSeverityINFO string = "INFO"
)

// prop value enum
func (m *Issue) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, issueTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Issue) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this issue based on context it is used
func (m *Issue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Issue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Issue) UnmarshalBinary(b []byte) error {
	var res Issue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
