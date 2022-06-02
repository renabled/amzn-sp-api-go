// Code generated by go-swagger; DO NOT EDIT.

package definitions

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

// NewGetDefinitionsProductTypeParams creates a new GetDefinitionsProductTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDefinitionsProductTypeParams() *GetDefinitionsProductTypeParams {
	return &GetDefinitionsProductTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefinitionsProductTypeParamsWithTimeout creates a new GetDefinitionsProductTypeParams object
// with the ability to set a timeout on a request.
func NewGetDefinitionsProductTypeParamsWithTimeout(timeout time.Duration) *GetDefinitionsProductTypeParams {
	return &GetDefinitionsProductTypeParams{
		timeout: timeout,
	}
}

// NewGetDefinitionsProductTypeParamsWithContext creates a new GetDefinitionsProductTypeParams object
// with the ability to set a context for a request.
func NewGetDefinitionsProductTypeParamsWithContext(ctx context.Context) *GetDefinitionsProductTypeParams {
	return &GetDefinitionsProductTypeParams{
		Context: ctx,
	}
}

// NewGetDefinitionsProductTypeParamsWithHTTPClient creates a new GetDefinitionsProductTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDefinitionsProductTypeParamsWithHTTPClient(client *http.Client) *GetDefinitionsProductTypeParams {
	return &GetDefinitionsProductTypeParams{
		HTTPClient: client,
	}
}

/* GetDefinitionsProductTypeParams contains all the parameters to send to the API endpoint
   for the get definitions product type operation.

   Typically these are written to a http.Request.
*/
type GetDefinitionsProductTypeParams struct {

	/* Locale.

	   Locale for retrieving display labels and other presentation details. Defaults to the default language of the first marketplace in the request.

	   Default: "DEFAULT"
	*/
	Locale *string

	/* MarketplaceIds.

	     A comma-delimited list of Amazon marketplace identifiers for the request.
	Note: This parameter is limited to one marketplaceId at this time.
	*/
	MarketplaceIds []string

	/* ProductType.

	   The Amazon product type name.
	*/
	ProductType string

	/* ProductTypeVersion.

	   The version of the Amazon product type to retrieve. Defaults to "LATEST",. Prerelease versions of product type definitions may be retrieved with "RELEASE_CANDIDATE". If no prerelease version is currently available, the "LATEST" live version will be provided.

	   Default: "LATEST"
	*/
	ProductTypeVersion *string

	/* Requirements.

	   The name of the requirements set to retrieve requirements for.

	   Default: "LISTING"
	*/
	Requirements *string

	/* RequirementsEnforced.

	   Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all the required attributes being present (such as for partial updates).

	   Default: "ENFORCED"
	*/
	RequirementsEnforced *string

	/* SellerID.

	   A selling partner identifier. When provided, seller-specific requirements and values are populated within the product type definition schema, such as brand names associated with the selling partner.
	*/
	SellerID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get definitions product type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDefinitionsProductTypeParams) WithDefaults() *GetDefinitionsProductTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get definitions product type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDefinitionsProductTypeParams) SetDefaults() {
	var (
		localeDefault = string("DEFAULT")

		productTypeVersionDefault = string("LATEST")

		requirementsDefault = string("LISTING")

		requirementsEnforcedDefault = string("ENFORCED")
	)

	val := GetDefinitionsProductTypeParams{
		Locale:               &localeDefault,
		ProductTypeVersion:   &productTypeVersionDefault,
		Requirements:         &requirementsDefault,
		RequirementsEnforced: &requirementsEnforcedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithTimeout(timeout time.Duration) *GetDefinitionsProductTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithContext(ctx context.Context) *GetDefinitionsProductTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithHTTPClient(client *http.Client) *GetDefinitionsProductTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocale adds the locale to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithLocale(locale *string) *GetDefinitionsProductTypeParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetLocale(locale *string) {
	o.Locale = locale
}

// WithMarketplaceIds adds the marketplaceIds to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithMarketplaceIds(marketplaceIds []string) *GetDefinitionsProductTypeParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithProductType adds the productType to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithProductType(productType string) *GetDefinitionsProductTypeParams {
	o.SetProductType(productType)
	return o
}

// SetProductType adds the productType to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetProductType(productType string) {
	o.ProductType = productType
}

