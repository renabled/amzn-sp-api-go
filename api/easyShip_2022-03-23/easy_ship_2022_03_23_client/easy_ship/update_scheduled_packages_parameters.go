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

// NewUpdateScheduledPackagesParams creates a new UpdateScheduledPackagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateScheduledPackagesParams() *UpdateScheduledPackagesParams {
	return &UpdateScheduledPackagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScheduledPackagesParamsWithTimeout creates a new UpdateScheduledPackagesParams object
// with the ability to set a timeout on a request.
func NewUpdateScheduledPackagesParamsWithTimeout(timeout time.Duration) *UpdateScheduledPackagesParams {
	return &UpdateScheduledPackagesParams{
		timeout: timeout,
	}
}

// NewUpdateScheduledPackagesParamsWithContext creates a new UpdateScheduledPackagesParams object
// with the ability to set a context for a request.
func NewUpdateScheduledPackagesParamsWithContext(ctx context.Context) *UpdateScheduledPackagesParams {
	return &UpdateScheduledPackagesParams{
		Context: ctx,
	}
}

// NewUpdateScheduledPackagesParamsWithHTTPClient creates a new UpdateScheduledPackagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateScheduledPackagesParamsWithHTTPClient(client *http.Client) *UpdateScheduledPackagesParams {
	return &UpdateScheduledPackagesParams{
		HTTPClient: client,
	}
}

/*
UpdateScheduledPackagesParams contains all the parameters to send to the API endpoint

	for the update scheduled packages operation.

	Typically these are written to a http.Request.
*/
type UpdateScheduledPackagesParams struct {

	// UpdateScheduledPackagesRequest.
	UpdateScheduledPackagesRequest *easy_ship_2022_03_23_models.UpdateScheduledPackagesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update scheduled packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScheduledPackagesParams) WithDefaults() *UpdateScheduledPackagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update scheduled packages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateScheduledPackagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) WithTimeout(timeout time.Duration) *UpdateScheduledPackagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) WithContext(ctx context.Context) *UpdateScheduledPackagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) WithHTTPClient(client *http.Client) *UpdateScheduledPackagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpdateScheduledPackagesRequest adds the updateScheduledPackagesRequest to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) WithUpdateScheduledPackagesRequest(updateScheduledPackagesRequest *easy_ship_2022_03_23_models.UpdateScheduledPackagesRequest) *UpdateScheduledPackagesParams {
	o.SetUpdateScheduledPackagesRequest(updateScheduledPackagesRequest)
	return o
}

// SetUpdateScheduledPackagesRequest adds the updateScheduledPackagesRequest to the update scheduled packages params
func (o *UpdateScheduledPackagesParams) SetUpdateScheduledPackagesRequest(updateScheduledPackagesRequest *easy_ship_2022_03_23_models.UpdateScheduledPackagesRequest) {
	o.UpdateScheduledPackagesRequest = updateScheduledPackagesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScheduledPackagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.UpdateScheduledPackagesRequest != nil {
		if err := r.SetBodyParam(o.UpdateScheduledPackagesRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
