// Code generated by go-swagger; DO NOT EDIT.

package applications

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

// NewRotateApplicationClientSecretParams creates a new RotateApplicationClientSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRotateApplicationClientSecretParams() *RotateApplicationClientSecretParams {
	return &RotateApplicationClientSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRotateApplicationClientSecretParamsWithTimeout creates a new RotateApplicationClientSecretParams object
// with the ability to set a timeout on a request.
func NewRotateApplicationClientSecretParamsWithTimeout(timeout time.Duration) *RotateApplicationClientSecretParams {
	return &RotateApplicationClientSecretParams{
		timeout: timeout,
	}
}

// NewRotateApplicationClientSecretParamsWithContext creates a new RotateApplicationClientSecretParams object
// with the ability to set a context for a request.
func NewRotateApplicationClientSecretParamsWithContext(ctx context.Context) *RotateApplicationClientSecretParams {
	return &RotateApplicationClientSecretParams{
		Context: ctx,
	}
}

// NewRotateApplicationClientSecretParamsWithHTTPClient creates a new RotateApplicationClientSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewRotateApplicationClientSecretParamsWithHTTPClient(client *http.Client) *RotateApplicationClientSecretParams {
	return &RotateApplicationClientSecretParams{
		HTTPClient: client,
	}
}

/*
RotateApplicationClientSecretParams contains all the parameters to send to the API endpoint

	for the rotate application client secret operation.

	Typically these are written to a http.Request.
*/
type RotateApplicationClientSecretParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rotate application client secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RotateApplicationClientSecretParams) WithDefaults() *RotateApplicationClientSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rotate application client secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RotateApplicationClientSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rotate application client secret params
func (o *RotateApplicationClientSecretParams) WithTimeout(timeout time.Duration) *RotateApplicationClientSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rotate application client secret params
func (o *RotateApplicationClientSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rotate application client secret params
func (o *RotateApplicationClientSecretParams) WithContext(ctx context.Context) *RotateApplicationClientSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rotate application client secret params
func (o *RotateApplicationClientSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rotate application client secret params
func (o *RotateApplicationClientSecretParams) WithHTTPClient(client *http.Client) *RotateApplicationClientSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rotate application client secret params
func (o *RotateApplicationClientSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RotateApplicationClientSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
