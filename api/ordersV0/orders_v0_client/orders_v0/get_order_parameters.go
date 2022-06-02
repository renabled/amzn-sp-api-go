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
)

// NewGetOrderParams creates a new GetOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrderParams() *GetOrderParams {
	return &GetOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderParamsWithTimeout creates a new GetOrderParams object
// with the ability to set a timeout on a request.
func NewGetOrderParamsWithTimeout(timeout time.Duration) *GetOrderParams {
	return &GetOrderParams{
		timeout: timeout,
	}
}

// NewGetOrderParamsWithContext creates a new GetOrderParams object
// with the ability to set a context for a request.
func NewGetOrderParamsWithContext(ctx context.Context) *GetOrderParams {
	return &GetOrderParams{
		Context: ctx,
	}
}

// NewGetOrderParamsWithHTTPClient creates a new GetOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrderParamsWithHTTPClient(client *http.Client) *GetOrderParams {
	return &GetOrderParams{
		HTTPClient: client,
	}
}

/* GetOrderParams contains all the parameters to send to the API endpoint
   for the get order operation.

   Typically these are written to a http.Request.
*/
type GetOrderParams struct {

	/* OrderID.

	   An Amazon-defined order identifier, in 3-7-7 format.
	*/
	OrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderParams) WithDefaults() *GetOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get order params
func (o *GetOrderParams) WithTimeout(timeout time.Duration) *GetOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order params
func (o *GetOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order params
func (o *GetOrderParams) WithContext(ctx context.Context) *GetOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order params
func (o *GetOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order params
func (o *GetOrderParams) WithHTTPClient(client *http.Client) *GetOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order params
func (o *GetOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the get order params
func (o *GetOrderParams) WithOrderID(orderID string) *GetOrderParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get order params
func (o *GetOrderParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param orderId
	if err := r.SetPathParam("orderId", o.OrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
