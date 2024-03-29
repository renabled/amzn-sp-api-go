// Code generated by go-swagger; DO NOT EDIT.

package fba_inventory

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

// NewGetInventorySummariesParams creates a new GetInventorySummariesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventorySummariesParams() *GetInventorySummariesParams {
	return &GetInventorySummariesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventorySummariesParamsWithTimeout creates a new GetInventorySummariesParams object
// with the ability to set a timeout on a request.
func NewGetInventorySummariesParamsWithTimeout(timeout time.Duration) *GetInventorySummariesParams {
	return &GetInventorySummariesParams{
		timeout: timeout,
	}
}

// NewGetInventorySummariesParamsWithContext creates a new GetInventorySummariesParams object
// with the ability to set a context for a request.
func NewGetInventorySummariesParamsWithContext(ctx context.Context) *GetInventorySummariesParams {
	return &GetInventorySummariesParams{
		Context: ctx,
	}
}

// NewGetInventorySummariesParamsWithHTTPClient creates a new GetInventorySummariesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventorySummariesParamsWithHTTPClient(client *http.Client) *GetInventorySummariesParams {
	return &GetInventorySummariesParams{
		HTTPClient: client,
	}
}

/*
GetInventorySummariesParams contains all the parameters to send to the API endpoint

	for the get inventory summaries operation.

	Typically these are written to a http.Request.
*/
type GetInventorySummariesParams struct {

	/* Details.

	   true to return inventory summaries with additional summarized inventory details and quantities. Otherwise, returns inventory summaries only (default value).
	*/
	Details *bool

	/* GranularityID.

	   The granularity ID for the inventory aggregation level.
	*/
	GranularityID string

	/* GranularityType.

	   The granularity type for the inventory aggregation level.
	*/
	GranularityType string

	/* MarketplaceIds.

	   The marketplace ID for the marketplace for which to return inventory summaries.
	*/
	MarketplaceIds []string

	/* NextToken.

	   String token returned in the response of your previous request. The string token will expire 30 seconds after being created.
	*/
	NextToken *string

	/* SellerSku.

	   A single seller SKU used for querying the specified seller SKU inventory summaries.
	*/
	SellerSku *string

	/* SellerSkus.

	   A list of seller SKUs for which to return inventory summaries. You may specify up to 50 SKUs.
	*/
	SellerSkus []string

	/* StartDateTime.

	   A start date and time in ISO8601 format. If specified, all inventory summaries that have changed since then are returned. You must specify a date and time that is no earlier than 18 months prior to the date and time when you call the API. Note: Changes in inboundWorkingQuantity, inboundShippedQuantity and inboundReceivingQuantity are not detected.

	   Format: date-time
	*/
	StartDateTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventorySummariesParams) WithDefaults() *GetInventorySummariesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory summaries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventorySummariesParams) SetDefaults() {
	var (
		detailsDefault = bool(false)
	)

	val := GetInventorySummariesParams{
		Details: &detailsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get inventory summaries params
func (o *GetInventorySummariesParams) WithTimeout(timeout time.Duration) *GetInventorySummariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory summaries params
func (o *GetInventorySummariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory summaries params
func (o *GetInventorySummariesParams) WithContext(ctx context.Context) *GetInventorySummariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory summaries params
func (o *GetInventorySummariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory summaries params
func (o *GetInventorySummariesParams) WithHTTPClient(client *http.Client) *GetInventorySummariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory summaries params
func (o *GetInventorySummariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetails adds the details to the get inventory summaries params
func (o *GetInventorySummariesParams) WithDetails(details *bool) *GetInventorySummariesParams {
	o.SetDetails(details)
	return o
}

// SetDetails adds the details to the get inventory summaries params
func (o *GetInventorySummariesParams) SetDetails(details *bool) {
	o.Details = details
}

// WithGranularityID adds the granularityID to the get inventory summaries params
func (o *GetInventorySummariesParams) WithGranularityID(granularityID string) *GetInventorySummariesParams {
	o.SetGranularityID(granularityID)
	return o
}

// SetGranularityID adds the granularityId to the get inventory summaries params
func (o *GetInventorySummariesParams) SetGranularityID(granularityID string) {
	o.GranularityID = granularityID
}

// WithGranularityType adds the granularityType to the get inventory summaries params
func (o *GetInventorySummariesParams) WithGranularityType(granularityType string) *GetInventorySummariesParams {
	o.SetGranularityType(granularityType)
	return o
}

// SetGranularityType adds the granularityType to the get inventory summaries params
func (o *GetInventorySummariesParams) SetGranularityType(granularityType string) {
	o.GranularityType = granularityType
}

// WithMarketplaceIds adds the marketplaceIds to the get inventory summaries params
func (o *GetInventorySummariesParams) WithMarketplaceIds(marketplaceIds []string) *GetInventorySummariesParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get inventory summaries params
func (o *GetInventorySummariesParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithNextToken adds the nextToken to the get inventory summaries params
func (o *GetInventorySummariesParams) WithNextToken(nextToken *string) *GetInventorySummariesParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get inventory summaries params
func (o *GetInventorySummariesParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithSellerSku adds the sellerSku to the get inventory summaries params
func (o *GetInventorySummariesParams) WithSellerSku(sellerSku *string) *GetInventorySummariesParams {
	o.SetSellerSku(sellerSku)
	return o
}

// SetSellerSku adds the sellerSku to the get inventory summaries params
func (o *GetInventorySummariesParams) SetSellerSku(sellerSku *string) {
	o.SellerSku = sellerSku
}

// WithSellerSkus adds the sellerSkus to the get inventory summaries params
func (o *GetInventorySummariesParams) WithSellerSkus(sellerSkus []string) *GetInventorySummariesParams {
	o.SetSellerSkus(sellerSkus)
	return o
}

// SetSellerSkus adds the sellerSkus to the get inventory summaries params
func (o *GetInventorySummariesParams) SetSellerSkus(sellerSkus []string) {
	o.SellerSkus = sellerSkus
}

// WithStartDateTime adds the startDateTime to the get inventory summaries params
func (o *GetInventorySummariesParams) WithStartDateTime(startDateTime *strfmt.DateTime) *GetInventorySummariesParams {
	o.SetStartDateTime(startDateTime)
	return o
}

// SetStartDateTime adds the startDateTime to the get inventory summaries params
func (o *GetInventorySummariesParams) SetStartDateTime(startDateTime *strfmt.DateTime) {
	o.StartDateTime = startDateTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventorySummariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Details != nil {

		// query param details
		var qrDetails bool

		if o.Details != nil {
			qrDetails = *o.Details
		}
		qDetails := swag.FormatBool(qrDetails)
		if qDetails != "" {

			if err := r.SetQueryParam("details", qDetails); err != nil {
				return err
			}
		}
	}

	// query param granularityId
	qrGranularityID := o.GranularityID
	qGranularityID := qrGranularityID
	if qGranularityID != "" {

		if err := r.SetQueryParam("granularityId", qGranularityID); err != nil {
			return err
		}
	}

	// query param granularityType
	qrGranularityType := o.GranularityType
	qGranularityType := qrGranularityType
	if qGranularityType != "" {

		if err := r.SetQueryParam("granularityType", qGranularityType); err != nil {
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

	if o.SellerSku != nil {

		// query param sellerSku
		var qrSellerSku string

		if o.SellerSku != nil {
			qrSellerSku = *o.SellerSku
		}
		qSellerSku := qrSellerSku
		if qSellerSku != "" {

			if err := r.SetQueryParam("sellerSku", qSellerSku); err != nil {
				return err
			}
		}
	}

	if o.SellerSkus != nil {

		// binding items for sellerSkus
		joinedSellerSkus := o.bindParamSellerSkus(reg)

		// query array param sellerSkus
		if err := r.SetQueryParam("sellerSkus", joinedSellerSkus...); err != nil {
			return err
		}
	}

	if o.StartDateTime != nil {

		// query param startDateTime
		var qrStartDateTime strfmt.DateTime

		if o.StartDateTime != nil {
			qrStartDateTime = *o.StartDateTime
		}
		qStartDateTime := qrStartDateTime.String()
		if qStartDateTime != "" {

			if err := r.SetQueryParam("startDateTime", qStartDateTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetInventorySummaries binds the parameter marketplaceIds
func (o *GetInventorySummariesParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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

// bindParamGetInventorySummaries binds the parameter sellerSkus
func (o *GetInventorySummariesParams) bindParamSellerSkus(formats strfmt.Registry) []string {
	sellerSkusIR := o.SellerSkus

	var sellerSkusIC []string
	for _, sellerSkusIIR := range sellerSkusIR { // explode []string

		sellerSkusIIV := sellerSkusIIR // string as string
		sellerSkusIC = append(sellerSkusIC, sellerSkusIIV)
	}

	// items.CollectionFormat: ""
	sellerSkusIS := swag.JoinByFormat(sellerSkusIC, "")

	return sellerSkusIS
}
