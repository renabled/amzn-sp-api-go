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

	"github.com/xamandar/amzn-sp-api-go/api/services/services_models"
)

// NewUpdateReservationParams creates a new UpdateReservationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateReservationParams() *UpdateReservationParams {
	return &UpdateReservationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReservationParamsWithTimeout creates a new UpdateReservationParams object
// with the ability to set a timeout on a request.
func NewUpdateReservationParamsWithTimeout(timeout time.Duration) *UpdateReservationParams {
	return &UpdateReservationParams{
		timeout: timeout,
	}
}

// NewUpdateReservationParamsWithContext creates a new UpdateReservationParams object
// with the ability to set a context for a request.
func NewUpdateReservationParamsWithContext(ctx context.Context) *UpdateReservationParams {
	return &UpdateReservationParams{
		Context: ctx,
	}
}

// NewUpdateReservationParamsWithHTTPClient creates a new UpdateReservationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateReservationParamsWithHTTPClient(client *http.Client) *UpdateReservationParams {
	return &UpdateReservationParams{
		HTTPClient: client,
	}
}

/*
UpdateReservationParams contains all the parameters to send to the API endpoint

	for the update reservation operation.

	Typically these are written to a http.Request.
*/
type UpdateReservationParams struct {

	/* Body.

	   Reservation details
	*/
	Body *services_models.UpdateReservationRequest

	/* MarketplaceIds.

	   An identifier for the marketplace in which the resource operates.
	*/
	MarketplaceIds []string

	/* ReservationID.

	   Reservation Identifier
	*/
	ReservationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReservationParams) WithDefaults() *UpdateReservationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateReservationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update reservation params
func (o *UpdateReservationParams) WithTimeout(timeout time.Duration) *UpdateReservationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update reservation params
func (o *UpdateReservationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update reservation params
func (o *UpdateReservationParams) WithContext(ctx context.Context) *UpdateReservationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update reservation params
func (o *UpdateReservationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update reservation params
func (o *UpdateReservationParams) WithHTTPClient(client *http.Client) *UpdateReservationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update reservation params
func (o *UpdateReservationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update reservation params
func (o *UpdateReservationParams) WithBody(body *services_models.UpdateReservationRequest) *UpdateReservationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update reservation params
func (o *UpdateReservationParams) SetBody(body *services_models.UpdateReservationRequest) {
	o.Body = body
}

// WithMarketplaceIds adds the marketplaceIds to the update reservation params
func (o *UpdateReservationParams) WithMarketplaceIds(marketplaceIds []string) *UpdateReservationParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the update reservation params
func (o *UpdateReservationParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithReservationID adds the reservationID to the update reservation params
func (o *UpdateReservationParams) WithReservationID(reservationID string) *UpdateReservationParams {
	o.SetReservationID(reservationID)
	return o
}

// SetReservationID adds the reservationId to the update reservation params
func (o *UpdateReservationParams) SetReservationID(reservationID string) {
	o.ReservationID = reservationID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReservationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

	// path param reservationId
	if err := r.SetPathParam("reservationId", o.ReservationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUpdateReservation binds the parameter marketplaceIds
func (o *UpdateReservationParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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
