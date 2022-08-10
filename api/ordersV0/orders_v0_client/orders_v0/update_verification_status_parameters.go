// Code generated by go-swagger; DO NOT EDIT.

package orders_v0

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

// NewUpdateVerificationStatusParams creates a new UpdateVerificationStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVerificationStatusParams() *UpdateVerificationStatusParams {
	return &UpdateVerificationStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVerificationStatusParamsWithTimeout creates a new UpdateVerificationStatusParams object
// with the ability to set a timeout on a request.
func NewUpdateVerificationStatusParamsWithTimeout(timeout time.Duration) *UpdateVerificationStatusParams {
	return &UpdateVerificationStatusParams{
		timeout: timeout,
	}
}

// NewUpdateVerificationStatusParamsWithContext creates a new UpdateVerificationStatusParams object
// with the ability to set a context for a request.
func NewUpdateVerificationStatusParamsWithContext(ctx context.Context) *UpdateVerificationStatusParams {
	return &UpdateVerificationStatusParams{
		Context: ctx,
	}
}

// NewUpdateVerificationStatusParamsWithHTTPClient creates a new UpdateVerificationStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVerificationStatusParamsWithHTTPClient(client *http.Client) *UpdateVerificationStatusParams {
	return &UpdateVerificationStatusParams{
		HTTPClient: client,
	}
}

/* UpdateVerificationStatusParams contains all the parameters to send to the API endpoint
   for the update verification status operation.

   Typically these are written to a http.Request.
*/
type UpdateVerificationStatusParams struct {

	/* OrderID.

	   An orderId is an Amazon-defined order identifier, in 3-7-7 format.
	*/
	OrderID string

	/* Payload.

	   The request body for the updateVerificationStatus operation.
	*/
	Payload *orders_v0_models.UpdateVerificationStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update verification status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVerificationStatusParams) WithDefaults() *UpdateVerificationStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update verification status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVerificationStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update verification status params
func (o *UpdateVerificationStatusParams) WithTimeout(timeout time.Duration) *UpdateVerificationStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update verification status params
func (o *UpdateVerificationStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update verification status params
func (o *UpdateVerificationStatusParams) WithContext(ctx context.Context) *UpdateVerificationStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update verification status params
func (o *UpdateVerificationStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update verification status params
func (o *UpdateVerificationStatusParams) WithHTTPClient(client *http.Client) *UpdateVerificationStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update verification status params
func (o *UpdateVerificationStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the update verification status params
func (o *UpdateVerificationStatusParams) WithOrderID(orderID string) *UpdateVerificationStatusParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the update verification status params
func (o *UpdateVerificationStatusParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WithPayload adds the payload to the update verification status params
func (o *UpdateVerificationStatusParams) WithPayload(payload *orders_v0_models.UpdateVerificationStatusRequest) *UpdateVerificationStatusParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the update verification status params
func (o *UpdateVerificationStatusParams) SetPayload(payload *orders_v0_models.UpdateVerificationStatusRequest) {
	o.Payload = payload
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVerificationStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
