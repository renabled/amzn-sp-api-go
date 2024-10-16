// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

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

// NewGetLabelsParams creates a new GetLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLabelsParams() *GetLabelsParams {
	return &GetLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelsParamsWithTimeout creates a new GetLabelsParams object
// with the ability to set a timeout on a request.
func NewGetLabelsParamsWithTimeout(timeout time.Duration) *GetLabelsParams {
	return &GetLabelsParams{
		timeout: timeout,
	}
}

// NewGetLabelsParamsWithContext creates a new GetLabelsParams object
// with the ability to set a context for a request.
func NewGetLabelsParamsWithContext(ctx context.Context) *GetLabelsParams {
	return &GetLabelsParams{
		Context: ctx,
	}
}

// NewGetLabelsParamsWithHTTPClient creates a new GetLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLabelsParamsWithHTTPClient(client *http.Client) *GetLabelsParams {
	return &GetLabelsParams{
		HTTPClient: client,
	}
}

/*
GetLabelsParams contains all the parameters to send to the API endpoint

	for the get labels operation.

	Typically these are written to a http.Request.
*/
type GetLabelsParams struct {

	/* LabelType.

	   The type of labels requested.
	*/
	LabelType string

	/* NumberOfPackages.

	   The number of packages in the shipment.
	*/
	NumberOfPackages *int64

	/* NumberOfPallets.

	   The number of pallets in the shipment. This returns four identical labels for each pallet.
	*/
	NumberOfPallets *int64

	/* PackageLabelsToPrint.

	     A list of identifiers that specify packages for which you want package labels printed.

	If you provide box content information with the [FBA Inbound Shipment Carton Information Feed](https://developer-docs.amazon.com/sp-api/docs/fulfillment-by-amazon-feed-type-values#fba-inbound-shipment-carton-information-feed), then `PackageLabelsToPrint` must match the `CartonId` values you provide through that feed. If you provide box content information with the Fulfillment Inbound API v2024-03-20, then `PackageLabelsToPrint` must match the `boxID` values from the [`listShipmentBoxes`](https://developer-docs.amazon.com/sp-api/docs/fulfillment-inbound-api-v2024-03-20-reference#listshipmentboxes) response. If these values do not match as required, the operation returns the `IncorrectPackageIdentifier` error code.
	*/
	PackageLabelsToPrint []string

	/* PageSize.

	   The page size for paginating through the total packages' labels. This is a required parameter for Non-Partnered LTL Shipments. Max value:1000.
	*/
	PageSize *int64

	/* PageStartIndex.

	   The page start index for paginating through the total packages' labels. This is a required parameter for Non-Partnered LTL Shipments.
	*/
	PageStartIndex *int64

	/* PageType.

	   The page type to use to print the labels. Submitting a PageType value that is not supported in your marketplace returns an error.
	*/
	PageType string

	/* ShipmentID.

	   A shipment identifier originally returned by the createInboundShipmentPlan operation.
	*/
	ShipmentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsParams) WithDefaults() *GetLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) WithTimeout(timeout time.Duration) *GetLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get labels params
func (o *GetLabelsParams) WithContext(ctx context.Context) *GetLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get labels params
func (o *GetLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) WithHTTPClient(client *http.Client) *GetLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabelType adds the labelType to the get labels params
func (o *GetLabelsParams) WithLabelType(labelType string) *GetLabelsParams {
	o.SetLabelType(labelType)
	return o
}

// SetLabelType adds the labelType to the get labels params
func (o *GetLabelsParams) SetLabelType(labelType string) {
	o.LabelType = labelType
}

// WithNumberOfPackages adds the numberOfPackages to the get labels params
func (o *GetLabelsParams) WithNumberOfPackages(numberOfPackages *int64) *GetLabelsParams {
	o.SetNumberOfPackages(numberOfPackages)
	return o
}

// SetNumberOfPackages adds the numberOfPackages to the get labels params
func (o *GetLabelsParams) SetNumberOfPackages(numberOfPackages *int64) {
	o.NumberOfPackages = numberOfPackages
}

// WithNumberOfPallets adds the numberOfPallets to the get labels params
func (o *GetLabelsParams) WithNumberOfPallets(numberOfPallets *int64) *GetLabelsParams {
	o.SetNumberOfPallets(numberOfPallets)
	return o
}

// SetNumberOfPallets adds the numberOfPallets to the get labels params
func (o *GetLabelsParams) SetNumberOfPallets(numberOfPallets *int64) {
	o.NumberOfPallets = numberOfPallets
}

// WithPackageLabelsToPrint adds the packageLabelsToPrint to the get labels params
func (o *GetLabelsParams) WithPackageLabelsToPrint(packageLabelsToPrint []string) *GetLabelsParams {
	o.SetPackageLabelsToPrint(packageLabelsToPrint)
	return o
}

