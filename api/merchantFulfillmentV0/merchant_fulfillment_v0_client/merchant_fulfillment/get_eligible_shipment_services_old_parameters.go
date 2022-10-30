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

// NewGetEligibleShipmentServicesOldParams creates a new GetEligibleShipmentServicesOldParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEligibleShipmentServicesOldParams() *GetEligibleShipmentServicesOldParams {
	return &GetEligibleShipmentServicesOldParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEligibleShipmentServicesOldParamsWithTimeout creates a new GetEligibleShipmentServicesOldParams object
// with the ability to set a timeout on a request.
func NewGetEligibleShipmentServicesOldParamsWithTimeout(timeout time.Duration) *GetEligibleShipmentServicesOldParams {
	return &GetEligibleShipmentServicesOldParams{
		timeout: timeout,
	}
}

// NewGetEligibleShipmentServicesOldParamsWithContext creates a new GetEligibleShipmentServicesOldParams object
// with the ability to set a context for a request.
func NewGetEligibleShipmentServicesOldParamsWithContext(ctx context.Context) *GetEligibleShipmentServicesOldParams {
	return &GetEligibleShipmentServicesOldParams{
		Context: ctx,
	}
}

// NewGetEligibleShipmentServicesOldParamsWithHTTPClient creates a new GetEligibleShipmentServicesOldParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEligibleShipmentServicesOldParamsWithHTTPClient(client *http.Client) *GetEligibleShipmentServicesOldParams {
	return &GetEligibleShipmentServicesOldParams{
		HTTPClient: client,
	}
}

/*
GetEligibleShipmentServicesOldParams contains all the parameters to send to the API endpoint

	for the get eligible shipment services old operation.

	Typically these are written to a http.Request.
*/
type GetEligibleShipmentServicesOldParams struct {

	// Body.
	Body *merchant_fulfillment_v0_models.GetEligibleShipmentServicesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get eligible shipment services old params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEligibleShipmentServicesOldParams) WithDefaults() *GetEligibleShipmentServicesOldParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get eligible shipment services old params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEligibleShipmentServicesOldParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) WithTimeout(timeout time.Duration) *GetEligibleShipmentServicesOldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) WithContext(ctx context.Context) *GetEligibleShipmentServicesOldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) WithHTTPClient(client *http.Client) *GetEligibleShipmentServicesOldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) WithBody(body *merchant_fulfillment_v0_models.GetEligibleShipmentServicesRequest) *GetEligibleShipmentServicesOldParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get eligible shipment services old params
func (o *GetEligibleShipmentServicesOldParams) SetBody(body *merchant_fulfillment_v0_models.GetEligibleShipmentServicesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetEligibleShipmentServicesOldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
