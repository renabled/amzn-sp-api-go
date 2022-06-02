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

// NewCancelReportScheduleParams creates a new CancelReportScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelReportScheduleParams() *CancelReportScheduleParams {
	return &CancelReportScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelReportScheduleParamsWithTimeout creates a new CancelReportScheduleParams object
// with the ability to set a timeout on a request.
func NewCancelReportScheduleParamsWithTimeout(timeout time.Duration) *CancelReportScheduleParams {
	return &CancelReportScheduleParams{
		timeout: timeout,
	}
}

// NewCancelReportScheduleParamsWithContext creates a new CancelReportScheduleParams object
// with the ability to set a context for a request.
func NewCancelReportScheduleParamsWithContext(ctx context.Context) *CancelReportScheduleParams {
	return &CancelReportScheduleParams{
		Context: ctx,
	}
}

// NewCancelReportScheduleParamsWithHTTPClient creates a new CancelReportScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelReportScheduleParamsWithHTTPClient(client *http.Client) *CancelReportScheduleParams {
	return &CancelReportScheduleParams{
		HTTPClient: client,
	}
}

/* CancelReportScheduleParams contains all the parameters to send to the API endpoint
   for the cancel report schedule operation.

   Typically these are written to a http.Request.
*/
type CancelReportScheduleParams struct {

	/* ReportScheduleID.

	   The identifier for the report schedule. This identifier is unique only in combination with a seller ID.
	*/
	ReportScheduleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel report schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelReportScheduleParams) WithDefaults() *CancelReportScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel report schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelReportScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel report schedule params
func (o *CancelReportScheduleParams) WithTimeout(timeout time.Duration) *CancelReportScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel report schedule params
func (o *CancelReportScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel report schedule params
func (o *CancelReportScheduleParams) WithContext(ctx context.Context) *CancelReportScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel report schedule params
func (o *CancelReportScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel report schedule params
func (o *CancelReportScheduleParams) WithHTTPClient(client *http.Client) *CancelReportScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel report schedule params
func (o *CancelReportScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportScheduleID adds the reportScheduleID to the cancel report schedule params
func (o *CancelReportScheduleParams) WithReportScheduleID(reportScheduleID string) *CancelReportScheduleParams {
	o.SetReportScheduleID(reportScheduleID)
	return o
}

// SetReportScheduleID adds the reportScheduleId to the cancel report schedule params
func (o *CancelReportScheduleParams) SetReportScheduleID(reportScheduleID string) {
	o.ReportScheduleID = reportScheduleID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelReportScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param reportScheduleId
	if err := r.SetPathParam("reportScheduleId", o.ReportScheduleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
