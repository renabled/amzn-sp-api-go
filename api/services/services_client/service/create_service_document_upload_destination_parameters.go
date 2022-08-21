// Code generated by go-swagger; DO NOT EDIT.

package service

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

	"github.com/xamandar/amzn-sp-api-go/api/services/services_models"
)

// NewCreateServiceDocumentUploadDestinationParams creates a new CreateServiceDocumentUploadDestinationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServiceDocumentUploadDestinationParams() *CreateServiceDocumentUploadDestinationParams {
	return &CreateServiceDocumentUploadDestinationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceDocumentUploadDestinationParamsWithTimeout creates a new CreateServiceDocumentUploadDestinationParams object
// with the ability to set a timeout on a request.
func NewCreateServiceDocumentUploadDestinationParamsWithTimeout(timeout time.Duration) *CreateServiceDocumentUploadDestinationParams {
	return &CreateServiceDocumentUploadDestinationParams{
		timeout: timeout,
	}
}

// NewCreateServiceDocumentUploadDestinationParamsWithContext creates a new CreateServiceDocumentUploadDestinationParams object
// with the ability to set a context for a request.
func NewCreateServiceDocumentUploadDestinationParamsWithContext(ctx context.Context) *CreateServiceDocumentUploadDestinationParams {
	return &CreateServiceDocumentUploadDestinationParams{
		Context: ctx,
	}
}

// NewCreateServiceDocumentUploadDestinationParamsWithHTTPClient creates a new CreateServiceDocumentUploadDestinationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServiceDocumentUploadDestinationParamsWithHTTPClient(client *http.Client) *CreateServiceDocumentUploadDestinationParams {
	return &CreateServiceDocumentUploadDestinationParams{
		HTTPClient: client,
	}
}

/* CreateServiceDocumentUploadDestinationParams contains all the parameters to send to the API endpoint
   for the create service document upload destination operation.

   Typically these are written to a http.Request.
*/
type CreateServiceDocumentUploadDestinationParams struct {

	/* Body.

	   Upload document operation input details.
	*/
	Body *services_models.ServiceUploadDocument

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create service document upload destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceDocumentUploadDestinationParams) WithDefaults() *CreateServiceDocumentUploadDestinationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create service document upload destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceDocumentUploadDestinationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) WithTimeout(timeout time.Duration) *CreateServiceDocumentUploadDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) WithContext(ctx context.Context) *CreateServiceDocumentUploadDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) WithHTTPClient(client *http.Client) *CreateServiceDocumentUploadDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) WithBody(body *services_models.ServiceUploadDocument) *CreateServiceDocumentUploadDestinationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create service document upload destination params
func (o *CreateServiceDocumentUploadDestinationParams) SetBody(body *services_models.ServiceUploadDocument) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceDocumentUploadDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
