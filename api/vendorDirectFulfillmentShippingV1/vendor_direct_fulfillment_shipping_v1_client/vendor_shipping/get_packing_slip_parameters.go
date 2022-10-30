// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping

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

// NewGetPackingSlipParams creates a new GetPackingSlipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPackingSlipParams() *GetPackingSlipParams {
	return &GetPackingSlipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackingSlipParamsWithTimeout creates a new GetPackingSlipParams object
// with the ability to set a timeout on a request.
func NewGetPackingSlipParamsWithTimeout(timeout time.Duration) *GetPackingSlipParams {
	return &GetPackingSlipParams{
		timeout: timeout,
	}
}

// NewGetPackingSlipParamsWithContext creates a new GetPackingSlipParams object
// with the ability to set a context for a request.
func NewGetPackingSlipParamsWithContext(ctx context.Context) *GetPackingSlipParams {
	return &GetPackingSlipParams{
		Context: ctx,
	}
}

// NewGetPackingSlipParamsWithHTTPClient creates a new GetPackingSlipParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPackingSlipParamsWithHTTPClient(client *http.Client) *GetPackingSlipParams {
	return &GetPackingSlipParams{
		HTTPClient: client,
	}
}

/*
GetPackingSlipParams contains all the parameters to send to the API endpoint

	for the get packing slip operation.

	Typically these are written to a http.Request.
*/
type GetPackingSlipParams struct {

	/* PurchaseOrderNumber.

	   The purchaseOrderNumber for the packing slip you want.
	*/
	PurchaseOrderNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get packing slip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackingSlipParams) WithDefaults() *GetPackingSlipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get packing slip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackingSlipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get packing slip params
func (o *GetPackingSlipParams) WithTimeout(timeout time.Duration) *GetPackingSlipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get packing slip params
func (o *GetPackingSlipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get packing slip params
func (o *GetPackingSlipParams) WithContext(ctx context.Context) *GetPackingSlipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get packing slip params
func (o *GetPackingSlipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get packing slip params
func (o *GetPackingSlipParams) WithHTTPClient(client *http.Client) *GetPackingSlipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get packing slip params
func (o *GetPackingSlipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPurchaseOrderNumber adds the purchaseOrderNumber to the get packing slip params
func (o *GetPackingSlipParams) WithPurchaseOrderNumber(purchaseOrderNumber string) *GetPackingSlipParams {
	o.SetPurchaseOrderNumber(purchaseOrderNumber)
	return o
}

// SetPurchaseOrderNumber adds the purchaseOrderNumber to the get packing slip params
func (o *GetPackingSlipParams) SetPurchaseOrderNumber(purchaseOrderNumber string) {
	o.PurchaseOrderNumber = purchaseOrderNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackingSlipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
