// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

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

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInbound_2024-03-20/fulfillment_inbound_2024_03_20_models"
)

// NewUpdateShipmentSourceAddressParams creates a new UpdateShipmentSourceAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateShipmentSourceAddressParams() *UpdateShipmentSourceAddressParams {
	return &UpdateShipmentSourceAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateShipmentSourceAddressParamsWithTimeout creates a new UpdateShipmentSourceAddressParams object
// with the ability to set a timeout on a request.
func NewUpdateShipmentSourceAddressParamsWithTimeout(timeout time.Duration) *UpdateShipmentSourceAddressParams {
	return &UpdateShipmentSourceAddressParams{
		timeout: timeout,
	}
}

// NewUpdateShipmentSourceAddressParamsWithContext creates a new UpdateShipmentSourceAddressParams object
// with the ability to set a context for a request.
func NewUpdateShipmentSourceAddressParamsWithContext(ctx context.Context) *UpdateShipmentSourceAddressParams {
	return &UpdateShipmentSourceAddressParams{
		Context: ctx,
	}
}

// NewUpdateShipmentSourceAddressParamsWithHTTPClient creates a new UpdateShipmentSourceAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateShipmentSourceAddressParamsWithHTTPClient(client *http.Client) *UpdateShipmentSourceAddressParams {
	return &UpdateShipmentSourceAddressParams{
		HTTPClient: client,
	}
}

/*
UpdateShipmentSourceAddressParams contains all the parameters to send to the API endpoint

	for the update shipment source address operation.

	Typically these are written to a http.Request.
*/
type UpdateShipmentSourceAddressParams struct {

	/* Body.

	   The body of the request to `updateShipmentSourceAddress`.
	*/
	Body *fulfillment_inbound_2024_03_20_models.UpdateShipmentSourceAddressRequest

	/* InboundPlanID.

	   Identifier of an inbound plan.
	*/
	InboundPlanID string

	/* ShipmentID.

	   Identifier of a shipment. A shipment contains the boxes and units being inbounded.
	*/
	ShipmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update shipment source address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateShipmentSourceAddressParams) WithDefaults() *UpdateShipmentSourceAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update shipment source address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateShipmentSourceAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) WithTimeout(timeout time.Duration) *UpdateShipmentSourceAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) WithContext(ctx context.Context) *UpdateShipmentSourceAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) WithHTTPClient(client *http.Client) *UpdateShipmentSourceAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) WithBody(body *fulfillment_inbound_2024_03_20_models.UpdateShipmentSourceAddressRequest) *UpdateShipmentSourceAddressParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) SetBody(body *fulfillment_inbound_2024_03_20_models.UpdateShipmentSourceAddressRequest) {
	o.Body = body
}

// WithInboundPlanID adds the inboundPlanID to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) WithInboundPlanID(inboundPlanID string) *UpdateShipmentSourceAddressParams {
	o.SetInboundPlanID(inboundPlanID)
	return o
}

// SetInboundPlanID adds the inboundPlanId to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) SetInboundPlanID(inboundPlanID string) {
	o.InboundPlanID = inboundPlanID
}

// WithShipmentID adds the shipmentID to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) WithShipmentID(shipmentID string) *UpdateShipmentSourceAddressParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the update shipment source address params
func (o *UpdateShipmentSourceAddressParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateShipmentSourceAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param inboundPlanId
	if err := r.SetPathParam("inboundPlanId", o.InboundPlanID); err != nil {
		return err
	}

	// path param shipmentId
	if err := r.SetPathParam("shipmentId", o.ShipmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
