// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new messaging API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for messaging API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAmazonMotors(params *CreateAmazonMotorsParams, opts ...ClientOption) (*CreateAmazonMotorsCreated, error)

	CreateWarranty(params *CreateWarrantyParams, opts ...ClientOption) (*CreateWarrantyCreated, error)

	GetAttributes(params *GetAttributesParams, opts ...ClientOption) (*GetAttributesOK, error)

	ConfirmCustomizationDetails(params *ConfirmCustomizationDetailsParams, opts ...ClientOption) (*ConfirmCustomizationDetailsCreated, error)

	CreateConfirmDeliveryDetails(params *CreateConfirmDeliveryDetailsParams, opts ...ClientOption) (*CreateConfirmDeliveryDetailsCreated, error)

	CreateConfirmOrderDetails(params *CreateConfirmOrderDetailsParams, opts ...ClientOption) (*CreateConfirmOrderDetailsCreated, error)

	CreateConfirmServiceDetails(params *CreateConfirmServiceDetailsParams, opts ...ClientOption) (*CreateConfirmServiceDetailsCreated, error)

	CreateDigitalAccessKey(params *CreateDigitalAccessKeyParams, opts ...ClientOption) (*CreateDigitalAccessKeyCreated, error)

	CreateLegalDisclosure(params *CreateLegalDisclosureParams, opts ...ClientOption) (*CreateLegalDisclosureCreated, error)

	CreateNegativeFeedbackRemoval(params *CreateNegativeFeedbackRemovalParams, opts ...ClientOption) (*CreateNegativeFeedbackRemovalCreated, error)

	CreateUnexpectedProblem(params *CreateUnexpectedProblemParams, opts ...ClientOption) (*CreateUnexpectedProblemCreated, error)

	GetMessagingActionsForOrder(params *GetMessagingActionsForOrderParams, opts ...ClientOption) (*GetMessagingActionsForOrderOK, error)

	SendInvoice(params *SendInvoiceParams, opts ...ClientOption) (*SendInvoiceCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateAmazonMotors Sends a message to a buyer to provide details about an Amazon Motors order. This message can only be sent by Amazon Motors sellers.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateAmazonMotors(params *CreateAmazonMotorsParams, opts ...ClientOption) (*CreateAmazonMotorsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAmazonMotorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateAmazonMotors",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/amazonMotors",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAmazonMotorsReader{formats: a.formats},
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
	success, ok := result.(*CreateAmazonMotorsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateAmazonMotors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateWarranty Sends a message to a buyer to provide details about warranty information on a purchase in their order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateWarranty(params *CreateWarrantyParams, opts ...ClientOption) (*CreateWarrantyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWarrantyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateWarranty",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/warranty",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateWarrantyReader{formats: a.formats},
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
	success, ok := result.(*CreateWarrantyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateWarranty: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetAttributes Returns a response containing attributes related to an order. This includes buyer preferences.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |
*/
func (a *Client) GetAttributes(params *GetAttributesParams, opts ...ClientOption) (*GetAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAttributesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAttributes",
		Method:             "GET",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/attributes",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAttributesReader{formats: a.formats},
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
	success, ok := result.(*GetAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAttributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	ConfirmCustomizationDetails Sends a message asking a buyer to provide or verify customization details such as name spelling, images, initials, etc.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) ConfirmCustomizationDetails(params *ConfirmCustomizationDetailsParams, opts ...ClientOption) (*ConfirmCustomizationDetailsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfirmCustomizationDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "confirmCustomizationDetails",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/confirmCustomizationDetails",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConfirmCustomizationDetailsReader{formats: a.formats},
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
	success, ok := result.(*ConfirmCustomizationDetailsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for confirmCustomizationDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateConfirmDeliveryDetails Sends a message to a buyer to arrange a delivery or to confirm contact information for making a delivery.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateConfirmDeliveryDetails(params *CreateConfirmDeliveryDetailsParams, opts ...ClientOption) (*CreateConfirmDeliveryDetailsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConfirmDeliveryDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createConfirmDeliveryDetails",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/confirmDeliveryDetails",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateConfirmDeliveryDetailsReader{formats: a.formats},
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
	success, ok := result.(*CreateConfirmDeliveryDetailsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createConfirmDeliveryDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateConfirmOrderDetails Sends a message to ask a buyer an order-related question prior to shipping their order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateConfirmOrderDetails(params *CreateConfirmOrderDetailsParams, opts ...ClientOption) (*CreateConfirmOrderDetailsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConfirmOrderDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createConfirmOrderDetails",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/confirmOrderDetails",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateConfirmOrderDetailsReader{formats: a.formats},
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
	success, ok := result.(*CreateConfirmOrderDetailsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createConfirmOrderDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateConfirmServiceDetails Sends a message to contact a Home Service customer to arrange a service call or to gather information prior to a service call.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateConfirmServiceDetails(params *CreateConfirmServiceDetailsParams, opts ...ClientOption) (*CreateConfirmServiceDetailsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateConfirmServiceDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createConfirmServiceDetails",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/confirmServiceDetails",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateConfirmServiceDetailsReader{formats: a.formats},
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
	success, ok := result.(*CreateConfirmServiceDetailsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createConfirmServiceDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateDigitalAccessKey Sends a message to a buyer to share a digital access key needed to utilize digital content in their order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateDigitalAccessKey(params *CreateDigitalAccessKeyParams, opts ...ClientOption) (*CreateDigitalAccessKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDigitalAccessKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDigitalAccessKey",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/digitalAccessKey",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDigitalAccessKeyReader{formats: a.formats},
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
	success, ok := result.(*CreateDigitalAccessKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDigitalAccessKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateLegalDisclosure Sends a critical message that contains documents that a seller is legally obligated to provide to the buyer. This message should only be used to deliver documents that are required by law.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateLegalDisclosure(params *CreateLegalDisclosureParams, opts ...ClientOption) (*CreateLegalDisclosureCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLegalDisclosureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createLegalDisclosure",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/legalDisclosure",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLegalDisclosureReader{formats: a.formats},
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
	success, ok := result.(*CreateLegalDisclosureCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createLegalDisclosure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateNegativeFeedbackRemoval Sends a non-critical message that asks a buyer to remove their negative feedback. This message should only be sent after the seller has resolved the buyer's problem.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateNegativeFeedbackRemoval(params *CreateNegativeFeedbackRemovalParams, opts ...ClientOption) (*CreateNegativeFeedbackRemovalCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNegativeFeedbackRemovalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createNegativeFeedbackRemoval",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/negativeFeedbackRemoval",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNegativeFeedbackRemovalReader{formats: a.formats},
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
	success, ok := result.(*CreateNegativeFeedbackRemovalCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createNegativeFeedbackRemoval: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	CreateUnexpectedProblem Sends a critical message to a buyer that an unexpected problem was encountered affecting the completion of the order.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateUnexpectedProblem(params *CreateUnexpectedProblemParams, opts ...ClientOption) (*CreateUnexpectedProblemCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUnexpectedProblemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUnexpectedProblem",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/unexpectedProblem",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUnexpectedProblemReader{formats: a.formats},
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
	success, ok := result.(*CreateUnexpectedProblemCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUnexpectedProblem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMessagingActionsForOrder Returns a list of message types that are available for an order that you specify. A message type is represented by an actions object, which contains a path and query parameter(s). You can use the path and parameter(s) to call an operation that sends a message.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](https://developer-docs.amazon.com/sp-api/docs/usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetMessagingActionsForOrder(params *GetMessagingActionsForOrderParams, opts ...ClientOption) (*GetMessagingActionsForOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMessagingActionsForOrderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMessagingActionsForOrder",
		Method:             "GET",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMessagingActionsForOrderReader{formats: a.formats},
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
	success, ok := result.(*GetMessagingActionsForOrderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMessagingActionsForOrder: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SendInvoice Sends a message providing the buyer an invoice
*/
func (a *Client) SendInvoice(params *SendInvoiceParams, opts ...ClientOption) (*SendInvoiceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendInvoiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sendInvoice",
		Method:             "POST",
		PathPattern:        "/messaging/v1/orders/{amazonOrderId}/messages/invoice",
		ProducesMediaTypes: []string{"application/hal+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendInvoiceReader{formats: a.formats},
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
	success, ok := result.(*SendInvoiceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendInvoice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
