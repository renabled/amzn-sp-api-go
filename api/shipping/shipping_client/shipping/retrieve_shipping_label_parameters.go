// Code generated by go-swagger; DO NOT EDIT.

package shipping

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

	"github.com/renabled/amzn-sp-api-go/api/shipping/shipping_models"
)

// NewRetrieveShippingLabelParams creates a new RetrieveShippingLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRetrieveShippingLabelParams() *RetrieveShippingLabelParams {
	return &RetrieveShippingLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveShippingLabelParamsWithTimeout creates a new RetrieveShippingLabelParams object
// with the ability to set a timeout on a request.
func NewRetrieveShippingLabelParamsWithTimeout(timeout time.Duration) *RetrieveShippingLabelParams {
	return &RetrieveShippingLabelParams{
		timeout: timeout,
	}
}

// NewRetrieveShippingLabelParamsWithContext creates a new RetrieveShippingLabelParams object
// with the ability to set a context for a request.
func NewRetrieveShippingLabelParamsWithContext(ctx context.Context) *RetrieveShippingLabelParams {
	return &RetrieveShippingLabelParams{
		Context: ctx,
	}
}

// NewRetrieveShippingLabelParamsWithHTTPClient creates a new RetrieveShippingLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewRetrieveShippingLabelParamsWithHTTPClient(client *http.Client) *RetrieveShippingLabelParams {
	return &RetrieveShippingLabelParams{
		HTTPClient: client,
	}
}

/*
RetrieveShippingLabelParams contains all the parameters to send to the API endpoint

	for the retrieve shipping label operation.

	Typically these are written to a http.Request.
*/
type RetrieveShippingLabelParams struct {

	// Body.
	Body *shipping_models.RetrieveShippingLabelRequest

	// ShipmentID.
	ShipmentID string

	// TrackingID.
	TrackingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the retrieve shipping label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveShippingLabelParams) WithDefaults() *RetrieveShippingLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the retrieve shipping label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RetrieveShippingLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) WithTimeout(timeout time.Duration) *RetrieveShippingLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) WithContext(ctx context.Context) *RetrieveShippingLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) WithHTTPClient(client *http.Client) *RetrieveShippingLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) WithBody(body *shipping_models.RetrieveShippingLabelRequest) *RetrieveShippingLabelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) SetBody(body *shipping_models.RetrieveShippingLabelRequest) {
	o.Body = body
}

// WithShipmentID adds the shipmentID to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) WithShipmentID(shipmentID string) *RetrieveShippingLabelParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WithTrackingID adds the trackingID to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) WithTrackingID(trackingID string) *RetrieveShippingLabelParams {
	o.SetTrackingID(trackingID)
	return o
}

// SetTrackingID adds the trackingId to the retrieve shipping label params
func (o *RetrieveShippingLabelParams) SetTrackingID(trackingID string) {
	o.TrackingID = trackingID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveShippingLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param shipmentId
	if err := r.SetPathParam("shipmentId", o.ShipmentID); err != nil {
		return err
	}

	// path param trackingId
	if err := r.SetPathParam("trackingId", o.TrackingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
