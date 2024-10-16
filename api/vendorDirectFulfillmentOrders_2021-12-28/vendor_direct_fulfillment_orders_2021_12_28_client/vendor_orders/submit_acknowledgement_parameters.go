// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders

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

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentOrders_2021-12-28/vendor_direct_fulfillment_orders_2021_12_28_models"
)

// NewSubmitAcknowledgementParams creates a new SubmitAcknowledgementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitAcknowledgementParams() *SubmitAcknowledgementParams {
	return &SubmitAcknowledgementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitAcknowledgementParamsWithTimeout creates a new SubmitAcknowledgementParams object
// with the ability to set a timeout on a request.
func NewSubmitAcknowledgementParamsWithTimeout(timeout time.Duration) *SubmitAcknowledgementParams {
	return &SubmitAcknowledgementParams{
		timeout: timeout,
	}
}

// NewSubmitAcknowledgementParamsWithContext creates a new SubmitAcknowledgementParams object
// with the ability to set a context for a request.
func NewSubmitAcknowledgementParamsWithContext(ctx context.Context) *SubmitAcknowledgementParams {
	return &SubmitAcknowledgementParams{
		Context: ctx,
	}
}

// NewSubmitAcknowledgementParamsWithHTTPClient creates a new SubmitAcknowledgementParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitAcknowledgementParamsWithHTTPClient(client *http.Client) *SubmitAcknowledgementParams {
	return &SubmitAcknowledgementParams{
		HTTPClient: client,
	}
}

/*
SubmitAcknowledgementParams contains all the parameters to send to the API endpoint

	for the submit acknowledgement operation.

	Typically these are written to a http.Request.
*/
type SubmitAcknowledgementParams struct {

	/* Body.

	   The request body containing the acknowledgement to an order
	*/
	Body *vendor_direct_fulfillment_orders_2021_12_28_models.SubmitAcknowledgementRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit acknowledgement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitAcknowledgementParams) WithDefaults() *SubmitAcknowledgementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit acknowledgement params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitAcknowledgementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) WithTimeout(timeout time.Duration) *SubmitAcknowledgementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) WithContext(ctx context.Context) *SubmitAcknowledgementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) WithHTTPClient(client *http.Client) *SubmitAcknowledgementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) WithBody(body *vendor_direct_fulfillment_orders_2021_12_28_models.SubmitAcknowledgementRequest) *SubmitAcknowledgementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit acknowledgement params
func (o *SubmitAcknowledgementParams) SetBody(body *vendor_direct_fulfillment_orders_2021_12_28_models.SubmitAcknowledgementRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitAcknowledgementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
