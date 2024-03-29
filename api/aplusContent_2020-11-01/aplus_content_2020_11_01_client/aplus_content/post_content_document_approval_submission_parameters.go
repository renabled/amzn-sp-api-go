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
)

// NewPostContentDocumentApprovalSubmissionParams creates a new PostContentDocumentApprovalSubmissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostContentDocumentApprovalSubmissionParams() *PostContentDocumentApprovalSubmissionParams {
	return &PostContentDocumentApprovalSubmissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostContentDocumentApprovalSubmissionParamsWithTimeout creates a new PostContentDocumentApprovalSubmissionParams object
// with the ability to set a timeout on a request.
func NewPostContentDocumentApprovalSubmissionParamsWithTimeout(timeout time.Duration) *PostContentDocumentApprovalSubmissionParams {
	return &PostContentDocumentApprovalSubmissionParams{
		timeout: timeout,
	}
}

// NewPostContentDocumentApprovalSubmissionParamsWithContext creates a new PostContentDocumentApprovalSubmissionParams object
// with the ability to set a context for a request.
func NewPostContentDocumentApprovalSubmissionParamsWithContext(ctx context.Context) *PostContentDocumentApprovalSubmissionParams {
	return &PostContentDocumentApprovalSubmissionParams{
		Context: ctx,
	}
}

// NewPostContentDocumentApprovalSubmissionParamsWithHTTPClient creates a new PostContentDocumentApprovalSubmissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostContentDocumentApprovalSubmissionParamsWithHTTPClient(client *http.Client) *PostContentDocumentApprovalSubmissionParams {
	return &PostContentDocumentApprovalSubmissionParams{
		HTTPClient: client,
	}
}

/*
PostContentDocumentApprovalSubmissionParams contains all the parameters to send to the API endpoint

	for the post content document approval submission operation.

	Typically these are written to a http.Request.
*/
type PostContentDocumentApprovalSubmissionParams struct {

	/* ContentReferenceKey.

	   The unique reference key for the A+ Content document. A content reference key cannot form a permalink and may change in the future. A content reference key is not guaranteed to match any A+ content identifier.
	*/
	ContentReferenceKey string

	/* MarketplaceID.

	   The identifier for the marketplace where the A+ Content is published.
	*/
	MarketplaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post content document approval submission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostContentDocumentApprovalSubmissionParams) WithDefaults() *PostContentDocumentApprovalSubmissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post content document approval submission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostContentDocumentApprovalSubmissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) WithTimeout(timeout time.Duration) *PostContentDocumentApprovalSubmissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) WithContext(ctx context.Context) *PostContentDocumentApprovalSubmissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) WithHTTPClient(client *http.Client) *PostContentDocumentApprovalSubmissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentReferenceKey adds the contentReferenceKey to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) WithContentReferenceKey(contentReferenceKey string) *PostContentDocumentApprovalSubmissionParams {
	o.SetContentReferenceKey(contentReferenceKey)
	return o
}

// SetContentReferenceKey adds the contentReferenceKey to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) SetContentReferenceKey(contentReferenceKey string) {
	o.ContentReferenceKey = contentReferenceKey
}

// WithMarketplaceID adds the marketplaceID to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) WithMarketplaceID(marketplaceID string) *PostContentDocumentApprovalSubmissionParams {
	o.SetMarketplaceID(marketplaceID)
	return o
}

// SetMarketplaceID adds the marketplaceId to the post content document approval submission params
func (o *PostContentDocumentApprovalSubmissionParams) SetMarketplaceID(marketplaceID string) {
	o.MarketplaceID = marketplaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostContentDocumentApprovalSubmissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
