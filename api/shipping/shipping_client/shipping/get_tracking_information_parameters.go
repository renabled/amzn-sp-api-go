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
)

// NewGetTrackingInformationParams creates a new GetTrackingInformationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTrackingInformationParams() *GetTrackingInformationParams {
	return &GetTrackingInformationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrackingInformationParamsWithTimeout creates a new GetTrackingInformationParams object
// with the ability to set a timeout on a request.
func NewGetTrackingInformationParamsWithTimeout(timeout time.Duration) *GetTrackingInformationParams {
	return &GetTrackingInformationParams{
		timeout: timeout,
	}
}

// NewGetTrackingInformationParamsWithContext creates a new GetTrackingInformationParams object
// with the ability to set a context for a request.
func NewGetTrackingInformationParamsWithContext(ctx context.Context) *GetTrackingInformationParams {
	return &GetTrackingInformationParams{
		Context: ctx,
	}
}

// NewGetTrackingInformationParamsWithHTTPClient creates a new GetTrackingInformationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTrackingInformationParamsWithHTTPClient(client *http.Client) *GetTrackingInformationParams {
	return &GetTrackingInformationParams{
		HTTPClient: client,
	}
}

/*
GetTrackingInformationParams contains all the parameters to send to the API endpoint

	for the get tracking information operation.

	Typically these are written to a http.Request.
*/
type GetTrackingInformationParams struct {

	// TrackingID.
	TrackingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tracking information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrackingInformationParams) WithDefaults() *GetTrackingInformationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tracking information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrackingInformationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tracking information params
func (o *GetTrackingInformationParams) WithTimeout(timeout time.Duration) *GetTrackingInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tracking information params
func (o *GetTrackingInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tracking information params
func (o *GetTrackingInformationParams) WithContext(ctx context.Context) *GetTrackingInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tracking information params
func (o *GetTrackingInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tracking information params
func (o *GetTrackingInformationParams) WithHTTPClient(client *http.Client) *GetTrackingInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tracking information params
func (o *GetTrackingInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTrackingID adds the trackingID to the get tracking information params
func (o *GetTrackingInformationParams) WithTrackingID(trackingID string) *GetTrackingInformationParams {
	o.SetTrackingID(trackingID)
	return o
}

// SetTrackingID adds the trackingId to the get tracking information params
func (o *GetTrackingInformationParams) SetTrackingID(trackingID string) {
	o.TrackingID = trackingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrackingInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param trackingId
	if err := r.SetPathParam("trackingId", o.TrackingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
