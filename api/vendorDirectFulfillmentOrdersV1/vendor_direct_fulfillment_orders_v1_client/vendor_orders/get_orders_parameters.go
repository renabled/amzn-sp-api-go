// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders

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

// NewGetOrdersParams creates a new GetOrdersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrdersParams() *GetOrdersParams {
	return &GetOrdersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrdersParamsWithTimeout creates a new GetOrdersParams object
// with the ability to set a timeout on a request.
func NewGetOrdersParamsWithTimeout(timeout time.Duration) *GetOrdersParams {
	return &GetOrdersParams{
		timeout: timeout,
	}
}

// NewGetOrdersParamsWithContext creates a new GetOrdersParams object
// with the ability to set a context for a request.
func NewGetOrdersParamsWithContext(ctx context.Context) *GetOrdersParams {
	return &GetOrdersParams{
		Context: ctx,
	}
}

// NewGetOrdersParamsWithHTTPClient creates a new GetOrdersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrdersParamsWithHTTPClient(client *http.Client) *GetOrdersParams {
	return &GetOrdersParams{
		HTTPClient: client,
	}
}

/*
GetOrdersParams contains all the parameters to send to the API endpoint

	for the get orders operation.

	Typically these are written to a http.Request.
*/
type GetOrdersParams struct {

	/* CreatedAfter.

	   Purchase orders that became available after this date and time will be included in the result. Must be in ISO-8601 date/time format.

	   Format: date-time
	*/
	CreatedAfter strfmt.DateTime

	/* CreatedBefore.

	   Purchase orders that became available before this date and time will be included in the result. Must be in ISO-8601 date/time format.

	   Format: date-time
	*/
	CreatedBefore strfmt.DateTime

	/* IncludeDetails.

	   When true, returns the complete purchase order details. Otherwise, only purchase order numbers are returned.

	   Format: boolean
	   Default: "true"
	*/
	IncludeDetails *string

	/* Limit.

	   The limit to the number of purchase orders returned.

	   Format: int64
	*/
	Limit *int64

	/* NextToken.

	   Used for pagination when there are more orders than the specified result size limit. The token value is returned in the previous API call.
	*/
	NextToken *string

	/* ShipFromPartyID.

	   The vendor warehouse identifier for the fulfillment warehouse. If not specified, the result will contain orders for all warehouses.
	*/
	ShipFromPartyID *string

	/* SortOrder.

	   Sort the list in ascending or descending order by order creation date.
	*/
	SortOrder *string

	/* Status.

	   Returns only the purchase orders that match the specified status. If not specified, the result will contain orders that match any status.
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrdersParams) WithDefaults() *GetOrdersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrdersParams) SetDefaults() {
	var (
		includeDetailsDefault = string("true")
	)

	val := GetOrdersParams{
		IncludeDetails: &includeDetailsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get orders params
func (o *GetOrdersParams) WithTimeout(timeout time.Duration) *GetOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orders params
func (o *GetOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orders params
func (o *GetOrdersParams) WithContext(ctx context.Context) *GetOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orders params
func (o *GetOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orders params
func (o *GetOrdersParams) WithHTTPClient(client *http.Client) *GetOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orders params
func (o *GetOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAfter adds the createdAfter to the get orders params
func (o *GetOrdersParams) WithCreatedAfter(createdAfter strfmt.DateTime) *GetOrdersParams {
	o.SetCreatedAfter(createdAfter)
	return o
}

// SetCreatedAfter adds the createdAfter to the get orders params
func (o *GetOrdersParams) SetCreatedAfter(createdAfter strfmt.DateTime) {
	o.CreatedAfter = createdAfter
}

// WithCreatedBefore adds the createdBefore to the get orders params
func (o *GetOrdersParams) WithCreatedBefore(createdBefore strfmt.DateTime) *GetOrdersParams {
	o.SetCreatedBefore(createdBefore)
	return o
}

// SetCreatedBefore adds the createdBefore to the get orders params
func (o *GetOrdersParams) SetCreatedBefore(createdBefore strfmt.DateTime) {
	o.CreatedBefore = createdBefore
}

// WithIncludeDetails adds the includeDetails to the get orders params
func (o *GetOrdersParams) WithIncludeDetails(includeDetails *string) *GetOrdersParams {
	o.SetIncludeDetails(includeDetails)
	return o
}

// SetIncludeDetails adds the includeDetails to the get orders params
func (o *GetOrdersParams) SetIncludeDetails(includeDetails *string) {
	o.IncludeDetails = includeDetails
}

// WithLimit adds the limit to the get orders params
func (o *GetOrdersParams) WithLimit(limit *int64) *GetOrdersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get orders params
func (o *GetOrdersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNextToken adds the nextToken to the get orders params
func (o *GetOrdersParams) WithNextToken(nextToken *string) *GetOrdersParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get orders params
func (o *GetOrdersParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithShipFromPartyID adds the shipFromPartyID to the get orders params
func (o *GetOrdersParams) WithShipFromPartyID(shipFromPartyID *string) *GetOrdersParams {
	o.SetShipFromPartyID(shipFromPartyID)
	return o
}

// SetShipFromPartyID adds the shipFromPartyId to the get orders params
func (o *GetOrdersParams) SetShipFromPartyID(shipFromPartyID *string) {
	o.ShipFromPartyID = shipFromPartyID
}

// WithSortOrder adds the sortOrder to the get orders params
func (o *GetOrdersParams) WithSortOrder(sortOrder *string) *GetOrdersParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get orders params
func (o *GetOrdersParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStatus adds the status to the get orders params
func (o *GetOrdersParams) WithStatus(status *string) *GetOrdersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get orders params
func (o *GetOrdersParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param createdAfter
	qrCreatedAfter := o.CreatedAfter
	qCreatedAfter := qrCreatedAfter.String()
	if qCreatedAfter != "" {

		if err := r.SetQueryParam("createdAfter", qCreatedAfter); err != nil {
			return err
		}
	}

	// query param createdBefore
	qrCreatedBefore := o.CreatedBefore
	qCreatedBefore := qrCreatedBefore.String()
	if qCreatedBefore != "" {

		if err := r.SetQueryParam("createdBefore", qCreatedBefore); err != nil {
			return err
		}
	}

	if o.IncludeDetails != nil {

		// query param includeDetails
		var qrIncludeDetails string

		if o.IncludeDetails != nil {
			qrIncludeDetails = *o.IncludeDetails
		}
		qIncludeDetails := qrIncludeDetails
		if qIncludeDetails != "" {

			if err := r.SetQueryParam("includeDetails", qIncludeDetails); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
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

	if o.ShipFromPartyID != nil {

		// query param shipFromPartyId
		var qrShipFromPartyID string

		if o.ShipFromPartyID != nil {
			qrShipFromPartyID = *o.ShipFromPartyID
		}
		qShipFromPartyID := qrShipFromPartyID
		if qShipFromPartyID != "" {

			if err := r.SetQueryParam("shipFromPartyId", qShipFromPartyID); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}