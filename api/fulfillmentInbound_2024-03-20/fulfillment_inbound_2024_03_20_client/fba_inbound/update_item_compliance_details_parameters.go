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

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInbound_2024-03-20/fulfillment_inbound_2024_03_20_models"
)

// NewUpdateItemComplianceDetailsParams creates a new UpdateItemComplianceDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateItemComplianceDetailsParams() *UpdateItemComplianceDetailsParams {
	return &UpdateItemComplianceDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateItemComplianceDetailsParamsWithTimeout creates a new UpdateItemComplianceDetailsParams object
// with the ability to set a timeout on a request.
func NewUpdateItemComplianceDetailsParamsWithTimeout(timeout time.Duration) *UpdateItemComplianceDetailsParams {
	return &UpdateItemComplianceDetailsParams{
		timeout: timeout,
	}
}

// NewUpdateItemComplianceDetailsParamsWithContext creates a new UpdateItemComplianceDetailsParams object
// with the ability to set a context for a request.
func NewUpdateItemComplianceDetailsParamsWithContext(ctx context.Context) *UpdateItemComplianceDetailsParams {
	return &UpdateItemComplianceDetailsParams{
		Context: ctx,
	}
}

// NewUpdateItemComplianceDetailsParamsWithHTTPClient creates a new UpdateItemComplianceDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateItemComplianceDetailsParamsWithHTTPClient(client *http.Client) *UpdateItemComplianceDetailsParams {
	return &UpdateItemComplianceDetailsParams{
		HTTPClient: client,
	}
}

/*
UpdateItemComplianceDetailsParams contains all the parameters to send to the API endpoint

	for the update item compliance details operation.

	Typically these are written to a http.Request.
*/
type UpdateItemComplianceDetailsParams struct {

	/* Body.

	   The body of the request to `updateItemComplianceDetails`.
	*/
	Body *fulfillment_inbound_2024_03_20_models.UpdateItemComplianceDetailsRequest

	/* MarketplaceID.

	   The Marketplace ID. For a list of possible values, refer to [Marketplace IDs](https://developer-docs.amazon.com/sp-api/docs/marketplace-ids).
	*/
	MarketplaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update item compliance details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateItemComplianceDetailsParams) WithDefaults() *UpdateItemComplianceDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update item compliance details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateItemComplianceDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) WithTimeout(timeout time.Duration) *UpdateItemComplianceDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) WithContext(ctx context.Context) *UpdateItemComplianceDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) WithHTTPClient(client *http.Client) *UpdateItemComplianceDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) WithBody(body *fulfillment_inbound_2024_03_20_models.UpdateItemComplianceDetailsRequest) *UpdateItemComplianceDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) SetBody(body *fulfillment_inbound_2024_03_20_models.UpdateItemComplianceDetailsRequest) {
	o.Body = body
}

// WithMarketplaceID adds the marketplaceID to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) WithMarketplaceID(marketplaceID string) *UpdateItemComplianceDetailsParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the update item compliance details params
func (o *UpdateItemComplianceDetailsParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateItemComplianceDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
