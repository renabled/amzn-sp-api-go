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
)

// NewCreateNegativeFeedbackRemovalParams creates a new CreateNegativeFeedbackRemovalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNegativeFeedbackRemovalParams() *CreateNegativeFeedbackRemovalParams {
	return &CreateNegativeFeedbackRemovalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNegativeFeedbackRemovalParamsWithTimeout creates a new CreateNegativeFeedbackRemovalParams object
// with the ability to set a timeout on a request.
func NewCreateNegativeFeedbackRemovalParamsWithTimeout(timeout time.Duration) *CreateNegativeFeedbackRemovalParams {
	return &CreateNegativeFeedbackRemovalParams{
		timeout: timeout,
	}
}

// NewCreateNegativeFeedbackRemovalParamsWithContext creates a new CreateNegativeFeedbackRemovalParams object
// with the ability to set a context for a request.
func NewCreateNegativeFeedbackRemovalParamsWithContext(ctx context.Context) *CreateNegativeFeedbackRemovalParams {
	return &CreateNegativeFeedbackRemovalParams{
		Context: ctx,
	}
}

// NewCreateNegativeFeedbackRemovalParamsWithHTTPClient creates a new CreateNegativeFeedbackRemovalParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNegativeFeedbackRemovalParamsWithHTTPClient(client *http.Client) *CreateNegativeFeedbackRemovalParams {
	return &CreateNegativeFeedbackRemovalParams{
		HTTPClient: client,
	}
}

/*
CreateNegativeFeedbackRemovalParams contains all the parameters to send to the API endpoint

	for the create negative feedback removal operation.

	Typically these are written to a http.Request.
*/
type CreateNegativeFeedbackRemovalParams struct {

	/* AmazonOrderID.

	   An Amazon order identifier. This specifies the order for which a message is sent.
	*/
	AmazonOrderID string

	/* MarketplaceIds.

	   A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	*/
	MarketplaceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create negative feedback removal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNegativeFeedbackRemovalParams) WithDefaults() *CreateNegativeFeedbackRemovalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create negative feedback removal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNegativeFeedbackRemovalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) WithTimeout(timeout time.Duration) *CreateNegativeFeedbackRemovalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) WithContext(ctx context.Context) *CreateNegativeFeedbackRemovalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) WithHTTPClient(client *http.Client) *CreateNegativeFeedbackRemovalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmazonOrderID adds the amazonOrderID to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) WithAmazonOrderID(amazonOrderID string) *CreateNegativeFeedbackRemovalParams {
	o.SetAmazonOrderID(amazonOrderID)
	return o
}

// SetAmazonOrderID adds the amazonOrderId to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) SetAmazonOrderID(amazonOrderID string) {
	o.AmazonOrderID = amazonOrderID
}

// WithMarketplaceIds adds the marketplaceIds to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) WithMarketplaceIds(marketplaceIds []string) *CreateNegativeFeedbackRemovalParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the create negative feedback removal params
func (o *CreateNegativeFeedbackRemovalParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNegativeFeedbackRemovalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param amazonOrderId
	if err := r.SetPathParam("amazonOrderId", o.AmazonOrderID); err != nil {
		return err
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

// bindParamCreateNegativeFeedbackRemoval binds the parameter marketplaceIds
func (o *CreateNegativeFeedbackRemovalParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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
