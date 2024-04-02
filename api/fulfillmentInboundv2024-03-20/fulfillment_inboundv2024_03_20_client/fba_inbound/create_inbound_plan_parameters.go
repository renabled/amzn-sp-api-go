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

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundv2024-03-20/fulfillment_inboundv2024_03_20_models"
)

// NewCreateInboundPlanParams creates a new CreateInboundPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInboundPlanParams() *CreateInboundPlanParams {
	return &CreateInboundPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInboundPlanParamsWithTimeout creates a new CreateInboundPlanParams object
// with the ability to set a timeout on a request.
func NewCreateInboundPlanParamsWithTimeout(timeout time.Duration) *CreateInboundPlanParams {
	return &CreateInboundPlanParams{
		timeout: timeout,
	}
}

// NewCreateInboundPlanParamsWithContext creates a new CreateInboundPlanParams object
// with the ability to set a context for a request.
func NewCreateInboundPlanParamsWithContext(ctx context.Context) *CreateInboundPlanParams {
	return &CreateInboundPlanParams{
		Context: ctx,
	}
}

// NewCreateInboundPlanParamsWithHTTPClient creates a new CreateInboundPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInboundPlanParamsWithHTTPClient(client *http.Client) *CreateInboundPlanParams {
	return &CreateInboundPlanParams{
		HTTPClient: client,
	}
}

/*
CreateInboundPlanParams contains all the parameters to send to the API endpoint

	for the create inbound plan operation.

	Typically these are written to a http.Request.
*/
type CreateInboundPlanParams struct {

	/* Body.

	   The body of the request to `createInboundPlan`.
	*/
	Body *fulfillment_inboundv2024_03_20_models.CreateInboundPlanRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create inbound plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInboundPlanParams) WithDefaults() *CreateInboundPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create inbound plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInboundPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create inbound plan params
func (o *CreateInboundPlanParams) WithTimeout(timeout time.Duration) *CreateInboundPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create inbound plan params
func (o *CreateInboundPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create inbound plan params
func (o *CreateInboundPlanParams) WithContext(ctx context.Context) *CreateInboundPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create inbound plan params
func (o *CreateInboundPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create inbound plan params
func (o *CreateInboundPlanParams) WithHTTPClient(client *http.Client) *CreateInboundPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create inbound plan params
func (o *CreateInboundPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create inbound plan params
func (o *CreateInboundPlanParams) WithBody(body *fulfillment_inboundv2024_03_20_models.CreateInboundPlanRequest) *CreateInboundPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create inbound plan params
func (o *CreateInboundPlanParams) SetBody(body *fulfillment_inboundv2024_03_20_models.CreateInboundPlanRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInboundPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
