// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping_labels

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

// NewGetShippingLabelParams creates a new GetShippingLabelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetShippingLabelParams() *GetShippingLabelParams {
	return &GetShippingLabelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetShippingLabelParamsWithTimeout creates a new GetShippingLabelParams object
// with the ability to set a timeout on a request.
func NewGetShippingLabelParamsWithTimeout(timeout time.Duration) *GetShippingLabelParams {
	return &GetShippingLabelParams{
		timeout: timeout,
	}
}

// NewGetShippingLabelParamsWithContext creates a new GetShippingLabelParams object
// with the ability to set a context for a request.
func NewGetShippingLabelParamsWithContext(ctx context.Context) *GetShippingLabelParams {
	return &GetShippingLabelParams{
		Context: ctx,
	}
}

// NewGetShippingLabelParamsWithHTTPClient creates a new GetShippingLabelParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetShippingLabelParamsWithHTTPClient(client *http.Client) *GetShippingLabelParams {
	return &GetShippingLabelParams{
		HTTPClient: client,
	}
}

/* GetShippingLabelParams contains all the parameters to send to the API endpoint
   for the get shipping label operation.

   Typically these are written to a http.Request.
*/
type GetShippingLabelParams struct {

	/* PurchaseOrderNumber.

	   The purchase order number for which you want to return the shipping label. It should be the same purchaseOrderNumber as received in the order.
	*/
	PurchaseOrderNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get shipping label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShippingLabelParams) WithDefaults() *GetShippingLabelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get shipping label params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShippingLabelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get shipping label params
func (o *GetShippingLabelParams) WithTimeout(timeout time.Duration) *GetShippingLabelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get shipping label params
func (o *GetShippingLabelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get shipping label params
func (o *GetShippingLabelParams) WithContext(ctx context.Context) *GetShippingLabelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get shipping label params
func (o *GetShippingLabelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get shipping label params
func (o *GetShippingLabelParams) WithHTTPClient(client *http.Client) *GetShippingLabelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get shipping label params
func (o *GetShippingLabelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPurchaseOrderNumber adds the purchaseOrderNumber to the get shipping label params
func (o *GetShippingLabelParams) WithPurchaseOrderNumber(purchaseOrderNumber string) *GetShippingLabelParams {
	o.SetPurchaseOrderNumber(purchaseOrderNumber)
	return o
}

// SetPurchaseOrderNumber adds the purchaseOrderNumber to the get shipping label params
func (o *GetShippingLabelParams) SetPurchaseOrderNumber(purchaseOrderNumber string) {
	o.PurchaseOrderNumber = purchaseOrderNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetShippingLabelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param purchaseOrderNumber
	if err := r.SetPathParam("purchaseOrderNumber", o.PurchaseOrderNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
