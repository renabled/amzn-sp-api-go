// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

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

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInbound_2024-03-20/fulfillment_inbound_2024_03_20_models"
)

// NewSetPrepDetailsParams creates a new SetPrepDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetPrepDetailsParams() *SetPrepDetailsParams {
	return &SetPrepDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetPrepDetailsParamsWithTimeout creates a new SetPrepDetailsParams object
// with the ability to set a timeout on a request.
func NewSetPrepDetailsParamsWithTimeout(timeout time.Duration) *SetPrepDetailsParams {
	return &SetPrepDetailsParams{
		timeout: timeout,
	}
}

// NewSetPrepDetailsParamsWithContext creates a new SetPrepDetailsParams object
// with the ability to set a context for a request.
func NewSetPrepDetailsParamsWithContext(ctx context.Context) *SetPrepDetailsParams {
	return &SetPrepDetailsParams{
		Context: ctx,
	}
}

// NewSetPrepDetailsParamsWithHTTPClient creates a new SetPrepDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetPrepDetailsParamsWithHTTPClient(client *http.Client) *SetPrepDetailsParams {
	return &SetPrepDetailsParams{
		HTTPClient: client,
	}
}

/*
SetPrepDetailsParams contains all the parameters to send to the API endpoint

	for the set prep details operation.

	Typically these are written to a http.Request.
*/
type SetPrepDetailsParams struct {

	/* Body.

	   The body of the request to `setPrepDetails`.
	*/
	Body *fulfillment_inbound_2024_03_20_models.SetPrepDetailsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set prep details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPrepDetailsParams) WithDefaults() *SetPrepDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set prep details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPrepDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set prep details params
func (o *SetPrepDetailsParams) WithTimeout(timeout time.Duration) *SetPrepDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set prep details params
func (o *SetPrepDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set prep details params
func (o *SetPrepDetailsParams) WithContext(ctx context.Context) *SetPrepDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set prep details params
func (o *SetPrepDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set prep details params
func (o *SetPrepDetailsParams) WithHTTPClient(client *http.Client) *SetPrepDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set prep details params
func (o *SetPrepDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set prep details params
func (o *SetPrepDetailsParams) WithBody(body *fulfillment_inbound_2024_03_20_models.SetPrepDetailsRequest) *SetPrepDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set prep details params
func (o *SetPrepDetailsParams) SetBody(body *fulfillment_inbound_2024_03_20_models.SetPrepDetailsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetPrepDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
