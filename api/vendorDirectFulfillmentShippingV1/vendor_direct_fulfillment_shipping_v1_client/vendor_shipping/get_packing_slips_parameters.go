// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping

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

// NewGetPackingSlipsParams creates a new GetPackingSlipsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPackingSlipsParams() *GetPackingSlipsParams {
	return &GetPackingSlipsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackingSlipsParamsWithTimeout creates a new GetPackingSlipsParams object
// with the ability to set a timeout on a request.
func NewGetPackingSlipsParamsWithTimeout(timeout time.Duration) *GetPackingSlipsParams {
	return &GetPackingSlipsParams{
		timeout: timeout,
	}
}

// NewGetPackingSlipsParamsWithContext creates a new GetPackingSlipsParams object
// with the ability to set a context for a request.
func NewGetPackingSlipsParamsWithContext(ctx context.Context) *GetPackingSlipsParams {
	return &GetPackingSlipsParams{
		Context: ctx,
	}
}

// NewGetPackingSlipsParamsWithHTTPClient creates a new GetPackingSlipsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPackingSlipsParamsWithHTTPClient(client *http.Client) *GetPackingSlipsParams {
	return &GetPackingSlipsParams{
		HTTPClient: client,
	}
}

/*
GetPackingSlipsParams contains all the parameters to send to the API endpoint

	for the get packing slips operation.

	Typically these are written to a http.Request.
*/
type GetPackingSlipsParams struct {

	/* CreatedAfter.

	   Packing slips that became available after this date and time will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.

	   Format: date-time
	*/
	CreatedAfter strfmt.DateTime

	/* CreatedBefore.

	   Packing slips that became available before this date and time will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.

	   Format: date-time
	*/
	CreatedBefore strfmt.DateTime

	/* Limit.

	   The limit to the number of records returned
	*/
	Limit *int64

	/* NextToken.

	   Used for pagination when there are more packing slips than the specified result size limit. The token value is returned in the previous API call.
	*/
	NextToken *string

	/* ShipFromPartyID.

	   The vendor `warehouseId` for order fulfillment. If not specified the result will contain orders for all warehouses.
	*/
	ShipFromPartyID *string

	/* SortOrder.

	   Sort ASC or DESC by packing slip creation date.

	   Default: "ASC"
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get packing slips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackingSlipsParams) WithDefaults() *GetPackingSlipsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get packing slips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackingSlipsParams) SetDefaults() {
	var (
		sortOrderDefault = string("ASC")
	)

	val := GetPackingSlipsParams{
		SortOrder: &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get packing slips params
func (o *GetPackingSlipsParams) WithTimeout(timeout time.Duration) *GetPackingSlipsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get packing slips params
func (o *GetPackingSlipsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get packing slips params
func (o *GetPackingSlipsParams) WithContext(ctx context.Context) *GetPackingSlipsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get packing slips params
func (o *GetPackingSlipsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get packing slips params
func (o *GetPackingSlipsParams) WithHTTPClient(client *http.Client) *GetPackingSlipsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get packing slips params
func (o *GetPackingSlipsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatedAfter adds the createdAfter to the get packing slips params
func (o *GetPackingSlipsParams) WithCreatedAfter(createdAfter strfmt.DateTime) *GetPackingSlipsParams {
	o.SetCreatedAfter(createdAfter)
	return o
}

// SetCreatedAfter adds the createdAfter to the get packing slips params
func (o *GetPackingSlipsParams) SetCreatedAfter(createdAfter strfmt.DateTime) {
	o.CreatedAfter = createdAfter
}

// WithCreatedBefore adds the createdBefore to the get packing slips params
func (o *GetPackingSlipsParams) WithCreatedBefore(createdBefore strfmt.DateTime) *GetPackingSlipsParams {
	o.SetCreatedBefore(createdBefore)
	return o
}

// SetCreatedBefore adds the createdBefore to the get packing slips params
func (o *GetPackingSlipsParams) SetCreatedBefore(createdBefore strfmt.DateTime) {
	o.CreatedBefore = createdBefore
}

// WithLimit adds the limit to the get packing slips params
func (o *GetPackingSlipsParams) WithLimit(limit *int64) *GetPackingSlipsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get packing slips params
func (o *GetPackingSlipsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNextToken adds the nextToken to the get packing slips params
func (o *GetPackingSlipsParams) WithNextToken(nextToken *string) *GetPackingSlipsParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get packing slips params
func (o *GetPackingSlipsParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithShipFromPartyID adds the shipFromPartyID to the get packing slips params
func (o *GetPackingSlipsParams) WithShipFromPartyID(shipFromPartyID *string) *GetPackingSlipsParams {
	o.SetShipFromPartyID(shipFromPartyID)
	return o
}

// SetShipFromPartyID adds the shipFromPartyId to the get packing slips params
func (o *GetPackingSlipsParams) SetShipFromPartyID(shipFromPartyID *string) {
	o.ShipFromPartyID = shipFromPartyID
}

// WithSortOrder adds the sortOrder to the get packing slips params
func (o *GetPackingSlipsParams) WithSortOrder(sortOrder *string) *GetPackingSlipsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get packing slips params
func (o *GetPackingSlipsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackingSlipsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
