// Code generated by go-swagger; DO NOT EDIT.

package aplus_content_2020_11_01_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StandardComparisonTableModule The standard product comparison table.
//
// swagger:model StandardComparisonTableModule
type StandardComparisonTableModule struct {

	// metric row labels
	// Max Items: 10
	// Min Items: 0
	MetricRowLabels []*PlainTextItem `json:"metricRowLabels"`

	// product columns
	// Max Items: 6
	// Min Items: 0
	ProductColumns []*StandardComparisonProductBlock `json:"productColumns"`
}

// Validate validates this standard comparison table module
func (m *StandardComparisonTableModule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetricRowLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductColumns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandardComparisonTableModule) validateMetricRowLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricRowLabels) { // not required
		return nil
	}

	iMetricRowLabelsSize := int64(len(m.MetricRowLabels))

	if err := validate.MinItems("metricRowLabels", "body", iMetricRowLabelsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("metricRowLabels", "body", iMetricRowLabelsSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.MetricRowLabels); i++ {
		if swag.IsZero(m.MetricRowLabels[i]) { // not required
			continue
		}

		if m.MetricRowLabels[i] != nil {
			if err := m.MetricRowLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metricRowLabels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metricRowLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StandardComparisonTableModule) validateProductColumns(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductColumns) { // not required
		return nil
	}

	iProductColumnsSize := int64(len(m.ProductColumns))

	if err := validate.MinItems("productColumns", "body", iProductColumnsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("productColumns", "body", iProductColumnsSize, 6); err != nil {
		return err
	}

	for i := 0; i < len(m.ProductColumns); i++ {
		if swag.IsZero(m.ProductColumns[i]) { // not required
			continue
		}

		if m.ProductColumns[i] != nil {
			if err := m.ProductColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productColumns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this standard comparison table module based on the context it is used
func (m *StandardComparisonTableModule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetricRowLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandardComparisonTableModule) contextValidateMetricRowLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetricRowLabels); i++ {

		if m.MetricRowLabels[i] != nil {
			if err := m.MetricRowLabels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metricRowLabels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metricRowLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StandardComparisonTableModule) contextValidateProductColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProductColumns); i++ {

		if m.ProductColumns[i] != nil {
			if err := m.ProductColumns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productColumns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandardComparisonTableModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandardComparisonTableModule) UnmarshalBinary(b []byte) error {
	var res StandardComparisonTableModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
