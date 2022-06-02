// Code generated by go-swagger; DO NOT EDIT.

package orders_v0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new orders v0 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for orders v0 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetOrder(params *GetOrderParams, opts ...ClientOption) (*GetOrderOK, error)

	GetOrderAddress(params *GetOrderAddressParams, opts ...ClientOption) (*GetOrderAddressOK, error)

	GetOrderBuyerInfo(params *GetOrderBuyerInfoParams, opts ...ClientOption) (*GetOrderBuyerInfoOK, error)

	GetOrderItems(params *GetOrderItemsParams, opts ...ClientOption) (*GetOrderItemsOK, error)

	GetOrderItemsBuyerInfo(params *GetOrderItemsBuyerInfoParams, opts ...ClientOption) (*GetOrderItemsBuyerInfoOK, error)

	GetOrderRegulatedInfo(params *GetOrderRegulatedInfoParams, opts ...ClientOption) (*GetOrderRegulatedInfoOK, error)

	GetOrders(params *GetOrdersParams, opts ...ClientOption) (*GetOrdersOK, error)

	UpdateVerificationStatus(params *UpdateVerificationStatusParams, opts ...ClientOption) (*UpdateVerificationStatusNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetOrder Returns the order indicated by the specified order ID.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.0167 | 20 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetOrder(params *GetOrderParams, opts ...ClientOption) (*GetOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrder",
		Method:             "GET",
		PathPattern:        "/orders/v0/orders/{orderId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrderReader{formats: a.formats},
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
	success, ok := result.(*GetOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrderAddress Returns the shipping address for the specified order.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.0167 | 20 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetOrderAddress(params *GetOrderAddressParams, opts ...ClientOption) (*GetOrderAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrderAddressParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrderAddress",
		Method:             "GET",
		PathPattern:        "/orders/v0/orders/{orderId}/address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrderAddressReader{formats: a.formats},
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
	success, ok := result.(*GetOrderAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrderAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrderBuyerInfo Returns buyer information for the specified order.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.0167 | 20 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetOrderBuyerInfo(params *GetOrderBuyerInfoParams, opts ...ClientOption) (*GetOrderBuyerInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrderBuyerInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrderBuyerInfo",
		Method:             "GET",
		PathPattern:        "/orders/v0/orders/{orderId}/buyerInfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrderBuyerInfoReader{formats: a.formats},
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
	success, ok := result.(*GetOrderBuyerInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrderBuyerInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrderItems Returns detailed order item information for the order indicated by the specified order ID. If NextToken is provided, it's used to retrieve the next page of order items.

Note: When an order is in the Pending state (the order has been placed but payment has not been authorized), the getOrderItems operation does not return information about pricing, taxes, shipping charges, gift status or promotions for the order items in the order. After an order leaves the Pending state (this occurs when payment has been authorized) and enters the Unshipped, Partially Shipped, or Shipped state, the getOrderItems operation returns information about pricing, taxes, shipping charges, gift status and promotions for the order items in the order.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.5 | 30 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetOrderItems(params *GetOrderItemsParams, opts ...ClientOption) (*GetOrderItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrderItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrderItems",
		Method:             "GET",
		PathPattern:        "/orders/v0/orders/{orderId}/orderItems",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrderItemsReader{formats: a.formats},
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
	success, ok := result.(*GetOrderItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrderItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrderItemsBuyerInfo Returns buyer information for the order items in the specified order.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.5 | 30 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetOrderItemsBuyerInfo(params *GetOrderItemsBuyerInfoParams, opts ...ClientOption) (*GetOrderItemsBuyerInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrderItemsBuyerInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrderItemsBuyerInfo",
		Method:             "GET",
		PathPattern:        "/orders/v0/orders/{orderId}/orderItems/buyerInfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrderItemsBuyerInfoReader{formats: a.formats},
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
	success, ok := result.(*GetOrderItemsBuyerInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrderItemsBuyerInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrderRegulatedInfo Returns regulated information for the order indicated by the specified order ID.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.5 | 30 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetOrderRegulatedInfo(params *GetOrderRegulatedInfoParams, opts ...ClientOption) (*GetOrderRegulatedInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrderRegulatedInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrderRegulatedInfo",
		Method:             "GET",
		PathPattern:        "/orders/v0/orders/{orderId}/regulatedInfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrderRegulatedInfoReader{formats: a.formats},
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
	success, ok := result.(*GetOrderRegulatedInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrderRegulatedInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOrders Returns orders created or updated during the time frame indicated by the specified parameters. You can also apply a range of filtering criteria to narrow the list of orders returned. If NextToken is present, that will be used to retrieve the orders instead of other criteria.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.0167 | 20 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetOrders(params *GetOrdersParams, opts ...ClientOption) (*GetOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOrders",
		Method:             "GET",
		PathPattern:        "/orders/v0/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrdersReader{formats: a.formats},
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
	success, ok := result.(*GetOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVerificationStatus Updates (approves or rejects) the verification status of an order containing regulated products.

**Usage Plans:**

| Plan type | Rate (requests per second) | Burst |
| ---- | ---- | ---- |
|Default| 0.5 | 30 |
|Selling partner specific| Variable | Variable |

The x-amzn-RateLimit-Limit response header returns the usage plan rate limits that were applied to the requested operation. Rate limits for some selling partners will vary from the default rate and burst shown in the table above. For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) UpdateVerificationStatus(params *UpdateVerificationStatusParams, opts ...ClientOption) (*UpdateVerificationStatusNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVerificationStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVerificationStatus",
		Method:             "PATCH",
		PathPattern:        "/orders/v0/orders/{orderId}/regulatedInfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVerificationStatusReader{formats: a.formats},
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
	success, ok := result.(*UpdateVerificationStatusNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVerificationStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
