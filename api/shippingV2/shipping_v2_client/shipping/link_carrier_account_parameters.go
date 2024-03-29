// Code generated by go-swagger; DO NOT EDIT.

package shipping

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

	"github.com/renabled/amzn-sp-api-go/api/shippingV2/shipping_v2_models"
)

// NewLinkCarrierAccountParams creates a new LinkCarrierAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLinkCarrierAccountParams() *LinkCarrierAccountParams {
	return &LinkCarrierAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLinkCarrierAccountParamsWithTimeout creates a new LinkCarrierAccountParams object
// with the ability to set a timeout on a request.
func NewLinkCarrierAccountParamsWithTimeout(timeout time.Duration) *LinkCarrierAccountParams {
	return &LinkCarrierAccountParams{
		timeout: timeout,
	}
}

// NewLinkCarrierAccountParamsWithContext creates a new LinkCarrierAccountParams object
// with the ability to set a context for a request.
func NewLinkCarrierAccountParamsWithContext(ctx context.Context) *LinkCarrierAccountParams {
	return &LinkCarrierAccountParams{
		Context: ctx,
	}
}

// NewLinkCarrierAccountParamsWithHTTPClient creates a new LinkCarrierAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewLinkCarrierAccountParamsWithHTTPClient(client *http.Client) *LinkCarrierAccountParams {
	return &LinkCarrierAccountParams{
		HTTPClient: client,
	}
}

/*
LinkCarrierAccountParams contains all the parameters to send to the API endpoint

	for the link carrier account operation.

	Typically these are written to a http.Request.
*/
type LinkCarrierAccountParams struct {

	// Body.
	Body *shipping_v2_models.LinkCarrierAccountRequest

	/* CarrierID.

	   The unique identifier associated with the carrier account.
	*/
	CarrierID string

	/* XAmznShippingBusinessID.

	   Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
	*/
	XAmznShippingBusinessID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the link carrier account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LinkCarrierAccountParams) WithDefaults() *LinkCarrierAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the link carrier account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LinkCarrierAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the link carrier account params
func (o *LinkCarrierAccountParams) WithTimeout(timeout time.Duration) *LinkCarrierAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the link carrier account params
func (o *LinkCarrierAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the link carrier account params
func (o *LinkCarrierAccountParams) WithContext(ctx context.Context) *LinkCarrierAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the link carrier account params
func (o *LinkCarrierAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the link carrier account params
func (o *LinkCarrierAccountParams) WithHTTPClient(client *http.Client) *LinkCarrierAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the link carrier account params
func (o *LinkCarrierAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the link carrier account params
func (o *LinkCarrierAccountParams) WithBody(body *shipping_v2_models.LinkCarrierAccountRequest) *LinkCarrierAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the link carrier account params
func (o *LinkCarrierAccountParams) SetBody(body *shipping_v2_models.LinkCarrierAccountRequest) {
	o.Body = body
}

// WithCarrierID adds the carrierID to the link carrier account params
func (o *LinkCarrierAccountParams) WithCarrierID(carrierID string) *LinkCarrierAccountParams {
	o.SetCarrierID(carrierID)
	return o
}

// SetCarrierID adds the carrierId to the link carrier account params
func (o *LinkCarrierAccountParams) SetCarrierID(carrierID string) {
	o.CarrierID = carrierID
}

// WithXAmznShippingBusinessID adds the xAmznShippingBusinessID to the link carrier account params
func (o *LinkCarrierAccountParams) WithXAmznShippingBusinessID(xAmznShippingBusinessID *string) *LinkCarrierAccountParams {
	o.SetXAmznShippingBusinessID(xAmznShippingBusinessID)
	return o
}

// SetXAmznShippingBusinessID adds the xAmznShippingBusinessId to the link carrier account params
func (o *LinkCarrierAccountParams) SetXAmznShippingBusinessID(xAmznShippingBusinessID *string) {
	o.XAmznShippingBusinessID = xAmznShippingBusinessID
}

// WriteToRequest writes these params to a swagger request
func (o *LinkCarrierAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param carrierId
	if err := r.SetPathParam("carrierId", o.CarrierID); err != nil {
		return err
	}

	if o.XAmznShippingBusinessID != nil {

		// header param x-amzn-shipping-business-id
		if err := r.SetHeaderParam("x-amzn-shipping-business-id", *o.XAmznShippingBusinessID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
