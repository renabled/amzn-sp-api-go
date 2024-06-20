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

// NewGetFeatureInventoryParams creates a new GetFeatureInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFeatureInventoryParams() *GetFeatureInventoryParams {
	return &GetFeatureInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeatureInventoryParamsWithTimeout creates a new GetFeatureInventoryParams object
// with the ability to set a timeout on a request.
func NewGetFeatureInventoryParamsWithTimeout(timeout time.Duration) *GetFeatureInventoryParams {
	return &GetFeatureInventoryParams{
		timeout: timeout,
	}
}

// NewGetFeatureInventoryParamsWithContext creates a new GetFeatureInventoryParams object
// with the ability to set a context for a request.
func NewGetFeatureInventoryParamsWithContext(ctx context.Context) *GetFeatureInventoryParams {
	return &GetFeatureInventoryParams{
		Context: ctx,
	}
}

// NewGetFeatureInventoryParamsWithHTTPClient creates a new GetFeatureInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFeatureInventoryParamsWithHTTPClient(client *http.Client) *GetFeatureInventoryParams {
	return &GetFeatureInventoryParams{
		HTTPClient: client,
	}
}

/*
GetFeatureInventoryParams contains all the parameters to send to the API endpoint

	for the get feature inventory operation.

	Typically these are written to a http.Request.
*/
type GetFeatureInventoryParams struct {

	/* FeatureName.

	   The name of the feature for which to return a list of eligible inventory.
	*/
	FeatureName string

	/* MarketplaceID.

	   The marketplace for which to return a list of the inventory that is eligible for the specified feature.
	*/
	MarketplaceID string

	/* NextToken.

	   A string token returned in the response to your previous request that is used to return the next response page. A value of null will return the first page.
	*/
	NextToken *string

	/* QueryStartDate.

	   A date that you can use to select inventory that has been updated since a specified date. An update is defined as any change in feature-enabled inventory availability. The date must be in the format yyyy-MM-ddTHH:mm:ss.sssZ

	   Format: date-time
	*/
	QueryStartDate *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get feature inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeatureInventoryParams) WithDefaults() *GetFeatureInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get feature inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeatureInventoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get feature inventory params
func (o *GetFeatureInventoryParams) WithTimeout(timeout time.Duration) *GetFeatureInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feature inventory params
func (o *GetFeatureInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feature inventory params
func (o *GetFeatureInventoryParams) WithContext(ctx context.Context) *GetFeatureInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feature inventory params
func (o *GetFeatureInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feature inventory params
func (o *GetFeatureInventoryParams) WithHTTPClient(client *http.Client) *GetFeatureInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feature inventory params
func (o *GetFeatureInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeatureName adds the featureName to the get feature inventory params
func (o *GetFeatureInventoryParams) WithFeatureName(featureName string) *GetFeatureInventoryParams {
	o.SetFeatureName(featureName)
	return o
}

// SetFeatureName adds the featureName to the get feature inventory params
func (o *GetFeatureInventoryParams) SetFeatureName(featureName string) {
	o.FeatureName = featureName
}

// WithMarketplaceID adds the marketplaceID to the get feature inventory params
func (o *GetFeatureInventoryParams) WithMarketplaceID(marketplaceID string) *GetFeatureInventoryParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the get feature inventory params
func (o *GetFeatureInventoryParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WithNextToken adds the nextToken to the get feature inventory params
func (o *GetFeatureInventoryParams) WithNextToken(nextToken *string) *GetFeatureInventoryParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get feature inventory params
func (o *GetFeatureInventoryParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithQueryStartDate adds the queryStartDate to the get feature inventory params
func (o *GetFeatureInventoryParams) WithQueryStartDate(queryStartDate *strfmt.DateTime) *GetFeatureInventoryParams {
	o.SetQueryStartDate(queryStartDate)
	return o
}

// SetQueryStartDate adds the queryStartDate to the get feature inventory params
func (o *GetFeatureInventoryParams) SetQueryStartDate(queryStartDate *strfmt.DateTime) {
	o.QueryStartDate = queryStartDate
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeatureInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param featureName
	if err := r.SetPathParam("featureName", o.FeatureName); err != nil {
		return err
	}

	// query param marketplaceId
	qrMarketplaceID := o.MarketplaceID
	qMarketplaceID := qrMarketplaceID
	if qMarketplaceID != "" {

		if err := r.SetQueryParam("marketplaceId", qMarketplaceID); err != nil {
			return err
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

	if o.QueryStartDate != nil {

		// query param queryStartDate
		var qrQueryStartDate strfmt.DateTime

		if o.QueryStartDate != nil {
			qrQueryStartDate = *o.QueryStartDate
		}
		qQueryStartDate := qrQueryStartDate.String()
		if qQueryStartDate != "" {

			if err := r.SetQueryParam("queryStartDate", qQueryStartDate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
