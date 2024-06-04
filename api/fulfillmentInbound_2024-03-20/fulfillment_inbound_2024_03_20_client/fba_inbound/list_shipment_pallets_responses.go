// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInbound_2024-03-20/fulfillment_inbound_2024_03_20_models"
)

// ListShipmentPalletsReader is a Reader for the ListShipmentPallets structure.
type ListShipmentPalletsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListShipmentPalletsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListShipmentPalletsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListShipmentPalletsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListShipmentPalletsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListShipmentPalletsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewListShipmentPalletsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewListShipmentPalletsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListShipmentPalletsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListShipmentPalletsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListShipmentPalletsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListShipmentPalletsOK creates a ListShipmentPalletsOK with default headers values
func NewListShipmentPalletsOK() *ListShipmentPalletsOK {
	return &ListShipmentPalletsOK{}
}

/*
ListShipmentPalletsOK describes a response with status code 200, with default header values.

ListShipmentPallets 200 response
*/
type ListShipmentPalletsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ListShipmentPalletsResponse
}

// IsSuccess returns true when this list shipment pallets o k response has a 2xx status code
func (o *ListShipmentPalletsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list shipment pallets o k response has a 3xx status code
func (o *ListShipmentPalletsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets o k response has a 4xx status code
func (o *ListShipmentPalletsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shipment pallets o k response has a 5xx status code
func (o *ListShipmentPalletsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment pallets o k response a status code equal to that given
func (o *ListShipmentPalletsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListShipmentPalletsOK) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsOK  %+v", 200, o.Payload)
}

func (o *ListShipmentPalletsOK) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsOK  %+v", 200, o.Payload)
}

func (o *ListShipmentPalletsOK) GetPayload() *fulfillment_inbound_2024_03_20_models.ListShipmentPalletsResponse {
	return o.Payload
}

func (o *ListShipmentPalletsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ListShipmentPalletsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsBadRequest creates a ListShipmentPalletsBadRequest with default headers values
func NewListShipmentPalletsBadRequest() *ListShipmentPalletsBadRequest {
	return &ListShipmentPalletsBadRequest{}
}

/*
ListShipmentPalletsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListShipmentPalletsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets bad request response has a 2xx status code
func (o *ListShipmentPalletsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets bad request response has a 3xx status code
func (o *ListShipmentPalletsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets bad request response has a 4xx status code
func (o *ListShipmentPalletsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment pallets bad request response has a 5xx status code
func (o *ListShipmentPalletsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment pallets bad request response a status code equal to that given
func (o *ListShipmentPalletsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListShipmentPalletsBadRequest) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsBadRequest  %+v", 400, o.Payload)
}

func (o *ListShipmentPalletsBadRequest) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsBadRequest  %+v", 400, o.Payload)
}

func (o *ListShipmentPalletsBadRequest) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsForbidden creates a ListShipmentPalletsForbidden with default headers values
func NewListShipmentPalletsForbidden() *ListShipmentPalletsForbidden {
	return &ListShipmentPalletsForbidden{}
}

/*
ListShipmentPalletsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListShipmentPalletsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets forbidden response has a 2xx status code
func (o *ListShipmentPalletsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets forbidden response has a 3xx status code
func (o *ListShipmentPalletsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets forbidden response has a 4xx status code
func (o *ListShipmentPalletsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment pallets forbidden response has a 5xx status code
func (o *ListShipmentPalletsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment pallets forbidden response a status code equal to that given
func (o *ListShipmentPalletsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListShipmentPalletsForbidden) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsForbidden  %+v", 403, o.Payload)
}

func (o *ListShipmentPalletsForbidden) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsForbidden  %+v", 403, o.Payload)
}

func (o *ListShipmentPalletsForbidden) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsNotFound creates a ListShipmentPalletsNotFound with default headers values
func NewListShipmentPalletsNotFound() *ListShipmentPalletsNotFound {
	return &ListShipmentPalletsNotFound{}
}

/*
ListShipmentPalletsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type ListShipmentPalletsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets not found response has a 2xx status code
func (o *ListShipmentPalletsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets not found response has a 3xx status code
func (o *ListShipmentPalletsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets not found response has a 4xx status code
func (o *ListShipmentPalletsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment pallets not found response has a 5xx status code
func (o *ListShipmentPalletsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment pallets not found response a status code equal to that given
func (o *ListShipmentPalletsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListShipmentPalletsNotFound) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsNotFound  %+v", 404, o.Payload)
}

func (o *ListShipmentPalletsNotFound) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsNotFound  %+v", 404, o.Payload)
}

func (o *ListShipmentPalletsNotFound) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RateLimit-Limit
	hdrXAmznRateLimitLimit := response.GetHeader("x-amzn-RateLimit-Limit")

	if hdrXAmznRateLimitLimit != "" {
		o.XAmznRateLimitLimit = hdrXAmznRateLimitLimit
	}

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsRequestEntityTooLarge creates a ListShipmentPalletsRequestEntityTooLarge with default headers values
func NewListShipmentPalletsRequestEntityTooLarge() *ListShipmentPalletsRequestEntityTooLarge {
	return &ListShipmentPalletsRequestEntityTooLarge{}
}

/*
ListShipmentPalletsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type ListShipmentPalletsRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets request entity too large response has a 2xx status code
func (o *ListShipmentPalletsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets request entity too large response has a 3xx status code
func (o *ListShipmentPalletsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets request entity too large response has a 4xx status code
func (o *ListShipmentPalletsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment pallets request entity too large response has a 5xx status code
func (o *ListShipmentPalletsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment pallets request entity too large response a status code equal to that given
func (o *ListShipmentPalletsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *ListShipmentPalletsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListShipmentPalletsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *ListShipmentPalletsRequestEntityTooLarge) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsUnsupportedMediaType creates a ListShipmentPalletsUnsupportedMediaType with default headers values
func NewListShipmentPalletsUnsupportedMediaType() *ListShipmentPalletsUnsupportedMediaType {
	return &ListShipmentPalletsUnsupportedMediaType{}
}

/*
ListShipmentPalletsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type ListShipmentPalletsUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets unsupported media type response has a 2xx status code
func (o *ListShipmentPalletsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets unsupported media type response has a 3xx status code
func (o *ListShipmentPalletsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets unsupported media type response has a 4xx status code
func (o *ListShipmentPalletsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment pallets unsupported media type response has a 5xx status code
func (o *ListShipmentPalletsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment pallets unsupported media type response a status code equal to that given
func (o *ListShipmentPalletsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *ListShipmentPalletsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListShipmentPalletsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListShipmentPalletsUnsupportedMediaType) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsTooManyRequests creates a ListShipmentPalletsTooManyRequests with default headers values
func NewListShipmentPalletsTooManyRequests() *ListShipmentPalletsTooManyRequests {
	return &ListShipmentPalletsTooManyRequests{}
}

/*
ListShipmentPalletsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListShipmentPalletsTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets too many requests response has a 2xx status code
func (o *ListShipmentPalletsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets too many requests response has a 3xx status code
func (o *ListShipmentPalletsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets too many requests response has a 4xx status code
func (o *ListShipmentPalletsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shipment pallets too many requests response has a 5xx status code
func (o *ListShipmentPalletsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list shipment pallets too many requests response a status code equal to that given
func (o *ListShipmentPalletsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListShipmentPalletsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListShipmentPalletsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListShipmentPalletsTooManyRequests) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsInternalServerError creates a ListShipmentPalletsInternalServerError with default headers values
func NewListShipmentPalletsInternalServerError() *ListShipmentPalletsInternalServerError {
	return &ListShipmentPalletsInternalServerError{}
}

/*
ListShipmentPalletsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListShipmentPalletsInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets internal server error response has a 2xx status code
func (o *ListShipmentPalletsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets internal server error response has a 3xx status code
func (o *ListShipmentPalletsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets internal server error response has a 4xx status code
func (o *ListShipmentPalletsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shipment pallets internal server error response has a 5xx status code
func (o *ListShipmentPalletsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list shipment pallets internal server error response a status code equal to that given
func (o *ListShipmentPalletsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListShipmentPalletsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListShipmentPalletsInternalServerError) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListShipmentPalletsInternalServerError) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListShipmentPalletsServiceUnavailable creates a ListShipmentPalletsServiceUnavailable with default headers values
func NewListShipmentPalletsServiceUnavailable() *ListShipmentPalletsServiceUnavailable {
	return &ListShipmentPalletsServiceUnavailable{}
}

/*
ListShipmentPalletsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListShipmentPalletsServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_2024_03_20_models.ErrorList
}

// IsSuccess returns true when this list shipment pallets service unavailable response has a 2xx status code
func (o *ListShipmentPalletsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shipment pallets service unavailable response has a 3xx status code
func (o *ListShipmentPalletsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shipment pallets service unavailable response has a 4xx status code
func (o *ListShipmentPalletsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shipment pallets service unavailable response has a 5xx status code
func (o *ListShipmentPalletsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list shipment pallets service unavailable response a status code equal to that given
func (o *ListShipmentPalletsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ListShipmentPalletsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListShipmentPalletsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /inbound/fba/2024-03-20/inboundPlans/{inboundPlanId}/shipments/{shipmentId}/pallets][%d] listShipmentPalletsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListShipmentPalletsServiceUnavailable) GetPayload() *fulfillment_inbound_2024_03_20_models.ErrorList {
	return o.Payload
}

func (o *ListShipmentPalletsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_2024_03_20_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
