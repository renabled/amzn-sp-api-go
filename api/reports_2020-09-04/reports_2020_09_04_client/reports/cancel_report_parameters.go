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
)

// NewCancelReportParams creates a new CancelReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelReportParams() *CancelReportParams {
	return &CancelReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelReportParamsWithTimeout creates a new CancelReportParams object
// with the ability to set a timeout on a request.
func NewCancelReportParamsWithTimeout(timeout time.Duration) *CancelReportParams {
	return &CancelReportParams{
		timeout: timeout,
	}
}

// NewCancelReportParamsWithContext creates a new CancelReportParams object
// with the ability to set a context for a request.
func NewCancelReportParamsWithContext(ctx context.Context) *CancelReportParams {
	return &CancelReportParams{
		Context: ctx,
	}
}

// NewCancelReportParamsWithHTTPClient creates a new CancelReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelReportParamsWithHTTPClient(client *http.Client) *CancelReportParams {
	return &CancelReportParams{
		HTTPClient: client,
	}
}

/* CancelReportParams contains all the parameters to send to the API endpoint
   for the cancel report operation.

   Typically these are written to a http.Request.
*/
type CancelReportParams struct {

	/* ReportID.

	   The identifier for the report. This identifier is unique only in combination with a seller ID.
	*/
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelReportParams) WithDefaults() *CancelReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel report params
func (o *CancelReportParams) WithTimeout(timeout time.Duration) *CancelReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel report params
func (o *CancelReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel report params
func (o *CancelReportParams) WithContext(ctx context.Context) *CancelReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel report params
func (o *CancelReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel report params
func (o *CancelReportParams) WithHTTPClient(client *http.Client) *CancelReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel report params
func (o *CancelReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportID adds the reportID to the cancel report params
func (o *CancelReportParams) WithReportID(reportID string) *CancelReportParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the cancel report params
func (o *CancelReportParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param reportId
	if err := r.SetPathParam("reportId", o.ReportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
