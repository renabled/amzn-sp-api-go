// Code generated by go-swagger; DO NOT EDIT.

package fees

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

	"github.com/xamandar/amzn-sp-api-go/api/productFeesV0/product_fees_v0_models"
)

// NewGetMyFeesEstimatesParams creates a new GetMyFeesEstimatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMyFeesEstimatesParams() *GetMyFeesEstimatesParams {
	return &GetMyFeesEstimatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMyFeesEstimatesParamsWithTimeout creates a new GetMyFeesEstimatesParams object
// with the ability to set a timeout on a request.
func NewGetMyFeesEstimatesParamsWithTimeout(timeout time.Duration) *GetMyFeesEstimatesParams {
	return &GetMyFeesEstimatesParams{
		timeout: timeout,
	}
}

// NewGetMyFeesEstimatesParamsWithContext creates a new GetMyFeesEstimatesParams object
// with the ability to set a context for a request.
func NewGetMyFeesEstimatesParamsWithContext(ctx context.Context) *GetMyFeesEstimatesParams {
	return &GetMyFeesEstimatesParams{
		Context: ctx,
	}
}

// NewGetMyFeesEstimatesParamsWithHTTPClient creates a new GetMyFeesEstimatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMyFeesEstimatesParamsWithHTTPClient(client *http.Client) *GetMyFeesEstimatesParams {
	return &GetMyFeesEstimatesParams{
		HTTPClient: client,
	}
}

/* GetMyFeesEstimatesParams contains all the parameters to send to the API endpoint
   for the get my fees estimates operation.

   Typically these are written to a http.Request.
*/
type GetMyFeesEstimatesParams struct {

	// Body.
	Body product_fees_v0_models.GetMyFeesEstimatesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get my fees estimates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMyFeesEstimatesParams) WithDefaults() *GetMyFeesEstimatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get my fees estimates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMyFeesEstimatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) WithTimeout(timeout time.Duration) *GetMyFeesEstimatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) WithContext(ctx context.Context) *GetMyFeesEstimatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) WithHTTPClient(client *http.Client) *GetMyFeesEstimatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) WithBody(body product_fees_v0_models.GetMyFeesEstimatesRequest) *GetMyFeesEstimatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get my fees estimates params
func (o *GetMyFeesEstimatesParams) SetBody(body product_fees_v0_models.GetMyFeesEstimatesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetMyFeesEstimatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
