// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewDeleteDestinationParams creates a new DeleteDestinationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDestinationParams() *DeleteDestinationParams {
	return &DeleteDestinationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDestinationParamsWithTimeout creates a new DeleteDestinationParams object
// with the ability to set a timeout on a request.
func NewDeleteDestinationParamsWithTimeout(timeout time.Duration) *DeleteDestinationParams {
	return &DeleteDestinationParams{
		timeout: timeout,
	}
}

// NewDeleteDestinationParamsWithContext creates a new DeleteDestinationParams object
// with the ability to set a context for a request.
func NewDeleteDestinationParamsWithContext(ctx context.Context) *DeleteDestinationParams {
	return &DeleteDestinationParams{
		Context: ctx,
	}
}

// NewDeleteDestinationParamsWithHTTPClient creates a new DeleteDestinationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDestinationParamsWithHTTPClient(client *http.Client) *DeleteDestinationParams {
	return &DeleteDestinationParams{
		HTTPClient: client,
	}
}

/*
DeleteDestinationParams contains all the parameters to send to the API endpoint

	for the delete destination operation.

	Typically these are written to a http.Request.
*/
type DeleteDestinationParams struct {

	/* DestinationID.

	   The identifier for the destination that you want to delete.
	*/
	DestinationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDestinationParams) WithDefaults() *DeleteDestinationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete destination params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDestinationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete destination params
func (o *DeleteDestinationParams) WithTimeout(timeout time.Duration) *DeleteDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete destination params
func (o *DeleteDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete destination params
func (o *DeleteDestinationParams) WithContext(ctx context.Context) *DeleteDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete destination params
func (o *DeleteDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete destination params
func (o *DeleteDestinationParams) WithHTTPClient(client *http.Client) *DeleteDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete destination params
func (o *DeleteDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestinationID adds the destinationID to the delete destination params
func (o *DeleteDestinationParams) WithDestinationID(destinationID string) *DeleteDestinationParams {
	o.SetDestinationID(destinationID)
	return o
}

// SetDestinationID adds the destinationId to the delete destination params
func (o *DeleteDestinationParams) SetDestinationID(destinationID string) {
	o.DestinationID = destinationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param destinationId
	if err := r.SetPathParam("destinationId", o.DestinationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
