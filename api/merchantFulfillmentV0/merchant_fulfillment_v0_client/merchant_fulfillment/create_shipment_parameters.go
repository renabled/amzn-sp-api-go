// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment

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

	"github.com/renabled/amzn-sp-api-go/api/merchantFulfillmentV0/merchant_fulfillment_v0_models"
)

// NewCreateShipmentParams creates a new CreateShipmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateShipmentParams() *CreateShipmentParams {
	return &CreateShipmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateShipmentParamsWithTimeout creates a new CreateShipmentParams object
// with the ability to set a timeout on a request.
func NewCreateShipmentParamsWithTimeout(timeout time.Duration) *CreateShipmentParams {
	return &CreateShipmentParams{
		timeout: timeout,
	}
}

// NewCreateShipmentParamsWithContext creates a new CreateShipmentParams object
// with the ability to set a context for a request.
func NewCreateShipmentParamsWithContext(ctx context.Context) *CreateShipmentParams {
	return &CreateShipmentParams{
		Context: ctx,
	}
}

// NewCreateShipmentParamsWithHTTPClient creates a new CreateShipmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateShipmentParamsWithHTTPClient(client *http.Client) *CreateShipmentParams {
	return &CreateShipmentParams{
		HTTPClient: client,
	}
}

/*
CreateShipmentParams contains all the parameters to send to the API endpoint

	for the create shipment operation.

	Typically these are written to a http.Request.
*/
type CreateShipmentParams struct {

	/* Body.

	   Request schema for `CreateShipment` operation.
	*/
	Body *merchant_fulfillment_v0_models.CreateShipmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create shipment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateShipmentParams) WithDefaults() *CreateShipmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create shipment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateShipmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create shipment params
func (o *CreateShipmentParams) WithTimeout(timeout time.Duration) *CreateShipmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create shipment params
func (o *CreateShipmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create shipment params
func (o *CreateShipmentParams) WithContext(ctx context.Context) *CreateShipmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create shipment params
func (o *CreateShipmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create shipment params
func (o *CreateShipmentParams) WithHTTPClient(client *http.Client) *CreateShipmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create shipment params
func (o *CreateShipmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create shipment params
func (o *CreateShipmentParams) WithBody(body *merchant_fulfillment_v0_models.CreateShipmentRequest) *CreateShipmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create shipment params
func (o *CreateShipmentParams) SetBody(body *merchant_fulfillment_v0_models.CreateShipmentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateShipmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
