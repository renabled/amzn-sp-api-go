// Code generated by go-swagger; DO NOT EDIT.

package orders_v0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/ordersV0/orders_v0_models"
)

// GetOrderItemsBuyerInfoReader is a Reader for the GetOrderItemsBuyerInfo structure.
type GetOrderItemsBuyerInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderItemsBuyerInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderItemsBuyerInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrderItemsBuyerInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrderItemsBuyerInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrderItemsBuyerInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrderItemsBuyerInfoTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderItemsBuyerInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrderItemsBuyerInfoServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrderItemsBuyerInfoOK creates a GetOrderItemsBuyerInfoOK with default headers values
func NewGetOrderItemsBuyerInfoOK() *GetOrderItemsBuyerInfoOK {
	return &GetOrderItemsBuyerInfoOK{}
}

/*
GetOrderItemsBuyerInfoOK describes a response with status code 200, with default header values.

Success.
*/
type GetOrderItemsBuyerInfoOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsBuyerInfoResponse
}

// IsSuccess returns true when this get order items buyer info o k response has a 2xx status code
func (o *GetOrderItemsBuyerInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get order items buyer info o k response has a 3xx status code
func (o *GetOrderItemsBuyerInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items buyer info o k response has a 4xx status code
func (o *GetOrderItemsBuyerInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order items buyer info o k response has a 5xx status code
func (o *GetOrderItemsBuyerInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items buyer info o k response a status code equal to that given
func (o *GetOrderItemsBuyerInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrderItemsBuyerInfoOK) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoOK  %+v", 200, o.Payload)
}

func (o *GetOrderItemsBuyerInfoOK) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoOK  %+v", 200, o.Payload)
}

func (o *GetOrderItemsBuyerInfoOK) GetPayload() *orders_v0_models.GetOrderItemsBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderItemsBuyerInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsBuyerInfoBadRequest creates a GetOrderItemsBuyerInfoBadRequest with default headers values
func NewGetOrderItemsBuyerInfoBadRequest() *GetOrderItemsBuyerInfoBadRequest {
	return &GetOrderItemsBuyerInfoBadRequest{}
}

/*
GetOrderItemsBuyerInfoBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetOrderItemsBuyerInfoBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsBuyerInfoResponse
}

// IsSuccess returns true when this get order items buyer info bad request response has a 2xx status code
func (o *GetOrderItemsBuyerInfoBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items buyer info bad request response has a 3xx status code
func (o *GetOrderItemsBuyerInfoBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items buyer info bad request response has a 4xx status code
func (o *GetOrderItemsBuyerInfoBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items buyer info bad request response has a 5xx status code
func (o *GetOrderItemsBuyerInfoBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items buyer info bad request response a status code equal to that given
func (o *GetOrderItemsBuyerInfoBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrderItemsBuyerInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrderItemsBuyerInfoBadRequest) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrderItemsBuyerInfoBadRequest) GetPayload() *orders_v0_models.GetOrderItemsBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderItemsBuyerInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsBuyerInfoForbidden creates a GetOrderItemsBuyerInfoForbidden with default headers values
func NewGetOrderItemsBuyerInfoForbidden() *GetOrderItemsBuyerInfoForbidden {
	return &GetOrderItemsBuyerInfoForbidden{}
}

/*
GetOrderItemsBuyerInfoForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetOrderItemsBuyerInfoForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsBuyerInfoResponse
}

// IsSuccess returns true when this get order items buyer info forbidden response has a 2xx status code
func (o *GetOrderItemsBuyerInfoForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items buyer info forbidden response has a 3xx status code
func (o *GetOrderItemsBuyerInfoForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items buyer info forbidden response has a 4xx status code
func (o *GetOrderItemsBuyerInfoForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items buyer info forbidden response has a 5xx status code
func (o *GetOrderItemsBuyerInfoForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items buyer info forbidden response a status code equal to that given
func (o *GetOrderItemsBuyerInfoForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrderItemsBuyerInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetOrderItemsBuyerInfoForbidden) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetOrderItemsBuyerInfoForbidden) GetPayload() *orders_v0_models.GetOrderItemsBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderItemsBuyerInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderItemsBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsBuyerInfoNotFound creates a GetOrderItemsBuyerInfoNotFound with default headers values
func NewGetOrderItemsBuyerInfoNotFound() *GetOrderItemsBuyerInfoNotFound {
	return &GetOrderItemsBuyerInfoNotFound{}
}

/*
GetOrderItemsBuyerInfoNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetOrderItemsBuyerInfoNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsBuyerInfoResponse
}

// IsSuccess returns true when this get order items buyer info not found response has a 2xx status code
func (o *GetOrderItemsBuyerInfoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items buyer info not found response has a 3xx status code
func (o *GetOrderItemsBuyerInfoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items buyer info not found response has a 4xx status code
func (o *GetOrderItemsBuyerInfoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items buyer info not found response has a 5xx status code
func (o *GetOrderItemsBuyerInfoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items buyer info not found response a status code equal to that given
func (o *GetOrderItemsBuyerInfoNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrderItemsBuyerInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetOrderItemsBuyerInfoNotFound) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetOrderItemsBuyerInfoNotFound) GetPayload() *orders_v0_models.GetOrderItemsBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderItemsBuyerInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsBuyerInfoTooManyRequests creates a GetOrderItemsBuyerInfoTooManyRequests with default headers values
func NewGetOrderItemsBuyerInfoTooManyRequests() *GetOrderItemsBuyerInfoTooManyRequests {
	return &GetOrderItemsBuyerInfoTooManyRequests{}
}

/*
GetOrderItemsBuyerInfoTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetOrderItemsBuyerInfoTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsBuyerInfoResponse
}

// IsSuccess returns true when this get order items buyer info too many requests response has a 2xx status code
func (o *GetOrderItemsBuyerInfoTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items buyer info too many requests response has a 3xx status code
func (o *GetOrderItemsBuyerInfoTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items buyer info too many requests response has a 4xx status code
func (o *GetOrderItemsBuyerInfoTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order items buyer info too many requests response has a 5xx status code
func (o *GetOrderItemsBuyerInfoTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get order items buyer info too many requests response a status code equal to that given
func (o *GetOrderItemsBuyerInfoTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrderItemsBuyerInfoTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrderItemsBuyerInfoTooManyRequests) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrderItemsBuyerInfoTooManyRequests) GetPayload() *orders_v0_models.GetOrderItemsBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderItemsBuyerInfoTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsBuyerInfoInternalServerError creates a GetOrderItemsBuyerInfoInternalServerError with default headers values
func NewGetOrderItemsBuyerInfoInternalServerError() *GetOrderItemsBuyerInfoInternalServerError {
	return &GetOrderItemsBuyerInfoInternalServerError{}
}

/*
GetOrderItemsBuyerInfoInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetOrderItemsBuyerInfoInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsBuyerInfoResponse
}

// IsSuccess returns true when this get order items buyer info internal server error response has a 2xx status code
func (o *GetOrderItemsBuyerInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items buyer info internal server error response has a 3xx status code
func (o *GetOrderItemsBuyerInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items buyer info internal server error response has a 4xx status code
func (o *GetOrderItemsBuyerInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order items buyer info internal server error response has a 5xx status code
func (o *GetOrderItemsBuyerInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get order items buyer info internal server error response a status code equal to that given
func (o *GetOrderItemsBuyerInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrderItemsBuyerInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderItemsBuyerInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderItemsBuyerInfoInternalServerError) GetPayload() *orders_v0_models.GetOrderItemsBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderItemsBuyerInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderItemsBuyerInfoServiceUnavailable creates a GetOrderItemsBuyerInfoServiceUnavailable with default headers values
func NewGetOrderItemsBuyerInfoServiceUnavailable() *GetOrderItemsBuyerInfoServiceUnavailable {
	return &GetOrderItemsBuyerInfoServiceUnavailable{}
}

/*
GetOrderItemsBuyerInfoServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetOrderItemsBuyerInfoServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderItemsBuyerInfoResponse
}

// IsSuccess returns true when this get order items buyer info service unavailable response has a 2xx status code
func (o *GetOrderItemsBuyerInfoServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order items buyer info service unavailable response has a 3xx status code
func (o *GetOrderItemsBuyerInfoServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order items buyer info service unavailable response has a 4xx status code
func (o *GetOrderItemsBuyerInfoServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order items buyer info service unavailable response has a 5xx status code
func (o *GetOrderItemsBuyerInfoServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get order items buyer info service unavailable response a status code equal to that given
func (o *GetOrderItemsBuyerInfoServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrderItemsBuyerInfoServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrderItemsBuyerInfoServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/orderItems/buyerInfo][%d] getOrderItemsBuyerInfoServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrderItemsBuyerInfoServiceUnavailable) GetPayload() *orders_v0_models.GetOrderItemsBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderItemsBuyerInfoServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderItemsBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
