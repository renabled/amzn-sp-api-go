// Code generated by go-swagger; DO NOT EDIT.

package listings

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

// NewGetListingsItemParams creates a new GetListingsItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetListingsItemParams() *GetListingsItemParams {
	return &GetListingsItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetListingsItemParamsWithTimeout creates a new GetListingsItemParams object
// with the ability to set a timeout on a request.
func NewGetListingsItemParamsWithTimeout(timeout time.Duration) *GetListingsItemParams {
	return &GetListingsItemParams{
		timeout: timeout,
	}
}

// NewGetListingsItemParamsWithContext creates a new GetListingsItemParams object
// with the ability to set a context for a request.
func NewGetListingsItemParamsWithContext(ctx context.Context) *GetListingsItemParams {
	return &GetListingsItemParams{
		Context: ctx,
	}
}

// NewGetListingsItemParamsWithHTTPClient creates a new GetListingsItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetListingsItemParamsWithHTTPClient(client *http.Client) *GetListingsItemParams {
	return &GetListingsItemParams{
		HTTPClient: client,
	}
}

/*
GetListingsItemParams contains all the parameters to send to the API endpoint

	for the get listings item operation.

	Typically these are written to a http.Request.
*/
type GetListingsItemParams struct {

	/* IncludedData.

	   A comma-delimited list of data sets to include in the response. Default: `summaries`.

	   Default: ["summaries"]
	*/
	IncludedData []string

	/* IssueLocale.

	   A locale for localization of issues. When not provided, the default language code of the first marketplace is used. Examples: `en_US`, `fr_CA`, `fr_FR`. Localized messages default to `en_US` when a localization is not available in the specified locale.
	*/
	IssueLocale *string

	/* MarketplaceIds.

	   A comma-delimited list of Amazon marketplace identifiers for the request.
	*/
	MarketplaceIds []string

	/* SellerID.

	   A selling partner identifier, such as a merchant account or vendor code.
	*/
	SellerID string

	/* Sku.

	   A selling partner provided identifier for an Amazon listing.
	*/
	Sku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get listings item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListingsItemParams) WithDefaults() *GetListingsItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get listings item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListingsItemParams) SetDefaults() {
	var (
		includedDataDefault = []string{"summaries"}
	)

	val := GetListingsItemParams{
		IncludedData: includedDataDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get listings item params
func (o *GetListingsItemParams) WithTimeout(timeout time.Duration) *GetListingsItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get listings item params
func (o *GetListingsItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get listings item params
func (o *GetListingsItemParams) WithContext(ctx context.Context) *GetListingsItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get listings item params
func (o *GetListingsItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get listings item params
func (o *GetListingsItemParams) WithHTTPClient(client *http.Client) *GetListingsItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get listings item params
func (o *GetListingsItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludedData adds the includedData to the get listings item params
func (o *GetListingsItemParams) WithIncludedData(includedData []string) *GetListingsItemParams {
	o.SetIncludedData(includedData)
	return o
}

// SetIncludedData adds the includedData to the get listings item params
func (o *GetListingsItemParams) SetIncludedData(includedData []string) {
	o.IncludedData = includedData
}

// WithIssueLocale adds the issueLocale to the get listings item params
func (o *GetListingsItemParams) WithIssueLocale(issueLocale *string) *GetListingsItemParams {
	o.SetIssueLocale(issueLocale)
	return o
}

// SetIssueLocale adds the issueLocale to the get listings item params
func (o *GetListingsItemParams) SetIssueLocale(issueLocale *string) {
	o.IssueLocale = issueLocale
}

// WithMarketplaceIds adds the marketplaceIds to the get listings item params
func (o *GetListingsItemParams) WithMarketplaceIds(marketplaceIds []string) *GetListingsItemParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get listings item params
func (o *GetListingsItemParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithSellerID adds the sellerID to the get listings item params
func (o *GetListingsItemParams) WithSellerID(sellerID string) *GetListingsItemParams {
	o.SetSellerID(sellerID)
	return o
}

// SetSellerID adds the sellerId to the get listings item params
func (o *GetListingsItemParams) SetSellerID(sellerID string) {
	o.SellerID = sellerID
}

// WithSku adds the sku to the get listings item params
func (o *GetListingsItemParams) WithSku(sku string) *GetListingsItemParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the get listings item params
func (o *GetListingsItemParams) SetSku(sku string) {
	o.Sku = sku
}

// WriteToRequest writes these params to a swagger request
func (o *GetListingsItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludedData != nil {

		// binding items for includedData
		joinedIncludedData := o.bindParamIncludedData(reg)

		// query array param includedData
		if err := r.SetQueryParam("includedData", joinedIncludedData...); err != nil {
			return err
		}
	}

	if o.IssueLocale != nil {

		// query param issueLocale
		var qrIssueLocale string

		if o.IssueLocale != nil {
			qrIssueLocale = *o.IssueLocale
		}
		qIssueLocale := qrIssueLocale
		if qIssueLocale != "" {

			if err := r.SetQueryParam("issueLocale", qIssueLocale); err != nil {
				return err
			}
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

	// path param sellerId
	if err := r.SetPathParam("sellerId", o.SellerID); err != nil {
		return err
	}

	// path param sku
	if err := r.SetPathParam("sku", o.Sku); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetListingsItem binds the parameter includedData
func (o *GetListingsItemParams) bindParamIncludedData(formats strfmt.Registry) []string {
	includedDataIR := o.IncludedData

	var includedDataIC []string
	for _, includedDataIIR := range includedDataIR { // explode []string

		includedDataIIV := includedDataIIR // string as string
		includedDataIC = append(includedDataIC, includedDataIIV)
	}

	// items.CollectionFormat: "csv"
	includedDataIS := swag.JoinByFormat(includedDataIC, "csv")

	return includedDataIS
}

// bindParamGetListingsItem binds the parameter marketplaceIds
func (o *GetListingsItemParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
	marketplaceIdsIR := o.MarketplaceIds

	var marketplaceIdsIC []string
	for _, marketplaceIdsIIR := range marketplaceIdsIR { // explode []string

		marketplaceIdsIIV := marketplaceIdsIIR // string as string
		marketplaceIdsIC = append(marketplaceIdsIC, marketplaceIdsIIV)
	}

	// items.CollectionFormat: "csv"
	marketplaceIdsIS := swag.JoinByFormat(marketplaceIdsIC, "csv")

	return marketplaceIdsIS
}
