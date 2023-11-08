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

// CompetitiveSummaryBatchResponse The response schema of the `competitiveSummaryBatch` operation.
//
// swagger:model CompetitiveSummaryBatchResponse
type CompetitiveSummaryBatchResponse struct {

	// The response list of the `competitiveSummaryBatch` operation.
	// Required: true
	Responses CompetitiveSummaryResponseList `json:"responses"`
}

// Validate validates this competitive summary batch response
func (m *CompetitiveSummaryBatchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompetitiveSummaryBatchResponse) validateResponses(formats strfmt.Registry) error {

	if err := validate.Required("responses", "body", m.Responses); err != nil {
		return err
	}

	if err := m.Responses.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("responses")
		}
		return err
	}

	return nil
}

// ContextValidate validate this competitive summary batch response based on the context it is used
func (m *CompetitiveSummaryBatchResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResponses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompetitiveSummaryBatchResponse) contextValidateResponses(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Responses.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responses")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("responses")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompetitiveSummaryBatchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompetitiveSummaryBatchResponse) UnmarshalBinary(b []byte) error {
	var res CompetitiveSummaryBatchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
