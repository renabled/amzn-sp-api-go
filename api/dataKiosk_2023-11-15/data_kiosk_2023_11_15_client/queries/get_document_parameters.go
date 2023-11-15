// Code generated by go-swagger; DO NOT EDIT.

package queries

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

// NewGetDocumentParams creates a new GetDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDocumentParams() *GetDocumentParams {
	return &GetDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentParamsWithTimeout creates a new GetDocumentParams object
// with the ability to set a timeout on a request.
func NewGetDocumentParamsWithTimeout(timeout time.Duration) *GetDocumentParams {
	return &GetDocumentParams{
		timeout: timeout,
	}
}

// NewGetDocumentParamsWithContext creates a new GetDocumentParams object
// with the ability to set a context for a request.
func NewGetDocumentParamsWithContext(ctx context.Context) *GetDocumentParams {
	return &GetDocumentParams{
		Context: ctx,
	}
}

// NewGetDocumentParamsWithHTTPClient creates a new GetDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDocumentParamsWithHTTPClient(client *http.Client) *GetDocumentParams {
	return &GetDocumentParams{
		HTTPClient: client,
	}
}

/*
GetDocumentParams contains all the parameters to send to the API endpoint

	for the get document operation.

	Typically these are written to a http.Request.
*/
type GetDocumentParams struct {

	/* DocumentID.

	   The identifier for the Data Kiosk document.
	*/
	DocumentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDocumentParams) WithDefaults() *GetDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get document params
func (o *GetDocumentParams) WithTimeout(timeout time.Duration) *GetDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get document params
func (o *GetDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get document params
func (o *GetDocumentParams) WithContext(ctx context.Context) *GetDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get document params
func (o *GetDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get document params
func (o *GetDocumentParams) WithHTTPClient(client *http.Client) *GetDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get document params
func (o *GetDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDocumentID adds the documentID to the get document params
func (o *GetDocumentParams) WithDocumentID(documentID string) *GetDocumentParams {
	o.SetDocumentID(documentID)
	return o
}

// SetDocumentID adds the documentId to the get document params
func (o *GetDocumentParams) SetDocumentID(documentID string) {
	o.DocumentID = documentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param documentId
	if err := r.SetPathParam("documentId", o.DocumentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
