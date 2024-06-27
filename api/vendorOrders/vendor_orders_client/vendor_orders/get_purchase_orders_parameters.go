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

// NewGetPurchaseOrdersParams creates a new GetPurchaseOrdersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPurchaseOrdersParams() *GetPurchaseOrdersParams {
	return &GetPurchaseOrdersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPurchaseOrdersParamsWithTimeout creates a new GetPurchaseOrdersParams object
// with the ability to set a timeout on a request.
func NewGetPurchaseOrdersParamsWithTimeout(timeout time.Duration) *GetPurchaseOrdersParams {
	return &GetPurchaseOrdersParams{
		timeout: timeout,
	}
}

// NewGetPurchaseOrdersParamsWithContext creates a new GetPurchaseOrdersParams object
// with the ability to set a context for a request.
func NewGetPurchaseOrdersParamsWithContext(ctx context.Context) *GetPurchaseOrdersParams {
	return &GetPurchaseOrdersParams{
		Context: ctx,
	}
}

// NewGetPurchaseOrdersParamsWithHTTPClient creates a new GetPurchaseOrdersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPurchaseOrdersParamsWithHTTPClient(client *http.Client) *GetPurchaseOrdersParams {
	return &GetPurchaseOrdersParams{
		HTTPClient: client,
	}
}

