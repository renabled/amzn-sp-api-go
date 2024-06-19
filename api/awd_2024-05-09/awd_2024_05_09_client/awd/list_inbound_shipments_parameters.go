// Code generated by go-swagger; DO NOT EDIT.

package awd

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

// NewListInboundShipmentsParams creates a new ListInboundShipmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListInboundShipmentsParams() *ListInboundShipmentsParams {
	return &ListInboundShipmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListInboundShipmentsParamsWithTimeout creates a new ListInboundShipmentsParams object
// with the ability to set a timeout on a request.
func NewListInboundShipmentsParamsWithTimeout(timeout time.Duration) *ListInboundShipmentsParams {
	return &ListInboundShipmentsParams{
		timeout: timeout,
	}
}

// NewListInboundShipmentsParamsWithContext creates a new ListInboundShipmentsParams object
// with the ability to set a context for a request.
func NewListInboundShipmentsParamsWithContext(ctx context.Context) *ListInboundShipmentsParams {
	return &ListInboundShipmentsParams{
		Context: ctx,
	}
}

// NewListInboundShipmentsParamsWithHTTPClient creates a new ListInboundShipmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListInboundShipmentsParamsWithHTTPClient(client *http.Client) *ListInboundShipmentsParams {
	return &ListInboundShipmentsParams{
		HTTPClient: client,
	}
}

/*
ListInboundShipmentsParams contains all the parameters to send to the API endpoint

	for the list inbound shipments operation.

	Typically these are written to a http.Request.
*/
type ListInboundShipmentsParams struct {

	/* MaxResults.

	   Maximum number of results to return.

	   Format: int32
	   Default: 25
	*/
	MaxResults *int32

	/* NextToken.

	   Token to retrieve the next set of paginated results.
	*/
	NextToken *string

	/* ShipmentStatus.

	   Filter by inbound shipment status.
	*/
	ShipmentStatus *string

	/* SortBy.

	   Field to sort results by. Required if `sortOrder` is provided.
	*/
	SortBy *string

	/* SortOrder.

	   Sort the response in `ASCENDING` or `DESCENDING` order.
	*/
	SortOrder *string

	/* UpdatedAfter.

	   List the inbound shipments that were updated after a certain time (inclusive). The date must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.

	   Format: date-time
	*/
	UpdatedAfter *strfmt.DateTime

	/* UpdatedBefore.

	   List the inbound shipments that were updated before a certain time (inclusive). The date must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> format.

	   Format: date-time
	*/
	UpdatedBefore *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list inbound shipments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInboundShipmentsParams) WithDefaults() *ListInboundShipmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list inbound shipments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListInboundShipmentsParams) SetDefaults() {
	var (
		maxResultsDefault = int32(25)
	)

	val := ListInboundShipmentsParams{
		MaxResults: &maxResultsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithTimeout(timeout time.Duration) *ListInboundShipmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithContext(ctx context.Context) *ListInboundShipmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithHTTPClient(client *http.Client) *ListInboundShipmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaxResults adds the maxResults to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithMaxResults(maxResults *int32) *ListInboundShipmentsParams {
	o.SetMaxResults(maxResults)
	return o
}

// SetMaxResults adds the maxResults to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetMaxResults(maxResults *int32) {
	o.MaxResults = maxResults
}

// WithNextToken adds the nextToken to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithNextToken(nextToken *string) *ListInboundShipmentsParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithShipmentStatus adds the shipmentStatus to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithShipmentStatus(shipmentStatus *string) *ListInboundShipmentsParams {
	o.SetShipmentStatus(shipmentStatus)
	return o
}

// SetShipmentStatus adds the shipmentStatus to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetShipmentStatus(shipmentStatus *string) {
	o.ShipmentStatus = shipmentStatus
}

// WithSortBy adds the sortBy to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithSortBy(sortBy *string) *ListInboundShipmentsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithSortOrder(sortOrder *string) *ListInboundShipmentsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithUpdatedAfter adds the updatedAfter to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithUpdatedAfter(updatedAfter *strfmt.DateTime) *ListInboundShipmentsParams {
	o.SetUpdatedAfter(updatedAfter)
	return o
}

// SetUpdatedAfter adds the updatedAfter to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetUpdatedAfter(updatedAfter *strfmt.DateTime) {
	o.UpdatedAfter = updatedAfter
}

// WithUpdatedBefore adds the updatedBefore to the list inbound shipments params
func (o *ListInboundShipmentsParams) WithUpdatedBefore(updatedBefore *strfmt.DateTime) *ListInboundShipmentsParams {
	o.SetUpdatedBefore(updatedBefore)
	return o
}

// SetUpdatedBefore adds the updatedBefore to the list inbound shipments params
func (o *ListInboundShipmentsParams) SetUpdatedBefore(updatedBefore *strfmt.DateTime) {
	o.UpdatedBefore = updatedBefore
}

// WriteToRequest writes these params to a swagger request
func (o *ListInboundShipmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MaxResults != nil {

		// query param maxResults
		var qrMaxResults int32

		if o.MaxResults != nil {
			qrMaxResults = *o.MaxResults
		}
		qMaxResults := swag.FormatInt32(qrMaxResults)
		if qMaxResults != "" {

			if err := r.SetQueryParam("maxResults", qMaxResults); err != nil {
				return err
			}
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

	if o.ShipmentStatus != nil {

		// query param shipmentStatus
		var qrShipmentStatus string

		if o.ShipmentStatus != nil {
			qrShipmentStatus = *o.ShipmentStatus
		}
		qShipmentStatus := qrShipmentStatus
		if qShipmentStatus != "" {

			if err := r.SetQueryParam("shipmentStatus", qShipmentStatus); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}
	}

	if o.UpdatedAfter != nil {

		// query param updatedAfter
		var qrUpdatedAfter strfmt.DateTime

		if o.UpdatedAfter != nil {
			qrUpdatedAfter = *o.UpdatedAfter
		}
		qUpdatedAfter := qrUpdatedAfter.String()
		if qUpdatedAfter != "" {

			if err := r.SetQueryParam("updatedAfter", qUpdatedAfter); err != nil {
				return err
			}
		}
	}

	if o.UpdatedBefore != nil {

		// query param updatedBefore
		var qrUpdatedBefore strfmt.DateTime

		if o.UpdatedBefore != nil {
			qrUpdatedBefore = *o.UpdatedBefore
		}
		qUpdatedBefore := qrUpdatedBefore.String()
		if qUpdatedBefore != "" {

			if err := r.SetQueryParam("updatedBefore", qUpdatedBefore); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
