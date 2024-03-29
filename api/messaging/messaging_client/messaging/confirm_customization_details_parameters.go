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

	"github.com/renabled/amzn-sp-api-go/api/messaging/messaging_models"
)

// NewConfirmCustomizationDetailsParams creates a new ConfirmCustomizationDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfirmCustomizationDetailsParams() *ConfirmCustomizationDetailsParams {
	return &ConfirmCustomizationDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfirmCustomizationDetailsParamsWithTimeout creates a new ConfirmCustomizationDetailsParams object
// with the ability to set a timeout on a request.
func NewConfirmCustomizationDetailsParamsWithTimeout(timeout time.Duration) *ConfirmCustomizationDetailsParams {
	return &ConfirmCustomizationDetailsParams{
		timeout: timeout,
	}
}

// NewConfirmCustomizationDetailsParamsWithContext creates a new ConfirmCustomizationDetailsParams object
// with the ability to set a context for a request.
func NewConfirmCustomizationDetailsParamsWithContext(ctx context.Context) *ConfirmCustomizationDetailsParams {
	return &ConfirmCustomizationDetailsParams{
		Context: ctx,
	}
}

// NewConfirmCustomizationDetailsParamsWithHTTPClient creates a new ConfirmCustomizationDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfirmCustomizationDetailsParamsWithHTTPClient(client *http.Client) *ConfirmCustomizationDetailsParams {
	return &ConfirmCustomizationDetailsParams{
		HTTPClient: client,
	}
}

/*
ConfirmCustomizationDetailsParams contains all the parameters to send to the API endpoint

	for the confirm customization details operation.

	Typically these are written to a http.Request.
*/
type ConfirmCustomizationDetailsParams struct {

	/* AmazonOrderID.

	   An Amazon order identifier. This specifies the order for which a message is sent.
	*/
	AmazonOrderID string

	// Body.
	Body *messaging_models.CreateConfirmCustomizationDetailsRequest

	/* MarketplaceIds.

	   A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	*/
	MarketplaceIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the confirm customization details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmCustomizationDetailsParams) WithDefaults() *ConfirmCustomizationDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the confirm customization details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfirmCustomizationDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) WithTimeout(timeout time.Duration) *ConfirmCustomizationDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) WithContext(ctx context.Context) *ConfirmCustomizationDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) WithHTTPClient(client *http.Client) *ConfirmCustomizationDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmazonOrderID adds the amazonOrderID to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) WithAmazonOrderID(amazonOrderID string) *ConfirmCustomizationDetailsParams {
	o.SetAmazonOrderID(amazonOrderID)
	return o
}

// SetAmazonOrderID adds the amazonOrderId to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) SetAmazonOrderID(amazonOrderID string) {
	o.AmazonOrderID = amazonOrderID
}

// WithBody adds the body to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) WithBody(body *messaging_models.CreateConfirmCustomizationDetailsRequest) *ConfirmCustomizationDetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) SetBody(body *messaging_models.CreateConfirmCustomizationDetailsRequest) {
	o.Body = body
}

// WithMarketplaceIds adds the marketplaceIds to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) WithMarketplaceIds(marketplaceIds []string) *ConfirmCustomizationDetailsParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the confirm customization details params
func (o *ConfirmCustomizationDetailsParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WriteToRequest writes these params to a swagger request
func (o *ConfirmCustomizationDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamConfirmCustomizationDetails binds the parameter marketplaceIds
func (o *ConfirmCustomizationDetailsParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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
