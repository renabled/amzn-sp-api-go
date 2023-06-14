// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping

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

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShippingV1/vendor_direct_fulfillment_shipping_v1_models"
)

// NewSubmitShipmentConfirmationsParams creates a new SubmitShipmentConfirmationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitShipmentConfirmationsParams() *SubmitShipmentConfirmationsParams {
	return &SubmitShipmentConfirmationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitShipmentConfirmationsParamsWithTimeout creates a new SubmitShipmentConfirmationsParams object
// with the ability to set a timeout on a request.
func NewSubmitShipmentConfirmationsParamsWithTimeout(timeout time.Duration) *SubmitShipmentConfirmationsParams {
	return &SubmitShipmentConfirmationsParams{
		timeout: timeout,
	}
}

// NewSubmitShipmentConfirmationsParamsWithContext creates a new SubmitShipmentConfirmationsParams object
// with the ability to set a context for a request.
func NewSubmitShipmentConfirmationsParamsWithContext(ctx context.Context) *SubmitShipmentConfirmationsParams {
	return &SubmitShipmentConfirmationsParams{
		Context: ctx,
	}
}

// NewSubmitShipmentConfirmationsParamsWithHTTPClient creates a new SubmitShipmentConfirmationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitShipmentConfirmationsParamsWithHTTPClient(client *http.Client) *SubmitShipmentConfirmationsParams {
	return &SubmitShipmentConfirmationsParams{
		HTTPClient: client,
	}
}

/*
SubmitShipmentConfirmationsParams contains all the parameters to send to the API endpoint

	for the submit shipment confirmations operation.

	Typically these are written to a http.Request.
*/
type SubmitShipmentConfirmationsParams struct {

	// Body.
	Body *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit shipment confirmations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitShipmentConfirmationsParams) WithDefaults() *SubmitShipmentConfirmationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit shipment confirmations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitShipmentConfirmationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) WithTimeout(timeout time.Duration) *SubmitShipmentConfirmationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) WithContext(ctx context.Context) *SubmitShipmentConfirmationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) WithHTTPClient(client *http.Client) *SubmitShipmentConfirmationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) WithBody(body *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsRequest) *SubmitShipmentConfirmationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit shipment confirmations params
func (o *SubmitShipmentConfirmationsParams) SetBody(body *vendor_direct_fulfillment_shipping_v1_models.SubmitShipmentConfirmationsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitShipmentConfirmationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
