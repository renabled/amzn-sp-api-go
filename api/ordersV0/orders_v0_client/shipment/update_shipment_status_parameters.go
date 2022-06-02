// Code generated by go-swagger; DO NOT EDIT.

package shipment

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

	"github.com/xamandar/amzn-sp-api-go/api/ordersV0/orders_v0_models"
)

// NewUpdateShipmentStatusParams creates a new UpdateShipmentStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateShipmentStatusParams() *UpdateShipmentStatusParams {
	return &UpdateShipmentStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateShipmentStatusParamsWithTimeout creates a new UpdateShipmentStatusParams object
// with the ability to set a timeout on a request.
func NewUpdateShipmentStatusParamsWithTimeout(timeout time.Duration) *UpdateShipmentStatusParams {
	return &UpdateShipmentStatusParams{
		timeout: timeout,
	}
}

// NewUpdateShipmentStatusParamsWithContext creates a new UpdateShipmentStatusParams object
// with the ability to set a context for a request.
func NewUpdateShipmentStatusParamsWithContext(ctx context.Context) *UpdateShipmentStatusParams {
	return &UpdateShipmentStatusParams{
		Context: ctx,
	}
}

// NewUpdateShipmentStatusParamsWithHTTPClient creates a new UpdateShipmentStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateShipmentStatusParamsWithHTTPClient(client *http.Client) *UpdateShipmentStatusParams {
	return &UpdateShipmentStatusParams{
		HTTPClient: client,
	}
}

/* UpdateShipmentStatusParams contains all the parameters to send to the API endpoint
   for the update shipment status operation.

   Typically these are written to a http.Request.
*/
type UpdateShipmentStatusParams struct {

	/* OrderID.

	   An Amazon-defined order identifier, in 3-7-7 format.
	*/
	OrderID string

	/* Payload.

	   Request to update the shipment status.
	*/
	Payload *orders_v0_models.UpdateShipmentStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update shipment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateShipmentStatusParams) WithDefaults() *UpdateShipmentStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update shipment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateShipmentStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update shipment status params
func (o *UpdateShipmentStatusParams) WithTimeout(timeout time.Duration) *UpdateShipmentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update shipment status params
func (o *UpdateShipmentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update shipment status params
func (o *UpdateShipmentStatusParams) WithContext(ctx context.Context) *UpdateShipmentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update shipment status params
func (o *UpdateShipmentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update shipment status params
func (o *UpdateShipmentStatusParams) WithHTTPClient(client *http.Client) *UpdateShipmentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update shipment status params
func (o *UpdateShipmentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the update shipment status params
func (o *UpdateShipmentStatusParams) WithOrderID(orderID string) *UpdateShipmentStatusParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the update shipment status params
func (o *UpdateShipmentStatusParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WithPayload adds the payload to the update shipment status params
func (o *UpdateShipmentStatusParams) WithPayload(payload *orders_v0_models.UpdateShipmentStatusRequest) *UpdateShipmentStatusParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the update shipment status params
func (o *UpdateShipmentStatusParams) SetPayload(payload *orders_v0_models.UpdateShipmentStatusRequest) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateShipmentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderId
	if err := r.SetPathParam("orderId", o.OrderID); err != nil {
		return err
	}
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
