// Code generated by go-swagger; DO NOT EDIT.

package queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new queries API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for queries API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelQuery(params *CancelQueryParams, opts ...ClientOption) (*CancelQueryNoContent, error)

	CreateQuery(params *CreateQueryParams, opts ...ClientOption) (*CreateQueryAccepted, error)

	GetDocument(params *GetDocumentParams, opts ...ClientOption) (*GetDocumentOK, error)

	GetQueries(params *GetQueriesParams, opts ...ClientOption) (*GetQueriesOK, error)

	GetQuery(params *GetQueryParams, opts ...ClientOption) (*GetQueryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CancelQuery Cancels the query specified by the `queryId` parameter. Only queries with a non-terminal `processingStatus` (`IN_QUEUE`, `IN_PROGRESS`) can be cancelled. Cancelling a query that already has a `processingStatus` of `CANCELLED` will no-op. Cancelled queries are returned in subsequent calls to the `getQuery` and `getQueries` operations.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CancelQuery(params *CancelQueryParams, opts ...ClientOption) (*CancelQueryNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelQuery",
		Method:             "DELETE",
		PathPattern:        "/dataKiosk/2023-11-15/queries/{queryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelQueryReader{formats: a.formats},
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
	success, ok := result.(*CancelQueryNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cancelQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateQuery Creates a Data Kiosk query request.

**Note:** The retention of a query varies based on the fields requested. Each field within a schema is annotated with a `@resultRetention` directive that defines how long a query containing that field will be retained. When a query contains multiple fields with different retentions, the shortest (minimum) retention is applied. The retention of a query's resulting documents always matches the retention of the query.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0167 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateQuery(params *CreateQueryParams, opts ...ClientOption) (*CreateQueryAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createQuery",
		Method:             "POST",
		PathPattern:        "/dataKiosk/2023-11-15/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateQueryReader{formats: a.formats},
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
	success, ok := result.(*CreateQueryAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetDocument Returns the information required for retrieving a Data Kiosk document's contents. See the `createQuery` operation for details about document retention.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0167 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetDocument(params *GetDocumentParams, opts ...ClientOption) (*GetDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDocument",
		Method:             "GET",
		PathPattern:        "/dataKiosk/2023-11-15/documents/{documentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDocumentReader{formats: a.formats},
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
	success, ok := result.(*GetDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDocument: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetQueries Returns details for the Data Kiosk queries that match the specified filters. See the `createQuery` operation for details about query retention.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetQueries(params *GetQueriesParams, opts ...ClientOption) (*GetQueriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getQueries",
		Method:             "GET",
		PathPattern:        "/dataKiosk/2023-11-15/queries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQueriesReader{formats: a.formats},
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
	success, ok := result.(*GetQueriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getQueries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetQuery Returns query details for the query specified by the `queryId` parameter. See the `createQuery` operation for details about query retention.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2.0 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetQuery(params *GetQueryParams, opts ...ClientOption) (*GetQueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getQuery",
		Method:             "GET",
		PathPattern:        "/dataKiosk/2023-11-15/queries/{queryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQueryReader{formats: a.formats},
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
	success, ok := result.(*GetQueryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getQuery: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
