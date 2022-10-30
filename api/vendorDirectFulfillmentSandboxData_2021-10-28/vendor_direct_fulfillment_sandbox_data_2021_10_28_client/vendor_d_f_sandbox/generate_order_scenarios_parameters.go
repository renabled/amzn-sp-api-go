// Code generated by go-swagger; DO NOT EDIT.

package vendor_d_f_sandbox

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

	"github.com/xamandar/amzn-sp-api-go/api/vendorDirectFulfillmentSandboxData_2021-10-28/vendor_direct_fulfillment_sandbox_data_2021_10_28_models"
)

// NewGenerateOrderScenariosParams creates a new GenerateOrderScenariosParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateOrderScenariosParams() *GenerateOrderScenariosParams {
	return &GenerateOrderScenariosParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateOrderScenariosParamsWithTimeout creates a new GenerateOrderScenariosParams object
// with the ability to set a timeout on a request.
func NewGenerateOrderScenariosParamsWithTimeout(timeout time.Duration) *GenerateOrderScenariosParams {
	return &GenerateOrderScenariosParams{
		timeout: timeout,
	}
}

// NewGenerateOrderScenariosParamsWithContext creates a new GenerateOrderScenariosParams object
// with the ability to set a context for a request.
func NewGenerateOrderScenariosParamsWithContext(ctx context.Context) *GenerateOrderScenariosParams {
	return &GenerateOrderScenariosParams{
		Context: ctx,
	}
}

// NewGenerateOrderScenariosParamsWithHTTPClient creates a new GenerateOrderScenariosParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateOrderScenariosParamsWithHTTPClient(client *http.Client) *GenerateOrderScenariosParams {
	return &GenerateOrderScenariosParams{
		HTTPClient: client,
	}
}

/*
GenerateOrderScenariosParams contains all the parameters to send to the API endpoint

	for the generate order scenarios operation.

	Typically these are written to a http.Request.
*/
type GenerateOrderScenariosParams struct {

	// Body.
	Body *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.GenerateOrderScenarioRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate order scenarios params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateOrderScenariosParams) WithDefaults() *GenerateOrderScenariosParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate order scenarios params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateOrderScenariosParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate order scenarios params
func (o *GenerateOrderScenariosParams) WithTimeout(timeout time.Duration) *GenerateOrderScenariosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate order scenarios params
func (o *GenerateOrderScenariosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate order scenarios params
func (o *GenerateOrderScenariosParams) WithContext(ctx context.Context) *GenerateOrderScenariosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate order scenarios params
func (o *GenerateOrderScenariosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate order scenarios params
func (o *GenerateOrderScenariosParams) WithHTTPClient(client *http.Client) *GenerateOrderScenariosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate order scenarios params
func (o *GenerateOrderScenariosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the generate order scenarios params
func (o *GenerateOrderScenariosParams) WithBody(body *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.GenerateOrderScenarioRequest) *GenerateOrderScenariosParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate order scenarios params
func (o *GenerateOrderScenariosParams) SetBody(body *vendor_direct_fulfillment_sandbox_data_2021_10_28_models.GenerateOrderScenarioRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateOrderScenariosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
