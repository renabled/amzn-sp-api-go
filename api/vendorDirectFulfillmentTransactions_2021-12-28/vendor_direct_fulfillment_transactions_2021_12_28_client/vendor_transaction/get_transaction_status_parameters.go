// Code generated by go-swagger; DO NOT EDIT.

package vendor_transaction

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

// NewGetTransactionStatusParams creates a new GetTransactionStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTransactionStatusParams() *GetTransactionStatusParams {
	return &GetTransactionStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionStatusParamsWithTimeout creates a new GetTransactionStatusParams object
// with the ability to set a timeout on a request.
func NewGetTransactionStatusParamsWithTimeout(timeout time.Duration) *GetTransactionStatusParams {
	return &GetTransactionStatusParams{
		timeout: timeout,
	}
}

// NewGetTransactionStatusParamsWithContext creates a new GetTransactionStatusParams object
// with the ability to set a context for a request.
func NewGetTransactionStatusParamsWithContext(ctx context.Context) *GetTransactionStatusParams {
	return &GetTransactionStatusParams{
		Context: ctx,
	}
}

// NewGetTransactionStatusParamsWithHTTPClient creates a new GetTransactionStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTransactionStatusParamsWithHTTPClient(client *http.Client) *GetTransactionStatusParams {
	return &GetTransactionStatusParams{
		HTTPClient: client,
	}
}

/*
GetTransactionStatusParams contains all the parameters to send to the API endpoint

	for the get transaction status operation.

	Typically these are written to a http.Request.
*/
type GetTransactionStatusParams struct {

	/* TransactionID.

	   Previously returned in the response to the POST request of a specific transaction.
	*/
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get transaction status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionStatusParams) WithDefaults() *GetTransactionStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get transaction status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get transaction status params
func (o *GetTransactionStatusParams) WithTimeout(timeout time.Duration) *GetTransactionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transaction status params
func (o *GetTransactionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transaction status params
func (o *GetTransactionStatusParams) WithContext(ctx context.Context) *GetTransactionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transaction status params
func (o *GetTransactionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transaction status params
func (o *GetTransactionStatusParams) WithHTTPClient(client *http.Client) *GetTransactionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transaction status params
func (o *GetTransactionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get transaction status params
func (o *GetTransactionStatusParams) WithTransactionID(transactionID string) *GetTransactionStatusParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get transaction status params
func (o *GetTransactionStatusParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
