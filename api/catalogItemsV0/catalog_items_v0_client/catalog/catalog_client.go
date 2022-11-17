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

	ListCatalogCategories(params *ListCatalogCategoriesParams, opts ...ClientOption) (*ListCatalogCategoriesOK, error)

	ListCatalogItems(params *ListCatalogItemsParams, opts ...ClientOption) (*ListCatalogItemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetCatalogItem Effective September 30, 2022, the `getCatalogItem` operation will no longer be available in the Selling Partner API for Catalog Items v0. This operation is available in the latest version of the [Selling Partner API for Catalog Items v2022-04-01](doc:catalog-items-api-v2022-04-01-reference). Integrations that rely on this operation should migrate to the latest version to avoid service disruption.

_Note:_ The [`listCatalogCategories`](#get-catalogv0categories) operation is not being deprecated and you can continue to make calls to it.
*/
func (a *Client) GetCatalogItem(params *GetCatalogItemParams, opts ...ClientOption) (*GetCatalogItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCatalogItem",
		Method:             "GET",
		PathPattern:        "/catalog/v0/items/{asin}",
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
	ListCatalogCategories Returns the parent categories to which an item belongs, based on the specified ASIN or SellerSKU.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) ListCatalogCategories(params *ListCatalogCategoriesParams, opts ...ClientOption) (*ListCatalogCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogCategoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listCatalogCategories",
		Method:             "GET",
		PathPattern:        "/catalog/v0/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCatalogCategoriesReader{formats: a.formats},
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
	success, ok := result.(*ListCatalogCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listCatalogCategories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ListCatalogItems Effective September 30, 2022, the `listCatalogItems` operation will no longer be available in the Selling Partner API for Catalog Items v0. As an alternative, `searchCatalogItems` is available in the latest version of the [Selling Partner API for Catalog Items v2022-04-01](doc:catalog-items-api-v2022-04-01-reference). Integrations that rely on the `listCatalogItems` operation should migrate to the `searchCatalogItems`operation to avoid service disruption.

_Note:_ The [`listCatalogCategories`](#get-catalogv0categories) operation is not being deprecated and you can continue to make calls to it.
*/
func (a *Client) ListCatalogItems(params *ListCatalogItemsParams, opts ...ClientOption) (*ListCatalogItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCatalogItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listCatalogItems",
		Method:             "GET",
		PathPattern:        "/catalog/v0/items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListCatalogItemsReader{formats: a.formats},
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
	success, ok := result.(*ListCatalogItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listCatalogItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
