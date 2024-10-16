// Code generated by go-swagger; DO NOT EDIT.

package listings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new listings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for listings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteListingsItem(params *DeleteListingsItemParams, opts ...ClientOption) (*DeleteListingsItemOK, error)

	GetListingsItem(params *GetListingsItemParams, opts ...ClientOption) (*GetListingsItemOK, error)

	PatchListingsItem(params *PatchListingsItemParams, opts ...ClientOption) (*PatchListingsItemOK, error)

	PutListingsItem(params *PutListingsItemParams, opts ...ClientOption) (*PutListingsItemOK, error)

	SearchListingsItems(params *SearchListingsItemsParams, opts ...ClientOption) (*SearchListingsItemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	DeleteListingsItem Delete a listings item for a selling partner.

**Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) DeleteListingsItem(params *DeleteListingsItemParams, opts ...ClientOption) (*DeleteListingsItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteListingsItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteListingsItem",
		Method:             "DELETE",
		PathPattern:        "/listings/2021-08-01/items/{sellerId}/{sku}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteListingsItemReader{formats: a.formats},
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
	success, ok := result.(*DeleteListingsItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteListingsItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetListingsItem Returns details about a listings item for a selling partner.

**Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetListingsItem(params *GetListingsItemParams, opts ...ClientOption) (*GetListingsItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetListingsItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getListingsItem",
		Method:             "GET",
		PathPattern:        "/listings/2021-08-01/items/{sellerId}/{sku}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetListingsItemReader{formats: a.formats},
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
	success, ok := result.(*GetListingsItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getListingsItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PatchListingsItem Partially update (patch) a listings item for a selling partner. Only top-level listings item attributes can be patched. Patching nested attributes is not supported.

**Note:** This operation has a throttling rate of one request per second when `mode` is `VALIDATION_PREVIEW`.

**Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) PatchListingsItem(params *PatchListingsItemParams, opts ...ClientOption) (*PatchListingsItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchListingsItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchListingsItem",
		Method:             "PATCH",
		PathPattern:        "/listings/2021-08-01/items/{sellerId}/{sku}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchListingsItemReader{formats: a.formats},
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
	success, ok := result.(*PatchListingsItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchListingsItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PutListingsItem Creates or fully updates an existing listings item for a selling partner.

**Note:** This operation has a throttling rate of one request per second when `mode` is `VALIDATION_PREVIEW`.

**Note:** The parameters associated with this operation may contain special characters that must be encoded to successfully call the API. To avoid errors with SKUs when encoding URLs, refer to [URL Encoding](https://developer-docs.amazon.com/sp-api/docs/url-encoding).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) PutListingsItem(params *PutListingsItemParams, opts ...ClientOption) (*PutListingsItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutListingsItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putListingsItem",
		Method:             "PUT",
		PathPattern:        "/listings/2021-08-01/items/{sellerId}/{sku}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutListingsItemReader{formats: a.formats},
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
	success, ok := result.(*PutListingsItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putListingsItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SearchListingsItems Search for and return list of listings items and respective details for a selling partner.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) SearchListingsItems(params *SearchListingsItemsParams, opts ...ClientOption) (*SearchListingsItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchListingsItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "searchListingsItems",
		Method:             "GET",
		PathPattern:        "/listings/2021-08-01/items/{sellerId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SearchListingsItemsReader{formats: a.formats},
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
	success, ok := result.(*SearchListingsItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchListingsItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
