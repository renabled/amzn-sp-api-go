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

// NewGetOrderItemsBuyerInfoParams creates a new GetOrderItemsBuyerInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrderItemsBuyerInfoParams() *GetOrderItemsBuyerInfoParams {
	return &GetOrderItemsBuyerInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderItemsBuyerInfoParamsWithTimeout creates a new GetOrderItemsBuyerInfoParams object
// with the ability to set a timeout on a request.
func NewGetOrderItemsBuyerInfoParamsWithTimeout(timeout time.Duration) *GetOrderItemsBuyerInfoParams {
	return &GetOrderItemsBuyerInfoParams{
		timeout: timeout,
	}
}

// NewGetOrderItemsBuyerInfoParamsWithContext creates a new GetOrderItemsBuyerInfoParams object
// with the ability to set a context for a request.
func NewGetOrderItemsBuyerInfoParamsWithContext(ctx context.Context) *GetOrderItemsBuyerInfoParams {
	return &GetOrderItemsBuyerInfoParams{
		Context: ctx,
	}
}

// NewGetOrderItemsBuyerInfoParamsWithHTTPClient creates a new GetOrderItemsBuyerInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrderItemsBuyerInfoParamsWithHTTPClient(client *http.Client) *GetOrderItemsBuyerInfoParams {
	return &GetOrderItemsBuyerInfoParams{
		HTTPClient: client,
	}
}

/* GetOrderItemsBuyerInfoParams contains all the parameters to send to the API endpoint
   for the get order items buyer info operation.

   Typically these are written to a http.Request.
*/
type GetOrderItemsBuyerInfoParams struct {

	/* NextToken.

	   A string token returned in the response of your previous request.
	*/
	NextToken *string

	/* OrderID.

	   An Amazon-defined order identifier, in 3-7-7 format.
	*/
	OrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get order items buyer info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderItemsBuyerInfoParams) WithDefaults() *GetOrderItemsBuyerInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get order items buyer info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderItemsBuyerInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) WithTimeout(timeout time.Duration) *GetOrderItemsBuyerInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) WithContext(ctx context.Context) *GetOrderItemsBuyerInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) WithHTTPClient(client *http.Client) *GetOrderItemsBuyerInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNextToken adds the nextToken to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) WithNextToken(nextToken *string) *GetOrderItemsBuyerInfoParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithOrderID adds the orderID to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) WithOrderID(orderID string) *GetOrderItemsBuyerInfoParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get order items buyer info params
func (o *GetOrderItemsBuyerInfoParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderItemsBuyerInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NextToken != nil {

		// query param NextToken
		var qrNextToken string

		if o.NextToken != nil {
			qrNextToken = *o.NextToken
		}
		qNextToken := qrNextToken
		if qNextToken != "" {

			if err := r.SetQueryParam("NextToken", qNextToken); err != nil {
				return err
			}
		}
	}

	// path param orderId
	if err := r.SetPathParam("orderId", o.OrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
