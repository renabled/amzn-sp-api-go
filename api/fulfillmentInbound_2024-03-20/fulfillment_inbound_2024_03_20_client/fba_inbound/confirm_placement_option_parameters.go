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

// NewConfirmPlacementOptionParams creates a new ConfirmPlacementOptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfirmPlacementOptionParams() *ConfirmPlacementOptionParams {
	return &ConfirmPlacementOptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfirmPlacementOptionParamsWithTimeout creates a new ConfirmPlacementOptionParams object
// with the ability to set a timeout on a request.
func NewConfirmPlacementOptionParamsWithTimeout(timeout time.Duration) *ConfirmPlacementOptionParams {
	return &ConfirmPlacementOptionParams{
		timeout: timeout,
	}
}

// NewConfirmPlacementOptionParamsWithContext creates a new ConfirmPlacementOptionParams object
// with the ability to set a context for a request.
func NewConfirmPlacementOptionParamsWithContext(ctx context.Context) *ConfirmPlacementOptionParams {
	return &ConfirmPlacementOptionParams{
		Context: ctx,
	}
}

// NewConfirmPlacementOptionParamsWithHTTPClient creates a new ConfirmPlacementOptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfirmPlacementOptionParamsWithHTTPClient(client *http.Client) *ConfirmPlacementOptionParams {
	return &ConfirmPlacementOptionParams{
		HTTPClient: client,
	}
}

/*
ConfirmPlacementOptionParams contains all the parameters to send to the API endpoint

	for the confirm placement option operation.

	Typically these are written to a http.Request.
*/
type ConfirmPlacementOptionParams struct {

	/* InboundPlanID.

	   Identifier of an inbound plan.
	*/
	InboundPlanID string

	/* PlacementOptionID.

	   The identifier of a placement option. A placement option represents the shipment splits and destinations of SKUs.
	*/
	PlacementOptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the confirm placement option params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmPlacementOptionParams) WithDefaults() *ConfirmPlacementOptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the confirm placement option params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmPlacementOptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the confirm placement option params
func (o *ConfirmPlacementOptionParams) WithTimeout(timeout time.Duration) *ConfirmPlacementOptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the confirm placement option params
func (o *ConfirmPlacementOptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the confirm placement option params
func (o *ConfirmPlacementOptionParams) WithContext(ctx context.Context) *ConfirmPlacementOptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the confirm placement option params
func (o *ConfirmPlacementOptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the confirm placement option params
func (o *ConfirmPlacementOptionParams) WithHTTPClient(client *http.Client) *ConfirmPlacementOptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the confirm placement option params
func (o *ConfirmPlacementOptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboundPlanID adds the inboundPlanID to the confirm placement option params
func (o *ConfirmPlacementOptionParams) WithInboundPlanID(inboundPlanID string) *ConfirmPlacementOptionParams {
	o.SetInboundPlanID(inboundPlanID)
	return o
}

// SetInboundPlanID adds the inboundPlanId to the confirm placement option params
func (o *ConfirmPlacementOptionParams) SetInboundPlanID(inboundPlanID string) {
	o.InboundPlanID = inboundPlanID
}

// WithPlacementOptionID adds the placementOptionID to the confirm placement option params
func (o *ConfirmPlacementOptionParams) WithPlacementOptionID(placementOptionID string) *ConfirmPlacementOptionParams {
	o.SetPlacementOptionID(placementOptionID)
	return o
}

// SetPlacementOptionID adds the placementOptionId to the confirm placement option params
func (o *ConfirmPlacementOptionParams) SetPlacementOptionID(placementOptionID string) {
	o.PlacementOptionID = placementOptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ConfirmPlacementOptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inboundPlanId
	if err := r.SetPathParam("inboundPlanId", o.InboundPlanID); err != nil {
		return err
	}

	// path param placementOptionId
	if err := r.SetPathParam("placementOptionId", o.PlacementOptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
