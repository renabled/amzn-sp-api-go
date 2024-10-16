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

// NewGetDeliveryChallanDocumentParams creates a new GetDeliveryChallanDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeliveryChallanDocumentParams() *GetDeliveryChallanDocumentParams {
	return &GetDeliveryChallanDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeliveryChallanDocumentParamsWithTimeout creates a new GetDeliveryChallanDocumentParams object
// with the ability to set a timeout on a request.
func NewGetDeliveryChallanDocumentParamsWithTimeout(timeout time.Duration) *GetDeliveryChallanDocumentParams {
	return &GetDeliveryChallanDocumentParams{
		timeout: timeout,
	}
}

// NewGetDeliveryChallanDocumentParamsWithContext creates a new GetDeliveryChallanDocumentParams object
// with the ability to set a context for a request.
func NewGetDeliveryChallanDocumentParamsWithContext(ctx context.Context) *GetDeliveryChallanDocumentParams {
	return &GetDeliveryChallanDocumentParams{
		Context: ctx,
	}
}

// NewGetDeliveryChallanDocumentParamsWithHTTPClient creates a new GetDeliveryChallanDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeliveryChallanDocumentParamsWithHTTPClient(client *http.Client) *GetDeliveryChallanDocumentParams {
	return &GetDeliveryChallanDocumentParams{
		HTTPClient: client,
	}
}

/*
GetDeliveryChallanDocumentParams contains all the parameters to send to the API endpoint

	for the get delivery challan document operation.

	Typically these are written to a http.Request.
*/
type GetDeliveryChallanDocumentParams struct {

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

// WithDefaults hydrates default values in the get delivery challan document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeliveryChallanDocumentParams) WithDefaults() *GetDeliveryChallanDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get delivery challan document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeliveryChallanDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) WithTimeout(timeout time.Duration) *GetDeliveryChallanDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) WithContext(ctx context.Context) *GetDeliveryChallanDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) WithHTTPClient(client *http.Client) *GetDeliveryChallanDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboundPlanID adds the inboundPlanID to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) WithInboundPlanID(inboundPlanID string) *GetDeliveryChallanDocumentParams {
	o.SetInboundPlanID(inboundPlanID)
	return o
}

// SetInboundPlanID adds the inboundPlanId to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) SetInboundPlanID(inboundPlanID string) {
	o.InboundPlanID = inboundPlanID
}

// WithShipmentID adds the shipmentID to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) WithShipmentID(shipmentID string) *GetDeliveryChallanDocumentParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the get delivery challan document params
func (o *GetDeliveryChallanDocumentParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeliveryChallanDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
