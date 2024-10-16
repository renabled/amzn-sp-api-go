// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

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

// NewListDeliveryWindowOptionsParams creates a new ListDeliveryWindowOptionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDeliveryWindowOptionsParams() *ListDeliveryWindowOptionsParams {
	return &ListDeliveryWindowOptionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDeliveryWindowOptionsParamsWithTimeout creates a new ListDeliveryWindowOptionsParams object
// with the ability to set a timeout on a request.
func NewListDeliveryWindowOptionsParamsWithTimeout(timeout time.Duration) *ListDeliveryWindowOptionsParams {
	return &ListDeliveryWindowOptionsParams{
		timeout: timeout,
	}
}

// NewListDeliveryWindowOptionsParamsWithContext creates a new ListDeliveryWindowOptionsParams object
// with the ability to set a context for a request.
func NewListDeliveryWindowOptionsParamsWithContext(ctx context.Context) *ListDeliveryWindowOptionsParams {
	return &ListDeliveryWindowOptionsParams{
		Context: ctx,
	}
}

// NewListDeliveryWindowOptionsParamsWithHTTPClient creates a new ListDeliveryWindowOptionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDeliveryWindowOptionsParamsWithHTTPClient(client *http.Client) *ListDeliveryWindowOptionsParams {
	return &ListDeliveryWindowOptionsParams{
		HTTPClient: client,
	}
}

/*
ListDeliveryWindowOptionsParams contains all the parameters to send to the API endpoint

	for the list delivery window options operation.

	Typically these are written to a http.Request.
*/
type ListDeliveryWindowOptionsParams struct {

	/* InboundPlanID.

	   Identifier of an inbound plan.
	*/
	InboundPlanID string

	/* PageSize.

	   The number of delivery window options to return in the response matching the given query.

	   Default: 10
	*/
	PageSize *int64

	/* PaginationToken.

	   A token to fetch a certain page when there are multiple pages worth of results. The value of this token is fetched from the `pagination` returned in the API response. In the absence of the token value from the query parameter the API returns the first page of the result.
	*/
	PaginationToken *string

	/* ShipmentID.

	   The shipment to get delivery window options for.
	*/
	ShipmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list delivery window options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDeliveryWindowOptionsParams) WithDefaults() *ListDeliveryWindowOptionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list delivery window options params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDeliveryWindowOptionsParams) SetDefaults() {
	var (
		pageSizeDefault = int64(10)
	)

	val := ListDeliveryWindowOptionsParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) WithTimeout(timeout time.Duration) *ListDeliveryWindowOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) WithContext(ctx context.Context) *ListDeliveryWindowOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) WithHTTPClient(client *http.Client) *ListDeliveryWindowOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboundPlanID adds the inboundPlanID to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) WithInboundPlanID(inboundPlanID string) *ListDeliveryWindowOptionsParams {
	o.SetInboundPlanID(inboundPlanID)
	return o
}

// SetInboundPlanID adds the inboundPlanId to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) SetInboundPlanID(inboundPlanID string) {
	o.InboundPlanID = inboundPlanID
}

// WithPageSize adds the pageSize to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) WithPageSize(pageSize *int64) *ListDeliveryWindowOptionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPaginationToken adds the paginationToken to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) WithPaginationToken(paginationToken *string) *ListDeliveryWindowOptionsParams {
	o.SetPaginationToken(paginationToken)
	return o
}

// SetPaginationToken adds the paginationToken to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) SetPaginationToken(paginationToken *string) {
	o.PaginationToken = paginationToken
}

// WithShipmentID adds the shipmentID to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) WithShipmentID(shipmentID string) *ListDeliveryWindowOptionsParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the list delivery window options params
func (o *ListDeliveryWindowOptionsParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *ListDeliveryWindowOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inboundPlanId
	if err := r.SetPathParam("inboundPlanId", o.InboundPlanID); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationToken != nil {

		// query param paginationToken
		var qrPaginationToken string

		if o.PaginationToken != nil {
			qrPaginationToken = *o.PaginationToken
		}
		qPaginationToken := qrPaginationToken
		if qPaginationToken != "" {

			if err := r.SetQueryParam("paginationToken", qPaginationToken); err != nil {
				return err
			}
		}
	}

	// path param shipmentId
	if err := r.SetPathParam("shipmentId", o.ShipmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
