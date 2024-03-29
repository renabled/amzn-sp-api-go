// Code generated by go-swagger; DO NOT EDIT.

package product_pricing_2022_05_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CompetitiveSummaryBatchRequest The `competitiveSummary` batch request data.
//
// swagger:model CompetitiveSummaryBatchRequest
type CompetitiveSummaryBatchRequest struct {

	// A batched list of `competitiveSummary` requests.
	// Required: true
	Requests CompetitiveSummaryRequestList `json:"requests"`
}

// Validate validates this competitive summary batch request
func (m *CompetitiveSummaryBatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompetitiveSummaryBatchRequest) validateRequests(formats strfmt.Registry) error {

	if err := validate.Required("requests", "body", m.Requests); err != nil {
		return err
	}

	if err := m.Requests.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requests")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requests")
		}
		return err
	}

	return nil
}

// ContextValidate validate this competitive summary batch request based on the context it is used
func (m *CompetitiveSummaryBatchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompetitiveSummaryBatchRequest) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Requests.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("requests")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("requests")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompetitiveSummaryBatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompetitiveSummaryBatchRequest) UnmarshalBinary(b []byte) error {
	var res CompetitiveSummaryBatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
