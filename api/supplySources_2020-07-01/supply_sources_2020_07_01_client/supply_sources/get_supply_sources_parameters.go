// Code generated by go-swagger; DO NOT EDIT.

package supply_sources

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

// NewGetSupplySourcesParams creates a new GetSupplySourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSupplySourcesParams() *GetSupplySourcesParams {
	return &GetSupplySourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSupplySourcesParamsWithTimeout creates a new GetSupplySourcesParams object
// with the ability to set a timeout on a request.
func NewGetSupplySourcesParamsWithTimeout(timeout time.Duration) *GetSupplySourcesParams {
	return &GetSupplySourcesParams{
		timeout: timeout,
	}
}

// NewGetSupplySourcesParamsWithContext creates a new GetSupplySourcesParams object
// with the ability to set a context for a request.
func NewGetSupplySourcesParamsWithContext(ctx context.Context) *GetSupplySourcesParams {
	return &GetSupplySourcesParams{
		Context: ctx,
	}
}

// NewGetSupplySourcesParamsWithHTTPClient creates a new GetSupplySourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSupplySourcesParamsWithHTTPClient(client *http.Client) *GetSupplySourcesParams {
	return &GetSupplySourcesParams{
		HTTPClient: client,
	}
}

/*
GetSupplySourcesParams contains all the parameters to send to the API endpoint

	for the get supply sources operation.

	Typically these are written to a http.Request.
*/
type GetSupplySourcesParams struct {

	/* NextPageToken.

	   The pagination token to retrieve a specific page of results.
	*/
	NextPageToken *string

	/* PageSize.

	   The number of supply sources to return per paginated request.

	   Default: 10
	*/
	PageSize *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get supply sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSupplySourcesParams) WithDefaults() *GetSupplySourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get supply sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSupplySourcesParams) SetDefaults() {
	var (
		pageSizeDefault = float64(10)
	)

	val := GetSupplySourcesParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get supply sources params
func (o *GetSupplySourcesParams) WithTimeout(timeout time.Duration) *GetSupplySourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get supply sources params
func (o *GetSupplySourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get supply sources params
func (o *GetSupplySourcesParams) WithContext(ctx context.Context) *GetSupplySourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get supply sources params
func (o *GetSupplySourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get supply sources params
func (o *GetSupplySourcesParams) WithHTTPClient(client *http.Client) *GetSupplySourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get supply sources params
func (o *GetSupplySourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNextPageToken adds the nextPageToken to the get supply sources params
func (o *GetSupplySourcesParams) WithNextPageToken(nextPageToken *string) *GetSupplySourcesParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the get supply sources params
func (o *GetSupplySourcesParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithPageSize adds the pageSize to the get supply sources params
func (o *GetSupplySourcesParams) WithPageSize(pageSize *float64) *GetSupplySourcesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get supply sources params
func (o *GetSupplySourcesParams) SetPageSize(pageSize *float64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetSupplySourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NextPageToken != nil {

		// query param nextPageToken
		var qrNextPageToken string

		if o.NextPageToken != nil {
			qrNextPageToken = *o.NextPageToken
		}
		qNextPageToken := qrNextPageToken
		if qNextPageToken != "" {

			if err := r.SetQueryParam("nextPageToken", qNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize float64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatFloat64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
