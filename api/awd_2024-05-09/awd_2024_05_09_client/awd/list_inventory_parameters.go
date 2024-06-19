// Code generated by go-swagger; DO NOT EDIT.

package awd

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

// NewListInventoryParams creates a new ListInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListInventoryParams() *ListInventoryParams {
	return &ListInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListInventoryParamsWithTimeout creates a new ListInventoryParams object
// with the ability to set a timeout on a request.
func NewListInventoryParamsWithTimeout(timeout time.Duration) *ListInventoryParams {
	return &ListInventoryParams{
		timeout: timeout,
	}
}

// NewListInventoryParamsWithContext creates a new ListInventoryParams object
// with the ability to set a context for a request.
func NewListInventoryParamsWithContext(ctx context.Context) *ListInventoryParams {
	return &ListInventoryParams{
		Context: ctx,
	}
}

// NewListInventoryParamsWithHTTPClient creates a new ListInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewListInventoryParamsWithHTTPClient(client *http.Client) *ListInventoryParams {
	return &ListInventoryParams{
		HTTPClient: client,
	}
}

/*
ListInventoryParams contains all the parameters to send to the API endpoint

	for the list inventory operation.

	Typically these are written to a http.Request.
*/
type ListInventoryParams struct {

	/* Details.

	   Set to `SHOW` to return summaries with additional inventory details. Defaults to `HIDE,` which returns only inventory summary totals.
	*/
	Details *string

	/* MaxResults.

	   Maximum number of results to return.

	   Format: int32
	   Default: 25
	*/
	MaxResults *int32

	/* NextToken.

	   Token to retrieve the next set of paginated results.
	*/
	NextToken *string

	/* Sku.

	   Filter by seller or merchant SKU for the item.
	*/
	Sku *string

	/* SortOrder.

	   Sort the response in `ASCENDING` or `DESCENDING` order.
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInventoryParams) WithDefaults() *ListInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInventoryParams) SetDefaults() {
	var (
		maxResultsDefault = int32(25)
	)

	val := ListInventoryParams{
		MaxResults: &maxResultsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list inventory params
func (o *ListInventoryParams) WithTimeout(timeout time.Duration) *ListInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list inventory params
func (o *ListInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list inventory params
func (o *ListInventoryParams) WithContext(ctx context.Context) *ListInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list inventory params
func (o *ListInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list inventory params
func (o *ListInventoryParams) WithHTTPClient(client *http.Client) *ListInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list inventory params
func (o *ListInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetails adds the details to the list inventory params
func (o *ListInventoryParams) WithDetails(details *string) *ListInventoryParams {
	o.SetDetails(details)
	return o
}

// SetDetails adds the details to the list inventory params
func (o *ListInventoryParams) SetDetails(details *string) {
	o.Details = details
}

// WithMaxResults adds the maxResults to the list inventory params
func (o *ListInventoryParams) WithMaxResults(maxResults *int32) *ListInventoryParams {
	o.SetMaxResults(maxResults)
	return o
}

// SetMaxResults adds the maxResults to the list inventory params
func (o *ListInventoryParams) SetMaxResults(maxResults *int32) {
	o.MaxResults = maxResults
}

// WithNextToken adds the nextToken to the list inventory params
func (o *ListInventoryParams) WithNextToken(nextToken *string) *ListInventoryParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list inventory params
func (o *ListInventoryParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithSku adds the sku to the list inventory params
func (o *ListInventoryParams) WithSku(sku *string) *ListInventoryParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the list inventory params
func (o *ListInventoryParams) SetSku(sku *string) {
	o.Sku = sku
}

// WithSortOrder adds the sortOrder to the list inventory params
func (o *ListInventoryParams) WithSortOrder(sortOrder *string) *ListInventoryParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the list inventory params
func (o *ListInventoryParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *ListInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Details != nil {

		// query param details
		var qrDetails string

		if o.Details != nil {
			qrDetails = *o.Details
		}
		qDetails := qrDetails
		if qDetails != "" {

			if err := r.SetQueryParam("details", qDetails); err != nil {
				return err
			}
		}
	}

	if o.MaxResults != nil {

		// query param maxResults
		var qrMaxResults int32

		if o.MaxResults != nil {
			qrMaxResults = *o.MaxResults
		}
		qMaxResults := swag.FormatInt32(qrMaxResults)
		if qMaxResults != "" {

			if err := r.SetQueryParam("maxResults", qMaxResults); err != nil {
				return err
			}
		}
	}

	if o.NextToken != nil {

		// query param nextToken
		var qrNextToken string

		if o.NextToken != nil {
			qrNextToken = *o.NextToken
		}
		qNextToken := qrNextToken
		if qNextToken != "" {

			if err := r.SetQueryParam("nextToken", qNextToken); err != nil {
				return err
			}
		}
	}

	if o.Sku != nil {

		// query param sku
		var qrSku string

		if o.Sku != nil {
			qrSku = *o.Sku
		}
		qSku := qrSku
		if qSku != "" {

			if err := r.SetQueryParam("sku", qSku); err != nil {
				return err
			}
		}
	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