// SetPackageLabelsToPrint adds the packageLabelsToPrint to the get labels params
func (o *GetLabelsParams) SetPackageLabelsToPrint(packageLabelsToPrint []string) {
	o.PackageLabelsToPrint = packageLabelsToPrint
}

// WithPageSize adds the pageSize to the get labels params
func (o *GetLabelsParams) WithPageSize(pageSize *int64) *GetLabelsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get labels params
func (o *GetLabelsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPageStartIndex adds the pageStartIndex to the get labels params
func (o *GetLabelsParams) WithPageStartIndex(pageStartIndex *int64) *GetLabelsParams {
	o.SetPageStartIndex(pageStartIndex)
	return o
}

// SetPageStartIndex adds the pageStartIndex to the get labels params
func (o *GetLabelsParams) SetPageStartIndex(pageStartIndex *int64) {
	o.PageStartIndex = pageStartIndex
}

// WithPageType adds the pageType to the get labels params
func (o *GetLabelsParams) WithPageType(pageType string) *GetLabelsParams {
	o.SetPageType(pageType)
	return o
}

// SetPageType adds the pageType to the get labels params
func (o *GetLabelsParams) SetPageType(pageType string) {
	o.PageType = pageType
}

// WithShipmentID adds the shipmentID to the get labels params
func (o *GetLabelsParams) WithShipmentID(shipmentID string) *GetLabelsParams {
	o.SetShipmentID(shipmentID)
	return o
}

// SetShipmentID adds the shipmentId to the get labels params
func (o *GetLabelsParams) SetShipmentID(shipmentID string) {
	o.ShipmentID = shipmentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param LabelType
	qrLabelType := o.LabelType
	qLabelType := qrLabelType
	if qLabelType != "" {

		if err := r.SetQueryParam("LabelType", qLabelType); err != nil {
			return err
		}
	}

	if o.NumberOfPackages != nil {

		// query param NumberOfPackages
		var qrNumberOfPackages int64

		if o.NumberOfPackages != nil {
			qrNumberOfPackages = *o.NumberOfPackages
		}
		qNumberOfPackages := swag.FormatInt64(qrNumberOfPackages)
		if qNumberOfPackages != "" {

			if err := r.SetQueryParam("NumberOfPackages", qNumberOfPackages); err != nil {
				return err
			}
		}
	}

	if o.NumberOfPallets != nil {

		// query param NumberOfPallets
		var qrNumberOfPallets int64

		if o.NumberOfPallets != nil {
			qrNumberOfPallets = *o.NumberOfPallets
		}
		qNumberOfPallets := swag.FormatInt64(qrNumberOfPallets)
		if qNumberOfPallets != "" {

			if err := r.SetQueryParam("NumberOfPallets", qNumberOfPallets); err != nil {
				return err
			}
		}
	}

	if o.PackageLabelsToPrint != nil {

		// binding items for PackageLabelsToPrint
		joinedPackageLabelsToPrint := o.bindParamPackageLabelsToPrint(reg)

		// query array param PackageLabelsToPrint
		if err := r.SetQueryParam("PackageLabelsToPrint", joinedPackageLabelsToPrint...); err != nil {
			return err
		}
	}

	if o.PageSize != nil {

		// query param PageSize
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("PageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PageStartIndex != nil {

		// query param PageStartIndex
		var qrPageStartIndex int64

		if o.PageStartIndex != nil {
			qrPageStartIndex = *o.PageStartIndex
		}
		qPageStartIndex := swag.FormatInt64(qrPageStartIndex)
		if qPageStartIndex != "" {

			if err := r.SetQueryParam("PageStartIndex", qPageStartIndex); err != nil {
				return err
			}
		}
	}

	// query param PageType
	qrPageType := o.PageType
	qPageType := qrPageType
	if qPageType != "" {

		if err := r.SetQueryParam("PageType", qPageType); err != nil {
			return err
		}
	}

	// path param shipmentId
	if err := r.SetPathParam("shipmentId", o.ShipmentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetLabels binds the parameter PackageLabelsToPrint
func (o *GetLabelsParams) bindParamPackageLabelsToPrint(formats strfmt.Registry) []string {
	packageLabelsToPrintIR := o.PackageLabelsToPrint

	var packageLabelsToPrintIC []string
	for _, packageLabelsToPrintIIR := range packageLabelsToPrintIR { // explode []string

		packageLabelsToPrintIIV := packageLabelsToPrintIIR // string as string
		packageLabelsToPrintIC = append(packageLabelsToPrintIC, packageLabelsToPrintIIV)
	}

	// items.CollectionFormat: ""
	packageLabelsToPrintIS := swag.JoinByFormat(packageLabelsToPrintIC, "")

	return packageLabelsToPrintIS
}
