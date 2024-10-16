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

	"github.com/renabled/amzn-sp-api-go/api/feeds_2021-06-30/feeds_2021_06_30_models"
)

// NewCreateFeedParams creates a new CreateFeedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFeedParams() *CreateFeedParams {
	return &CreateFeedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFeedParamsWithTimeout creates a new CreateFeedParams object
// with the ability to set a timeout on a request.
func NewCreateFeedParamsWithTimeout(timeout time.Duration) *CreateFeedParams {
	return &CreateFeedParams{
		timeout: timeout,
	}
}

// NewCreateFeedParamsWithContext creates a new CreateFeedParams object
// with the ability to set a context for a request.
func NewCreateFeedParamsWithContext(ctx context.Context) *CreateFeedParams {
	return &CreateFeedParams{
		Context: ctx,
	}
}

// NewCreateFeedParamsWithHTTPClient creates a new CreateFeedParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFeedParamsWithHTTPClient(client *http.Client) *CreateFeedParams {
	return &CreateFeedParams{
		HTTPClient: client,
	}
}

/*
CreateFeedParams contains all the parameters to send to the API endpoint

	for the create feed operation.

	Typically these are written to a http.Request.
*/
type CreateFeedParams struct {

	/* Body.

	   Information required to create the feed.
	*/
	Body *feeds_2021_06_30_models.CreateFeedSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create feed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFeedParams) WithDefaults() *CreateFeedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create feed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFeedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create feed params
func (o *CreateFeedParams) WithTimeout(timeout time.Duration) *CreateFeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create feed params
func (o *CreateFeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create feed params
func (o *CreateFeedParams) WithContext(ctx context.Context) *CreateFeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create feed params
func (o *CreateFeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create feed params
func (o *CreateFeedParams) WithHTTPClient(client *http.Client) *CreateFeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create feed params
func (o *CreateFeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create feed params
func (o *CreateFeedParams) WithBody(body *feeds_2021_06_30_models.CreateFeedSpecification) *CreateFeedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create feed params
func (o *CreateFeedParams) SetBody(body *feeds_2021_06_30_models.CreateFeedSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
