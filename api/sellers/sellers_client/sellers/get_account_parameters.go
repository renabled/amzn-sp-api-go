// Code generated by go-swagger; DO NOT EDIT.

package sellers

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

// NewGetAccountParams creates a new GetAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountParams() *GetAccountParams {
	return &GetAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountParamsWithTimeout creates a new GetAccountParams object
// with the ability to set a timeout on a request.
func NewGetAccountParamsWithTimeout(timeout time.Duration) *GetAccountParams {
	return &GetAccountParams{
		timeout: timeout,
	}
}

// NewGetAccountParamsWithContext creates a new GetAccountParams object
// with the ability to set a context for a request.
func NewGetAccountParamsWithContext(ctx context.Context) *GetAccountParams {
	return &GetAccountParams{
		Context: ctx,
	}
}

// NewGetAccountParamsWithHTTPClient creates a new GetAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountParamsWithHTTPClient(client *http.Client) *GetAccountParams {
	return &GetAccountParams{
		HTTPClient: client,
	}
}

/*
GetAccountParams contains all the parameters to send to the API endpoint

	for the get account operation.

	Typically these are written to a http.Request.
*/
type GetAccountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountParams) WithDefaults() *GetAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account params
func (o *GetAccountParams) WithTimeout(timeout time.Duration) *GetAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account params
func (o *GetAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account params
func (o *GetAccountParams) WithContext(ctx context.Context) *GetAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account params
func (o *GetAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account params
func (o *GetAccountParams) WithHTTPClient(client *http.Client) *GetAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account params
func (o *GetAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
