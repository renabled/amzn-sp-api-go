// Code generated by go-swagger; DO NOT EDIT.

package supply_sources

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

	"github.com/renabled/amzn-sp-api-go/api/supplySources_2020-07-01/supply_sources_2020_07_01_models"
)

// NewUpdateSupplySourceStatusParams creates a new UpdateSupplySourceStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSupplySourceStatusParams() *UpdateSupplySourceStatusParams {
	return &UpdateSupplySourceStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSupplySourceStatusParamsWithTimeout creates a new UpdateSupplySourceStatusParams object
// with the ability to set a timeout on a request.
func NewUpdateSupplySourceStatusParamsWithTimeout(timeout time.Duration) *UpdateSupplySourceStatusParams {
	return &UpdateSupplySourceStatusParams{
		timeout: timeout,
	}
}

// NewUpdateSupplySourceStatusParamsWithContext creates a new UpdateSupplySourceStatusParams object
// with the ability to set a context for a request.
func NewUpdateSupplySourceStatusParamsWithContext(ctx context.Context) *UpdateSupplySourceStatusParams {
	return &UpdateSupplySourceStatusParams{
		Context: ctx,
	}
}

// NewUpdateSupplySourceStatusParamsWithHTTPClient creates a new UpdateSupplySourceStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSupplySourceStatusParamsWithHTTPClient(client *http.Client) *UpdateSupplySourceStatusParams {
	return &UpdateSupplySourceStatusParams{
		HTTPClient: client,
	}
}

/*
UpdateSupplySourceStatusParams contains all the parameters to send to the API endpoint

	for the update supply source status operation.

	Typically these are written to a http.Request.
*/
type UpdateSupplySourceStatusParams struct {

	// Payload.
	Payload *supply_sources_2020_07_01_models.UpdateSupplySourceStatusRequest

	/* SupplySourceID.

	   The unique identifier of a supply source.
	*/
	SupplySourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update supply source status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSupplySourceStatusParams) WithDefaults() *UpdateSupplySourceStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update supply source status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSupplySourceStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update supply source status params
func (o *UpdateSupplySourceStatusParams) WithTimeout(timeout time.Duration) *UpdateSupplySourceStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update supply source status params
func (o *UpdateSupplySourceStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update supply source status params
func (o *UpdateSupplySourceStatusParams) WithContext(ctx context.Context) *UpdateSupplySourceStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update supply source status params
func (o *UpdateSupplySourceStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update supply source status params
func (o *UpdateSupplySourceStatusParams) WithHTTPClient(client *http.Client) *UpdateSupplySourceStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update supply source status params
func (o *UpdateSupplySourceStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the update supply source status params
func (o *UpdateSupplySourceStatusParams) WithPayload(payload *supply_sources_2020_07_01_models.UpdateSupplySourceStatusRequest) *UpdateSupplySourceStatusParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the update supply source status params
func (o *UpdateSupplySourceStatusParams) SetPayload(payload *supply_sources_2020_07_01_models.UpdateSupplySourceStatusRequest) {
	o.Payload = payload
}

// WithSupplySourceID adds the supplySourceID to the update supply source status params
func (o *UpdateSupplySourceStatusParams) WithSupplySourceID(supplySourceID string) *UpdateSupplySourceStatusParams {
	o.SetSupplySourceID(supplySourceID)
	return o
}

// SetSupplySourceID adds the supplySourceId to the update supply source status params
func (o *UpdateSupplySourceStatusParams) SetSupplySourceID(supplySourceID string) {
	o.SupplySourceID = supplySourceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSupplySourceStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param supplySourceId
	if err := r.SetPathParam("supplySourceId", o.SupplySourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
