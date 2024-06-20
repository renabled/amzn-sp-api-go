// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentOutbound_2020-07-01/fulfillment_outbound_2020_07_01_models"
)

// CancelFulfillmentOrderReader is a Reader for the CancelFulfillmentOrder structure.
type CancelFulfillmentOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelFulfillmentOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelFulfillmentOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelFulfillmentOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelFulfillmentOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelFulfillmentOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelFulfillmentOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCancelFulfillmentOrderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelFulfillmentOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelFulfillmentOrderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelFulfillmentOrderOK creates a CancelFulfillmentOrderOK with default headers values
func NewCancelFulfillmentOrderOK() *CancelFulfillmentOrderOK {
	return &CancelFulfillmentOrderOK{}
}

/*
CancelFulfillmentOrderOK describes a response with status code 200, with default header values.

Success.
*/
type CancelFulfillmentOrderOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order o k response has a 2xx status code
func (o *CancelFulfillmentOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel fulfillment order o k response has a 3xx status code
func (o *CancelFulfillmentOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order o k response has a 4xx status code
func (o *CancelFulfillmentOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel fulfillment order o k response has a 5xx status code
func (o *CancelFulfillmentOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel fulfillment order o k response a status code equal to that given
func (o *CancelFulfillmentOrderOK) IsCode(code int) bool {
	return code == 200
}

func (o *CancelFulfillmentOrderOK) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderOK  %+v", 200, o.Payload)
}

func (o *CancelFulfillmentOrderOK) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderOK  %+v", 200, o.Payload)
}

func (o *CancelFulfillmentOrderOK) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFulfillmentOrderBadRequest creates a CancelFulfillmentOrderBadRequest with default headers values
func NewCancelFulfillmentOrderBadRequest() *CancelFulfillmentOrderBadRequest {
	return &CancelFulfillmentOrderBadRequest{}
}

/*
CancelFulfillmentOrderBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CancelFulfillmentOrderBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order bad request response has a 2xx status code
func (o *CancelFulfillmentOrderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel fulfillment order bad request response has a 3xx status code
func (o *CancelFulfillmentOrderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order bad request response has a 4xx status code
func (o *CancelFulfillmentOrderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel fulfillment order bad request response has a 5xx status code
func (o *CancelFulfillmentOrderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel fulfillment order bad request response a status code equal to that given
func (o *CancelFulfillmentOrderBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CancelFulfillmentOrderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderBadRequest  %+v", 400, o.Payload)
}

func (o *CancelFulfillmentOrderBadRequest) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderBadRequest  %+v", 400, o.Payload)
}

func (o *CancelFulfillmentOrderBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFulfillmentOrderUnauthorized creates a CancelFulfillmentOrderUnauthorized with default headers values
func NewCancelFulfillmentOrderUnauthorized() *CancelFulfillmentOrderUnauthorized {
	return &CancelFulfillmentOrderUnauthorized{}
}

/*
CancelFulfillmentOrderUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CancelFulfillmentOrderUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order unauthorized response has a 2xx status code
func (o *CancelFulfillmentOrderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel fulfillment order unauthorized response has a 3xx status code
func (o *CancelFulfillmentOrderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order unauthorized response has a 4xx status code
func (o *CancelFulfillmentOrderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel fulfillment order unauthorized response has a 5xx status code
func (o *CancelFulfillmentOrderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel fulfillment order unauthorized response a status code equal to that given
func (o *CancelFulfillmentOrderUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CancelFulfillmentOrderUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelFulfillmentOrderUnauthorized) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelFulfillmentOrderUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFulfillmentOrderForbidden creates a CancelFulfillmentOrderForbidden with default headers values
func NewCancelFulfillmentOrderForbidden() *CancelFulfillmentOrderForbidden {
	return &CancelFulfillmentOrderForbidden{}
}

/*
CancelFulfillmentOrderForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CancelFulfillmentOrderForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order forbidden response has a 2xx status code
func (o *CancelFulfillmentOrderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel fulfillment order forbidden response has a 3xx status code
func (o *CancelFulfillmentOrderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order forbidden response has a 4xx status code
func (o *CancelFulfillmentOrderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel fulfillment order forbidden response has a 5xx status code
func (o *CancelFulfillmentOrderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel fulfillment order forbidden response a status code equal to that given
func (o *CancelFulfillmentOrderForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CancelFulfillmentOrderForbidden) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderForbidden  %+v", 403, o.Payload)
}

func (o *CancelFulfillmentOrderForbidden) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderForbidden  %+v", 403, o.Payload)
}

func (o *CancelFulfillmentOrderForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFulfillmentOrderNotFound creates a CancelFulfillmentOrderNotFound with default headers values
func NewCancelFulfillmentOrderNotFound() *CancelFulfillmentOrderNotFound {
	return &CancelFulfillmentOrderNotFound{}
}

/*
CancelFulfillmentOrderNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CancelFulfillmentOrderNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order not found response has a 2xx status code
func (o *CancelFulfillmentOrderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel fulfillment order not found response has a 3xx status code
func (o *CancelFulfillmentOrderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order not found response has a 4xx status code
func (o *CancelFulfillmentOrderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel fulfillment order not found response has a 5xx status code
func (o *CancelFulfillmentOrderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel fulfillment order not found response a status code equal to that given
func (o *CancelFulfillmentOrderNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CancelFulfillmentOrderNotFound) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderNotFound  %+v", 404, o.Payload)
}

func (o *CancelFulfillmentOrderNotFound) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderNotFound  %+v", 404, o.Payload)
}

func (o *CancelFulfillmentOrderNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFulfillmentOrderTooManyRequests creates a CancelFulfillmentOrderTooManyRequests with default headers values
func NewCancelFulfillmentOrderTooManyRequests() *CancelFulfillmentOrderTooManyRequests {
	return &CancelFulfillmentOrderTooManyRequests{}
}

/*
CancelFulfillmentOrderTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CancelFulfillmentOrderTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order too many requests response has a 2xx status code
func (o *CancelFulfillmentOrderTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel fulfillment order too many requests response has a 3xx status code
func (o *CancelFulfillmentOrderTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order too many requests response has a 4xx status code
func (o *CancelFulfillmentOrderTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel fulfillment order too many requests response has a 5xx status code
func (o *CancelFulfillmentOrderTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel fulfillment order too many requests response a status code equal to that given
func (o *CancelFulfillmentOrderTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CancelFulfillmentOrderTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelFulfillmentOrderTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelFulfillmentOrderTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFulfillmentOrderInternalServerError creates a CancelFulfillmentOrderInternalServerError with default headers values
func NewCancelFulfillmentOrderInternalServerError() *CancelFulfillmentOrderInternalServerError {
	return &CancelFulfillmentOrderInternalServerError{}
}

/*
CancelFulfillmentOrderInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CancelFulfillmentOrderInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order internal server error response has a 2xx status code
func (o *CancelFulfillmentOrderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel fulfillment order internal server error response has a 3xx status code
func (o *CancelFulfillmentOrderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order internal server error response has a 4xx status code
func (o *CancelFulfillmentOrderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel fulfillment order internal server error response has a 5xx status code
func (o *CancelFulfillmentOrderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel fulfillment order internal server error response a status code equal to that given
func (o *CancelFulfillmentOrderInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CancelFulfillmentOrderInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelFulfillmentOrderInternalServerError) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelFulfillmentOrderInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFulfillmentOrderServiceUnavailable creates a CancelFulfillmentOrderServiceUnavailable with default headers values
func NewCancelFulfillmentOrderServiceUnavailable() *CancelFulfillmentOrderServiceUnavailable {
	return &CancelFulfillmentOrderServiceUnavailable{}
}

/*
CancelFulfillmentOrderServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CancelFulfillmentOrderServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse
}

// IsSuccess returns true when this cancel fulfillment order service unavailable response has a 2xx status code
func (o *CancelFulfillmentOrderServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel fulfillment order service unavailable response has a 3xx status code
func (o *CancelFulfillmentOrderServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel fulfillment order service unavailable response has a 4xx status code
func (o *CancelFulfillmentOrderServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel fulfillment order service unavailable response has a 5xx status code
func (o *CancelFulfillmentOrderServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel fulfillment order service unavailable response a status code equal to that given
func (o *CancelFulfillmentOrderServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CancelFulfillmentOrderServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelFulfillmentOrderServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/cancel][%d] cancelFulfillmentOrderServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelFulfillmentOrderServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse {
	return o.Payload
}

func (o *CancelFulfillmentOrderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CancelFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
