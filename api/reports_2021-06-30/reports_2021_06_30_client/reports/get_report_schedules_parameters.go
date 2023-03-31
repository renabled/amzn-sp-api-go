// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetReportSchedulesParams creates a new GetReportSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReportSchedulesParams() *GetReportSchedulesParams {
	return &GetReportSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportSchedulesParamsWithTimeout creates a new GetReportSchedulesParams object
// with the ability to set a timeout on a request.
func NewGetReportSchedulesParamsWithTimeout(timeout time.Duration) *GetReportSchedulesParams {
	return &GetReportSchedulesParams{
		timeout: timeout,
	}
}

// NewGetReportSchedulesParamsWithContext creates a new GetReportSchedulesParams object
// with the ability to set a context for a request.
func NewGetReportSchedulesParamsWithContext(ctx context.Context) *GetReportSchedulesParams {
	return &GetReportSchedulesParams{
		Context: ctx,
	}
}

// NewGetReportSchedulesParamsWithHTTPClient creates a new GetReportSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReportSchedulesParamsWithHTTPClient(client *http.Client) *GetReportSchedulesParams {
	return &GetReportSchedulesParams{
		HTTPClient: client,
	}
}

/*
GetReportSchedulesParams contains all the parameters to send to the API endpoint

	for the get report schedules operation.

	Typically these are written to a http.Request.
*/
type GetReportSchedulesParams struct {

	/* ReportTypes.

	   A list of report types used to filter report schedules. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information.
	*/
	ReportTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get report schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportSchedulesParams) WithDefaults() *GetReportSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get report schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportSchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get report schedules params
func (o *GetReportSchedulesParams) WithTimeout(timeout time.Duration) *GetReportSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get report schedules params
func (o *GetReportSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get report schedules params
func (o *GetReportSchedulesParams) WithContext(ctx context.Context) *GetReportSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get report schedules params
func (o *GetReportSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get report schedules params
func (o *GetReportSchedulesParams) WithHTTPClient(client *http.Client) *GetReportSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get report schedules params
func (o *GetReportSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportTypes adds the reportTypes to the get report schedules params
func (o *GetReportSchedulesParams) WithReportTypes(reportTypes []string) *GetReportSchedulesParams {
	o.SetReportTypes(reportTypes)
	return o
}

// SetReportTypes adds the reportTypes to the get report schedules params
func (o *GetReportSchedulesParams) SetReportTypes(reportTypes []string) {
	o.ReportTypes = reportTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReportTypes != nil {

		// binding items for reportTypes
		joinedReportTypes := o.bindParamReportTypes(reg)

		// query array param reportTypes
		if err := r.SetQueryParam("reportTypes", joinedReportTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetReportSchedules binds the parameter reportTypes
func (o *GetReportSchedulesParams) bindParamReportTypes(formats strfmt.Registry) []string {
	reportTypesIR := o.ReportTypes

	var reportTypesIC []string
	for _, reportTypesIIR := range reportTypesIR { // explode []string

		reportTypesIIV := reportTypesIIR // string as string
		reportTypesIC = append(reportTypesIC, reportTypesIIV)
	}

	// items.CollectionFormat: ""
	reportTypesIS := swag.JoinByFormat(reportTypesIC, "")

	return reportTypesIS
}
