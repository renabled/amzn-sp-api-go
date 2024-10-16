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
)

// NewConfirmDeliveryWindowOptionsParams creates a new ConfirmDeliveryWindowOptionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfirmDeliveryWindowOptionsParams() *ConfirmDeliveryWindowOptionsParams {
	return &ConfirmDeliveryWindowOptionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfirmDeliveryWindowOptionsParamsWithTimeout creates a new ConfirmDeliveryWindowOptionsParams object
// with the ability to set a timeout on a request.
func NewConfirmDeliveryWindowOptionsParamsWithTimeout(timeout time.Duration) *ConfirmDeliveryWindowOptionsParams {
	return &ConfirmDeliveryWindowOptionsParams{
		timeout: timeout,
	}
}

// NewConfirmDeliveryWindowOptionsParamsWithContext creates a new ConfirmDeliveryWindowOptionsParams object
// with the ability to set a context for a request.
func NewConfirmDeliveryWindowOptionsParamsWithContext(ctx context.Context) *ConfirmDeliveryWindowOptionsParams {
	return &ConfirmDeliveryWindowOptionsParams{
		Context: ctx,
	}
}

// NewConfirmDeliveryWindowOptionsParamsWithHTTPClient creates a new ConfirmDeliveryWindowOptionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfirmDeliveryWindowOptionsParamsWithHTTPClient(client *http.Client) *ConfirmDeliveryWindowOptionsParams {
	return &ConfirmDeliveryWindowOptionsParams{
		HTTPClient: client,
	}
}

/*
ConfirmDeliveryWindowOptionsParams contains all the parameters to send to the API endpoint

	for the confirm delivery window options operation.

	Typically these are written to a http.Request.
*/
type ConfirmDeliveryWindowOptionsParams struct {

	/* DeliveryWindowOptionID.

	   The id of the delivery window option to be confirmed.
	*/
	DeliveryWindowOptionID string

	/* InboundPlanID.

	   Identifier of an inbound plan.
	*/
	InboundPlanID string

	/* ShipmentID.

	   The shipment to confirm the delivery window option for.
	*/
	ShipmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the confirm delivery window options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmDeliveryWindowOptionsParams) WithDefaults() *ConfirmDeliveryWindowOptionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the confirm delivery window options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmDeliveryWindowOptionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) WithTimeout(timeout time.Duration) *ConfirmDeliveryWindowOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) WithContext(ctx context.Context) *ConfirmDeliveryWindowOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) WithHTTPClient(client *http.Client) *ConfirmDeliveryWindowOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeliveryWindowOptionID adds the deliveryWindowOptionID to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) WithDeliveryWindowOptionID(deliveryWindowOptionID string) *ConfirmDeliveryWindowOptionsParams {
	o.SetDeliveryWindowOptionID(deliveryWindowOptionID)
	return o
}

// SetDeliveryWindowOptionID adds the deliveryWindowOptionId to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) SetDeliveryWindowOptionID(deliveryWindowOptionID string) {
	o.DeliveryWindowOptionID = deliveryWindowOptionID
}

// WithInboundPlanID adds the inboundPlanID to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) WithInboundPlanID(inboundPlanID string) *ConfirmDeliveryWindowOptionsParams {
	o.SetInboundPlanID(inboundPlanID)
	return o
}

// SetInboundPlanID adds the inboundPlanId to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) SetInboundPlanID(inboundPlanID string) {
	o.InboundPlanID = inboundPlanID
}

// WithShipmentID adds the shipmentID to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) WithShipmentID(shipmentID string) *ConfirmDeliveryWindowOptionsParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the confirm delivery window options params
func (o *ConfirmDeliveryWindowOptionsParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *ConfirmDeliveryWindowOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deliveryWindowOptionId
	if err := r.SetPathParam("deliveryWindowOptionId", o.DeliveryWindowOptionID); err != nil {
		return err
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
