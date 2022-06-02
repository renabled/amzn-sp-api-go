// Code generated by go-swagger; DO NOT EDIT.

package messaging

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

// NewGetAttributesParams creates a new GetAttributesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAttributesParams() *GetAttributesParams {
	return &GetAttributesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAttributesParamsWithTimeout creates a new GetAttributesParams object
// with the ability to set a timeout on a request.
func NewGetAttributesParamsWithTimeout(timeout time.Duration) *GetAttributesParams {
	return &GetAttributesParams{
		timeout: timeout,
	}
}

// NewGetAttributesParamsWithContext creates a new GetAttributesParams object
// with the ability to set a context for a request.
func NewGetAttributesParamsWithContext(ctx context.Context) *GetAttributesParams {
	return &GetAttributesParams{
		Context: ctx,
	}
}

// NewGetAttributesParamsWithHTTPClient creates a new GetAttributesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAttributesParamsWithHTTPClient(client *http.Client) *GetAttributesParams {
	return &GetAttributesParams{
		HTTPClient: client,
	}
}

/* GetAttributesParams contains all the parameters to send to the API endpoint
   for the get attributes operation.

   Typically these are written to a http.Request.
*/
type GetAttributesParams struct {

	/* AmazonOrderID.

	   An Amazon order identifier. This specifies the order for which a message is sent.
	*/
	AmazonOrderID string

	/* MarketplaceIds.

	   A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	*/
	MarketplaceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAttributesParams) WithDefaults() *GetAttributesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAttributesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get attributes params
func (o *GetAttributesParams) WithTimeout(timeout time.Duration) *GetAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get attributes params
func (o *GetAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get attributes params
func (o *GetAttributesParams) WithContext(ctx context.Context) *GetAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get attributes params
func (o *GetAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get attributes params
func (o *GetAttributesParams) WithHTTPClient(client *http.Client) *GetAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get attributes params
func (o *GetAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmazonOrderID adds the amazonOrderID to the get attributes params
func (o *GetAttributesParams) WithAmazonOrderID(amazonOrderID string) *GetAttributesParams {
	o.SetAmazonOrderID(amazonOrderID)
	return o
}

// SetAmazonOrderID adds the amazonOrderId to the get attributes params
func (o *GetAttributesParams) SetAmazonOrderID(amazonOrderID string) {
	o.AmazonOrderID = amazonOrderID
}

// WithMarketplaceIds adds the marketplaceIds to the get attributes params
func (o *GetAttributesParams) WithMarketplaceIds(marketplaceIds []string) *GetAttributesParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get attributes params
func (o *GetAttributesParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param amazonOrderId
	if err := r.SetPathParam("amazonOrderId", o.AmazonOrderID); err != nil {
		return err
	}

	if o.MarketplaceIds != nil {

		// binding items for marketplaceIds
		joinedMarketplaceIds := o.bindParamMarketplaceIds(reg)

		// query array param marketplaceIds
		if err := r.SetQueryParam("marketplaceIds", joinedMarketplaceIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAttributes binds the parameter marketplaceIds
func (o *GetAttributesParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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
