// Code generated by go-swagger; DO NOT EDIT.

package product_pricing

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

	"github.com/xamandar/amzn-sp-api-go/api/productPricing_2022-05-01/product_pricing_2022_05_01_models"
)

// NewGetFeaturedOfferExpectedPriceBatchParams creates a new GetFeaturedOfferExpectedPriceBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFeaturedOfferExpectedPriceBatchParams() *GetFeaturedOfferExpectedPriceBatchParams {
	return &GetFeaturedOfferExpectedPriceBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeaturedOfferExpectedPriceBatchParamsWithTimeout creates a new GetFeaturedOfferExpectedPriceBatchParams object
// with the ability to set a timeout on a request.
func NewGetFeaturedOfferExpectedPriceBatchParamsWithTimeout(timeout time.Duration) *GetFeaturedOfferExpectedPriceBatchParams {
	return &GetFeaturedOfferExpectedPriceBatchParams{
		timeout: timeout,
	}
}

// NewGetFeaturedOfferExpectedPriceBatchParamsWithContext creates a new GetFeaturedOfferExpectedPriceBatchParams object
// with the ability to set a context for a request.
func NewGetFeaturedOfferExpectedPriceBatchParamsWithContext(ctx context.Context) *GetFeaturedOfferExpectedPriceBatchParams {
	return &GetFeaturedOfferExpectedPriceBatchParams{
		Context: ctx,
	}
}

// NewGetFeaturedOfferExpectedPriceBatchParamsWithHTTPClient creates a new GetFeaturedOfferExpectedPriceBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFeaturedOfferExpectedPriceBatchParamsWithHTTPClient(client *http.Client) *GetFeaturedOfferExpectedPriceBatchParams {
	return &GetFeaturedOfferExpectedPriceBatchParams{
		HTTPClient: client,
	}
}

/*
GetFeaturedOfferExpectedPriceBatchParams contains all the parameters to send to the API endpoint

	for the get featured offer expected price batch operation.

	Typically these are written to a http.Request.
*/
type GetFeaturedOfferExpectedPriceBatchParams struct {

	// GetFeaturedOfferExpectedPriceBatchRequestBody.
	GetFeaturedOfferExpectedPriceBatchRequestBody *product_pricing_2022_05_01_models.GetFeaturedOfferExpectedPriceBatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get featured offer expected price batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeaturedOfferExpectedPriceBatchParams) WithDefaults() *GetFeaturedOfferExpectedPriceBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get featured offer expected price batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFeaturedOfferExpectedPriceBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) WithTimeout(timeout time.Duration) *GetFeaturedOfferExpectedPriceBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) WithContext(ctx context.Context) *GetFeaturedOfferExpectedPriceBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) WithHTTPClient(client *http.Client) *GetFeaturedOfferExpectedPriceBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetFeaturedOfferExpectedPriceBatchRequestBody adds the getFeaturedOfferExpectedPriceBatchRequestBody to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) WithGetFeaturedOfferExpectedPriceBatchRequestBody(getFeaturedOfferExpectedPriceBatchRequestBody *product_pricing_2022_05_01_models.GetFeaturedOfferExpectedPriceBatchRequest) *GetFeaturedOfferExpectedPriceBatchParams {
	o.SetGetFeaturedOfferExpectedPriceBatchRequestBody(getFeaturedOfferExpectedPriceBatchRequestBody)
	return o
}

// SetGetFeaturedOfferExpectedPriceBatchRequestBody adds the getFeaturedOfferExpectedPriceBatchRequestBody to the get featured offer expected price batch params
func (o *GetFeaturedOfferExpectedPriceBatchParams) SetGetFeaturedOfferExpectedPriceBatchRequestBody(getFeaturedOfferExpectedPriceBatchRequestBody *product_pricing_2022_05_01_models.GetFeaturedOfferExpectedPriceBatchRequest) {
	o.GetFeaturedOfferExpectedPriceBatchRequestBody = getFeaturedOfferExpectedPriceBatchRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeaturedOfferExpectedPriceBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GetFeaturedOfferExpectedPriceBatchRequestBody != nil {
		if err := r.SetBodyParam(o.GetFeaturedOfferExpectedPriceBatchRequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
