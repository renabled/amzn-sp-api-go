// Code generated by go-swagger; DO NOT EDIT.

package vendor_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vendor orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vendor orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPurchaseOrder(params *GetPurchaseOrderParams, opts ...ClientOption) (*GetPurchaseOrderOK, error)

	GetPurchaseOrders(params *GetPurchaseOrdersParams, opts ...ClientOption) (*GetPurchaseOrdersOK, error)

	GetPurchaseOrdersStatus(params *GetPurchaseOrdersStatusParams, opts ...ClientOption) (*GetPurchaseOrdersStatusOK, error)

	SubmitAcknowledgement(params *SubmitAcknowledgementParams, opts ...ClientOption) (*SubmitAcknowledgementAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetPurchaseOrder Returns a purchase order based on the purchaseOrderNumber value that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetPurchaseOrder(params *GetPurchaseOrderParams, opts ...ClientOption) (*GetPurchaseOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPurchaseOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPurchaseOrder",
		Method:             "GET",
		PathPattern:        "/vendor/orders/v1/purchaseOrders/{purchaseOrderNumber}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPurchaseOrderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPurchaseOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPurchaseOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetPurchaseOrders Returns a list of purchase orders created or changed during the time frame that you specify. You define the time frame using the createdAfter, createdBefore, changedAfter and changedBefore parameters. The date range to search must not be more than 7 days. You can choose to get only the purchase order numbers by setting includeDetails to false. You can then use the getPurchaseOrder operation to receive details for a specific purchase order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetPurchaseOrders(params *GetPurchaseOrdersParams, opts ...ClientOption) (*GetPurchaseOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPurchaseOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPurchaseOrders",
		Method:             "GET",
		PathPattern:        "/vendor/orders/v1/purchaseOrders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPurchaseOrdersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPurchaseOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPurchaseOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetPurchaseOrdersStatus Returns purchase order statuses based on the filters that you specify. Date range to search must not be more than 7 days. You can return a list of purchase order statuses using the available filters, or a single purchase order status by providing the purchase order number.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetPurchaseOrdersStatus(params *GetPurchaseOrdersStatusParams, opts ...ClientOption) (*GetPurchaseOrdersStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPurchaseOrdersStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPurchaseOrdersStatus",
		Method:             "GET",
		PathPattern:        "/vendor/orders/v1/purchaseOrdersStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPurchaseOrdersStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPurchaseOrdersStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPurchaseOrdersStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SubmitAcknowledgement Submits acknowledgements for one or more purchase orders.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) SubmitAcknowledgement(params *SubmitAcknowledgementParams, opts ...ClientOption) (*SubmitAcknowledgementAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitAcknowledgementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "submitAcknowledgement",
		Method:             "POST",
		PathPattern:        "/vendor/orders/v1/acknowledgements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubmitAcknowledgementReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubmitAcknowledgementAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for submitAcknowledgement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