/*
GetPurchaseOrdersParams contains all the parameters to send to the API endpoint

	for the get purchase orders operation.

	Typically these are written to a http.Request.
*/
type GetPurchaseOrdersParams struct {

	/* ChangedAfter.

	   Purchase orders that changed after this timestamp will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.

	   Format: date-time
	*/
	ChangedAfter *strfmt.DateTime

	/* ChangedBefore.

	   Purchase orders that changed before this timestamp will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.

	   Format: date-time
	*/
	ChangedBefore *strfmt.DateTime

	/* CreatedAfter.

	   Purchase orders that became available after this time will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.

	   Format: date-time
	*/
	CreatedAfter *strfmt.DateTime

	/* CreatedBefore.

	   Purchase orders that became available before this time will be included in the result. Must be in <a href='https://developer-docs.amazon.com/sp-api/docs/iso-8601'>ISO 8601</a> date/time format.

	   Format: date-time
	*/
	CreatedBefore *strfmt.DateTime

	/* IncludeDetails.

	   When `true`, returns purchase orders with complete details. Otherwise, only purchase order numbers are returned. Default value is `true`.

	   Format: boolean
	*/
	IncludeDetails *string

	/* IsPOChanged.

	   When `true`, returns purchase orders which were modified after the order was placed. Vendors are required to pull the changed purchase order and fulfill the updated purchase order and not the original one. Default value is `false`.

	   Format: boolean
	*/
	IsPOChanged *string

	/* Limit.

	   The limit to the number of records returned. Default value is 100 records.

	   Format: int64
	*/
	Limit *int64

	/* NextToken.

	   Used for pagination when there is more purchase orders than the specified result size limit. The token value is returned in the previous API call
	*/
	NextToken *string

	/* OrderingVendorCode.

	   Filters purchase orders based on the specified ordering vendor code. This value should be same as 'sellingParty.partyId' in the purchase order. If not included in the filter, all purchase orders for all of the vendor codes that exist in the vendor group used to authorize the API client application are returned.
	*/
	OrderingVendorCode *string

	/* PoItemState.

	   Current state of the purchase order item. If this value is `Cancelled`, this API will return purchase orders which have one or more items cancelled by Amazon with updated item quantity as zero.
	*/
	PoItemState *string

	/* PurchaseOrderState.

	   Filters purchase orders based on the purchase order state.
	*/
	PurchaseOrderState *string

	/* SortOrder.

	   Sort in ascending or descending order by purchase order creation date.
	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get purchase orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPurchaseOrdersParams) WithDefaults() *GetPurchaseOrdersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get purchase orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPurchaseOrdersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithTimeout(timeout time.Duration) *GetPurchaseOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithContext(ctx context.Context) *GetPurchaseOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithHTTPClient(client *http.Client) *GetPurchaseOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangedAfter adds the changedAfter to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithChangedAfter(changedAfter *strfmt.DateTime) *GetPurchaseOrdersParams {
	o.SetChangedAfter(changedAfter)
	return o
}

// SetChangedAfter adds the changedAfter to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetChangedAfter(changedAfter *strfmt.DateTime) {
	o.ChangedAfter = changedAfter
}

// WithChangedBefore adds the changedBefore to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithChangedBefore(changedBefore *strfmt.DateTime) *GetPurchaseOrdersParams {
	o.SetChangedBefore(changedBefore)
	return o
}

// SetChangedBefore adds the changedBefore to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetChangedBefore(changedBefore *strfmt.DateTime) {
	o.ChangedBefore = changedBefore
}

// WithCreatedAfter adds the createdAfter to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithCreatedAfter(createdAfter *strfmt.DateTime) *GetPurchaseOrdersParams {
	o.SetCreatedAfter(createdAfter)
	return o
}

// SetCreatedAfter adds the createdAfter to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetCreatedAfter(createdAfter *strfmt.DateTime) {
	o.CreatedAfter = createdAfter
}

// WithCreatedBefore adds the createdBefore to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithCreatedBefore(createdBefore *strfmt.DateTime) *GetPurchaseOrdersParams {
	o.SetCreatedBefore(createdBefore)
	return o
}

// SetCreatedBefore adds the createdBefore to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetCreatedBefore(createdBefore *strfmt.DateTime) {
	o.CreatedBefore = createdBefore
}

// WithIncludeDetails adds the includeDetails to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithIncludeDetails(includeDetails *string) *GetPurchaseOrdersParams {
	o.SetIncludeDetails(includeDetails)
	return o
}

// SetIncludeDetails adds the includeDetails to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetIncludeDetails(includeDetails *string) {
	o.IncludeDetails = includeDetails
}

// WithIsPOChanged adds the isPOChanged to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithIsPOChanged(isPOChanged *string) *GetPurchaseOrdersParams {
	o.SetIsPOChanged(isPOChanged)
	return o
}

// SetIsPOChanged adds the isPOChanged to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetIsPOChanged(isPOChanged *string) {
	o.IsPOChanged = isPOChanged
}

// WithLimit adds the limit to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithLimit(limit *int64) *GetPurchaseOrdersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithNextToken adds the nextToken to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithNextToken(nextToken *string) *GetPurchaseOrdersParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithOrderingVendorCode adds the orderingVendorCode to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithOrderingVendorCode(orderingVendorCode *string) *GetPurchaseOrdersParams {
	o.SetOrderingVendorCode(orderingVendorCode)
	return o
}

// SetOrderingVendorCode adds the orderingVendorCode to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetOrderingVendorCode(orderingVendorCode *string) {
	o.OrderingVendorCode = orderingVendorCode
}

// WithPoItemState adds the poItemState to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithPoItemState(poItemState *string) *GetPurchaseOrdersParams {
	o.SetPoItemState(poItemState)
	return o
}

// SetPoItemState adds the poItemState to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetPoItemState(poItemState *string) {
	o.PoItemState = poItemState
}

// WithPurchaseOrderState adds the purchaseOrderState to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithPurchaseOrderState(purchaseOrderState *string) *GetPurchaseOrdersParams {
	o.SetPurchaseOrderState(purchaseOrderState)
	return o
}

// SetPurchaseOrderState adds the purchaseOrderState to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetPurchaseOrderState(purchaseOrderState *string) {
	o.PurchaseOrderState = purchaseOrderState
}

// WithSortOrder adds the sortOrder to the get purchase orders params
func (o *GetPurchaseOrdersParams) WithSortOrder(sortOrder *string) *GetPurchaseOrdersParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get purchase orders params
func (o *GetPurchaseOrdersParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetPurchaseOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangedAfter != nil {

		// query param changedAfter
		var qrChangedAfter strfmt.DateTime

		if o.ChangedAfter != nil {
			qrChangedAfter = *o.ChangedAfter
		}
		qChangedAfter := qrChangedAfter.String()
		if qChangedAfter != "" {

			if err := r.SetQueryParam("changedAfter", qChangedAfter); err != nil {
				return err
			}
		}
	}

	if o.ChangedBefore != nil {

		// query param changedBefore
		var qrChangedBefore strfmt.DateTime

		if o.ChangedBefore != nil {
			qrChangedBefore = *o.ChangedBefore
		}
		qChangedBefore := qrChangedBefore.String()
		if qChangedBefore != "" {

			if err := r.SetQueryParam("changedBefore", qChangedBefore); err != nil {
				return err
			}
		}
	}

	if o.CreatedAfter != nil {

		// query param createdAfter
		var qrCreatedAfter strfmt.DateTime

		if o.CreatedAfter != nil {
			qrCreatedAfter = *o.CreatedAfter
		}
		qCreatedAfter := qrCreatedAfter.String()
		if qCreatedAfter != "" {

			if err := r.SetQueryParam("createdAfter", qCreatedAfter); err != nil {
				return err
			}
		}
	}

	if o.CreatedBefore != nil {

		// query param createdBefore
		var qrCreatedBefore strfmt.DateTime

		if o.CreatedBefore != nil {
			qrCreatedBefore = *o.CreatedBefore
		}
		qCreatedBefore := qrCreatedBefore.String()
		if qCreatedBefore != "" {

			if err := r.SetQueryParam("createdBefore", qCreatedBefore); err != nil {
				return err
			}
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

	if o.IsPOChanged != nil {

		// query param isPOChanged
		var qrIsPOChanged string

		if o.IsPOChanged != nil {
			qrIsPOChanged = *o.IsPOChanged
		}
		qIsPOChanged := qrIsPOChanged
		if qIsPOChanged != "" {

			if err := r.SetQueryParam("isPOChanged", qIsPOChanged); err != nil {
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

	if o.OrderingVendorCode != nil {

		// query param orderingVendorCode
		var qrOrderingVendorCode string

		if o.OrderingVendorCode != nil {
			qrOrderingVendorCode = *o.OrderingVendorCode
		}
		qOrderingVendorCode := qrOrderingVendorCode
		if qOrderingVendorCode != "" {

			if err := r.SetQueryParam("orderingVendorCode", qOrderingVendorCode); err != nil {
				return err
			}
		}
	}

	if o.PoItemState != nil {

		// query param poItemState
		var qrPoItemState string

		if o.PoItemState != nil {
			qrPoItemState = *o.PoItemState
		}
		qPoItemState := qrPoItemState
		if qPoItemState != "" {

			if err := r.SetQueryParam("poItemState", qPoItemState); err != nil {
				return err
			}
		}
	}

	if o.PurchaseOrderState != nil {

		// query param purchaseOrderState
		var qrPurchaseOrderState string

		if o.PurchaseOrderState != nil {
			qrPurchaseOrderState = *o.PurchaseOrderState
		}
		qPurchaseOrderState := qrPurchaseOrderState
		if qPurchaseOrderState != "" {

			if err := r.SetQueryParam("purchaseOrderState", qPurchaseOrderState); err != nil {
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
