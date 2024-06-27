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

	"github.com/renabled/amzn-sp-api-go/api/vendorDirectFulfillmentShipping_2021-12-28/vendor_direct_fulfillment_shipping_2021_12_28_models"
)

// NewSubmitShipmentStatusUpdatesParams creates a new SubmitShipmentStatusUpdatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubmitShipmentStatusUpdatesParams() *SubmitShipmentStatusUpdatesParams {
	return &SubmitShipmentStatusUpdatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitShipmentStatusUpdatesParamsWithTimeout creates a new SubmitShipmentStatusUpdatesParams object
// with the ability to set a timeout on a request.
func NewSubmitShipmentStatusUpdatesParamsWithTimeout(timeout time.Duration) *SubmitShipmentStatusUpdatesParams {
	return &SubmitShipmentStatusUpdatesParams{
		timeout: timeout,
	}
}

// NewSubmitShipmentStatusUpdatesParamsWithContext creates a new SubmitShipmentStatusUpdatesParams object
// with the ability to set a context for a request.
func NewSubmitShipmentStatusUpdatesParamsWithContext(ctx context.Context) *SubmitShipmentStatusUpdatesParams {
	return &SubmitShipmentStatusUpdatesParams{
		Context: ctx,
	}
}

// NewSubmitShipmentStatusUpdatesParamsWithHTTPClient creates a new SubmitShipmentStatusUpdatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubmitShipmentStatusUpdatesParamsWithHTTPClient(client *http.Client) *SubmitShipmentStatusUpdatesParams {
	return &SubmitShipmentStatusUpdatesParams{
		HTTPClient: client,
	}
}

/*
SubmitShipmentStatusUpdatesParams contains all the parameters to send to the API endpoint

	for the submit shipment status updates operation.

	Typically these are written to a http.Request.
*/
type SubmitShipmentStatusUpdatesParams struct {

	/* Body.

	   Request body that contains the shipment status update data.
	*/
	Body *vendor_direct_fulfillment_shipping_2021_12_28_models.SubmitShipmentStatusUpdatesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the submit shipment status updates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitShipmentStatusUpdatesParams) WithDefaults() *SubmitShipmentStatusUpdatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the submit shipment status updates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubmitShipmentStatusUpdatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) WithTimeout(timeout time.Duration) *SubmitShipmentStatusUpdatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) WithContext(ctx context.Context) *SubmitShipmentStatusUpdatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) WithHTTPClient(client *http.Client) *SubmitShipmentStatusUpdatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) WithBody(body *vendor_direct_fulfillment_shipping_2021_12_28_models.SubmitShipmentStatusUpdatesRequest) *SubmitShipmentStatusUpdatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the submit shipment status updates params
func (o *SubmitShipmentStatusUpdatesParams) SetBody(body *vendor_direct_fulfillment_shipping_2021_12_28_models.SubmitShipmentStatusUpdatesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitShipmentStatusUpdatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
