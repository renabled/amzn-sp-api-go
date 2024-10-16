// Code generated by go-swagger; DO NOT EDIT.

package reports_2021_06_30_models

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

// Report Detailed information about the report.
//
// swagger:model Report
type Report struct {

	// The date and time when the report was created.
	// Required: true
	// Format: date-time
	CreatedTime *strfmt.DateTime `json:"createdTime"`

	// The end of a date and time range used for selecting the data to report.
	// Format: date-time
	DataEndTime strfmt.DateTime `json:"dataEndTime,omitempty"`

	// The start of a date and time range used for selecting the data to report.
	// Format: date-time
	DataStartTime strfmt.DateTime `json:"dataStartTime,omitempty"`

	// A list of marketplace identifiers for the report.
	MarketplaceIds []string `json:"marketplaceIds"`

	// The date and time when the report processing completed, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format.
	// Format: date-time
	ProcessingEndTime strfmt.DateTime `json:"processingEndTime,omitempty"`

	// The date and time when the report processing started, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format.
	// Format: date-time
	ProcessingStartTime strfmt.DateTime `json:"processingStartTime,omitempty"`

	// The processing status of the report.
	// Required: true
	// Enum: [CANCELLED DONE FATAL IN_PROGRESS IN_QUEUE]
	ProcessingStatus *string `json:"processingStatus"`

	// The identifier for the report document. Pass this into the `getReportDocument` operation to get the information you will need to retrieve the report document's contents.
	ReportDocumentID string `json:"reportDocumentId,omitempty"`

	// The identifier for the report. This identifier is unique only in combination with a seller ID.
	// Required: true
	ReportID *string `json:"reportId"`

	// The identifier of the report schedule that created this report (if any). This identifier is unique only in combination with a seller ID.
	ReportScheduleID string `json:"reportScheduleId,omitempty"`

	// The report type. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information.
	// Required: true
	ReportType *string `json:"reportType"`
}

// Validate validates this report
func (m *Report) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Report) validateCreatedTime(formats strfmt.Registry) error {

	if err := validate.Required("createdTime", "body", m.CreatedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateDataEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DataEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dataEndTime", "body", "date-time", m.DataEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateDataStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.DataStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dataStartTime", "body", "date-time", m.DataStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateProcessingEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("processingEndTime", "body", "date-time", m.ProcessingEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateProcessingStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("processingStartTime", "body", "date-time", m.ProcessingStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var reportTypeProcessingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CANCELLED","DONE","FATAL","IN_PROGRESS","IN_QUEUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportTypeProcessingStatusPropEnum = append(reportTypeProcessingStatusPropEnum, v)
	}
}

const (

	// ReportProcessingStatusCANCELLED captures enum value "CANCELLED"
	ReportProcessingStatusCANCELLED string = "CANCELLED"

	// ReportProcessingStatusDONE captures enum value "DONE"
	ReportProcessingStatusDONE string = "DONE"

	// ReportProcessingStatusFATAL captures enum value "FATAL"
	ReportProcessingStatusFATAL string = "FATAL"

	// ReportProcessingStatusINPROGRESS captures enum value "IN_PROGRESS"
	ReportProcessingStatusINPROGRESS string = "IN_PROGRESS"

	// ReportProcessingStatusINQUEUE captures enum value "IN_QUEUE"
	ReportProcessingStatusINQUEUE string = "IN_QUEUE"
)

// prop value enum
func (m *Report) validateProcessingStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportTypeProcessingStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Report) validateProcessingStatus(formats strfmt.Registry) error {

	if err := validate.Required("processingStatus", "body", m.ProcessingStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateProcessingStatusEnum("processingStatus", "body", *m.ProcessingStatus); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateReportID(formats strfmt.Registry) error {

	if err := validate.Required("reportId", "body", m.ReportID); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateReportType(formats strfmt.Registry) error {

	if err := validate.Required("reportType", "body", m.ReportType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report based on context it is used
func (m *Report) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Report) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Report) UnmarshalBinary(b []byte) error {
	var res Report
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
