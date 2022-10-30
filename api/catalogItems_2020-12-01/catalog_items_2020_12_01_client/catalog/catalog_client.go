// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCatalogItem(params *GetCatalogItemParams, opts ...ClientOption) (*GetCatalogItemOK, error)

	SearchCatalogItems(params *SearchCatalogItemsParams, opts ...ClientOption) (*SearchCatalogItemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetCatalogItem Retrieves details for an item in the Amazon catalog.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetCatalogItem(params *GetCatalogItemParams, opts ...ClientOption) (*GetCatalogItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCatalogItem",
		Method:             "GET",
		PathPattern:        "/catalog/2020-12-01/items/{asin}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCatalogItemReader{formats: a.formats},
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
	success, ok := result.(*GetCatalogItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCatalogItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SearchCatalogItems Search for and return a list of Amazon catalog items and associated information.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) SearchCatalogItems(params *SearchCatalogItemsParams, opts ...ClientOption) (*SearchCatalogItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchCatalogItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchCatalogItems",
		Method:             "GET",
		PathPattern:        "/catalog/2020-12-01/items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchCatalogItemsReader{formats: a.formats},
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
	success, ok := result.(*SearchCatalogItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchCatalogItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
