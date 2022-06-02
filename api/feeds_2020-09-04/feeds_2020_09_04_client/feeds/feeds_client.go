// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new feeds API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for feeds API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelFeed(params *CancelFeedParams, opts ...ClientOption) (*CancelFeedOK, error)

	CreateFeed(params *CreateFeedParams, opts ...ClientOption) (*CreateFeedAccepted, error)

	CreateFeedDocument(params *CreateFeedDocumentParams, opts ...ClientOption) (*CreateFeedDocumentCreated, error)

	GetFeed(params *GetFeedParams, opts ...ClientOption) (*GetFeedOK, error)

	GetFeedDocument(params *GetFeedDocumentParams, opts ...ClientOption) (*GetFeedDocumentOK, error)

	GetFeeds(params *GetFeedsParams, opts ...ClientOption) (*GetFeedsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelFeed Cancels the feed that you specify. Only feeds with processingStatus=IN_QUEUE can be cancelled. Cancelled feeds are returned in subsequent calls to the getFeed and getFeeds operations.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) CancelFeed(params *CancelFeedParams, opts ...ClientOption) (*CancelFeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelFeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelFeed",
		Method:             "DELETE",
		PathPattern:        "/feeds/2020-09-04/feeds/{feedId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelFeedReader{formats: a.formats},
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
	success, ok := result.(*CancelFeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for cancelFeed: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateFeed Creates a feed. Encrypt and upload the contents of the feed document before calling this operation.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0083 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) CreateFeed(params *CreateFeedParams, opts ...ClientOption) (*CreateFeedAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFeed",
		Method:             "POST",
		PathPattern:        "/feeds/2020-09-04/feeds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFeedReader{formats: a.formats},
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
	success, ok := result.(*CreateFeedAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFeed: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateFeedDocument Creates a feed document for the feed type that you specify. This operation returns encryption details for encrypting the contents of the document, as well as a presigned URL for uploading the encrypted feed document contents. It also returns a feedDocumentId value that you can pass in with a subsequent call to the createFeed operation.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0083 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) CreateFeedDocument(params *CreateFeedDocumentParams, opts ...ClientOption) (*CreateFeedDocumentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFeedDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFeedDocument",
		Method:             "POST",
		PathPattern:        "/feeds/2020-09-04/documents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFeedDocumentReader{formats: a.formats},
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
	success, ok := result.(*CreateFeedDocumentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFeedDocument: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFeed Returns feed details (including the resultDocumentId, if available) for the feed that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 2.0 | 15 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFeed(params *GetFeedParams, opts ...ClientOption) (*GetFeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFeed",
		Method:             "GET",
		PathPattern:        "/feeds/2020-09-04/feeds/{feedId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeedReader{formats: a.formats},
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
	success, ok := result.(*GetFeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFeed: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFeedDocument Returns the information required for retrieving a feed document's contents. This includes a presigned URL for the feed document as well as the information required to decrypt the document's contents.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFeedDocument(params *GetFeedDocumentParams, opts ...ClientOption) (*GetFeedDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeedDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFeedDocument",
		Method:             "GET",
		PathPattern:        "/feeds/2020-09-04/documents/{feedDocumentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeedDocumentReader{formats: a.formats},
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
	success, ok := result.(*GetFeedDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFeedDocument: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFeeds Returns feed details for the feeds that match the filters that you specify.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 0.0222 | 10 |

For more information, see "Usage Plans and Rate Limits" in the Selling Partner API documentation.
*/
func (a *Client) GetFeeds(params *GetFeedsParams, opts ...ClientOption) (*GetFeedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFeedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFeeds",
		Method:             "GET",
		PathPattern:        "/feeds/2020-09-04/feeds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFeedsReader{formats: a.formats},
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
	success, ok := result.(*GetFeedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFeeds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
