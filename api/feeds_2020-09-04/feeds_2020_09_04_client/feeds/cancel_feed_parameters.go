// Code generated by go-swagger; DO NOT EDIT.

package feeds

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

// NewCancelFeedParams creates a new CancelFeedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelFeedParams() *CancelFeedParams {
	return &CancelFeedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelFeedParamsWithTimeout creates a new CancelFeedParams object
// with the ability to set a timeout on a request.
func NewCancelFeedParamsWithTimeout(timeout time.Duration) *CancelFeedParams {
	return &CancelFeedParams{
		timeout: timeout,
	}
}

// NewCancelFeedParamsWithContext creates a new CancelFeedParams object
// with the ability to set a context for a request.
func NewCancelFeedParamsWithContext(ctx context.Context) *CancelFeedParams {
	return &CancelFeedParams{
		Context: ctx,
	}
}

// NewCancelFeedParamsWithHTTPClient creates a new CancelFeedParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelFeedParamsWithHTTPClient(client *http.Client) *CancelFeedParams {
	return &CancelFeedParams{
		HTTPClient: client,
	}
}

/* CancelFeedParams contains all the parameters to send to the API endpoint
   for the cancel feed operation.

   Typically these are written to a http.Request.
*/
type CancelFeedParams struct {

	/* FeedID.

	   The identifier for the feed. This identifier is unique only in combination with a seller ID.
	*/
	FeedID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel feed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelFeedParams) WithDefaults() *CancelFeedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel feed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelFeedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel feed params
func (o *CancelFeedParams) WithTimeout(timeout time.Duration) *CancelFeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel feed params
func (o *CancelFeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel feed params
func (o *CancelFeedParams) WithContext(ctx context.Context) *CancelFeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel feed params
func (o *CancelFeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel feed params
func (o *CancelFeedParams) WithHTTPClient(client *http.Client) *CancelFeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel feed params
func (o *CancelFeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeedID adds the feedID to the cancel feed params
func (o *CancelFeedParams) WithFeedID(feedID string) *CancelFeedParams {
	o.SetFeedID(feedID)
	return o
}

// SetFeedID adds the feedId to the cancel feed params
func (o *CancelFeedParams) SetFeedID(feedID string) {
	o.FeedID = feedID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelFeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param feedId
	if err := r.SetPathParam("feedId", o.FeedID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
