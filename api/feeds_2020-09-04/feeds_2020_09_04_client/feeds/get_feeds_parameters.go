// Code generated by go-swagger; DO NOT EDIT.

package feeds

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

// NewGetFeedsParams creates a new GetFeedsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFeedsParams() *GetFeedsParams {
	return &GetFeedsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeedsParamsWithTimeout creates a new GetFeedsParams object
// with the ability to set a timeout on a request.
func NewGetFeedsParamsWithTimeout(timeout time.Duration) *GetFeedsParams {
	return &GetFeedsParams{
		timeout: timeout,
	}
}

// NewGetFeedsParamsWithContext creates a new GetFeedsParams object
// with the ability to set a context for a request.
func NewGetFeedsParamsWithContext(ctx context.Context) *GetFeedsParams {
	return &GetFeedsParams{
		Context: ctx,
	}
}

// NewGetFeedsParamsWithHTTPClient creates a new GetFeedsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFeedsParamsWithHTTPClient(client *http.Client) *GetFeedsParams {
	return &GetFeedsParams{
		HTTPClient: client,
	}
}

/* GetFeedsParams contains all the parameters to send to the API endpoint
   for the get feeds operation.

   Typically these are written to a http.Request.
*/
type GetFeedsParams struct {

	/* CreatedSince.

	   The earliest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is 90 days ago. Feeds are retained for a maximum of 90 days.

	   Format: date-time
	*/
	CreatedSince *strfmt.DateTime

	/* CreatedUntil.

	   The latest feed creation date and time for feeds included in the response, in ISO 8601 format. The default is now.

	   Format: date-time
	*/
	CreatedUntil *strfmt.DateTime

	/* FeedTypes.

	   A list of feed types used to filter feeds. When feedTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either feedTypes or nextToken is required.
	*/
	FeedTypes []string

	/* MarketplaceIds.

	   A list of marketplace identifiers used to filter feeds. The feeds returned will match at least one of the marketplaces that you specify.
	*/
	MarketplaceIds []string

	/* NextToken.

	   A string token returned in the response to your previous request. nextToken is returned when the number of results exceeds the specified pageSize value. To get the next page of results, call the getFeeds operation and include this token as the only parameter. Specifying nextToken with any other parameters will cause the request to fail.
	*/
	NextToken *string

	/* PageSize.

	   The maximum number of feeds to return in a single call.

	   Default: 10
	*/
	PageSize *int64

	/* ProcessingStatuses.

	   A list of processing statuses used to filter feeds.
	*/
	ProcessingStatuses []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get feeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeedsParams) WithDefaults() *GetFeedsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get feeds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeedsParams) SetDefaults() {
	var (
		pageSizeDefault = int64(10)
	)

	val := GetFeedsParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get feeds params
func (o *GetFeedsParams) WithTimeout(timeout time.Duration) *GetFeedsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get feeds params
func (o *GetFeedsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get feeds params
func (o *GetFeedsParams) WithContext(ctx context.Context) *GetFeedsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get feeds params
func (o *GetFeedsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get feeds params
func (o *GetFeedsParams) WithHTTPClient(client *http.Client) *GetFeedsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get feeds params
func (o *GetFeedsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedSince adds the createdSince to the get feeds params
func (o *GetFeedsParams) WithCreatedSince(createdSince *strfmt.DateTime) *GetFeedsParams {
	o.SetCreatedSince(createdSince)
	return o
}

// SetCreatedSince adds the createdSince to the get feeds params
func (o *GetFeedsParams) SetCreatedSince(createdSince *strfmt.DateTime) {
	o.CreatedSince = createdSince
}

// WithCreatedUntil adds the createdUntil to the get feeds params
func (o *GetFeedsParams) WithCreatedUntil(createdUntil *strfmt.DateTime) *GetFeedsParams {
	o.SetCreatedUntil(createdUntil)
	return o
}

// SetCreatedUntil adds the createdUntil to the get feeds params
func (o *GetFeedsParams) SetCreatedUntil(createdUntil *strfmt.DateTime) {
	o.CreatedUntil = createdUntil
}

// WithFeedTypes adds the feedTypes to the get feeds params
func (o *GetFeedsParams) WithFeedTypes(feedTypes []string) *GetFeedsParams {
	o.SetFeedTypes(feedTypes)
	return o
}

// SetFeedTypes adds the feedTypes to the get feeds params
func (o *GetFeedsParams) SetFeedTypes(feedTypes []string) {
	o.FeedTypes = feedTypes
}

// WithMarketplaceIds adds the marketplaceIds to the get feeds params
func (o *GetFeedsParams) WithMarketplaceIds(marketplaceIds []string) *GetFeedsParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get feeds params
func (o *GetFeedsParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithNextToken adds the nextToken to the get feeds params
func (o *GetFeedsParams) WithNextToken(nextToken *string) *GetFeedsParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get feeds params
func (o *GetFeedsParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithPageSize adds the pageSize to the get feeds params
func (o *GetFeedsParams) WithPageSize(pageSize *int64) *GetFeedsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get feeds params
func (o *GetFeedsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProcessingStatuses adds the processingStatuses to the get feeds params
func (o *GetFeedsParams) WithProcessingStatuses(processingStatuses []string) *GetFeedsParams {
	o.SetProcessingStatuses(processingStatuses)
	return o
}

// SetProcessingStatuses adds the processingStatuses to the get feeds params
func (o *GetFeedsParams) SetProcessingStatuses(processingStatuses []string) {
	o.ProcessingStatuses = processingStatuses
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeedsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatedSince != nil {

		// query param createdSince
		var qrCreatedSince strfmt.DateTime

		if o.CreatedSince != nil {
			qrCreatedSince = *o.CreatedSince
		}
		qCreatedSince := qrCreatedSince.String()
		if qCreatedSince != "" {

			if err := r.SetQueryParam("createdSince", qCreatedSince); err != nil {
				return err
			}
		}
	}

	if o.CreatedUntil != nil {

		// query param createdUntil
		var qrCreatedUntil strfmt.DateTime

		if o.CreatedUntil != nil {
			qrCreatedUntil = *o.CreatedUntil
		}
		qCreatedUntil := qrCreatedUntil.String()
		if qCreatedUntil != "" {

			if err := r.SetQueryParam("createdUntil", qCreatedUntil); err != nil {
				return err
			}
		}
	}

	if o.FeedTypes != nil {

		// binding items for feedTypes
		joinedFeedTypes := o.bindParamFeedTypes(reg)

		// query array param feedTypes
		if err := r.SetQueryParam("feedTypes", joinedFeedTypes...); err != nil {
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

	if o.ProcessingStatuses != nil {

		// binding items for processingStatuses
		joinedProcessingStatuses := o.bindParamProcessingStatuses(reg)

		// query array param processingStatuses
		if err := r.SetQueryParam("processingStatuses", joinedProcessingStatuses...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetFeeds binds the parameter feedTypes
func (o *GetFeedsParams) bindParamFeedTypes(formats strfmt.Registry) []string {
	feedTypesIR := o.FeedTypes

	var feedTypesIC []string
	for _, feedTypesIIR := range feedTypesIR { // explode []string

		feedTypesIIV := feedTypesIIR // string as string
		feedTypesIC = append(feedTypesIC, feedTypesIIV)
	}

	// items.CollectionFormat: ""
	feedTypesIS := swag.JoinByFormat(feedTypesIC, "")

	return feedTypesIS
}

// bindParamGetFeeds binds the parameter marketplaceIds
func (o *GetFeedsParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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

// bindParamGetFeeds binds the parameter processingStatuses
func (o *GetFeedsParams) bindParamProcessingStatuses(formats strfmt.Registry) []string {
	processingStatusesIR := o.ProcessingStatuses

	var processingStatusesIC []string
	for _, processingStatusesIIR := range processingStatusesIR { // explode []string

		processingStatusesIIV := processingStatusesIIR // string as string
		processingStatusesIC = append(processingStatusesIC, processingStatusesIIV)
	}

	// items.CollectionFormat: ""
	processingStatusesIS := swag.JoinByFormat(processingStatusesIC, "")

	return processingStatusesIS
}
