// Code generated by go-swagger; DO NOT EDIT.

package vendor_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vendor transaction API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vendor transaction API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetTransactionStatus(params *GetTransactionStatusParams, opts ...ClientOption) (*GetTransactionStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTransactionStatus Returns the status of the transaction indicated by the specified transactionId.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 10 | 10 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values then those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetTransactionStatus(params *GetTransactionStatusParams, opts ...ClientOption) (*GetTransactionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTransactionStatus",
		Method:             "GET",
		PathPattern:        "/vendor/directFulfillment/transactions/2021-12-28/transactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionStatusReader{formats: a.formats},
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
	success, ok := result.(*GetTransactionStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