// WithProductTypeVersion adds the productTypeVersion to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithProductTypeVersion(productTypeVersion *string) *GetDefinitionsProductTypeParams {
	o.SetProductTypeVersion(productTypeVersion)
	return o
}

// SetProductTypeVersion adds the productTypeVersion to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetProductTypeVersion(productTypeVersion *string) {
	o.ProductTypeVersion = productTypeVersion
}

// WithRequirements adds the requirements to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithRequirements(requirements *string) *GetDefinitionsProductTypeParams {
	o.SetRequirements(requirements)
	return o
}

// SetRequirements adds the requirements to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetRequirements(requirements *string) {
	o.Requirements = requirements
}

// WithRequirementsEnforced adds the requirementsEnforced to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithRequirementsEnforced(requirementsEnforced *string) *GetDefinitionsProductTypeParams {
	o.SetRequirementsEnforced(requirementsEnforced)
	return o
}

// SetRequirementsEnforced adds the requirementsEnforced to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetRequirementsEnforced(requirementsEnforced *string) {
	o.RequirementsEnforced = requirementsEnforced
}

// WithSellerID adds the sellerID to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) WithSellerID(sellerID *string) *GetDefinitionsProductTypeParams {
	o.SetSellerID(sellerID)
	return o
}

// SetSellerID adds the sellerId to the get definitions product type params
func (o *GetDefinitionsProductTypeParams) SetSellerID(sellerID *string) {
	o.SellerID = sellerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefinitionsProductTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Locale != nil {

		// query param locale
		var qrLocale string

		if o.Locale != nil {
			qrLocale = *o.Locale
		}
		qLocale := qrLocale
		if qLocale != "" {

			if err := r.SetQueryParam("locale", qLocale); err != nil {
				return err
			}
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

	// path param productType
	if err := r.SetPathParam("productType", o.ProductType); err != nil {
		return err
	}

	if o.ProductTypeVersion != nil {

		// query param productTypeVersion
		var qrProductTypeVersion string

		if o.ProductTypeVersion != nil {
			qrProductTypeVersion = *o.ProductTypeVersion
		}
		qProductTypeVersion := qrProductTypeVersion
		if qProductTypeVersion != "" {

			if err := r.SetQueryParam("productTypeVersion", qProductTypeVersion); err != nil {
				return err
			}
		}
	}

	if o.Requirements != nil {

		// query param requirements
		var qrRequirements string

		if o.Requirements != nil {
			qrRequirements = *o.Requirements
		}
		qRequirements := qrRequirements
		if qRequirements != "" {

			if err := r.SetQueryParam("requirements", qRequirements); err != nil {
				return err
			}
		}
	}

	if o.RequirementsEnforced != nil {

		// query param requirementsEnforced
		var qrRequirementsEnforced string

		if o.RequirementsEnforced != nil {
			qrRequirementsEnforced = *o.RequirementsEnforced
		}
		qRequirementsEnforced := qrRequirementsEnforced
		if qRequirementsEnforced != "" {

			if err := r.SetQueryParam("requirementsEnforced", qRequirementsEnforced); err != nil {
				return err
			}
		}
	}

	if o.SellerID != nil {

		// query param sellerId
		var qrSellerID string

		if o.SellerID != nil {
			qrSellerID = *o.SellerID
		}
		qSellerID := qrSellerID
		if qSellerID != "" {

			if err := r.SetQueryParam("sellerId", qSellerID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDefinitionsProductType binds the parameter marketplaceIds
func (o *GetDefinitionsProductTypeParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
	marketplaceIdsIR := o.MarketplaceIds

	var marketplaceIdsIC []string
	for _, marketplaceIdsIIR := range marketplaceIdsIR { // explode []string

		marketplaceIdsIIV := marketplaceIdsIIR // string as string
		marketplaceIdsIC = append(marketplaceIdsIC, marketplaceIdsIIV)
	}

	// items.CollectionFormat: "csv"
	marketplaceIdsIS := swag.JoinByFormat(marketplaceIdsIC, "csv")

	return marketplaceIdsIS
}
