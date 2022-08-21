// Code generated by go-swagger; DO NOT EDIT.

package easy_ship

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new easy ship API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for easy ship API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateScheduledPackage(params *CreateScheduledPackageParams, opts ...ClientOption) (*CreateScheduledPackageOK, error)

	CreateScheduledPackageBulk(params *CreateScheduledPackageBulkParams, opts ...ClientOption) (*CreateScheduledPackageBulkOK, error)

	GetScheduledPackage(params *GetScheduledPackageParams, opts ...ClientOption) (*GetScheduledPackageOK, error)

	ListHandoverSlots(params *ListHandoverSlotsParams, opts ...ClientOption) (*ListHandoverSlotsOK, error)

	UpdateScheduledPackages(params *UpdateScheduledPackagesParams, opts ...ClientOption) (*UpdateScheduledPackagesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateScheduledPackage Schedules an Easy Ship order and returns the scheduled package information.

This operation does the following:

*  Specifies the time slot and handover method for the order to be scheduled for delivery.

* Updates the Easy Ship order status.

* Generates a shipping label and an invoice. Calling `createScheduledPackage` also generates a warranty document if you specify a `SerialNumber` value. To get these documents, see [How to get invoice, shipping label, and warranty documents](doc:easy-ship-api-v2022-03-23-use-case-guide).

* Shows the status of Easy Ship orders when you call the `getOrders` operation of the Selling Partner API for Orders and examine the `EasyShipShipmentStatus` property in the response body.

See the **Shipping Label**, **Invoice**, and **Warranty** columns in the [Marketplace Support Table](doc:easyship-api-v2022-03-23-use-case-guide#marketplace-support-table) to see which documents are supported in each marketplace.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateScheduledPackage(params *CreateScheduledPackageParams, opts ...ClientOption) (*CreateScheduledPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateScheduledPackageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createScheduledPackage",
		Method:             "POST",
		PathPattern:        "/easyShip/2022-03-23/package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateScheduledPackageReader{formats: a.formats},
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
	success, ok := result.(*CreateScheduledPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createScheduledPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateScheduledPackageBulk This operation automatically schedules a time slot for all the `amazonOrderId`s given as input, generating the associated shipping labels, along with other compliance documents according to the marketplace (refer to the [marketplace document support table](doc:easy-ship-api-v2022-03-23-marketplace-document-support-table)).

Developers calling this operation may optionally assign a `packageDetails` object, allowing them to input a preferred time slot for each order in ther request. In this case, Amazon will try to schedule the respective packages using their optional settings. On the other hand, *i.e.*, if the time slot is not provided, Amazon will then pick the earliest time slot possible.

Regarding the shipping label's file format, external developers are able to choose between PDF or ZPL, and Amazon will create the label accordingly.

This operation returns an array composed of the scheduled packages, and a short-lived URL pointing to a zip file containing the generated shipping labels and the other documents enabled for your marketplace. If at least an order couldn't be scheduled, then Amazon adds the `rejectedOrders` list into the response, which contains an entry for each order we couldn't process. Each entry is composed of an error message describing the reason of the failure, so that sellers can take action.

The table below displays the supported request and burst maximum rates:
| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |
The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) CreateScheduledPackageBulk(params *CreateScheduledPackageBulkParams, opts ...ClientOption) (*CreateScheduledPackageBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateScheduledPackageBulkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createScheduledPackageBulk",
		Method:             "POST",
		PathPattern:        "/easyShip/2022-03-23/packages/bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateScheduledPackageBulkReader{formats: a.formats},
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
	success, ok := result.(*CreateScheduledPackageBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createScheduledPackageBulk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetScheduledPackage Returns information about a package, including dimensions, weight, time slot information for handover, invoice and item information, and status.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) GetScheduledPackage(params *GetScheduledPackageParams, opts ...ClientOption) (*GetScheduledPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduledPackageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getScheduledPackage",
		Method:             "GET",
		PathPattern:        "/easyShip/2022-03-23/package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduledPackageReader{formats: a.formats},
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
	success, ok := result.(*GetScheduledPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getScheduledPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListHandoverSlots Returns time slots available for Easy Ship orders to be scheduled based on the package weight and dimensions that the seller specifies.

This operation is available for scheduled and unscheduled orders based on marketplace support. See **Get Time Slots** in the [Marketplace Support Table](doc:easyship-api-v2022-03-23-use-case-guide#marketplace-support-table).

This operation can return time slots that have either pickup or drop-off handover methods - see **Supported Handover Methods** in the [Marketplace Support Table](doc:easyship-api-v2022-03-23-use-case-guide#marketplace-support-table).

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) ListHandoverSlots(params *ListHandoverSlotsParams, opts ...ClientOption) (*ListHandoverSlotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListHandoverSlotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listHandoverSlots",
		Method:             "POST",
		PathPattern:        "/easyShip/2022-03-23/timeSlot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListHandoverSlotsReader{formats: a.formats},
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
	success, ok := result.(*ListHandoverSlotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listHandoverSlots: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateScheduledPackages Updates the time slot for handing over the package indicated by the specified `scheduledPackageId`. You can get the new `slotId` value for the time slot by calling the `listHandoverSlots` operation before making another `patch` call.

See the **Update Package** column in the [Marketplace Support Table](doc:easyship-api-v2022-03-23-use-case-guide#marketplace-support-table) to see which marketplaces this operation is supported in.

**Usage Plan:**

| Rate (requests per second) | Burst |
| ---- | ---- |
| 1 | 5 |

The `x-amzn-RateLimit-Limit` response header returns the usage plan rate limits that were applied to the requested operation, when available. The table above indicates the default rate and burst values for this operation. Selling partners whose business demands require higher throughput may see higher rate and burst values than those shown here. For more information, see [Usage Plans and Rate Limits in the Selling Partner API](doc:usage-plans-and-rate-limits-in-the-sp-api).
*/
func (a *Client) UpdateScheduledPackages(params *UpdateScheduledPackagesParams, opts ...ClientOption) (*UpdateScheduledPackagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateScheduledPackagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateScheduledPackages",
		Method:             "PATCH",
		PathPattern:        "/easyShip/2022-03-23/package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateScheduledPackagesReader{formats: a.formats},
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
	success, ok := result.(*UpdateScheduledPackagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateScheduledPackages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
