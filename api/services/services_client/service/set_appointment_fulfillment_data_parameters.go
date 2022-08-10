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

// NewSetAppointmentFulfillmentDataParams creates a new SetAppointmentFulfillmentDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetAppointmentFulfillmentDataParams() *SetAppointmentFulfillmentDataParams {
	return &SetAppointmentFulfillmentDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetAppointmentFulfillmentDataParamsWithTimeout creates a new SetAppointmentFulfillmentDataParams object
// with the ability to set a timeout on a request.
func NewSetAppointmentFulfillmentDataParamsWithTimeout(timeout time.Duration) *SetAppointmentFulfillmentDataParams {
	return &SetAppointmentFulfillmentDataParams{
		timeout: timeout,
	}
}

// NewSetAppointmentFulfillmentDataParamsWithContext creates a new SetAppointmentFulfillmentDataParams object
// with the ability to set a context for a request.
func NewSetAppointmentFulfillmentDataParamsWithContext(ctx context.Context) *SetAppointmentFulfillmentDataParams {
	return &SetAppointmentFulfillmentDataParams{
		Context: ctx,
	}
}

// NewSetAppointmentFulfillmentDataParamsWithHTTPClient creates a new SetAppointmentFulfillmentDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetAppointmentFulfillmentDataParamsWithHTTPClient(client *http.Client) *SetAppointmentFulfillmentDataParams {
	return &SetAppointmentFulfillmentDataParams{
		HTTPClient: client,
	}
}

/* SetAppointmentFulfillmentDataParams contains all the parameters to send to the API endpoint
   for the set appointment fulfillment data operation.

   Typically these are written to a http.Request.
*/
type SetAppointmentFulfillmentDataParams struct {

	/* AppointmentID.

	   An Amazon-defined identifier of active service job appointment.
	*/
	AppointmentID string

	/* Body.

	   Appointment fulfillment data collection details.
	*/
	Body *services_models.SetAppointmentFulfillmentDataRequest

	/* ServiceJobID.

	   An Amazon-defined service job identifier. Get this value by calling the `getServiceJobs` operation of the Services API.
	*/
	ServiceJobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set appointment fulfillment data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAppointmentFulfillmentDataParams) WithDefaults() *SetAppointmentFulfillmentDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set appointment fulfillment data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAppointmentFulfillmentDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) WithTimeout(timeout time.Duration) *SetAppointmentFulfillmentDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) WithContext(ctx context.Context) *SetAppointmentFulfillmentDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) WithHTTPClient(client *http.Client) *SetAppointmentFulfillmentDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppointmentID adds the appointmentID to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) WithAppointmentID(appointmentID string) *SetAppointmentFulfillmentDataParams {
	o.SetAppointmentID(appointmentID)
	return o
}

// SetAppointmentID adds the appointmentId to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) SetAppointmentID(appointmentID string) {
	o.AppointmentID = appointmentID
}

// WithBody adds the body to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) WithBody(body *services_models.SetAppointmentFulfillmentDataRequest) *SetAppointmentFulfillmentDataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) SetBody(body *services_models.SetAppointmentFulfillmentDataRequest) {
	o.Body = body
}

// WithServiceJobID adds the serviceJobID to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) WithServiceJobID(serviceJobID string) *SetAppointmentFulfillmentDataParams {
	o.SetServiceJobID(serviceJobID)
	return o
}

// SetServiceJobID adds the serviceJobId to the set appointment fulfillment data params
func (o *SetAppointmentFulfillmentDataParams) SetServiceJobID(serviceJobID string) {
	o.ServiceJobID = serviceJobID
}

// WriteToRequest writes these params to a swagger request
func (o *SetAppointmentFulfillmentDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
