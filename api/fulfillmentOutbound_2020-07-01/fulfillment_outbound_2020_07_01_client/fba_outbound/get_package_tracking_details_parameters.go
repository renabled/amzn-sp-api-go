// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

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
	"github.com/go-openapi/swag"
)

// NewGetPackageTrackingDetailsParams creates a new GetPackageTrackingDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPackageTrackingDetailsParams() *GetPackageTrackingDetailsParams {
	return &GetPackageTrackingDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackageTrackingDetailsParamsWithTimeout creates a new GetPackageTrackingDetailsParams object
// with the ability to set a timeout on a request.
func NewGetPackageTrackingDetailsParamsWithTimeout(timeout time.Duration) *GetPackageTrackingDetailsParams {
	return &GetPackageTrackingDetailsParams{
		timeout: timeout,
	}
}

// NewGetPackageTrackingDetailsParamsWithContext creates a new GetPackageTrackingDetailsParams object
// with the ability to set a context for a request.
func NewGetPackageTrackingDetailsParamsWithContext(ctx context.Context) *GetPackageTrackingDetailsParams {
	return &GetPackageTrackingDetailsParams{
		Context: ctx,
	}
}

// NewGetPackageTrackingDetailsParamsWithHTTPClient creates a new GetPackageTrackingDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPackageTrackingDetailsParamsWithHTTPClient(client *http.Client) *GetPackageTrackingDetailsParams {
	return &GetPackageTrackingDetailsParams{
		HTTPClient: client,
	}
}

/*
GetPackageTrackingDetailsParams contains all the parameters to send to the API endpoint

	for the get package tracking details operation.

	Typically these are written to a http.Request.
*/
type GetPackageTrackingDetailsParams struct {

	/* PackageNumber.

	   The unencrypted package identifier returned by the `getFulfillmentOrder` operation.

	   Format: int32
	*/
	PackageNumber int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get package tracking details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackageTrackingDetailsParams) WithDefaults() *GetPackageTrackingDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get package tracking details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackageTrackingDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) WithTimeout(timeout time.Duration) *GetPackageTrackingDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) WithContext(ctx context.Context) *GetPackageTrackingDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) WithHTTPClient(client *http.Client) *GetPackageTrackingDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageNumber adds the packageNumber to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) WithPackageNumber(packageNumber int32) *GetPackageTrackingDetailsParams {
	o.SetPackageNumber(packageNumber)
	return o
}

// SetPackageNumber adds the packageNumber to the get package tracking details params
func (o *GetPackageTrackingDetailsParams) SetPackageNumber(packageNumber int32) {
	o.PackageNumber = packageNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackageTrackingDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param packageNumber
	qrPackageNumber := o.PackageNumber
	qPackageNumber := swag.FormatInt32(qrPackageNumber)
	if qPackageNumber != "" {

		if err := r.SetQueryParam("packageNumber", qPackageNumber); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
