// Code generated by go-swagger; DO NOT EDIT.

package replenishment_2022_11_07_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TimeInterval A date-time interval in ISO 8601 format which is used to compute metrics. Only the date is required, but you must pass the complete date and time value. For example, November 11, 2022 should be passed as "2022-11-07T00:00:00Z". Note that only data for the trailing 2 years is supported.
//
//	**Note**: The `listOfferMetrics` operation only supports a time interval which covers a single unit of the aggregation frequency. For example, for a MONTH aggregation frequency, the duration of the interval between the startDate and endDate can not be more than 1 month.
//
// swagger:model TimeInterval
type TimeInterval struct {

	// When this object is used as a request parameter, the specified endDate is adjusted based on the aggregation frequency.
	//
	// * For WEEK the metric is computed up to the last day of the week (that is, Sunday based on ISO 8601) that contains the endDate.
	// * For MONTH, the metric is computed up to the last day that contains the endDate.
	// * For QUARTER the metric is computed up to the last day of the quarter that contains the endDate.
	// * For YEAR the metric is computed up to the last day of the year that contains the endDate.
	//  Note: The end date may be adjusted to a lower value based on the data available in our system.
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate"`

	// When this object is used as a request parameter, the specified startDate is adjusted based on the aggregation frequency.
	//
	// * For WEEK the metric is computed from the first day of the week (that is, Sunday based on ISO 8601) that contains the startDate.
	// * For MONTH the metric is computed from the first day of the month that contains the startDate.
	// * For QUARTER the metric is computed from the first day of the quarter that contains the startDate.
	// * For YEAR the metric is computed from the first day of the year that contains the startDate.
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate"`
}

// Validate validates this time interval
func (m *TimeInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimeInterval) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("endDate", "body", m.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimeInterval) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", m.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this time interval based on context it is used
func (m *TimeInterval) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeInterval) UnmarshalBinary(b []byte) error {
	var res TimeInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}