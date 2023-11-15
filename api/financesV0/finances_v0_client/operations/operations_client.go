// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListFinancialEventGroups(params *ListFinancialEventGroupsParams, opts ...ClientOption) (*ListFinancialEventGroupsOK, error)

	ListFinancialEvents(params *ListFinancialEventsParams, opts ...ClientOption) (*ListFinancialEventsOK, error)

	ListFinancialEventsByGroupID(params *ListFinancialEventsByGroupIDParams, opts ...ClientOption) (*ListFinancialEventsByGroupIDOK, error)

	ListFinancialEventsByOrderID(params *ListFinancialEventsByOrderIDParams, opts ...ClientOption) (*ListFinancialEventsByOrderIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	ListFinancialEventGroups Returns financial event groups for a given date range. It may take up to 48 hours for orders to appear in your financial events.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) ListFinancialEventGroups(params *ListFinancialEventGroupsParams, opts ...ClientOption) (*ListFinancialEventGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFinancialEventGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listFinancialEventGroups",
		Method:             "GET",
		PathPattern:        "/finances/v0/financialEventGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFinancialEventGroupsReader{formats: a.formats},
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
	success, ok := result.(*ListFinancialEventGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listFinancialEventGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ListFinancialEvents Returns financial events for the specified data range. It may take up to 48 hours for orders to appear in your financial events. **Note:** in `ListFinancialEvents`, deferred events don't show up in responses until in they are released.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) ListFinancialEvents(params *ListFinancialEventsParams, opts ...ClientOption) (*ListFinancialEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFinancialEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listFinancialEvents",
		Method:             "GET",
		PathPattern:        "/finances/v0/financialEvents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFinancialEventsReader{formats: a.formats},
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
	success, ok := result.(*ListFinancialEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listFinancialEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ListFinancialEventsByGroupID Returns all financial events for the specified financial event group. It may take up to 48 hours for orders to appear in your financial events.

**Note:** This operation will only retrieve group's data for the past two years. If a request is submitted for data spanning more than two years, an empty response is returned.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) ListFinancialEventsByGroupID(params *ListFinancialEventsByGroupIDParams, opts ...ClientOption) (*ListFinancialEventsByGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFinancialEventsByGroupIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listFinancialEventsByGroupId",
		Method:             "GET",
		PathPattern:        "/finances/v0/financialEventGroups/{eventGroupId}/financialEvents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFinancialEventsByGroupIDReader{formats: a.formats},
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
	success, ok := result.(*ListFinancialEventsByGroupIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listFinancialEventsByGroupId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ListFinancialEventsByOrderID Returns all financial events for the specified order. It may take up to 48 hours for orders to appear in your financial events.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 30 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) ListFinancialEventsByOrderID(params *ListFinancialEventsByOrderIDParams, opts ...ClientOption) (*ListFinancialEventsByOrderIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFinancialEventsByOrderIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listFinancialEventsByOrderId",
		Method:             "GET",
		PathPattern:        "/finances/v0/orders/{orderId}/financialEvents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFinancialEventsByOrderIDReader{formats: a.formats},
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
	success, ok := result.(*ListFinancialEventsByOrderIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listFinancialEventsByOrderId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
