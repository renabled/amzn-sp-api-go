// Code generated by go-swagger; DO NOT EDIT.

package data_kiosk_2023_11_15_models

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

// Query Detailed information about the query.
//
// swagger:model Query
type Query struct {

	// The date and time when the query was created, in ISO 8601 date time format.
	// Required: true
	// Format: date-time
	CreatedTime *strfmt.DateTime `json:"createdTime"`

	// The data document identifier. This identifier is only present when there is data available as a result of the query. This identifier is unique only in combination with a selling partner account ID. Pass this identifier into the `getDocument` operation to get the information required to retrieve the data document's contents.
	DataDocumentID string `json:"dataDocumentId,omitempty"`

	// The error document identifier. This identifier is only present when an error occurs during query processing. This identifier is unique only in combination with a selling partner account ID. Pass this identifier into the `getDocument` operation to get the information required to retrieve the error document's contents.
	ErrorDocumentID string `json:"errorDocumentId,omitempty"`

	// pagination
	Pagination *QueryPagination `json:"pagination,omitempty"`

	// The date and time when the query processing completed, in ISO 8601 date time format.
	// Format: date-time
	ProcessingEndTime strfmt.DateTime `json:"processingEndTime,omitempty"`

	// The date and time when the query processing started, in ISO 8601 date time format.
	// Format: date-time
	ProcessingStartTime strfmt.DateTime `json:"processingStartTime,omitempty"`

	// The processing status of the query.
	// Required: true
	// Enum: [CANCELLED DONE FATAL IN_PROGRESS IN_QUEUE]
	ProcessingStatus *string `json:"processingStatus"`

	// The submitted query.
	// Required: true
	Query *string `json:"query"`

	// The query identifier. This identifier is unique only in combination with a selling partner account ID.
	// Required: true
	QueryID *string `json:"queryId"`
}

// Validate validates this query
func (m *Query) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
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

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Query) validateCreatedTime(formats strfmt.Registry) error {

	if err := validate.Required("createdTime", "body", m.CreatedTime); err != nil {
		return err
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Query) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *Query) validateProcessingEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("processingEndTime", "body", "date-time", m.ProcessingEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Query) validateProcessingStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcessingStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("processingStartTime", "body", "date-time", m.ProcessingStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var queryTypeProcessingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CANCELLED","DONE","FATAL","IN_PROGRESS","IN_QUEUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queryTypeProcessingStatusPropEnum = append(queryTypeProcessingStatusPropEnum, v)
	}
}

const (

	// QueryProcessingStatusCANCELLED captures enum value "CANCELLED"
	QueryProcessingStatusCANCELLED string = "CANCELLED"

	// QueryProcessingStatusDONE captures enum value "DONE"
	QueryProcessingStatusDONE string = "DONE"

	// QueryProcessingStatusFATAL captures enum value "FATAL"
	QueryProcessingStatusFATAL string = "FATAL"

	// QueryProcessingStatusINPROGRESS captures enum value "IN_PROGRESS"
	QueryProcessingStatusINPROGRESS string = "IN_PROGRESS"

	// QueryProcessingStatusINQUEUE captures enum value "IN_QUEUE"
	QueryProcessingStatusINQUEUE string = "IN_QUEUE"
)

// prop value enum
func (m *Query) validateProcessingStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, queryTypeProcessingStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Query) validateProcessingStatus(formats strfmt.Registry) error {

	if err := validate.Required("processingStatus", "body", m.ProcessingStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateProcessingStatusEnum("processingStatus", "body", *m.ProcessingStatus); err != nil {
		return err
	}

	return nil
}

func (m *Query) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}

func (m *Query) validateQueryID(formats strfmt.Registry) error {

	if err := validate.Required("queryId", "body", m.QueryID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this query based on the context it is used
func (m *Query) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Query) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Query) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Query) UnmarshalBinary(b []byte) error {
	var res Query
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QueryPagination When a query produces results that are not included in the data document, pagination occurs. This means the results are divided into pages. To retrieve the next page, you must pass a `CreateQuerySpecification` object with `paginationToken` set to this object's `nextToken` and with `query` set to this object's `query` in the subsequent `createQuery` request. When there are no more pages to fetch, the `nextToken` field will be absent.
//
// swagger:model QueryPagination
type QueryPagination struct {

	// A token that can be used to fetch the next page of results.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this query pagination
func (m *QueryPagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this query pagination based on context it is used
func (m *QueryPagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryPagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryPagination) UnmarshalBinary(b []byte) error {
	var res QueryPagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
