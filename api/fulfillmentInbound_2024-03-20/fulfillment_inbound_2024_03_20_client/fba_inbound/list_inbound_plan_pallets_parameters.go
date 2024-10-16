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

// NewListInboundPlanPalletsParams creates a new ListInboundPlanPalletsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListInboundPlanPalletsParams() *ListInboundPlanPalletsParams {
	return &ListInboundPlanPalletsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListInboundPlanPalletsParamsWithTimeout creates a new ListInboundPlanPalletsParams object
// with the ability to set a timeout on a request.
func NewListInboundPlanPalletsParamsWithTimeout(timeout time.Duration) *ListInboundPlanPalletsParams {
	return &ListInboundPlanPalletsParams{
		timeout: timeout,
	}
}

// NewListInboundPlanPalletsParamsWithContext creates a new ListInboundPlanPalletsParams object
// with the ability to set a context for a request.
func NewListInboundPlanPalletsParamsWithContext(ctx context.Context) *ListInboundPlanPalletsParams {
	return &ListInboundPlanPalletsParams{
		Context: ctx,
	}
}

// NewListInboundPlanPalletsParamsWithHTTPClient creates a new ListInboundPlanPalletsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListInboundPlanPalletsParamsWithHTTPClient(client *http.Client) *ListInboundPlanPalletsParams {
	return &ListInboundPlanPalletsParams{
		HTTPClient: client,
	}
}

/*
ListInboundPlanPalletsParams contains all the parameters to send to the API endpoint

	for the list inbound plan pallets operation.

	Typically these are written to a http.Request.
*/
type ListInboundPlanPalletsParams struct {

	/* InboundPlanID.

	   Identifier of an inbound plan.
	*/
	InboundPlanID string

	/* PageSize.

	   The number of pallets to return in the response matching the given query.

	   Default: 10
	*/
	PageSize *int64

	/* PaginationToken.

	   A token to fetch a certain page when there are multiple pages worth of results. The value of this token is fetched from the `pagination` returned in the API response. In the absence of the token value from the query parameter the API returns the first page of the result.
	*/
	PaginationToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list inbound plan pallets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInboundPlanPalletsParams) WithDefaults() *ListInboundPlanPalletsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list inbound plan pallets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInboundPlanPalletsParams) SetDefaults() {
	var (
		pageSizeDefault = int64(10)
	)

	val := ListInboundPlanPalletsParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) WithTimeout(timeout time.Duration) *ListInboundPlanPalletsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) WithContext(ctx context.Context) *ListInboundPlanPalletsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) WithHTTPClient(client *http.Client) *ListInboundPlanPalletsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInboundPlanID adds the inboundPlanID to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) WithInboundPlanID(inboundPlanID string) *ListInboundPlanPalletsParams {
	o.SetInboundPlanID(inboundPlanID)
	return o
}

// SetInboundPlanID adds the inboundPlanId to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) SetInboundPlanID(inboundPlanID string) {
	o.InboundPlanID = inboundPlanID
}

// WithPageSize adds the pageSize to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) WithPageSize(pageSize *int64) *ListInboundPlanPalletsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPaginationToken adds the paginationToken to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) WithPaginationToken(paginationToken *string) *ListInboundPlanPalletsParams {
	o.SetPaginationToken(paginationToken)
	return o
}

// SetPaginationToken adds the paginationToken to the list inbound plan pallets params
func (o *ListInboundPlanPalletsParams) SetPaginationToken(paginationToken *string) {
	o.PaginationToken = paginationToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListInboundPlanPalletsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
