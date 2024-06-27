// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping_labels

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

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_models"
)

// NewSubmitShippingLabelRequestParams creates a new SubmitShippingLabelRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitShippingLabelRequestParams() *SubmitShippingLabelRequestParams {
	return &SubmitShippingLabelRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitShippingLabelRequestParamsWithTimeout creates a new SubmitShippingLabelRequestParams object
// with the ability to set a timeout on a request.
func NewSubmitShippingLabelRequestParamsWithTimeout(timeout time.Duration) *SubmitShippingLabelRequestParams {
	return &SubmitShippingLabelRequestParams{
		timeout: timeout,
	}
}

// NewSubmitShippingLabelRequestParamsWithContext creates a new SubmitShippingLabelRequestParams object
// with the ability to set a context for a request.
func NewSubmitShippingLabelRequestParamsWithContext(ctx context.Context) *SubmitShippingLabelRequestParams {
	return &SubmitShippingLabelRequestParams{
		Context: ctx,
	}
}

// NewSubmitShippingLabelRequestParamsWithHTTPClient creates a new SubmitShippingLabelRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitShippingLabelRequestParamsWithHTTPClient(client *http.Client) *SubmitShippingLabelRequestParams {
	return &SubmitShippingLabelRequestParams{
		HTTPClient: client,
	}
}

/*
SubmitShippingLabelRequestParams contains all the parameters to send to the API endpoint

	for the submit shipping label request operation.

	Typically these are written to a http.Request.
*/
type SubmitShippingLabelRequestParams struct {

	/* Body.

	   Request body that contains the shipping labels data.
	*/
	Body *vendor_direct_fulfillment_shipping_2021_12_28_models.SubmitShippingLabelsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit shipping label request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitShippingLabelRequestParams) WithDefaults() *SubmitShippingLabelRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit shipping label request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitShippingLabelRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) WithTimeout(timeout time.Duration) *SubmitShippingLabelRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) WithContext(ctx context.Context) *SubmitShippingLabelRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) WithHTTPClient(client *http.Client) *SubmitShippingLabelRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) WithBody(body *vendor_direct_fulfillment_shipping_2021_12_28_models.SubmitShippingLabelsRequest) *SubmitShippingLabelRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit shipping label request params
func (o *SubmitShippingLabelRequestParams) SetBody(body *vendor_direct_fulfillment_shipping_2021_12_28_models.SubmitShippingLabelsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitShippingLabelRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
