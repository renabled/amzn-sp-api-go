// Code generated by go-swagger; DO NOT EDIT.

package aplus_content

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

	"github.com/xamandar/amzn-sp-api-go/api/aplusContent_2020-11-01/aplus_content_2020_11_01_models"
)

// NewValidateContentDocumentAsinRelationsParams creates a new ValidateContentDocumentAsinRelationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateContentDocumentAsinRelationsParams() *ValidateContentDocumentAsinRelationsParams {
	return &ValidateContentDocumentAsinRelationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateContentDocumentAsinRelationsParamsWithTimeout creates a new ValidateContentDocumentAsinRelationsParams object
// with the ability to set a timeout on a request.
func NewValidateContentDocumentAsinRelationsParamsWithTimeout(timeout time.Duration) *ValidateContentDocumentAsinRelationsParams {
	return &ValidateContentDocumentAsinRelationsParams{
		timeout: timeout,
	}
}

// NewValidateContentDocumentAsinRelationsParamsWithContext creates a new ValidateContentDocumentAsinRelationsParams object
// with the ability to set a context for a request.
func NewValidateContentDocumentAsinRelationsParamsWithContext(ctx context.Context) *ValidateContentDocumentAsinRelationsParams {
	return &ValidateContentDocumentAsinRelationsParams{
		Context: ctx,
	}
}

// NewValidateContentDocumentAsinRelationsParamsWithHTTPClient creates a new ValidateContentDocumentAsinRelationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateContentDocumentAsinRelationsParamsWithHTTPClient(client *http.Client) *ValidateContentDocumentAsinRelationsParams {
	return &ValidateContentDocumentAsinRelationsParams{
		HTTPClient: client,
	}
}

/*
ValidateContentDocumentAsinRelationsParams contains all the parameters to send to the API endpoint

	for the validate content document asin relations operation.

	Typically these are written to a http.Request.
*/
type ValidateContentDocumentAsinRelationsParams struct {

	/* AsinSet.

	   The set of ASINs.
	*/
	AsinSet []string

	/* MarketplaceID.

	   The identifier for the marketplace where the A+ Content is published.
	*/
	MarketplaceID string

	/* PostContentDocumentRequest.

	   The content document request details.
	*/
	PostContentDocumentRequest *aplus_content_2020_11_01_models.PostContentDocumentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate content document asin relations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateContentDocumentAsinRelationsParams) WithDefaults() *ValidateContentDocumentAsinRelationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate content document asin relations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateContentDocumentAsinRelationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) WithTimeout(timeout time.Duration) *ValidateContentDocumentAsinRelationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) WithContext(ctx context.Context) *ValidateContentDocumentAsinRelationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) WithHTTPClient(client *http.Client) *ValidateContentDocumentAsinRelationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsinSet adds the asinSet to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) WithAsinSet(asinSet []string) *ValidateContentDocumentAsinRelationsParams {
	o.SetAsinSet(asinSet)
	return o
}

// SetAsinSet adds the asinSet to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) SetAsinSet(asinSet []string) {
	o.AsinSet = asinSet
}

// WithMarketplaceID adds the marketplaceID to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) WithMarketplaceID(marketplaceID string) *ValidateContentDocumentAsinRelationsParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WithPostContentDocumentRequest adds the postContentDocumentRequest to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) WithPostContentDocumentRequest(postContentDocumentRequest *aplus_content_2020_11_01_models.PostContentDocumentRequest) *ValidateContentDocumentAsinRelationsParams {
	o.SetPostContentDocumentRequest(postContentDocumentRequest)
	return o
}

// SetPostContentDocumentRequest adds the postContentDocumentRequest to the validate content document asin relations params
func (o *ValidateContentDocumentAsinRelationsParams) SetPostContentDocumentRequest(postContentDocumentRequest *aplus_content_2020_11_01_models.PostContentDocumentRequest) {
	o.PostContentDocumentRequest = postContentDocumentRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateContentDocumentAsinRelationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AsinSet != nil {

		// binding items for asinSet
		joinedAsinSet := o.bindParamAsinSet(reg)

		// query array param asinSet
		if err := r.SetQueryParam("asinSet", joinedAsinSet...); err != nil {
			return err
		}
	}

	// query param marketplaceId
	qrMarketplaceID := o.MarketplaceID
	qMarketplaceID := qrMarketplaceID
	if qMarketplaceID != "" {

		if err := r.SetQueryParam("marketplaceId", qMarketplaceID); err != nil {
			return err
		}
	}
	if o.PostContentDocumentRequest != nil {
		if err := r.SetBodyParam(o.PostContentDocumentRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamValidateContentDocumentAsinRelations binds the parameter asinSet
func (o *ValidateContentDocumentAsinRelationsParams) bindParamAsinSet(formats strfmt.Registry) []string {
	asinSetIR := o.AsinSet

	var asinSetIC []string
	for _, asinSetIIR := range asinSetIR { // explode []string

		asinSetIIV := asinSetIIR // string as string
		asinSetIC = append(asinSetIC, asinSetIIV)
	}

	// items.CollectionFormat: "csv"
	asinSetIS := swag.JoinByFormat(asinSetIC, "csv")

	return asinSetIS
}
