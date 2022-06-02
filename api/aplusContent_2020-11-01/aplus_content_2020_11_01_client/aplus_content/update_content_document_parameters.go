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

	"github.com/xamandar/amzn-sp-api-go/api/aplusContent_2020-11-01/aplus_content_2020_11_01_models"
)

// NewUpdateContentDocumentParams creates a new UpdateContentDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateContentDocumentParams() *UpdateContentDocumentParams {
	return &UpdateContentDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateContentDocumentParamsWithTimeout creates a new UpdateContentDocumentParams object
// with the ability to set a timeout on a request.
func NewUpdateContentDocumentParamsWithTimeout(timeout time.Duration) *UpdateContentDocumentParams {
	return &UpdateContentDocumentParams{
		timeout: timeout,
	}
}

// NewUpdateContentDocumentParamsWithContext creates a new UpdateContentDocumentParams object
// with the ability to set a context for a request.
func NewUpdateContentDocumentParamsWithContext(ctx context.Context) *UpdateContentDocumentParams {
	return &UpdateContentDocumentParams{
		Context: ctx,
	}
}

// NewUpdateContentDocumentParamsWithHTTPClient creates a new UpdateContentDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateContentDocumentParamsWithHTTPClient(client *http.Client) *UpdateContentDocumentParams {
	return &UpdateContentDocumentParams{
		HTTPClient: client,
	}
}

/* UpdateContentDocumentParams contains all the parameters to send to the API endpoint
   for the update content document operation.

   Typically these are written to a http.Request.
*/
type UpdateContentDocumentParams struct {

	/* ContentReferenceKey.

	   The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ Content identifier.
	*/
	ContentReferenceKey string

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

// WithDefaults hydrates default values in the update content document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContentDocumentParams) WithDefaults() *UpdateContentDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update content document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateContentDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update content document params
func (o *UpdateContentDocumentParams) WithTimeout(timeout time.Duration) *UpdateContentDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update content document params
func (o *UpdateContentDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update content document params
func (o *UpdateContentDocumentParams) WithContext(ctx context.Context) *UpdateContentDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update content document params
func (o *UpdateContentDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update content document params
func (o *UpdateContentDocumentParams) WithHTTPClient(client *http.Client) *UpdateContentDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update content document params
func (o *UpdateContentDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentReferenceKey adds the contentReferenceKey to the update content document params
func (o *UpdateContentDocumentParams) WithContentReferenceKey(contentReferenceKey string) *UpdateContentDocumentParams {
	o.SetContentReferenceKey(contentReferenceKey)
	return o
}

// SetContentReferenceKey adds the contentReferenceKey to the update content document params
func (o *UpdateContentDocumentParams) SetContentReferenceKey(contentReferenceKey string) {
	o.ContentReferenceKey = contentReferenceKey
}

// WithMarketplaceID adds the marketplaceID to the update content document params
func (o *UpdateContentDocumentParams) WithMarketplaceID(marketplaceID string) *UpdateContentDocumentParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the update content document params
func (o *UpdateContentDocumentParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WithPostContentDocumentRequest adds the postContentDocumentRequest to the update content document params
func (o *UpdateContentDocumentParams) WithPostContentDocumentRequest(postContentDocumentRequest *aplus_content_2020_11_01_models.PostContentDocumentRequest) *UpdateContentDocumentParams {
	o.SetPostContentDocumentRequest(postContentDocumentRequest)
	return o
}

// SetPostContentDocumentRequest adds the postContentDocumentRequest to the update content document params
func (o *UpdateContentDocumentParams) SetPostContentDocumentRequest(postContentDocumentRequest *aplus_content_2020_11_01_models.PostContentDocumentRequest) {
	o.PostContentDocumentRequest = postContentDocumentRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateContentDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contentReferenceKey
	if err := r.SetPathParam("contentReferenceKey", o.ContentReferenceKey); err != nil {
		return err
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
