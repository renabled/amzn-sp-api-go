// Code generated by go-swagger; DO NOT EDIT.

package service

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

	"github.com/xamandar/amzn-sp-api-go/api/services/services_models"
)

// NewAddAppointmentForServiceJobByServiceJobIDParams creates a new AddAppointmentForServiceJobByServiceJobIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAppointmentForServiceJobByServiceJobIDParams() *AddAppointmentForServiceJobByServiceJobIDParams {
	return &AddAppointmentForServiceJobByServiceJobIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAppointmentForServiceJobByServiceJobIDParamsWithTimeout creates a new AddAppointmentForServiceJobByServiceJobIDParams object
// with the ability to set a timeout on a request.
func NewAddAppointmentForServiceJobByServiceJobIDParamsWithTimeout(timeout time.Duration) *AddAppointmentForServiceJobByServiceJobIDParams {
	return &AddAppointmentForServiceJobByServiceJobIDParams{
		timeout: timeout,
	}
}

// NewAddAppointmentForServiceJobByServiceJobIDParamsWithContext creates a new AddAppointmentForServiceJobByServiceJobIDParams object
// with the ability to set a context for a request.
func NewAddAppointmentForServiceJobByServiceJobIDParamsWithContext(ctx context.Context) *AddAppointmentForServiceJobByServiceJobIDParams {
	return &AddAppointmentForServiceJobByServiceJobIDParams{
		Context: ctx,
	}
}

// NewAddAppointmentForServiceJobByServiceJobIDParamsWithHTTPClient creates a new AddAppointmentForServiceJobByServiceJobIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAppointmentForServiceJobByServiceJobIDParamsWithHTTPClient(client *http.Client) *AddAppointmentForServiceJobByServiceJobIDParams {
	return &AddAppointmentForServiceJobByServiceJobIDParams{
		HTTPClient: client,
	}
}

/*
AddAppointmentForServiceJobByServiceJobIDParams contains all the parameters to send to the API endpoint

	for the add appointment for service job by service job Id operation.

	Typically these are written to a http.Request.
*/
type AddAppointmentForServiceJobByServiceJobIDParams struct {

	/* Body.

	   Add appointment operation input details.
	*/
	Body *services_models.AddAppointmentRequest

	/* ServiceJobID.

	   An Amazon defined service job identifier.
	*/
	ServiceJobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add appointment for service job by service job Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAppointmentForServiceJobByServiceJobIDParams) WithDefaults() *AddAppointmentForServiceJobByServiceJobIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add appointment for service job by service job Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAppointmentForServiceJobByServiceJobIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) WithTimeout(timeout time.Duration) *AddAppointmentForServiceJobByServiceJobIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) WithContext(ctx context.Context) *AddAppointmentForServiceJobByServiceJobIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) WithHTTPClient(client *http.Client) *AddAppointmentForServiceJobByServiceJobIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) WithBody(body *services_models.AddAppointmentRequest) *AddAppointmentForServiceJobByServiceJobIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) SetBody(body *services_models.AddAppointmentRequest) {
	o.Body = body
}

// WithServiceJobID adds the serviceJobID to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) WithServiceJobID(serviceJobID string) *AddAppointmentForServiceJobByServiceJobIDParams {
	o.SetServiceJobID(serviceJobID)
	return o
}

// SetServiceJobID adds the serviceJobId to the add appointment for service job by service job Id params
func (o *AddAppointmentForServiceJobByServiceJobIDParams) SetServiceJobID(serviceJobID string) {
	o.ServiceJobID = serviceJobID
}

// WriteToRequest writes these params to a swagger request
func (o *AddAppointmentForServiceJobByServiceJobIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param serviceJobId
	if err := r.SetPathParam("serviceJobId", o.ServiceJobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
