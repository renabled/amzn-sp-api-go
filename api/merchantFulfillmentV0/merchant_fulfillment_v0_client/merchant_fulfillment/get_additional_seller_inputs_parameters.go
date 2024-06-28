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

	"github.com/renabled/amzn-sp-api-go/api/merchantFulfillmentV0/merchant_fulfillment_v0_models"
)

// NewGetAdditionalSellerInputsParams creates a new GetAdditionalSellerInputsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAdditionalSellerInputsParams() *GetAdditionalSellerInputsParams {
	return &GetAdditionalSellerInputsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdditionalSellerInputsParamsWithTimeout creates a new GetAdditionalSellerInputsParams object
// with the ability to set a timeout on a request.
func NewGetAdditionalSellerInputsParamsWithTimeout(timeout time.Duration) *GetAdditionalSellerInputsParams {
	return &GetAdditionalSellerInputsParams{
		timeout: timeout,
	}
}

// NewGetAdditionalSellerInputsParamsWithContext creates a new GetAdditionalSellerInputsParams object
// with the ability to set a context for a request.
func NewGetAdditionalSellerInputsParamsWithContext(ctx context.Context) *GetAdditionalSellerInputsParams {
	return &GetAdditionalSellerInputsParams{
		Context: ctx,
	}
}

// NewGetAdditionalSellerInputsParamsWithHTTPClient creates a new GetAdditionalSellerInputsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAdditionalSellerInputsParamsWithHTTPClient(client *http.Client) *GetAdditionalSellerInputsParams {
	return &GetAdditionalSellerInputsParams{
		HTTPClient: client,
	}
}

/*
GetAdditionalSellerInputsParams contains all the parameters to send to the API endpoint

	for the get additional seller inputs operation.

	Typically these are written to a http.Request.
*/
type GetAdditionalSellerInputsParams struct {

	/* Body.

	   Request schema for the `GetAdditionalSellerInputs` operation.
	*/
	Body *merchant_fulfillment_v0_models.GetAdditionalSellerInputsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get additional seller inputs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdditionalSellerInputsParams) WithDefaults() *GetAdditionalSellerInputsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get additional seller inputs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAdditionalSellerInputsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) WithTimeout(timeout time.Duration) *GetAdditionalSellerInputsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) WithContext(ctx context.Context) *GetAdditionalSellerInputsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) WithHTTPClient(client *http.Client) *GetAdditionalSellerInputsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) WithBody(body *merchant_fulfillment_v0_models.GetAdditionalSellerInputsRequest) *GetAdditionalSellerInputsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get additional seller inputs params
func (o *GetAdditionalSellerInputsParams) SetBody(body *merchant_fulfillment_v0_models.GetAdditionalSellerInputsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdditionalSellerInputsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
