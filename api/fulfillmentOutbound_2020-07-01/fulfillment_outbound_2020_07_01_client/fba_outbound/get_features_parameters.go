// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

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
)

// NewGetFeaturesParams creates a new GetFeaturesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFeaturesParams() *GetFeaturesParams {
	return &GetFeaturesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeaturesParamsWithTimeout creates a new GetFeaturesParams object
// with the ability to set a timeout on a request.
func NewGetFeaturesParamsWithTimeout(timeout time.Duration) *GetFeaturesParams {
	return &GetFeaturesParams{
		timeout: timeout,
	}
}

// NewGetFeaturesParamsWithContext creates a new GetFeaturesParams object
// with the ability to set a context for a request.
func NewGetFeaturesParamsWithContext(ctx context.Context) *GetFeaturesParams {
	return &GetFeaturesParams{
		Context: ctx,
	}
}

// NewGetFeaturesParamsWithHTTPClient creates a new GetFeaturesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFeaturesParamsWithHTTPClient(client *http.Client) *GetFeaturesParams {
	return &GetFeaturesParams{
		HTTPClient: client,
	}
}

/* GetFeaturesParams contains all the parameters to send to the API endpoint
   for the get features operation.

   Typically these are written to a http.Request.
*/
type GetFeaturesParams struct {

	/* MarketplaceID.

	   The marketplace for which to return the list of features.
	*/
	MarketplaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeaturesParams) WithDefaults() *GetFeaturesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get features params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeaturesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get features params
func (o *GetFeaturesParams) WithTimeout(timeout time.Duration) *GetFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get features params
func (o *GetFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get features params
func (o *GetFeaturesParams) WithContext(ctx context.Context) *GetFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get features params
func (o *GetFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get features params
func (o *GetFeaturesParams) WithHTTPClient(client *http.Client) *GetFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get features params
func (o *GetFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMarketplaceID adds the marketplaceID to the get features params
func (o *GetFeaturesParams) WithMarketplaceID(marketplaceID string) *GetFeaturesParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the get features params
func (o *GetFeaturesParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param marketplaceId
	qrMarketplaceID := o.MarketplaceID
	qMarketplaceID := qrMarketplaceID
	if qMarketplaceID != "" {

		if err := r.SetQueryParam("marketplaceId", qMarketplaceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
