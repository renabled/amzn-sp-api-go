// Code generated by go-swagger; DO NOT EDIT.

package vendor_shipping_labels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vendor shipping labels API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vendor shipping labels API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetShippingLabel(params *GetShippingLabelParams, opts ...ClientOption) (*GetShippingLabelOK, error)

	GetShippingLabels(params *GetShippingLabelsParams, opts ...ClientOption) (*GetShippingLabelsOK, error)

	SubmitShippingLabelRequest(params *SubmitShippingLabelRequestParams, opts ...ClientOption) (*SubmitShippingLabelRequestAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetShippingLabel Returns a shipping label for the `purchaseOrderNumber` that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetShippingLabel(params *GetShippingLabelParams, opts ...ClientOption) (*GetShippingLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShippingLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getShippingLabel",
		Method:             "GET",
		PathPattern:        "/vendor/directFulfillment/shipping/v1/shippingLabels/{purchaseOrderNumber}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetShippingLabelReader{formats: a.formats},
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
	success, ok := result.(*GetShippingLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getShippingLabel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetShippingLabels Returns a list of shipping labels created during the time frame that you specify. You define that time frame using the `createdAfter` and `createdBefore` parameters. You must use both of these parameters. The date range to search must not be more than 7 days.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetShippingLabels(params *GetShippingLabelsParams, opts ...ClientOption) (*GetShippingLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShippingLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getShippingLabels",
		Method:             "GET",
		PathPattern:        "/vendor/directFulfillment/shipping/v1/shippingLabels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetShippingLabelsReader{formats: a.formats},
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
	success, ok := result.(*GetShippingLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getShippingLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	SubmitShippingLabelRequest Creates a shipping label for a purchase order and returns a `transactionId` for reference.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, refer to [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) SubmitShippingLabelRequest(params *SubmitShippingLabelRequestParams, opts ...ClientOption) (*SubmitShippingLabelRequestAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitShippingLabelRequestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "submitShippingLabelRequest",
		Method:             "POST",
		PathPattern:        "/vendor/directFulfillment/shipping/v1/shippingLabels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubmitShippingLabelRequestReader{formats: a.formats},
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
	success, ok := result.(*SubmitShippingLabelRequestAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for submitShippingLabelRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
