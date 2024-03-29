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

// NewGetMessagingActionsForOrderParams creates a new GetMessagingActionsForOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMessagingActionsForOrderParams() *GetMessagingActionsForOrderParams {
	return &GetMessagingActionsForOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMessagingActionsForOrderParamsWithTimeout creates a new GetMessagingActionsForOrderParams object
// with the ability to set a timeout on a request.
func NewGetMessagingActionsForOrderParamsWithTimeout(timeout time.Duration) *GetMessagingActionsForOrderParams {
	return &GetMessagingActionsForOrderParams{
		timeout: timeout,
	}
}

// NewGetMessagingActionsForOrderParamsWithContext creates a new GetMessagingActionsForOrderParams object
// with the ability to set a context for a request.
func NewGetMessagingActionsForOrderParamsWithContext(ctx context.Context) *GetMessagingActionsForOrderParams {
	return &GetMessagingActionsForOrderParams{
		Context: ctx,
	}
}

// NewGetMessagingActionsForOrderParamsWithHTTPClient creates a new GetMessagingActionsForOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMessagingActionsForOrderParamsWithHTTPClient(client *http.Client) *GetMessagingActionsForOrderParams {
	return &GetMessagingActionsForOrderParams{
		HTTPClient: client,
	}
}

/*
GetMessagingActionsForOrderParams contains all the parameters to send to the API endpoint

	for the get messaging actions for order operation.

	Typically these are written to a http.Request.
*/
type GetMessagingActionsForOrderParams struct {

	/* AmazonOrderID.

	   An Amazon order identifier. This specifies the order for which you want a list of available message types.
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

// WithDefaults hydrates default values in the get messaging actions for order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessagingActionsForOrderParams) WithDefaults() *GetMessagingActionsForOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get messaging actions for order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMessagingActionsForOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) WithTimeout(timeout time.Duration) *GetMessagingActionsForOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) WithContext(ctx context.Context) *GetMessagingActionsForOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) WithHTTPClient(client *http.Client) *GetMessagingActionsForOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmazonOrderID adds the amazonOrderID to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) WithAmazonOrderID(amazonOrderID string) *GetMessagingActionsForOrderParams {
	o.SetAmazonOrderID(amazonOrderID)
	return o
}

// SetAmazonOrderID adds the amazonOrderId to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) SetAmazonOrderID(amazonOrderID string) {
	o.AmazonOrderID = amazonOrderID
}

// WithMarketplaceIds adds the marketplaceIds to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) WithMarketplaceIds(marketplaceIds []string) *GetMessagingActionsForOrderParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get messaging actions for order params
func (o *GetMessagingActionsForOrderParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetMessagingActionsForOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamGetMessagingActionsForOrder binds the parameter marketplaceIds
func (o *GetMessagingActionsForOrderParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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
