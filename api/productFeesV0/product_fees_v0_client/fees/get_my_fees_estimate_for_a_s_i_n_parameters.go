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

// NewGetMyFeesEstimateForASINParams creates a new GetMyFeesEstimateForASINParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMyFeesEstimateForASINParams() *GetMyFeesEstimateForASINParams {
	return &GetMyFeesEstimateForASINParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMyFeesEstimateForASINParamsWithTimeout creates a new GetMyFeesEstimateForASINParams object
// with the ability to set a timeout on a request.
func NewGetMyFeesEstimateForASINParamsWithTimeout(timeout time.Duration) *GetMyFeesEstimateForASINParams {
	return &GetMyFeesEstimateForASINParams{
		timeout: timeout,
	}
}

// NewGetMyFeesEstimateForASINParamsWithContext creates a new GetMyFeesEstimateForASINParams object
// with the ability to set a context for a request.
func NewGetMyFeesEstimateForASINParamsWithContext(ctx context.Context) *GetMyFeesEstimateForASINParams {
	return &GetMyFeesEstimateForASINParams{
		Context: ctx,
	}
}

// NewGetMyFeesEstimateForASINParamsWithHTTPClient creates a new GetMyFeesEstimateForASINParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMyFeesEstimateForASINParamsWithHTTPClient(client *http.Client) *GetMyFeesEstimateForASINParams {
	return &GetMyFeesEstimateForASINParams{
		HTTPClient: client,
	}
}

/* GetMyFeesEstimateForASINParams contains all the parameters to send to the API endpoint
   for the get my fees estimate for a s i n operation.

   Typically these are written to a http.Request.
*/
type GetMyFeesEstimateForASINParams struct {

	/* Asin.

	   The Amazon Standard Identification Number (ASIN) of the item.
	*/
	Asin string

	// Body.
	Body *product_fees_v0_models.GetMyFeesEstimateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get my fees estimate for a s i n params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMyFeesEstimateForASINParams) WithDefaults() *GetMyFeesEstimateForASINParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get my fees estimate for a s i n params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMyFeesEstimateForASINParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) WithTimeout(timeout time.Duration) *GetMyFeesEstimateForASINParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) WithContext(ctx context.Context) *GetMyFeesEstimateForASINParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) WithHTTPClient(client *http.Client) *GetMyFeesEstimateForASINParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsin adds the asin to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) WithAsin(asin string) *GetMyFeesEstimateForASINParams {
	o.SetAsin(asin)
	return o
}

// SetAsin adds the asin to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) SetAsin(asin string) {
	o.Asin = asin
}

// WithBody adds the body to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) WithBody(body *product_fees_v0_models.GetMyFeesEstimateRequest) *GetMyFeesEstimateForASINParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get my fees estimate for a s i n params
func (o *GetMyFeesEstimateForASINParams) SetBody(body *product_fees_v0_models.GetMyFeesEstimateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetMyFeesEstimateForASINParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Asin
	if err := r.SetPathParam("Asin", o.Asin); err != nil {
		return err
	}
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
