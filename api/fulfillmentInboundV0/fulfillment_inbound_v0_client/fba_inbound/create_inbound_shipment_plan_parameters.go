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

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundV0/fulfillment_inbound_v0_models"
)

// NewCreateInboundShipmentPlanParams creates a new CreateInboundShipmentPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInboundShipmentPlanParams() *CreateInboundShipmentPlanParams {
	return &CreateInboundShipmentPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInboundShipmentPlanParamsWithTimeout creates a new CreateInboundShipmentPlanParams object
// with the ability to set a timeout on a request.
func NewCreateInboundShipmentPlanParamsWithTimeout(timeout time.Duration) *CreateInboundShipmentPlanParams {
	return &CreateInboundShipmentPlanParams{
		timeout: timeout,
	}
}

// NewCreateInboundShipmentPlanParamsWithContext creates a new CreateInboundShipmentPlanParams object
// with the ability to set a context for a request.
func NewCreateInboundShipmentPlanParamsWithContext(ctx context.Context) *CreateInboundShipmentPlanParams {
	return &CreateInboundShipmentPlanParams{
		Context: ctx,
	}
}

// NewCreateInboundShipmentPlanParamsWithHTTPClient creates a new CreateInboundShipmentPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInboundShipmentPlanParamsWithHTTPClient(client *http.Client) *CreateInboundShipmentPlanParams {
	return &CreateInboundShipmentPlanParams{
		HTTPClient: client,
	}
}

/*
CreateInboundShipmentPlanParams contains all the parameters to send to the API endpoint

	for the create inbound shipment plan operation.

	Typically these are written to a http.Request.
*/
type CreateInboundShipmentPlanParams struct {

	// Body.
	Body *fulfillment_inbound_v0_models.CreateInboundShipmentPlanRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create inbound shipment plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInboundShipmentPlanParams) WithDefaults() *CreateInboundShipmentPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create inbound shipment plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInboundShipmentPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) WithTimeout(timeout time.Duration) *CreateInboundShipmentPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) WithContext(ctx context.Context) *CreateInboundShipmentPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) WithHTTPClient(client *http.Client) *CreateInboundShipmentPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) WithBody(body *fulfillment_inbound_v0_models.CreateInboundShipmentPlanRequest) *CreateInboundShipmentPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create inbound shipment plan params
func (o *CreateInboundShipmentPlanParams) SetBody(body *fulfillment_inbound_v0_models.CreateInboundShipmentPlanRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInboundShipmentPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
