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

	"github.com/renabled/amzn-sp-api-go/api/productPricingV0/product_pricing_v0_models"
)

// NewGetListingOffersBatchParams creates a new GetListingOffersBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetListingOffersBatchParams() *GetListingOffersBatchParams {
	return &GetListingOffersBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetListingOffersBatchParamsWithTimeout creates a new GetListingOffersBatchParams object
// with the ability to set a timeout on a request.
func NewGetListingOffersBatchParamsWithTimeout(timeout time.Duration) *GetListingOffersBatchParams {
	return &GetListingOffersBatchParams{
		timeout: timeout,
	}
}

// NewGetListingOffersBatchParamsWithContext creates a new GetListingOffersBatchParams object
// with the ability to set a context for a request.
func NewGetListingOffersBatchParamsWithContext(ctx context.Context) *GetListingOffersBatchParams {
	return &GetListingOffersBatchParams{
		Context: ctx,
	}
}

// NewGetListingOffersBatchParamsWithHTTPClient creates a new GetListingOffersBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetListingOffersBatchParamsWithHTTPClient(client *http.Client) *GetListingOffersBatchParams {
	return &GetListingOffersBatchParams{
		HTTPClient: client,
	}
}

/*
GetListingOffersBatchParams contains all the parameters to send to the API endpoint

	for the get listing offers batch operation.

	Typically these are written to a http.Request.
*/
type GetListingOffersBatchParams struct {

	// GetListingOffersBatchRequestBody.
	GetListingOffersBatchRequestBody *product_pricing_v0_models.GetListingOffersBatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get listing offers batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListingOffersBatchParams) WithDefaults() *GetListingOffersBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get listing offers batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetListingOffersBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get listing offers batch params
func (o *GetListingOffersBatchParams) WithTimeout(timeout time.Duration) *GetListingOffersBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get listing offers batch params
func (o *GetListingOffersBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get listing offers batch params
func (o *GetListingOffersBatchParams) WithContext(ctx context.Context) *GetListingOffersBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get listing offers batch params
func (o *GetListingOffersBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get listing offers batch params
func (o *GetListingOffersBatchParams) WithHTTPClient(client *http.Client) *GetListingOffersBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get listing offers batch params
func (o *GetListingOffersBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGetListingOffersBatchRequestBody adds the getListingOffersBatchRequestBody to the get listing offers batch params
func (o *GetListingOffersBatchParams) WithGetListingOffersBatchRequestBody(getListingOffersBatchRequestBody *product_pricing_v0_models.GetListingOffersBatchRequest) *GetListingOffersBatchParams {
	o.SetGetListingOffersBatchRequestBody(getListingOffersBatchRequestBody)
	return o
}

// SetGetListingOffersBatchRequestBody adds the getListingOffersBatchRequestBody to the get listing offers batch params
func (o *GetListingOffersBatchParams) SetGetListingOffersBatchRequestBody(getListingOffersBatchRequestBody *product_pricing_v0_models.GetListingOffersBatchRequest) {
	o.GetListingOffersBatchRequestBody = getListingOffersBatchRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetListingOffersBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GetListingOffersBatchRequestBody != nil {
		if err := r.SetBodyParam(o.GetListingOffersBatchRequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
