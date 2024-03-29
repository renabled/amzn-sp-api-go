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

// NewGetInboundGuidanceParams creates a new GetInboundGuidanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInboundGuidanceParams() *GetInboundGuidanceParams {
	return &GetInboundGuidanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInboundGuidanceParamsWithTimeout creates a new GetInboundGuidanceParams object
// with the ability to set a timeout on a request.
func NewGetInboundGuidanceParamsWithTimeout(timeout time.Duration) *GetInboundGuidanceParams {
	return &GetInboundGuidanceParams{
		timeout: timeout,
	}
}

// NewGetInboundGuidanceParamsWithContext creates a new GetInboundGuidanceParams object
// with the ability to set a context for a request.
func NewGetInboundGuidanceParamsWithContext(ctx context.Context) *GetInboundGuidanceParams {
	return &GetInboundGuidanceParams{
		Context: ctx,
	}
}

// NewGetInboundGuidanceParamsWithHTTPClient creates a new GetInboundGuidanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInboundGuidanceParamsWithHTTPClient(client *http.Client) *GetInboundGuidanceParams {
	return &GetInboundGuidanceParams{
		HTTPClient: client,
	}
}

/*
GetInboundGuidanceParams contains all the parameters to send to the API endpoint

	for the get inbound guidance operation.

	Typically these are written to a http.Request.
*/
type GetInboundGuidanceParams struct {

	/* ASINList.

	   A list of ASIN values. Used to identify items for which you want inbound guidance for shipment to Amazon's fulfillment network. Note: If you specify a ASIN that identifies a variation parent ASIN, this operation returns an error. A variation parent ASIN represents a generic product that cannot be sold. Variation child ASINs represent products that have specific characteristics (such as size and color) and can be sold.
	*/
	ASINList []string

	/* MarketplaceID.

	   A marketplace identifier. Specifies the marketplace where the product would be stored.
	*/
	MarketplaceID string

	/* SellerSKUList.

	   A list of SellerSKU values. Used to identify items for which you want inbound guidance for shipment to Amazon's fulfillment network. Note: SellerSKU is qualified by the SellerId, which is included with every Selling Partner API operation that you submit. If you specify a SellerSKU that identifies a variation parent ASIN, this operation returns an error. A variation parent ASIN represents a generic product that cannot be sold. Variation child ASINs represent products that have specific characteristics (such as size and color) and can be sold.
	*/
	SellerSKUList []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inbound guidance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInboundGuidanceParams) WithDefaults() *GetInboundGuidanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inbound guidance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInboundGuidanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inbound guidance params
func (o *GetInboundGuidanceParams) WithTimeout(timeout time.Duration) *GetInboundGuidanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inbound guidance params
func (o *GetInboundGuidanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inbound guidance params
func (o *GetInboundGuidanceParams) WithContext(ctx context.Context) *GetInboundGuidanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inbound guidance params
func (o *GetInboundGuidanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inbound guidance params
func (o *GetInboundGuidanceParams) WithHTTPClient(client *http.Client) *GetInboundGuidanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inbound guidance params
func (o *GetInboundGuidanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithASINList adds the aSINList to the get inbound guidance params
func (o *GetInboundGuidanceParams) WithASINList(aSINList []string) *GetInboundGuidanceParams {
	o.SetASINList(aSINList)
	return o
}

// SetASINList adds the aSINList to the get inbound guidance params
func (o *GetInboundGuidanceParams) SetASINList(aSINList []string) {
	o.ASINList = aSINList
}

// WithMarketplaceID adds the marketplaceID to the get inbound guidance params
func (o *GetInboundGuidanceParams) WithMarketplaceID(marketplaceID string) *GetInboundGuidanceParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the get inbound guidance params
func (o *GetInboundGuidanceParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WithSellerSKUList adds the sellerSKUList to the get inbound guidance params
func (o *GetInboundGuidanceParams) WithSellerSKUList(sellerSKUList []string) *GetInboundGuidanceParams {
	o.SetSellerSKUList(sellerSKUList)
	return o
}

// SetSellerSKUList adds the sellerSKUList to the get inbound guidance params
func (o *GetInboundGuidanceParams) SetSellerSKUList(sellerSKUList []string) {
	o.SellerSKUList = sellerSKUList
}

// WriteToRequest writes these params to a swagger request
func (o *GetInboundGuidanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ASINList != nil {

		// binding items for ASINList
		joinedASINList := o.bindParamASINList(reg)

		// query array param ASINList
		if err := r.SetQueryParam("ASINList", joinedASINList...); err != nil {
			return err
		}
	}

	// query param MarketplaceId
	qrMarketplaceID := o.MarketplaceID
	qMarketplaceID := qrMarketplaceID
	if qMarketplaceID != "" {

		if err := r.SetQueryParam("MarketplaceId", qMarketplaceID); err != nil {
			return err
		}
	}

	if o.SellerSKUList != nil {

		// binding items for SellerSKUList
		joinedSellerSKUList := o.bindParamSellerSKUList(reg)

		// query array param SellerSKUList
		if err := r.SetQueryParam("SellerSKUList", joinedSellerSKUList...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetInboundGuidance binds the parameter ASINList
func (o *GetInboundGuidanceParams) bindParamASINList(formats strfmt.Registry) []string {
	aSINListIR := o.ASINList

	var aSINListIC []string
	for _, aSINListIIR := range aSINListIR { // explode []string

		aSINListIIV := aSINListIIR // string as string
		aSINListIC = append(aSINListIC, aSINListIIV)
	}

	// items.CollectionFormat: ""
	aSINListIS := swag.JoinByFormat(aSINListIC, "")

	return aSINListIS
}

// bindParamGetInboundGuidance binds the parameter SellerSKUList
func (o *GetInboundGuidanceParams) bindParamSellerSKUList(formats strfmt.Registry) []string {
	sellerSKUListIR := o.SellerSKUList

	var sellerSKUListIC []string
	for _, sellerSKUListIIR := range sellerSKUListIR { // explode []string

		sellerSKUListIIV := sellerSKUListIIR // string as string
		sellerSKUListIC = append(sellerSKUListIC, sellerSKUListIIV)
	}

	// items.CollectionFormat: ""
	sellerSKUListIS := swag.JoinByFormat(sellerSKUListIC, "")

	return sellerSKUListIS
}
