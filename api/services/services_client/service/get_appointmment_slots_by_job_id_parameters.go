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
	"github.com/go-openapi/swag"
)

// NewGetAppointmmentSlotsByJobIDParams creates a new GetAppointmmentSlotsByJobIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppointmmentSlotsByJobIDParams() *GetAppointmmentSlotsByJobIDParams {
	return &GetAppointmmentSlotsByJobIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppointmmentSlotsByJobIDParamsWithTimeout creates a new GetAppointmmentSlotsByJobIDParams object
// with the ability to set a timeout on a request.
func NewGetAppointmmentSlotsByJobIDParamsWithTimeout(timeout time.Duration) *GetAppointmmentSlotsByJobIDParams {
	return &GetAppointmmentSlotsByJobIDParams{
		timeout: timeout,
	}
}

// NewGetAppointmmentSlotsByJobIDParamsWithContext creates a new GetAppointmmentSlotsByJobIDParams object
// with the ability to set a context for a request.
func NewGetAppointmmentSlotsByJobIDParamsWithContext(ctx context.Context) *GetAppointmmentSlotsByJobIDParams {
	return &GetAppointmmentSlotsByJobIDParams{
		Context: ctx,
	}
}

// NewGetAppointmmentSlotsByJobIDParamsWithHTTPClient creates a new GetAppointmmentSlotsByJobIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppointmmentSlotsByJobIDParamsWithHTTPClient(client *http.Client) *GetAppointmmentSlotsByJobIDParams {
	return &GetAppointmmentSlotsByJobIDParams{
		HTTPClient: client,
	}
}

/* GetAppointmmentSlotsByJobIDParams contains all the parameters to send to the API endpoint
   for the get appointmment slots by job Id operation.

   Typically these are written to a http.Request.
*/
type GetAppointmmentSlotsByJobIDParams struct {

	/* EndTime.

	   A time up to which the appointment slots will be retrieved. The specified time must be in ISO 8601 format. If `endTime` is provided, `startTime` should also be provided. Default value is as per business configuration. Maximum range of appointment slots can be 90 days.
	*/
	EndTime *string

	/* MarketplaceIds.

	   An identifier for the marketplace in which the resource operates.
	*/
	MarketplaceIds []string

	/* ServiceJobID.

	   A service job identifier to retrive appointment slots for associated service.
	*/
	ServiceJobID string

	/* StartTime.

	   A time from which the appointment slots will be retrieved. The specified time must be in ISO 8601 format. If `startTime` is provided, `endTime` should also be provided. Default value is as per business configuration.
	*/
	StartTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get appointmment slots by job Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppointmmentSlotsByJobIDParams) WithDefaults() *GetAppointmmentSlotsByJobIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get appointmment slots by job Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppointmmentSlotsByJobIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) WithTimeout(timeout time.Duration) *GetAppointmmentSlotsByJobIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) WithContext(ctx context.Context) *GetAppointmmentSlotsByJobIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) WithHTTPClient(client *http.Client) *GetAppointmmentSlotsByJobIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndTime adds the endTime to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) WithEndTime(endTime *string) *GetAppointmmentSlotsByJobIDParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithMarketplaceIds adds the marketplaceIds to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) WithMarketplaceIds(marketplaceIds []string) *GetAppointmmentSlotsByJobIDParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithServiceJobID adds the serviceJobID to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) WithServiceJobID(serviceJobID string) *GetAppointmmentSlotsByJobIDParams {
	o.SetServiceJobID(serviceJobID)
	return o
}

// SetServiceJobID adds the serviceJobId to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) SetServiceJobID(serviceJobID string) {
	o.ServiceJobID = serviceJobID
}

// WithStartTime adds the startTime to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) WithStartTime(startTime *string) *GetAppointmmentSlotsByJobIDParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get appointmment slots by job Id params
func (o *GetAppointmmentSlotsByJobIDParams) SetStartTime(startTime *string) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppointmmentSlotsByJobIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime string

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.MarketplaceIds != nil {

		// binding items for marketplaceIds
		joinedMarketplaceIds := o.bindParamMarketplaceIds(reg)

		// query array param marketplaceIds
		if err := r.SetQueryParam("marketplaceIds", joinedMarketplaceIds...); err != nil {
			return err
		}
	}

	// path param serviceJobId
	if err := r.SetPathParam("serviceJobId", o.ServiceJobID); err != nil {
		return err
	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime string

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime
		if qStartTime != "" {

			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAppointmmentSlotsByJobID binds the parameter marketplaceIds
func (o *GetAppointmmentSlotsByJobIDParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
	marketplaceIdsIR := o.MarketplaceIds

	var marketplaceIdsIC []string
	for _, marketplaceIdsIIR := range marketplaceIdsIR { // explode []string

		marketplaceIdsIIV := marketplaceIdsIIR // string as string
		marketplaceIdsIC = append(marketplaceIdsIC, marketplaceIdsIIV)
	}

	// items.CollectionFormat: ""
	marketplaceIdsIS := swag.JoinByFormat(marketplaceIdsIC, "")

	return marketplaceIdsIS
}
