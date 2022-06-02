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

// NewListReturnReasonCodesParams creates a new ListReturnReasonCodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListReturnReasonCodesParams() *ListReturnReasonCodesParams {
	return &ListReturnReasonCodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListReturnReasonCodesParamsWithTimeout creates a new ListReturnReasonCodesParams object
// with the ability to set a timeout on a request.
func NewListReturnReasonCodesParamsWithTimeout(timeout time.Duration) *ListReturnReasonCodesParams {
	return &ListReturnReasonCodesParams{
		timeout: timeout,
	}
}

// NewListReturnReasonCodesParamsWithContext creates a new ListReturnReasonCodesParams object
// with the ability to set a context for a request.
func NewListReturnReasonCodesParamsWithContext(ctx context.Context) *ListReturnReasonCodesParams {
	return &ListReturnReasonCodesParams{
		Context: ctx,
	}
}

// NewListReturnReasonCodesParamsWithHTTPClient creates a new ListReturnReasonCodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListReturnReasonCodesParamsWithHTTPClient(client *http.Client) *ListReturnReasonCodesParams {
	return &ListReturnReasonCodesParams{
		HTTPClient: client,
	}
}

/* ListReturnReasonCodesParams contains all the parameters to send to the API endpoint
   for the list return reason codes operation.

   Typically these are written to a http.Request.
*/
type ListReturnReasonCodesParams struct {

	/* Language.

	   The language that the TranslatedDescription property of the ReasonCodeDetails response object should be translated into.
	*/
	Language string

	/* MarketplaceID.

	   The marketplace for which the seller wants return reason codes.
	*/
	MarketplaceID *string

	/* SellerFulfillmentOrderID.

	   The identifier assigned to the item by the seller when the fulfillment order was created. The service uses this value to determine the marketplace for which the seller wants return reason codes.
	*/
	SellerFulfillmentOrderID *string

	/* SellerSku.

	   The seller SKU for which return reason codes are required.
	*/
	SellerSku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list return reason codes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListReturnReasonCodesParams) WithDefaults() *ListReturnReasonCodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list return reason codes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListReturnReasonCodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list return reason codes params
func (o *ListReturnReasonCodesParams) WithTimeout(timeout time.Duration) *ListReturnReasonCodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list return reason codes params
func (o *ListReturnReasonCodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list return reason codes params
func (o *ListReturnReasonCodesParams) WithContext(ctx context.Context) *ListReturnReasonCodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list return reason codes params
func (o *ListReturnReasonCodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list return reason codes params
func (o *ListReturnReasonCodesParams) WithHTTPClient(client *http.Client) *ListReturnReasonCodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list return reason codes params
func (o *ListReturnReasonCodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the list return reason codes params
func (o *ListReturnReasonCodesParams) WithLanguage(language string) *ListReturnReasonCodesParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the list return reason codes params
func (o *ListReturnReasonCodesParams) SetLanguage(language string) {
	o.Language = language
}

// WithMarketplaceID adds the marketplaceID to the list return reason codes params
func (o *ListReturnReasonCodesParams) WithMarketplaceID(marketplaceID *string) *ListReturnReasonCodesParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the list return reason codes params
func (o *ListReturnReasonCodesParams) SetMarketplaceID(marketplaceID *string) {
	o.MarketplaceID = marketplaceID
}

// WithSellerFulfillmentOrderID adds the sellerFulfillmentOrderID to the list return reason codes params
func (o *ListReturnReasonCodesParams) WithSellerFulfillmentOrderID(sellerFulfillmentOrderID *string) *ListReturnReasonCodesParams {
	o.SetSellerFulfillmentOrderID(sellerFulfillmentOrderID)
	return o
}

// SetSellerFulfillmentOrderID adds the sellerFulfillmentOrderId to the list return reason codes params
func (o *ListReturnReasonCodesParams) SetSellerFulfillmentOrderID(sellerFulfillmentOrderID *string) {
	o.SellerFulfillmentOrderID = sellerFulfillmentOrderID
}

// WithSellerSku adds the sellerSku to the list return reason codes params
func (o *ListReturnReasonCodesParams) WithSellerSku(sellerSku string) *ListReturnReasonCodesParams {
	o.SetSellerSku(sellerSku)
	return o
}

// SetSellerSku adds the sellerSku to the list return reason codes params
func (o *ListReturnReasonCodesParams) SetSellerSku(sellerSku string) {
	o.SellerSku = sellerSku
}

// WriteToRequest writes these params to a swagger request
func (o *ListReturnReasonCodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param language
	qrLanguage := o.Language
	qLanguage := qrLanguage
	if qLanguage != "" {

		if err := r.SetQueryParam("language", qLanguage); err != nil {
			return err
		}
	}

	if o.MarketplaceID != nil {

		// query param marketplaceId
		var qrMarketplaceID string

		if o.MarketplaceID != nil {
			qrMarketplaceID = *o.MarketplaceID
		}
		qMarketplaceID := qrMarketplaceID
		if qMarketplaceID != "" {

			if err := r.SetQueryParam("marketplaceId", qMarketplaceID); err != nil {
				return err
			}
		}
	}

	if o.SellerFulfillmentOrderID != nil {

		// query param sellerFulfillmentOrderId
		var qrSellerFulfillmentOrderID string

		if o.SellerFulfillmentOrderID != nil {
			qrSellerFulfillmentOrderID = *o.SellerFulfillmentOrderID
		}
		qSellerFulfillmentOrderID := qrSellerFulfillmentOrderID
		if qSellerFulfillmentOrderID != "" {

			if err := r.SetQueryParam("sellerFulfillmentOrderId", qSellerFulfillmentOrderID); err != nil {
				return err
			}
		}
	}

	// query param sellerSku
	qrSellerSku := o.SellerSku
	qSellerSku := qrSellerSku
	if qSellerSku != "" {

		if err := r.SetQueryParam("sellerSku", qSellerSku); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
