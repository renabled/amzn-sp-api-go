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

	"github.com/xamandar/amzn-sp-api-go/api/messaging/messaging_models"
)

// NewCreateUnexpectedProblemParams creates a new CreateUnexpectedProblemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUnexpectedProblemParams() *CreateUnexpectedProblemParams {
	return &CreateUnexpectedProblemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUnexpectedProblemParamsWithTimeout creates a new CreateUnexpectedProblemParams object
// with the ability to set a timeout on a request.
func NewCreateUnexpectedProblemParamsWithTimeout(timeout time.Duration) *CreateUnexpectedProblemParams {
	return &CreateUnexpectedProblemParams{
		timeout: timeout,
	}
}

// NewCreateUnexpectedProblemParamsWithContext creates a new CreateUnexpectedProblemParams object
// with the ability to set a context for a request.
func NewCreateUnexpectedProblemParamsWithContext(ctx context.Context) *CreateUnexpectedProblemParams {
	return &CreateUnexpectedProblemParams{
		Context: ctx,
	}
}

// NewCreateUnexpectedProblemParamsWithHTTPClient creates a new CreateUnexpectedProblemParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUnexpectedProblemParamsWithHTTPClient(client *http.Client) *CreateUnexpectedProblemParams {
	return &CreateUnexpectedProblemParams{
		HTTPClient: client,
	}
}

/*
CreateUnexpectedProblemParams contains all the parameters to send to the API endpoint

	for the create unexpected problem operation.

	Typically these are written to a http.Request.
*/
type CreateUnexpectedProblemParams struct {

	/* AmazonOrderID.

	   An Amazon order identifier. This specifies the order for which a message is sent.
	*/
	AmazonOrderID string

	// Body.
	Body *messaging_models.CreateUnexpectedProblemRequest

	/* MarketplaceIds.

	   A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	*/
	MarketplaceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create unexpected problem params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUnexpectedProblemParams) WithDefaults() *CreateUnexpectedProblemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create unexpected problem params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUnexpectedProblemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) WithTimeout(timeout time.Duration) *CreateUnexpectedProblemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) WithContext(ctx context.Context) *CreateUnexpectedProblemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) WithHTTPClient(client *http.Client) *CreateUnexpectedProblemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmazonOrderID adds the amazonOrderID to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) WithAmazonOrderID(amazonOrderID string) *CreateUnexpectedProblemParams {
	o.SetAmazonOrderID(amazonOrderID)
	return o
}

// SetAmazonOrderID adds the amazonOrderId to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) SetAmazonOrderID(amazonOrderID string) {
	o.AmazonOrderID = amazonOrderID
}

// WithBody adds the body to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) WithBody(body *messaging_models.CreateUnexpectedProblemRequest) *CreateUnexpectedProblemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) SetBody(body *messaging_models.CreateUnexpectedProblemRequest) {
	o.Body = body
}

// WithMarketplaceIds adds the marketplaceIds to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) WithMarketplaceIds(marketplaceIds []string) *CreateUnexpectedProblemParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the create unexpected problem params
func (o *CreateUnexpectedProblemParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUnexpectedProblemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param amazonOrderId
	if err := r.SetPathParam("amazonOrderId", o.AmazonOrderID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

// bindParamCreateUnexpectedProblem binds the parameter marketplaceIds
func (o *CreateUnexpectedProblemParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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
