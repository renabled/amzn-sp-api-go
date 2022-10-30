// Code generated by go-swagger; DO NOT EDIT.

package fees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fees API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fees API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetMyFeesEstimateForASIN(params *GetMyFeesEstimateForASINParams, opts ...ClientOption) (*GetMyFeesEstimateForASINOK, error)

	GetMyFeesEstimateForSKU(params *GetMyFeesEstimateForSKUParams, opts ...ClientOption) (*GetMyFeesEstimateForSKUOK, error)

	GetMyFeesEstimates(params *GetMyFeesEstimatesParams, opts ...ClientOption) (*GetMyFeesEstimatesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetMyFeesEstimateForASIN Returns the estimated fees for the item indicated by the specified ASIN in the marketplace specified in the request body.

You can call `getMyFeesEstimateForASIN` for an item on behalf of a selling partner before the selling partner sets the item's price. The selling partner can then take estimated fees into account. Each fees request must include an original identifier. This identifier is included in the fees estimate so you can correlate a fees estimate with the original request.

**Note:** This identifier value is only an estimate, actual costs may vary. Search "fees" in [Seller Central](https://sellercentral.amazon.com/) and consult the store-specific fee schedule for the most up-to-date information.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetMyFeesEstimateForASIN(params *GetMyFeesEstimateForASINParams, opts ...ClientOption) (*GetMyFeesEstimateForASINOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMyFeesEstimateForASINParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMyFeesEstimateForASIN",
		Method:             "POST",
		PathPattern:        "/products/fees/v0/items/{Asin}/feesEstimate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMyFeesEstimateForASINReader{formats: a.formats},
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
	success, ok := result.(*GetMyFeesEstimateForASINOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMyFeesEstimateForASIN: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMyFeesEstimateForSKU Returns the estimated fees for the item indicated by the specified seller SKU in the marketplace specified in the request body.

You can call `getMyFeesEstimateForSKU` for an item on behalf of a selling partner before the selling partner sets the item's price. The selling partner can then take any estimated fees into account. Each fees estimate request must include an original identifier. This identifier is included in the fees estimate so that you can correlate a fees estimate with the original request.

**Note:** The identifier value is only an estimate, actual costs may vary. Search "fees" in [Seller Central](https://sellercentral.amazon.com/) and consult the store-specific fee schedule for the most up-to-date information.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 2 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetMyFeesEstimateForSKU(params *GetMyFeesEstimateForSKUParams, opts ...ClientOption) (*GetMyFeesEstimateForSKUOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMyFeesEstimateForSKUParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMyFeesEstimateForSKU",
		Method:             "POST",
		PathPattern:        "/products/fees/v0/listings/{SellerSKU}/feesEstimate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMyFeesEstimateForSKUReader{formats: a.formats},
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
	success, ok := result.(*GetMyFeesEstimateForSKUOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMyFeesEstimateForSKU: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMyFeesEstimates Returns the estimated fees for a list of products.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.5 | 1 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetMyFeesEstimates(params *GetMyFeesEstimatesParams, opts ...ClientOption) (*GetMyFeesEstimatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMyFeesEstimatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMyFeesEstimates",
		Method:             "POST",
		PathPattern:        "/products/fees/v0/feesEstimate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMyFeesEstimatesReader{formats: a.formats},
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
	success, ok := result.(*GetMyFeesEstimatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMyFeesEstimates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
