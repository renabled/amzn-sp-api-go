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

// NewUpdateShipmentTrackingDetailsParams creates a new UpdateShipmentTrackingDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateShipmentTrackingDetailsParams() *UpdateShipmentTrackingDetailsParams {
	return &UpdateShipmentTrackingDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateShipmentTrackingDetailsParamsWithTimeout creates a new UpdateShipmentTrackingDetailsParams object
// with the ability to set a timeout on a request.
func NewUpdateShipmentTrackingDetailsParamsWithTimeout(timeout time.Duration) *UpdateShipmentTrackingDetailsParams {
	return &UpdateShipmentTrackingDetailsParams{
		timeout: timeout,
	}
}

// NewUpdateShipmentTrackingDetailsParamsWithContext creates a new UpdateShipmentTrackingDetailsParams object
// with the ability to set a context for a request.
func NewUpdateShipmentTrackingDetailsParamsWithContext(ctx context.Context) *UpdateShipmentTrackingDetailsParams {
	return &UpdateShipmentTrackingDetailsParams{
		Context: ctx,
	}
}

// NewUpdateShipmentTrackingDetailsParamsWithHTTPClient creates a new UpdateShipmentTrackingDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateShipmentTrackingDetailsParamsWithHTTPClient(client *http.Client) *UpdateShipmentTrackingDetailsParams {
	return &UpdateShipmentTrackingDetailsParams{
		HTTPClient: client,
	}
}

/*
UpdateShipmentTrackingDetailsParams contains all the parameters to send to the API endpoint

	for the update shipment tracking details operation.

	Typically these are written to a http.Request.
*/
type UpdateShipmentTrackingDetailsParams struct {

	/* Body.

	   The body of the request to `updateShipmentTrackingDetails`.
	*/
	Body *fulfillment_inbound_2024_03_20_models.UpdateShipmentTrackingDetailsRequest

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

// WithDefaults hydrates default values in the update shipment tracking details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateShipmentTrackingDetailsParams) WithDefaults() *UpdateShipmentTrackingDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update shipment tracking details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateShipmentTrackingDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) WithTimeout(timeout time.Duration) *UpdateShipmentTrackingDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) WithContext(ctx context.Context) *UpdateShipmentTrackingDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) WithHTTPClient(client *http.Client) *UpdateShipmentTrackingDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) WithBody(body *fulfillment_inbound_2024_03_20_models.UpdateShipmentTrackingDetailsRequest) *UpdateShipmentTrackingDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) SetBody(body *fulfillment_inbound_2024_03_20_models.UpdateShipmentTrackingDetailsRequest) {
	o.Body = body
}

// WithInboundPlanID adds the inboundPlanID to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) WithInboundPlanID(inboundPlanID string) *UpdateShipmentTrackingDetailsParams {
	o.SetInboundPlanID(inboundPlanID)
	return o
}

// SetInboundPlanID adds the inboundPlanId to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) SetInboundPlanID(inboundPlanID string) {
	o.InboundPlanID = inboundPlanID
}

// WithShipmentID adds the shipmentID to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) WithShipmentID(shipmentID string) *UpdateShipmentTrackingDetailsParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the update shipment tracking details params
func (o *UpdateShipmentTrackingDetailsParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateShipmentTrackingDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
