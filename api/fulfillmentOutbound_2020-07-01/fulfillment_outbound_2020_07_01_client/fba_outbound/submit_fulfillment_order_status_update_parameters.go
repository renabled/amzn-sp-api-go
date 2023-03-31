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

	"github.com/xamandar/amzn-sp-api-go/api/fulfillmentOutbound_2020-07-01/fulfillment_outbound_2020_07_01_models"
)

// NewSubmitFulfillmentOrderStatusUpdateParams creates a new SubmitFulfillmentOrderStatusUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitFulfillmentOrderStatusUpdateParams() *SubmitFulfillmentOrderStatusUpdateParams {
	return &SubmitFulfillmentOrderStatusUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitFulfillmentOrderStatusUpdateParamsWithTimeout creates a new SubmitFulfillmentOrderStatusUpdateParams object
// with the ability to set a timeout on a request.
func NewSubmitFulfillmentOrderStatusUpdateParamsWithTimeout(timeout time.Duration) *SubmitFulfillmentOrderStatusUpdateParams {
	return &SubmitFulfillmentOrderStatusUpdateParams{
		timeout: timeout,
	}
}

// NewSubmitFulfillmentOrderStatusUpdateParamsWithContext creates a new SubmitFulfillmentOrderStatusUpdateParams object
// with the ability to set a context for a request.
func NewSubmitFulfillmentOrderStatusUpdateParamsWithContext(ctx context.Context) *SubmitFulfillmentOrderStatusUpdateParams {
	return &SubmitFulfillmentOrderStatusUpdateParams{
		Context: ctx,
	}
}

// NewSubmitFulfillmentOrderStatusUpdateParamsWithHTTPClient creates a new SubmitFulfillmentOrderStatusUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitFulfillmentOrderStatusUpdateParamsWithHTTPClient(client *http.Client) *SubmitFulfillmentOrderStatusUpdateParams {
	return &SubmitFulfillmentOrderStatusUpdateParams{
		HTTPClient: client,
	}
}

/*
SubmitFulfillmentOrderStatusUpdateParams contains all the parameters to send to the API endpoint

	for the submit fulfillment order status update operation.

	Typically these are written to a http.Request.
*/
type SubmitFulfillmentOrderStatusUpdateParams struct {

	// Body.
	Body *fulfillment_outbound_2020_07_01_models.SubmitFulfillmentOrderStatusUpdateRequest

	/* SellerFulfillmentOrderID.

	   The identifier assigned to the item by the seller when the fulfillment order was created.
	*/
	SellerFulfillmentOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit fulfillment order status update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitFulfillmentOrderStatusUpdateParams) WithDefaults() *SubmitFulfillmentOrderStatusUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit fulfillment order status update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitFulfillmentOrderStatusUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) WithTimeout(timeout time.Duration) *SubmitFulfillmentOrderStatusUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) WithContext(ctx context.Context) *SubmitFulfillmentOrderStatusUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) WithHTTPClient(client *http.Client) *SubmitFulfillmentOrderStatusUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) WithBody(body *fulfillment_outbound_2020_07_01_models.SubmitFulfillmentOrderStatusUpdateRequest) *SubmitFulfillmentOrderStatusUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) SetBody(body *fulfillment_outbound_2020_07_01_models.SubmitFulfillmentOrderStatusUpdateRequest) {
	o.Body = body
}

// WithSellerFulfillmentOrderID adds the sellerFulfillmentOrderID to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) WithSellerFulfillmentOrderID(sellerFulfillmentOrderID string) *SubmitFulfillmentOrderStatusUpdateParams {
	o.SetSellerFulfillmentOrderID(sellerFulfillmentOrderID)
	return o
}

// SetSellerFulfillmentOrderID adds the sellerFulfillmentOrderId to the submit fulfillment order status update params
func (o *SubmitFulfillmentOrderStatusUpdateParams) SetSellerFulfillmentOrderID(sellerFulfillmentOrderID string) {
	o.SellerFulfillmentOrderID = sellerFulfillmentOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitFulfillmentOrderStatusUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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