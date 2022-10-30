// Code generated by go-swagger; DO NOT EDIT.

package merchant_fulfillment

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

	"github.com/xamandar/amzn-sp-api-go/api/merchantFulfillmentV0/merchant_fulfillment_v0_models"
)

// NewGetAdditionalSellerInputsOldParams creates a new GetAdditionalSellerInputsOldParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAdditionalSellerInputsOldParams() *GetAdditionalSellerInputsOldParams {
	return &GetAdditionalSellerInputsOldParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdditionalSellerInputsOldParamsWithTimeout creates a new GetAdditionalSellerInputsOldParams object
// with the ability to set a timeout on a request.
func NewGetAdditionalSellerInputsOldParamsWithTimeout(timeout time.Duration) *GetAdditionalSellerInputsOldParams {
	return &GetAdditionalSellerInputsOldParams{
		timeout: timeout,
	}
}

// NewGetAdditionalSellerInputsOldParamsWithContext creates a new GetAdditionalSellerInputsOldParams object
// with the ability to set a context for a request.
func NewGetAdditionalSellerInputsOldParamsWithContext(ctx context.Context) *GetAdditionalSellerInputsOldParams {
	return &GetAdditionalSellerInputsOldParams{
		Context: ctx,
	}
}

// NewGetAdditionalSellerInputsOldParamsWithHTTPClient creates a new GetAdditionalSellerInputsOldParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAdditionalSellerInputsOldParamsWithHTTPClient(client *http.Client) *GetAdditionalSellerInputsOldParams {
	return &GetAdditionalSellerInputsOldParams{
		HTTPClient: client,
	}
}

/*
GetAdditionalSellerInputsOldParams contains all the parameters to send to the API endpoint

	for the get additional seller inputs old operation.

	Typically these are written to a http.Request.
*/
type GetAdditionalSellerInputsOldParams struct {

	// Body.
	Body *merchant_fulfillment_v0_models.GetAdditionalSellerInputsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get additional seller inputs old params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdditionalSellerInputsOldParams) WithDefaults() *GetAdditionalSellerInputsOldParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get additional seller inputs old params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdditionalSellerInputsOldParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) WithTimeout(timeout time.Duration) *GetAdditionalSellerInputsOldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) WithContext(ctx context.Context) *GetAdditionalSellerInputsOldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) WithHTTPClient(client *http.Client) *GetAdditionalSellerInputsOldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) WithBody(body *merchant_fulfillment_v0_models.GetAdditionalSellerInputsRequest) *GetAdditionalSellerInputsOldParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get additional seller inputs old params
func (o *GetAdditionalSellerInputsOldParams) SetBody(body *merchant_fulfillment_v0_models.GetAdditionalSellerInputsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdditionalSellerInputsOldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
