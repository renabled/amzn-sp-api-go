// Code generated by go-swagger; DO NOT EDIT.

package easy_ship

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

	"github.com/xamandar/amzn-sp-api-go/api/easyShip_2022-03-23/easy_ship_2022_03_23_models"
)

// NewCreateScheduledPackageParams creates a new CreateScheduledPackageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateScheduledPackageParams() *CreateScheduledPackageParams {
	return &CreateScheduledPackageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScheduledPackageParamsWithTimeout creates a new CreateScheduledPackageParams object
// with the ability to set a timeout on a request.
func NewCreateScheduledPackageParamsWithTimeout(timeout time.Duration) *CreateScheduledPackageParams {
	return &CreateScheduledPackageParams{
		timeout: timeout,
	}
}

// NewCreateScheduledPackageParamsWithContext creates a new CreateScheduledPackageParams object
// with the ability to set a context for a request.
func NewCreateScheduledPackageParamsWithContext(ctx context.Context) *CreateScheduledPackageParams {
	return &CreateScheduledPackageParams{
		Context: ctx,
	}
}

// NewCreateScheduledPackageParamsWithHTTPClient creates a new CreateScheduledPackageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateScheduledPackageParamsWithHTTPClient(client *http.Client) *CreateScheduledPackageParams {
	return &CreateScheduledPackageParams{
		HTTPClient: client,
	}
}

/*
CreateScheduledPackageParams contains all the parameters to send to the API endpoint

	for the create scheduled package operation.

	Typically these are written to a http.Request.
*/
type CreateScheduledPackageParams struct {

	// CreateScheduledPackageRequest.
	CreateScheduledPackageRequest *easy_ship_2022_03_23_models.CreateScheduledPackageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create scheduled package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateScheduledPackageParams) WithDefaults() *CreateScheduledPackageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create scheduled package params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateScheduledPackageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create scheduled package params
func (o *CreateScheduledPackageParams) WithTimeout(timeout time.Duration) *CreateScheduledPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scheduled package params
func (o *CreateScheduledPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scheduled package params
func (o *CreateScheduledPackageParams) WithContext(ctx context.Context) *CreateScheduledPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scheduled package params
func (o *CreateScheduledPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scheduled package params
func (o *CreateScheduledPackageParams) WithHTTPClient(client *http.Client) *CreateScheduledPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scheduled package params
func (o *CreateScheduledPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateScheduledPackageRequest adds the createScheduledPackageRequest to the create scheduled package params
func (o *CreateScheduledPackageParams) WithCreateScheduledPackageRequest(createScheduledPackageRequest *easy_ship_2022_03_23_models.CreateScheduledPackageRequest) *CreateScheduledPackageParams {
	o.SetCreateScheduledPackageRequest(createScheduledPackageRequest)
	return o
}

// SetCreateScheduledPackageRequest adds the createScheduledPackageRequest to the create scheduled package params
func (o *CreateScheduledPackageParams) SetCreateScheduledPackageRequest(createScheduledPackageRequest *easy_ship_2022_03_23_models.CreateScheduledPackageRequest) {
	o.CreateScheduledPackageRequest = createScheduledPackageRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScheduledPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CreateScheduledPackageRequest != nil {
		if err := r.SetBodyParam(o.CreateScheduledPackageRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
