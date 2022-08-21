// Code generated by go-swagger; DO NOT EDIT.

package shipping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new shipping API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for shipping API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelShipment(params *CancelShipmentParams, opts ...ClientOption) (*CancelShipmentOK, error)

	CreateShipment(params *CreateShipmentParams, opts ...ClientOption) (*CreateShipmentOK, error)

	GetAccount(params *GetAccountParams, opts ...ClientOption) (*GetAccountOK, error)

	GetRates(params *GetRatesParams, opts ...ClientOption) (*GetRatesOK, error)

	GetShipment(params *GetShipmentParams, opts ...ClientOption) (*GetShipmentOK, error)

	GetTrackingInformation(params *GetTrackingInformationParams, opts ...ClientOption) (*GetTrackingInformationOK, error)

	PurchaseLabels(params *PurchaseLabelsParams, opts ...ClientOption) (*PurchaseLabelsOK, error)

	PurchaseShipment(params *PurchaseShipmentParams, opts ...ClientOption) (*PurchaseShipmentOK, error)

	RetrieveShippingLabel(params *RetrieveShippingLabelParams, opts ...ClientOption) (*RetrieveShippingLabelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelShipment Cancel a shipment by the given shipmentId.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) CancelShipment(params *CancelShipmentParams, opts ...ClientOption) (*CancelShipmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelShipmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelShipment",
		Method:             "POST",
		PathPattern:        "/shipping/v1/shipments/{shipmentId}/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelShipmentReader{formats: a.formats},
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
	success, ok := result.(*CancelShipmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cancelShipment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateShipment Create a new shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateShipment(params *CreateShipmentParams, opts ...ClientOption) (*CreateShipmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateShipmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createShipment",
		Method:             "POST",
		PathPattern:        "/shipping/v1/shipments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateShipmentReader{formats: a.formats},
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
	success, ok := result.(*CreateShipmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createShipment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccount Verify if the current account is valid.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetAccount(params *GetAccountParams, opts ...ClientOption) (*GetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccount",
		Method:             "GET",
		PathPattern:        "/shipping/v1/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountReader{formats: a.formats},
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
	success, ok := result.(*GetAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRates Get service rates.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetRates(params *GetRatesParams, opts ...ClientOption) (*GetRatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRates",
		Method:             "POST",
		PathPattern:        "/shipping/v1/rates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRatesReader{formats: a.formats},
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
	success, ok := result.(*GetRatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetShipment Return the entire shipment object for the shipmentId.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetShipment(params *GetShipmentParams, opts ...ClientOption) (*GetShipmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShipmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getShipment",
		Method:             "GET",
		PathPattern:        "/shipping/v1/shipments/{shipmentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetShipmentReader{formats: a.formats},
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
	success, ok := result.(*GetShipmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getShipment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrackingInformation Return the tracking information of a shipment.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 1 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetTrackingInformation(params *GetTrackingInformationParams, opts ...ClientOption) (*GetTrackingInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrackingInformationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTrackingInformation",
		Method:             "GET",
		PathPattern:        "/shipping/v1/tracking/{trackingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrackingInformationReader{formats: a.formats},
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
	success, ok := result.(*GetTrackingInformationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTrackingInformation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurchaseLabels Purchase shipping labels based on a given rate.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) PurchaseLabels(params *PurchaseLabelsParams, opts ...ClientOption) (*PurchaseLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "purchaseLabels",
		Method:             "POST",
		PathPattern:        "/shipping/v1/shipments/{shipmentId}/purchaseLabels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseLabelsReader{formats: a.formats},
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
	success, ok := result.(*PurchaseLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for purchaseLabels: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PurchaseShipment Purchase shipping labels.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) PurchaseShipment(params *PurchaseShipmentParams, opts ...ClientOption) (*PurchaseShipmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPurchaseShipmentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "purchaseShipment",
		Method:             "POST",
		PathPattern:        "/shipping/v1/purchaseShipment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PurchaseShipmentReader{formats: a.formats},
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
	success, ok := result.(*PurchaseShipmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for purchaseShipment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RetrieveShippingLabel Retrieve shipping label based on the shipment id and tracking id.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 5 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) RetrieveShippingLabel(params *RetrieveShippingLabelParams, opts ...ClientOption) (*RetrieveShippingLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveShippingLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "retrieveShippingLabel",
		Method:             "POST",
		PathPattern:        "/shipping/v1/shipments/{shipmentId}/containers/{trackingId}/label",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveShippingLabelReader{formats: a.formats},
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
	success, ok := result.(*RetrieveShippingLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retrieveShippingLabel: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
