// Code generated by go-swagger; DO NOT EDIT.

package shipping

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

// NewGetShipmentDocumentsParams creates a new GetShipmentDocumentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetShipmentDocumentsParams() *GetShipmentDocumentsParams {
	return &GetShipmentDocumentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetShipmentDocumentsParamsWithTimeout creates a new GetShipmentDocumentsParams object
// with the ability to set a timeout on a request.
func NewGetShipmentDocumentsParamsWithTimeout(timeout time.Duration) *GetShipmentDocumentsParams {
	return &GetShipmentDocumentsParams{
		timeout: timeout,
	}
}

// NewGetShipmentDocumentsParamsWithContext creates a new GetShipmentDocumentsParams object
// with the ability to set a context for a request.
func NewGetShipmentDocumentsParamsWithContext(ctx context.Context) *GetShipmentDocumentsParams {
	return &GetShipmentDocumentsParams{
		Context: ctx,
	}
}

// NewGetShipmentDocumentsParamsWithHTTPClient creates a new GetShipmentDocumentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetShipmentDocumentsParamsWithHTTPClient(client *http.Client) *GetShipmentDocumentsParams {
	return &GetShipmentDocumentsParams{
		HTTPClient: client,
	}
}

/*
GetShipmentDocumentsParams contains all the parameters to send to the API endpoint

	for the get shipment documents operation.

	Typically these are written to a http.Request.
*/
type GetShipmentDocumentsParams struct {

	/* Dpi.

	   The resolution of the document (for example, 300 means 300 dots per inch). Must be one of the supported resolutions returned in the response to the getRates operation.
	*/
	Dpi *float64

	/* Format.

	   The file format of the document. Must be one of the supported formats returned by the getRates operation.
	*/
	Format *string

	/* PackageClientReferenceID.

	   The package client reference identifier originally provided in the request body parameter for the getRates operation.
	*/
	PackageClientReferenceID string

	/* ShipmentID.

	   The shipment identifier originally returned by the purchaseShipment operation.
	*/
	ShipmentID string

	/* XAmznShippingBusinessID.

	   Amazon shipping business to assume for this request. The default is AmazonShipping_UK.
	*/
	XAmznShippingBusinessID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get shipment documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShipmentDocumentsParams) WithDefaults() *GetShipmentDocumentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get shipment documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetShipmentDocumentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithTimeout(timeout time.Duration) *GetShipmentDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithContext(ctx context.Context) *GetShipmentDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithHTTPClient(client *http.Client) *GetShipmentDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDpi adds the dpi to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithDpi(dpi *float64) *GetShipmentDocumentsParams {
	o.SetDpi(dpi)
	return o
}

// SetDpi adds the dpi to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetDpi(dpi *float64) {
	o.Dpi = dpi
}

// WithFormat adds the format to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithFormat(format *string) *GetShipmentDocumentsParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetFormat(format *string) {
	o.Format = format
}

// WithPackageClientReferenceID adds the packageClientReferenceID to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithPackageClientReferenceID(packageClientReferenceID string) *GetShipmentDocumentsParams {
	o.SetPackageClientReferenceID(packageClientReferenceID)
	return o
}

// SetPackageClientReferenceID adds the packageClientReferenceId to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetPackageClientReferenceID(packageClientReferenceID string) {
	o.PackageClientReferenceID = packageClientReferenceID
}

// WithShipmentID adds the shipmentID to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithShipmentID(shipmentID string) *GetShipmentDocumentsParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WithXAmznShippingBusinessID adds the xAmznShippingBusinessID to the get shipment documents params
func (o *GetShipmentDocumentsParams) WithXAmznShippingBusinessID(xAmznShippingBusinessID *string) *GetShipmentDocumentsParams {
	o.SetXAmznShippingBusinessID(xAmznShippingBusinessID)
	return o
}

// SetXAmznShippingBusinessID adds the xAmznShippingBusinessId to the get shipment documents params
func (o *GetShipmentDocumentsParams) SetXAmznShippingBusinessID(xAmznShippingBusinessID *string) {
	o.XAmznShippingBusinessID = xAmznShippingBusinessID
}

// WriteToRequest writes these params to a swagger request
func (o *GetShipmentDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dpi != nil {

		// query param dpi
		var qrDpi float64

		if o.Dpi != nil {
			qrDpi = *o.Dpi
		}
		qDpi := swag.FormatFloat64(qrDpi)
		if qDpi != "" {

			if err := r.SetQueryParam("dpi", qDpi); err != nil {
				return err
			}
		}
	}

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}
	}

	// query param packageClientReferenceId
	qrPackageClientReferenceID := o.PackageClientReferenceID
	qPackageClientReferenceID := qrPackageClientReferenceID
	if qPackageClientReferenceID != "" {

		if err := r.SetQueryParam("packageClientReferenceId", qPackageClientReferenceID); err != nil {
			return err
		}
	}

	// path param shipmentId
	if err := r.SetPathParam("shipmentId", o.ShipmentID); err != nil {
		return err
	}

	if o.XAmznShippingBusinessID != nil {

		// header param x-amzn-shipping-business-id
		if err := r.SetHeaderParam("x-amzn-shipping-business-id", *o.XAmznShippingBusinessID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}