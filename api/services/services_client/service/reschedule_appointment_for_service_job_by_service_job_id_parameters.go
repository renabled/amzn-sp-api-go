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

	"github.com/renabled/amzn-sp-api-go/api/services/services_models"
)

// NewRescheduleAppointmentForServiceJobByServiceJobIDParams creates a new RescheduleAppointmentForServiceJobByServiceJobIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRescheduleAppointmentForServiceJobByServiceJobIDParams() *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	return &RescheduleAppointmentForServiceJobByServiceJobIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDParamsWithTimeout creates a new RescheduleAppointmentForServiceJobByServiceJobIDParams object
// with the ability to set a timeout on a request.
func NewRescheduleAppointmentForServiceJobByServiceJobIDParamsWithTimeout(timeout time.Duration) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	return &RescheduleAppointmentForServiceJobByServiceJobIDParams{
		timeout: timeout,
	}
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDParamsWithContext creates a new RescheduleAppointmentForServiceJobByServiceJobIDParams object
// with the ability to set a context for a request.
func NewRescheduleAppointmentForServiceJobByServiceJobIDParamsWithContext(ctx context.Context) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	return &RescheduleAppointmentForServiceJobByServiceJobIDParams{
		Context: ctx,
	}
}

// NewRescheduleAppointmentForServiceJobByServiceJobIDParamsWithHTTPClient creates a new RescheduleAppointmentForServiceJobByServiceJobIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewRescheduleAppointmentForServiceJobByServiceJobIDParamsWithHTTPClient(client *http.Client) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	return &RescheduleAppointmentForServiceJobByServiceJobIDParams{
		HTTPClient: client,
	}
}

/*
RescheduleAppointmentForServiceJobByServiceJobIDParams contains all the parameters to send to the API endpoint

	for the reschedule appointment for service job by service job Id operation.

	Typically these are written to a http.Request.
*/
type RescheduleAppointmentForServiceJobByServiceJobIDParams struct {

	/* AppointmentID.

	   An existing appointment identifier for the Service Job.
	*/
	AppointmentID string

	/* Body.

	   Reschedule appointment operation input details.
	*/
	Body *services_models.RescheduleAppointmentRequest

	/* ServiceJobID.

	   An Amazon defined service job identifier.
	*/
	ServiceJobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reschedule appointment for service job by service job Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WithDefaults() *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reschedule appointment for service job by service job Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WithTimeout(timeout time.Duration) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WithContext(ctx context.Context) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WithHTTPClient(client *http.Client) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppointmentID adds the appointmentID to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WithAppointmentID(appointmentID string) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	o.SetAppointmentID(appointmentID)
	return o
}

// SetAppointmentID adds the appointmentId to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) SetAppointmentID(appointmentID string) {
	o.AppointmentID = appointmentID
}

// WithBody adds the body to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WithBody(body *services_models.RescheduleAppointmentRequest) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) SetBody(body *services_models.RescheduleAppointmentRequest) {
	o.Body = body
}

// WithServiceJobID adds the serviceJobID to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WithServiceJobID(serviceJobID string) *RescheduleAppointmentForServiceJobByServiceJobIDParams {
	o.SetServiceJobID(serviceJobID)
	return o
}

// SetServiceJobID adds the serviceJobId to the reschedule appointment for service job by service job Id params
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) SetServiceJobID(serviceJobID string) {
	o.ServiceJobID = serviceJobID
}

// WriteToRequest writes these params to a swagger request
func (o *RescheduleAppointmentForServiceJobByServiceJobIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appointmentId
	if err := r.SetPathParam("appointmentId", o.AppointmentID); err != nil {
		return err
	}
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
