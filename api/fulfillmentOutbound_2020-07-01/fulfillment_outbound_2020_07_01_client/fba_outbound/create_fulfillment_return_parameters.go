// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

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

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentOutbound_2020-07-01/fulfillment_outbound_2020_07_01_models"
)

// NewCreateFulfillmentReturnParams creates a new CreateFulfillmentReturnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFulfillmentReturnParams() *CreateFulfillmentReturnParams {
	return &CreateFulfillmentReturnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFulfillmentReturnParamsWithTimeout creates a new CreateFulfillmentReturnParams object
// with the ability to set a timeout on a request.
func NewCreateFulfillmentReturnParamsWithTimeout(timeout time.Duration) *CreateFulfillmentReturnParams {
	return &CreateFulfillmentReturnParams{
		timeout: timeout,
	}
}

// NewCreateFulfillmentReturnParamsWithContext creates a new CreateFulfillmentReturnParams object
// with the ability to set a context for a request.
func NewCreateFulfillmentReturnParamsWithContext(ctx context.Context) *CreateFulfillmentReturnParams {
	return &CreateFulfillmentReturnParams{
		Context: ctx,
	}
}

// NewCreateFulfillmentReturnParamsWithHTTPClient creates a new CreateFulfillmentReturnParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFulfillmentReturnParamsWithHTTPClient(client *http.Client) *CreateFulfillmentReturnParams {
	return &CreateFulfillmentReturnParams{
		HTTPClient: client,
	}
}

/*
CreateFulfillmentReturnParams contains all the parameters to send to the API endpoint

	for the create fulfillment return operation.

	Typically these are written to a http.Request.
*/
type CreateFulfillmentReturnParams struct {

	/* Body.

	   CreateFulfillmentReturnRequest parameter
	*/
	Body *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnRequest

	/* SellerFulfillmentOrderID.

	   An identifier assigned by the seller to the fulfillment order at the time it was created. The seller uses their own records to find the correct `SellerFulfillmentOrderId` value based on the buyer's request to return items.
	*/
	SellerFulfillmentOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create fulfillment return params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFulfillmentReturnParams) WithDefaults() *CreateFulfillmentReturnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create fulfillment return params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFulfillmentReturnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) WithTimeout(timeout time.Duration) *CreateFulfillmentReturnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) WithContext(ctx context.Context) *CreateFulfillmentReturnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) WithHTTPClient(client *http.Client) *CreateFulfillmentReturnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) WithBody(body *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnRequest) *CreateFulfillmentReturnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) SetBody(body *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnRequest) {
	o.Body = body
}

// WithSellerFulfillmentOrderID adds the sellerFulfillmentOrderID to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) WithSellerFulfillmentOrderID(sellerFulfillmentOrderID string) *CreateFulfillmentReturnParams {
	o.SetSellerFulfillmentOrderID(sellerFulfillmentOrderID)
	return o
}

// SetSellerFulfillmentOrderID adds the sellerFulfillmentOrderId to the create fulfillment return params
func (o *CreateFulfillmentReturnParams) SetSellerFulfillmentOrderID(sellerFulfillmentOrderID string) {
	o.SellerFulfillmentOrderID = sellerFulfillmentOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFulfillmentReturnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param sellerFulfillmentOrderId
	if err := r.SetPathParam("sellerFulfillmentOrderId", o.SellerFulfillmentOrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
