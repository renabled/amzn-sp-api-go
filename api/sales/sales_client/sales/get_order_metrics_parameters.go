// Code generated by go-swagger; DO NOT EDIT.

package sales

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

// NewGetOrderMetricsParams creates a new GetOrderMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrderMetricsParams() *GetOrderMetricsParams {
	return &GetOrderMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderMetricsParamsWithTimeout creates a new GetOrderMetricsParams object
// with the ability to set a timeout on a request.
func NewGetOrderMetricsParamsWithTimeout(timeout time.Duration) *GetOrderMetricsParams {
	return &GetOrderMetricsParams{
		timeout: timeout,
	}
}

// NewGetOrderMetricsParamsWithContext creates a new GetOrderMetricsParams object
// with the ability to set a context for a request.
func NewGetOrderMetricsParamsWithContext(ctx context.Context) *GetOrderMetricsParams {
	return &GetOrderMetricsParams{
		Context: ctx,
	}
}

// NewGetOrderMetricsParamsWithHTTPClient creates a new GetOrderMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrderMetricsParamsWithHTTPClient(client *http.Client) *GetOrderMetricsParams {
	return &GetOrderMetricsParams{
		HTTPClient: client,
	}
}

/*
GetOrderMetricsParams contains all the parameters to send to the API endpoint

	for the get order metrics operation.

	Typically these are written to a http.Request.
*/
type GetOrderMetricsParams struct {

	/* Asin.

	   Filters the results by the ASIN that you specify. Specifying both ASIN and SKU returns an error. Do not include this filter if you want the response to include order metrics for all ASINs. Example: B0792R1RSN, if you want the response to include order metrics for only ASIN B0792R1RSN.
	*/
	Asin *string

	/* BuyerType.

	   Filters the results by the buyer type that you specify, B2B (business to business) or B2C (business to customer). Example: B2B, if you want the response to include order metrics for only B2B buyers.

	   Default: "All"
	*/
	BuyerType *string

	/* FirstDayOfWeek.

	   Specifies the day that the week starts on when granularity=Week, either Monday or Sunday. Default: Monday. Example: Sunday, if you want the week to start on a Sunday.

	   Default: "Monday"
	*/
	FirstDayOfWeek *string

	/* FulfillmentNetwork.

	   Filters the results by the fulfillment network that you specify, MFN (merchant fulfillment network) or AFN (Amazon fulfillment network). Do not include this filter if you want the response to include order metrics for all fulfillment networks. Example: AFN, if you want the response to include order metrics for only Amazon fulfillment network.
	*/
	FulfillmentNetwork *string

	/* Granularity.

	   The granularity of the grouping of order metrics, based on a unit of time. Specifying granularity=Hour results in a successful request only if the interval specified is less than or equal to 30 days from now. For all other granularities, the interval specified must be less or equal to 2 years from now. Specifying granularity=Total results in order metrics that are aggregated over the entire interval that you specify. If the interval start and end date don’t align with the specified granularity, the head and tail end of the response interval will contain partial data. Example: Day to get a daily breakdown of the request interval, where the day boundary is defined by the granularityTimeZone.
	*/
	Granularity string

	/* GranularityTimeZone.

	   An IANA-compatible time zone for determining the day boundary. Required when specifying a granularity value greater than Hour. The granularityTimeZone value must align with the offset of the specified interval value. For example, if the interval value uses Z notation, then granularityTimeZone must be UTC. If the interval value uses an offset, then granularityTimeZone must be an IANA-compatible time zone that matches the offset. Example: US/Pacific to compute day boundaries, accounting for daylight time savings, for US/Pacific zone.
	*/
	GranularityTimeZone *string

	/* Interval.

	   A time interval used for selecting order metrics. This takes the form of two dates separated by two hyphens (first date is inclusive; second date is exclusive). Dates are in ISO8601 format and must represent absolute time (either Z notation or offset notation). Example: 2018-09-01T00:00:00-07:00--2018-09-04T00:00:00-07:00 requests order metrics for Sept 1st, 2nd and 3rd in the -07:00 zone.
	*/
	Interval string

	/* MarketplaceIds.

	     A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.

	For example, ATVPDKIKX0DER indicates the US marketplace.
	*/
	MarketplaceIds []string

	/* Sku.

	   Filters the results by the SKU that you specify. Specifying both ASIN and SKU returns an error. Do not include this filter if you want the response to include order metrics for all SKUs. Example: TestSKU, if you want the response to include order metrics for only SKU TestSKU.
	*/
	Sku *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get order metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderMetricsParams) WithDefaults() *GetOrderMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get order metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderMetricsParams) SetDefaults() {
	var (
		buyerTypeDefault = string("All")

		firstDayOfWeekDefault = string("Monday")
	)

	val := GetOrderMetricsParams{
		BuyerType:      &buyerTypeDefault,
		FirstDayOfWeek: &firstDayOfWeekDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get order metrics params
func (o *GetOrderMetricsParams) WithTimeout(timeout time.Duration) *GetOrderMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order metrics params
func (o *GetOrderMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order metrics params
func (o *GetOrderMetricsParams) WithContext(ctx context.Context) *GetOrderMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order metrics params
func (o *GetOrderMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order metrics params
func (o *GetOrderMetricsParams) WithHTTPClient(client *http.Client) *GetOrderMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order metrics params
func (o *GetOrderMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsin adds the asin to the get order metrics params
func (o *GetOrderMetricsParams) WithAsin(asin *string) *GetOrderMetricsParams {
	o.SetAsin(asin)
	return o
}

// SetAsin adds the asin to the get order metrics params
func (o *GetOrderMetricsParams) SetAsin(asin *string) {
	o.Asin = asin
}

// WithBuyerType adds the buyerType to the get order metrics params
func (o *GetOrderMetricsParams) WithBuyerType(buyerType *string) *GetOrderMetricsParams {
	o.SetBuyerType(buyerType)
	return o
}

// SetBuyerType adds the buyerType to the get order metrics params
func (o *GetOrderMetricsParams) SetBuyerType(buyerType *string) {
	o.BuyerType = buyerType
}

// WithFirstDayOfWeek adds the firstDayOfWeek to the get order metrics params
func (o *GetOrderMetricsParams) WithFirstDayOfWeek(firstDayOfWeek *string) *GetOrderMetricsParams {
	o.SetFirstDayOfWeek(firstDayOfWeek)
	return o
}

// SetFirstDayOfWeek adds the firstDayOfWeek to the get order metrics params
func (o *GetOrderMetricsParams) SetFirstDayOfWeek(firstDayOfWeek *string) {
	o.FirstDayOfWeek = firstDayOfWeek
}

// WithFulfillmentNetwork adds the fulfillmentNetwork to the get order metrics params
func (o *GetOrderMetricsParams) WithFulfillmentNetwork(fulfillmentNetwork *string) *GetOrderMetricsParams {
	o.SetFulfillmentNetwork(fulfillmentNetwork)
	return o
}

// SetFulfillmentNetwork adds the fulfillmentNetwork to the get order metrics params
func (o *GetOrderMetricsParams) SetFulfillmentNetwork(fulfillmentNetwork *string) {
	o.FulfillmentNetwork = fulfillmentNetwork
}

// WithGranularity adds the granularity to the get order metrics params
func (o *GetOrderMetricsParams) WithGranularity(granularity string) *GetOrderMetricsParams {
	o.SetGranularity(granularity)
	return o
}

// SetGranularity adds the granularity to the get order metrics params
func (o *GetOrderMetricsParams) SetGranularity(granularity string) {
	o.Granularity = granularity
}

// WithGranularityTimeZone adds the granularityTimeZone to the get order metrics params
func (o *GetOrderMetricsParams) WithGranularityTimeZone(granularityTimeZone *string) *GetOrderMetricsParams {
	o.SetGranularityTimeZone(granularityTimeZone)
	return o
}

// SetGranularityTimeZone adds the granularityTimeZone to the get order metrics params
func (o *GetOrderMetricsParams) SetGranularityTimeZone(granularityTimeZone *string) {
	o.GranularityTimeZone = granularityTimeZone
}

// WithInterval adds the interval to the get order metrics params
func (o *GetOrderMetricsParams) WithInterval(interval string) *GetOrderMetricsParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the get order metrics params
func (o *GetOrderMetricsParams) SetInterval(interval string) {
	o.Interval = interval
}

// WithMarketplaceIds adds the marketplaceIds to the get order metrics params
func (o *GetOrderMetricsParams) WithMarketplaceIds(marketplaceIds []string) *GetOrderMetricsParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get order metrics params
func (o *GetOrderMetricsParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithSku adds the sku to the get order metrics params
func (o *GetOrderMetricsParams) WithSku(sku *string) *GetOrderMetricsParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the get order metrics params
func (o *GetOrderMetricsParams) SetSku(sku *string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Asin != nil {

		// query param asin
		var qrAsin string

		if o.Asin != nil {
			qrAsin = *o.Asin
		}
		qAsin := qrAsin
		if qAsin != "" {

			if err := r.SetQueryParam("asin", qAsin); err != nil {
				return err
			}
		}
	}

	if o.BuyerType != nil {

		// query param buyerType
		var qrBuyerType string

		if o.BuyerType != nil {
			qrBuyerType = *o.BuyerType
		}
		qBuyerType := qrBuyerType
		if qBuyerType != "" {

			if err := r.SetQueryParam("buyerType", qBuyerType); err != nil {
				return err
			}
		}
	}

	if o.FirstDayOfWeek != nil {

		// query param firstDayOfWeek
		var qrFirstDayOfWeek string

		if o.FirstDayOfWeek != nil {
			qrFirstDayOfWeek = *o.FirstDayOfWeek
		}
		qFirstDayOfWeek := qrFirstDayOfWeek
		if qFirstDayOfWeek != "" {

			if err := r.SetQueryParam("firstDayOfWeek", qFirstDayOfWeek); err != nil {
				return err
			}
		}
	}

	if o.FulfillmentNetwork != nil {

		// query param fulfillmentNetwork
		var qrFulfillmentNetwork string

		if o.FulfillmentNetwork != nil {
			qrFulfillmentNetwork = *o.FulfillmentNetwork
		}
		qFulfillmentNetwork := qrFulfillmentNetwork
		if qFulfillmentNetwork != "" {

			if err := r.SetQueryParam("fulfillmentNetwork", qFulfillmentNetwork); err != nil {
				return err
			}
		}
	}

	// query param granularity
	qrGranularity := o.Granularity
	qGranularity := qrGranularity
	if qGranularity != "" {

		if err := r.SetQueryParam("granularity", qGranularity); err != nil {
			return err
		}
	}

	if o.GranularityTimeZone != nil {

		// query param granularityTimeZone
		var qrGranularityTimeZone string

		if o.GranularityTimeZone != nil {
			qrGranularityTimeZone = *o.GranularityTimeZone
		}
		qGranularityTimeZone := qrGranularityTimeZone
		if qGranularityTimeZone != "" {

			if err := r.SetQueryParam("granularityTimeZone", qGranularityTimeZone); err != nil {
				return err
			}
		}
	}

	// query param interval
	qrInterval := o.Interval
	qInterval := qrInterval
	if qInterval != "" {

		if err := r.SetQueryParam("interval", qInterval); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrderMetrics binds the parameter marketplaceIds
func (o *GetOrderMetricsParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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
