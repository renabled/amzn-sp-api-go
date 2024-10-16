// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// NewGetReportsParams creates a new GetReportsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReportsParams() *GetReportsParams {
	return &GetReportsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReportsParamsWithTimeout creates a new GetReportsParams object
// with the ability to set a timeout on a request.
func NewGetReportsParamsWithTimeout(timeout time.Duration) *GetReportsParams {
	return &GetReportsParams{
		timeout: timeout,
	}
}

// NewGetReportsParamsWithContext creates a new GetReportsParams object
// with the ability to set a context for a request.
func NewGetReportsParamsWithContext(ctx context.Context) *GetReportsParams {
	return &GetReportsParams{
		Context: ctx,
	}
}

// NewGetReportsParamsWithHTTPClient creates a new GetReportsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReportsParamsWithHTTPClient(client *http.Client) *GetReportsParams {
	return &GetReportsParams{
		HTTPClient: client,
	}
}

/*
GetReportsParams contains all the parameters to send to the API endpoint

	for the get reports operation.

	Typically these are written to a http.Request.
*/
type GetReportsParams struct {

	/* CreatedSince.

	   The earliest report creation date and time for reports to include in the response, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. The default is 90 days ago. Reports are retained for a maximum of 90 days.

	   Format: date-time
	*/
	CreatedSince *strfmt.DateTime

	/* CreatedUntil.

	   The latest report creation date and time for reports to include in the response, in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date time format. The default is now.

	   Format: date-time
	*/
	CreatedUntil *strfmt.DateTime

	/* MarketplaceIds.

	   A list of marketplace identifiers used to filter reports. The reports returned will match at least one of the marketplaces that you specify.
	*/
	MarketplaceIds []string

	/* NextToken.

	   A string token returned in the response to your previous request. `nextToken` is returned when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the `getReports` operation and include this token as the only parameter. Specifying `nextToken` with any other parameters will cause the request to fail.
	*/
	NextToken *string

	/* PageSize.

	   The maximum number of reports to return in a single call.

	   Default: 10
	*/
	PageSize *int64

	/* ProcessingStatuses.

	   A list of processing statuses used to filter reports.
	*/
	ProcessingStatuses []string

	/* ReportTypes.

	   A list of report types used to filter reports. Refer to [Report Type Values](https://developer-docs.amazon.com/sp-api/docs/report-type-values) for more information. When reportTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either reportTypes or nextToken is required.
	*/
	ReportTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportsParams) WithDefaults() *GetReportsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get reports params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReportsParams) SetDefaults() {
	var (
		pageSizeDefault = int64(10)
	)

	val := GetReportsParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get reports params
func (o *GetReportsParams) WithTimeout(timeout time.Duration) *GetReportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get reports params
func (o *GetReportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get reports params
func (o *GetReportsParams) WithContext(ctx context.Context) *GetReportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get reports params
func (o *GetReportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get reports params
func (o *GetReportsParams) WithHTTPClient(client *http.Client) *GetReportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get reports params
func (o *GetReportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedSince adds the createdSince to the get reports params
func (o *GetReportsParams) WithCreatedSince(createdSince *strfmt.DateTime) *GetReportsParams {
	o.SetCreatedSince(createdSince)
	return o
}

// SetCreatedSince adds the createdSince to the get reports params
func (o *GetReportsParams) SetCreatedSince(createdSince *strfmt.DateTime) {
	o.CreatedSince = createdSince
}

// WithCreatedUntil adds the createdUntil to the get reports params
func (o *GetReportsParams) WithCreatedUntil(createdUntil *strfmt.DateTime) *GetReportsParams {
	o.SetCreatedUntil(createdUntil)
	return o
}

// SetCreatedUntil adds the createdUntil to the get reports params
func (o *GetReportsParams) SetCreatedUntil(createdUntil *strfmt.DateTime) {
	o.CreatedUntil = createdUntil
}

// WithMarketplaceIds adds the marketplaceIds to the get reports params
func (o *GetReportsParams) WithMarketplaceIds(marketplaceIds []string) *GetReportsParams {
	o.SetMarketplaceIds(marketplaceIds)
	return o
}

// SetMarketplaceIds adds the marketplaceIds to the get reports params
func (o *GetReportsParams) SetMarketplaceIds(marketplaceIds []string) {
	o.MarketplaceIds = marketplaceIds
}

// WithNextToken adds the nextToken to the get reports params
func (o *GetReportsParams) WithNextToken(nextToken *string) *GetReportsParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get reports params
func (o *GetReportsParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithPageSize adds the pageSize to the get reports params
func (o *GetReportsParams) WithPageSize(pageSize *int64) *GetReportsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get reports params
func (o *GetReportsParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithProcessingStatuses adds the processingStatuses to the get reports params
func (o *GetReportsParams) WithProcessingStatuses(processingStatuses []string) *GetReportsParams {
	o.SetProcessingStatuses(processingStatuses)
	return o
}

// SetProcessingStatuses adds the processingStatuses to the get reports params
func (o *GetReportsParams) SetProcessingStatuses(processingStatuses []string) {
	o.ProcessingStatuses = processingStatuses
}

// WithReportTypes adds the reportTypes to the get reports params
func (o *GetReportsParams) WithReportTypes(reportTypes []string) *GetReportsParams {
	o.SetReportTypes(reportTypes)
	return o
}

// SetReportTypes adds the reportTypes to the get reports params
func (o *GetReportsParams) SetReportTypes(reportTypes []string) {
	o.ReportTypes = reportTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetReportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreatedSince != nil {

		// query param createdSince
		var qrCreatedSince strfmt.DateTime

		if o.CreatedSince != nil {
			qrCreatedSince = *o.CreatedSince
		}
		qCreatedSince := qrCreatedSince.String()
		if qCreatedSince != "" {

			if err := r.SetQueryParam("createdSince", qCreatedSince); err != nil {
				return err
			}
		}
	}

	if o.CreatedUntil != nil {

		// query param createdUntil
		var qrCreatedUntil strfmt.DateTime

		if o.CreatedUntil != nil {
			qrCreatedUntil = *o.CreatedUntil
		}
		qCreatedUntil := qrCreatedUntil.String()
		if qCreatedUntil != "" {

			if err := r.SetQueryParam("createdUntil", qCreatedUntil); err != nil {
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

	if o.NextToken != nil {

		// query param nextToken
		var qrNextToken string

		if o.NextToken != nil {
			qrNextToken = *o.NextToken
		}
		qNextToken := qrNextToken
		if qNextToken != "" {

			if err := r.SetQueryParam("nextToken", qNextToken); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.ProcessingStatuses != nil {

		// binding items for processingStatuses
		joinedProcessingStatuses := o.bindParamProcessingStatuses(reg)

		// query array param processingStatuses
		if err := r.SetQueryParam("processingStatuses", joinedProcessingStatuses...); err != nil {
			return err
		}
	}

	if o.ReportTypes != nil {

		// binding items for reportTypes
		joinedReportTypes := o.bindParamReportTypes(reg)

		// query array param reportTypes
		if err := r.SetQueryParam("reportTypes", joinedReportTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetReports binds the parameter marketplaceIds
func (o *GetReportsParams) bindParamMarketplaceIds(formats strfmt.Registry) []string {
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

// bindParamGetReports binds the parameter processingStatuses
func (o *GetReportsParams) bindParamProcessingStatuses(formats strfmt.Registry) []string {
	processingStatusesIR := o.ProcessingStatuses

	var processingStatusesIC []string
	for _, processingStatusesIIR := range processingStatusesIR { // explode []string

		processingStatusesIIV := processingStatusesIIR // string as string
		processingStatusesIC = append(processingStatusesIC, processingStatusesIIV)
	}

	// items.CollectionFormat: ""
	processingStatusesIS := swag.JoinByFormat(processingStatusesIC, "")

	return processingStatusesIS
}

// bindParamGetReports binds the parameter reportTypes
func (o *GetReportsParams) bindParamReportTypes(formats strfmt.Registry) []string {
	reportTypesIR := o.ReportTypes

	var reportTypesIC []string
	for _, reportTypesIIR := range reportTypesIR { // explode []string

		reportTypesIIV := reportTypesIIR // string as string
		reportTypesIC = append(reportTypesIC, reportTypesIIV)
	}

	// items.CollectionFormat: ""
	reportTypesIS := swag.JoinByFormat(reportTypesIC, "")

	return reportTypesIS
}
