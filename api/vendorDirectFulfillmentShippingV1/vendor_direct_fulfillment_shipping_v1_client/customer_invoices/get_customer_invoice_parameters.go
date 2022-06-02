// Code generated by go-swagger; DO NOT EDIT.

package customer_invoices

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

// NewGetCustomerInvoiceParams creates a new GetCustomerInvoiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCustomerInvoiceParams() *GetCustomerInvoiceParams {
	return &GetCustomerInvoiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomerInvoiceParamsWithTimeout creates a new GetCustomerInvoiceParams object
// with the ability to set a timeout on a request.
func NewGetCustomerInvoiceParamsWithTimeout(timeout time.Duration) *GetCustomerInvoiceParams {
	return &GetCustomerInvoiceParams{
		timeout: timeout,
	}
}

// NewGetCustomerInvoiceParamsWithContext creates a new GetCustomerInvoiceParams object
// with the ability to set a context for a request.
func NewGetCustomerInvoiceParamsWithContext(ctx context.Context) *GetCustomerInvoiceParams {
	return &GetCustomerInvoiceParams{
		Context: ctx,
	}
}

// NewGetCustomerInvoiceParamsWithHTTPClient creates a new GetCustomerInvoiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCustomerInvoiceParamsWithHTTPClient(client *http.Client) *GetCustomerInvoiceParams {
	return &GetCustomerInvoiceParams{
		HTTPClient: client,
	}
}

/* GetCustomerInvoiceParams contains all the parameters to send to the API endpoint
   for the get customer invoice operation.

   Typically these are written to a http.Request.
*/
type GetCustomerInvoiceParams struct {

	/* PurchaseOrderNumber.

	   Purchase order number of the shipment for which to return the invoice.
	*/
	PurchaseOrderNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get customer invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomerInvoiceParams) WithDefaults() *GetCustomerInvoiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get customer invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCustomerInvoiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get customer invoice params
func (o *GetCustomerInvoiceParams) WithTimeout(timeout time.Duration) *GetCustomerInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get customer invoice params
func (o *GetCustomerInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get customer invoice params
func (o *GetCustomerInvoiceParams) WithContext(ctx context.Context) *GetCustomerInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get customer invoice params
func (o *GetCustomerInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get customer invoice params
func (o *GetCustomerInvoiceParams) WithHTTPClient(client *http.Client) *GetCustomerInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get customer invoice params
func (o *GetCustomerInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPurchaseOrderNumber adds the purchaseOrderNumber to the get customer invoice params
func (o *GetCustomerInvoiceParams) WithPurchaseOrderNumber(purchaseOrderNumber string) *GetCustomerInvoiceParams {
	o.SetPurchaseOrderNumber(purchaseOrderNumber)
	return o
}

// SetPurchaseOrderNumber adds the purchaseOrderNumber to the get customer invoice params
func (o *GetCustomerInvoiceParams) SetPurchaseOrderNumber(purchaseOrderNumber string) {
	o.PurchaseOrderNumber = purchaseOrderNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomerInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
