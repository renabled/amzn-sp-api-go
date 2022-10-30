// Code generated by go-swagger; DO NOT EDIT.

package update_inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new update inventory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for update inventory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SubmitInventoryUpdate(params *SubmitInventoryUpdateParams, opts ...ClientOption) (*SubmitInventoryUpdateAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	SubmitInventoryUpdate Submits inventory updates for the specified warehouse for either a partial or full feed of inventory items.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 10 | 10 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) SubmitInventoryUpdate(params *SubmitInventoryUpdateParams, opts ...ClientOption) (*SubmitInventoryUpdateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubmitInventoryUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "submitInventoryUpdate",
		Method:             "POST",
		PathPattern:        "/vendor/directFulfillment/inventory/v1/warehouses/{warehouseId}/items",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubmitInventoryUpdateReader{formats: a.formats},
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
	success, ok := result.(*SubmitInventoryUpdateAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for submitInventoryUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
