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
)

// NewCancelShipmentOldParams creates a new CancelShipmentOldParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelShipmentOldParams() *CancelShipmentOldParams {
	return &CancelShipmentOldParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelShipmentOldParamsWithTimeout creates a new CancelShipmentOldParams object
// with the ability to set a timeout on a request.
func NewCancelShipmentOldParamsWithTimeout(timeout time.Duration) *CancelShipmentOldParams {
	return &CancelShipmentOldParams{
		timeout: timeout,
	}
}

// NewCancelShipmentOldParamsWithContext creates a new CancelShipmentOldParams object
// with the ability to set a context for a request.
func NewCancelShipmentOldParamsWithContext(ctx context.Context) *CancelShipmentOldParams {
	return &CancelShipmentOldParams{
		Context: ctx,
	}
}

// NewCancelShipmentOldParamsWithHTTPClient creates a new CancelShipmentOldParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelShipmentOldParamsWithHTTPClient(client *http.Client) *CancelShipmentOldParams {
	return &CancelShipmentOldParams{
		HTTPClient: client,
	}
}

/*
CancelShipmentOldParams contains all the parameters to send to the API endpoint

	for the cancel shipment old operation.

	Typically these are written to a http.Request.
*/
type CancelShipmentOldParams struct {

	/* ShipmentID.

	   The Amazon-defined shipment identifier for the shipment to cancel.
	*/
	ShipmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel shipment old params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelShipmentOldParams) WithDefaults() *CancelShipmentOldParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel shipment old params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelShipmentOldParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel shipment old params
func (o *CancelShipmentOldParams) WithTimeout(timeout time.Duration) *CancelShipmentOldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel shipment old params
func (o *CancelShipmentOldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel shipment old params
func (o *CancelShipmentOldParams) WithContext(ctx context.Context) *CancelShipmentOldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel shipment old params
func (o *CancelShipmentOldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel shipment old params
func (o *CancelShipmentOldParams) WithHTTPClient(client *http.Client) *CancelShipmentOldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel shipment old params
func (o *CancelShipmentOldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithShipmentID adds the shipmentID to the cancel shipment old params
func (o *CancelShipmentOldParams) WithShipmentID(shipmentID string) *CancelShipmentOldParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the cancel shipment old params
func (o *CancelShipmentOldParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *CancelShipmentOldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param shipmentId
	if err := r.SetPathParam("shipmentId", o.ShipmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
