// Code generated by go-swagger; DO NOT EDIT.

package orders_v0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/ordersV0/orders_v0_models"
)

// GetOrderItemsReader is a Reader for the GetOrderItems structure.
type GetOrderItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrderItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrderItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrderItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrderItemsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrderItemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrderItemsOK creates a GetOrderItemsOK with default headers values
func NewGetOrderItemsOK() *GetOrderItemsOK {
	return &GetOrderItemsOK{}
}

/*
GetOrderItemsOK describes a response with status code 200, with default header values.

Success.
*/
type GetOrderItemsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsResponse
}

// IsSuccess returns true when this get order items o k response has a 2xx status code
func (o *GetOrderItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get order items o k response has a 3xx status code
func (o *GetOrderItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items o k response has a 4xx status code
func (o *GetOrderItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order items o k response has a 5xx status code
func (o *GetOrderItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items o k response a status code equal to that given
func (o *GetOrderItemsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrderItemsOK) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsOK  %+v", 200, o.Payload)
}

func (o *GetOrderItemsOK) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsOK  %+v", 200, o.Payload)
}

func (o *GetOrderItemsOK) GetPayload() *orders_v0_models.GetOrderItemsResponse {
	return o.Payload
}

func (o *GetOrderItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsBadRequest creates a GetOrderItemsBadRequest with default headers values
func NewGetOrderItemsBadRequest() *GetOrderItemsBadRequest {
	return &GetOrderItemsBadRequest{}
}

/*
GetOrderItemsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetOrderItemsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsResponse
}

// IsSuccess returns true when this get order items bad request response has a 2xx status code
func (o *GetOrderItemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items bad request response has a 3xx status code
func (o *GetOrderItemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items bad request response has a 4xx status code
func (o *GetOrderItemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items bad request response has a 5xx status code
func (o *GetOrderItemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items bad request response a status code equal to that given
func (o *GetOrderItemsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrderItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrderItemsBadRequest) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrderItemsBadRequest) GetPayload() *orders_v0_models.GetOrderItemsResponse {
	return o.Payload
}

func (o *GetOrderItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsForbidden creates a GetOrderItemsForbidden with default headers values
func NewGetOrderItemsForbidden() *GetOrderItemsForbidden {
	return &GetOrderItemsForbidden{}
}

/*
GetOrderItemsForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetOrderItemsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsResponse
}

// IsSuccess returns true when this get order items forbidden response has a 2xx status code
func (o *GetOrderItemsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items forbidden response has a 3xx status code
func (o *GetOrderItemsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items forbidden response has a 4xx status code
func (o *GetOrderItemsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items forbidden response has a 5xx status code
func (o *GetOrderItemsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items forbidden response a status code equal to that given
func (o *GetOrderItemsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrderItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrderItemsForbidden) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrderItemsForbidden) GetPayload() *orders_v0_models.GetOrderItemsResponse {
	return o.Payload
}

func (o *GetOrderItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsNotFound creates a GetOrderItemsNotFound with default headers values
func NewGetOrderItemsNotFound() *GetOrderItemsNotFound {
	return &GetOrderItemsNotFound{}
}

/*
GetOrderItemsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetOrderItemsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsResponse
}

// IsSuccess returns true when this get order items not found response has a 2xx status code
func (o *GetOrderItemsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items not found response has a 3xx status code
func (o *GetOrderItemsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items not found response has a 4xx status code
func (o *GetOrderItemsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items not found response has a 5xx status code
func (o *GetOrderItemsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items not found response a status code equal to that given
func (o *GetOrderItemsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrderItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrderItemsNotFound) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrderItemsNotFound) GetPayload() *orders_v0_models.GetOrderItemsResponse {
	return o.Payload
}

func (o *GetOrderItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsTooManyRequests creates a GetOrderItemsTooManyRequests with default headers values
func NewGetOrderItemsTooManyRequests() *GetOrderItemsTooManyRequests {
	return &GetOrderItemsTooManyRequests{}
}

/*
GetOrderItemsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetOrderItemsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsResponse
}

// IsSuccess returns true when this get order items too many requests response has a 2xx status code
func (o *GetOrderItemsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items too many requests response has a 3xx status code
func (o *GetOrderItemsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items too many requests response has a 4xx status code
func (o *GetOrderItemsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items too many requests response has a 5xx status code
func (o *GetOrderItemsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items too many requests response a status code equal to that given
func (o *GetOrderItemsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrderItemsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrderItemsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrderItemsTooManyRequests) GetPayload() *orders_v0_models.GetOrderItemsResponse {
	return o.Payload
}

func (o *GetOrderItemsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsInternalServerError creates a GetOrderItemsInternalServerError with default headers values
func NewGetOrderItemsInternalServerError() *GetOrderItemsInternalServerError {
	return &GetOrderItemsInternalServerError{}
}

/*
GetOrderItemsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetOrderItemsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsResponse
}

// IsSuccess returns true when this get order items internal server error response has a 2xx status code
func (o *GetOrderItemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items internal server error response has a 3xx status code
func (o *GetOrderItemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items internal server error response has a 4xx status code
func (o *GetOrderItemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order items internal server error response has a 5xx status code
func (o *GetOrderItemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get order items internal server error response a status code equal to that given
func (o *GetOrderItemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrderItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderItemsInternalServerError) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderItemsInternalServerError) GetPayload() *orders_v0_models.GetOrderItemsResponse {
	return o.Payload
}

func (o *GetOrderItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsServiceUnavailable creates a GetOrderItemsServiceUnavailable with default headers values
func NewGetOrderItemsServiceUnavailable() *GetOrderItemsServiceUnavailable {
	return &GetOrderItemsServiceUnavailable{}
}

/*
GetOrderItemsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetOrderItemsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsResponse
}

// IsSuccess returns true when this get order items service unavailable response has a 2xx status code
func (o *GetOrderItemsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items service unavailable response has a 3xx status code
func (o *GetOrderItemsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items service unavailable response has a 4xx status code
func (o *GetOrderItemsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order items service unavailable response has a 5xx status code
func (o *GetOrderItemsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get order items service unavailable response a status code equal to that given
func (o *GetOrderItemsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrderItemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrderItemsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems][%d] getOrderItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrderItemsServiceUnavailable) GetPayload() *orders_v0_models.GetOrderItemsResponse {
	return o.Payload
}

func (o *GetOrderItemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
