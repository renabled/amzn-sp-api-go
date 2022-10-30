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

	"github.com/xamandar/amzn-sp-api-go/api/feeds_2021-06-30/feeds_2021_06_30_models"
)

// NewCreateFeedDocumentParams creates a new CreateFeedDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFeedDocumentParams() *CreateFeedDocumentParams {
	return &CreateFeedDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFeedDocumentParamsWithTimeout creates a new CreateFeedDocumentParams object
// with the ability to set a timeout on a request.
func NewCreateFeedDocumentParamsWithTimeout(timeout time.Duration) *CreateFeedDocumentParams {
	return &CreateFeedDocumentParams{
		timeout: timeout,
	}
}

// NewCreateFeedDocumentParamsWithContext creates a new CreateFeedDocumentParams object
// with the ability to set a context for a request.
func NewCreateFeedDocumentParamsWithContext(ctx context.Context) *CreateFeedDocumentParams {
	return &CreateFeedDocumentParams{
		Context: ctx,
	}
}

// NewCreateFeedDocumentParamsWithHTTPClient creates a new CreateFeedDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFeedDocumentParamsWithHTTPClient(client *http.Client) *CreateFeedDocumentParams {
	return &CreateFeedDocumentParams{
		HTTPClient: client,
	}
}

/*
CreateFeedDocumentParams contains all the parameters to send to the API endpoint

	for the create feed document operation.

	Typically these are written to a http.Request.
*/
type CreateFeedDocumentParams struct {

	// Body.
	Body *feeds_2021_06_30_models.CreateFeedDocumentSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create feed document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFeedDocumentParams) WithDefaults() *CreateFeedDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create feed document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFeedDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create feed document params
func (o *CreateFeedDocumentParams) WithTimeout(timeout time.Duration) *CreateFeedDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create feed document params
func (o *CreateFeedDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create feed document params
func (o *CreateFeedDocumentParams) WithContext(ctx context.Context) *CreateFeedDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create feed document params
func (o *CreateFeedDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create feed document params
func (o *CreateFeedDocumentParams) WithHTTPClient(client *http.Client) *CreateFeedDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create feed document params
func (o *CreateFeedDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create feed document params
func (o *CreateFeedDocumentParams) WithBody(body *feeds_2021_06_30_models.CreateFeedDocumentSpecification) *CreateFeedDocumentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create feed document params
func (o *CreateFeedDocumentParams) SetBody(body *feeds_2021_06_30_models.CreateFeedDocumentSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFeedDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
