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

// GetOrderRegulatedInfoReader is a Reader for the GetOrderRegulatedInfo structure.
type GetOrderRegulatedInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderRegulatedInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderRegulatedInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrderRegulatedInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrderRegulatedInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrderRegulatedInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrderRegulatedInfoTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderRegulatedInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrderRegulatedInfoServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrderRegulatedInfoOK creates a GetOrderRegulatedInfoOK with default headers values
func NewGetOrderRegulatedInfoOK() *GetOrderRegulatedInfoOK {
	return &GetOrderRegulatedInfoOK{}
}

/*
GetOrderRegulatedInfoOK describes a response with status code 200, with default header values.

Success.
*/
type GetOrderRegulatedInfoOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderRegulatedInfoResponse
}

// IsSuccess returns true when this get order regulated info o k response has a 2xx status code
func (o *GetOrderRegulatedInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get order regulated info o k response has a 3xx status code
func (o *GetOrderRegulatedInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order regulated info o k response has a 4xx status code
func (o *GetOrderRegulatedInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order regulated info o k response has a 5xx status code
func (o *GetOrderRegulatedInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get order regulated info o k response a status code equal to that given
func (o *GetOrderRegulatedInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrderRegulatedInfoOK) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoOK  %+v", 200, o.Payload)
}

func (o *GetOrderRegulatedInfoOK) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoOK  %+v", 200, o.Payload)
}

func (o *GetOrderRegulatedInfoOK) GetPayload() *orders_v0_models.GetOrderRegulatedInfoResponse {
	return o.Payload
}

func (o *GetOrderRegulatedInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderRegulatedInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderRegulatedInfoBadRequest creates a GetOrderRegulatedInfoBadRequest with default headers values
func NewGetOrderRegulatedInfoBadRequest() *GetOrderRegulatedInfoBadRequest {
	return &GetOrderRegulatedInfoBadRequest{}
}

/*
GetOrderRegulatedInfoBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetOrderRegulatedInfoBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderRegulatedInfoResponse
}

// IsSuccess returns true when this get order regulated info bad request response has a 2xx status code
func (o *GetOrderRegulatedInfoBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order regulated info bad request response has a 3xx status code
func (o *GetOrderRegulatedInfoBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order regulated info bad request response has a 4xx status code
func (o *GetOrderRegulatedInfoBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order regulated info bad request response has a 5xx status code
func (o *GetOrderRegulatedInfoBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get order regulated info bad request response a status code equal to that given
func (o *GetOrderRegulatedInfoBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrderRegulatedInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrderRegulatedInfoBadRequest) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrderRegulatedInfoBadRequest) GetPayload() *orders_v0_models.GetOrderRegulatedInfoResponse {
	return o.Payload
}

func (o *GetOrderRegulatedInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderRegulatedInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderRegulatedInfoForbidden creates a GetOrderRegulatedInfoForbidden with default headers values
func NewGetOrderRegulatedInfoForbidden() *GetOrderRegulatedInfoForbidden {
	return &GetOrderRegulatedInfoForbidden{}
}

/*
GetOrderRegulatedInfoForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetOrderRegulatedInfoForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderRegulatedInfoResponse
}

// IsSuccess returns true when this get order regulated info forbidden response has a 2xx status code
func (o *GetOrderRegulatedInfoForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order regulated info forbidden response has a 3xx status code
func (o *GetOrderRegulatedInfoForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order regulated info forbidden response has a 4xx status code
func (o *GetOrderRegulatedInfoForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order regulated info forbidden response has a 5xx status code
func (o *GetOrderRegulatedInfoForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get order regulated info forbidden response a status code equal to that given
func (o *GetOrderRegulatedInfoForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrderRegulatedInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetOrderRegulatedInfoForbidden) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoForbidden  %+v", 403, o.Payload)
}

func (o *GetOrderRegulatedInfoForbidden) GetPayload() *orders_v0_models.GetOrderRegulatedInfoResponse {
	return o.Payload
}

func (o *GetOrderRegulatedInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderRegulatedInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderRegulatedInfoNotFound creates a GetOrderRegulatedInfoNotFound with default headers values
func NewGetOrderRegulatedInfoNotFound() *GetOrderRegulatedInfoNotFound {
	return &GetOrderRegulatedInfoNotFound{}
}

/*
GetOrderRegulatedInfoNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetOrderRegulatedInfoNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderRegulatedInfoResponse
}

// IsSuccess returns true when this get order regulated info not found response has a 2xx status code
func (o *GetOrderRegulatedInfoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order regulated info not found response has a 3xx status code
func (o *GetOrderRegulatedInfoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order regulated info not found response has a 4xx status code
func (o *GetOrderRegulatedInfoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order regulated info not found response has a 5xx status code
func (o *GetOrderRegulatedInfoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get order regulated info not found response a status code equal to that given
func (o *GetOrderRegulatedInfoNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrderRegulatedInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetOrderRegulatedInfoNotFound) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetOrderRegulatedInfoNotFound) GetPayload() *orders_v0_models.GetOrderRegulatedInfoResponse {
	return o.Payload
}

func (o *GetOrderRegulatedInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderRegulatedInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderRegulatedInfoTooManyRequests creates a GetOrderRegulatedInfoTooManyRequests with default headers values
func NewGetOrderRegulatedInfoTooManyRequests() *GetOrderRegulatedInfoTooManyRequests {
	return &GetOrderRegulatedInfoTooManyRequests{}
}

/*
GetOrderRegulatedInfoTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetOrderRegulatedInfoTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderRegulatedInfoResponse
}

// IsSuccess returns true when this get order regulated info too many requests response has a 2xx status code
func (o *GetOrderRegulatedInfoTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order regulated info too many requests response has a 3xx status code
func (o *GetOrderRegulatedInfoTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order regulated info too many requests response has a 4xx status code
func (o *GetOrderRegulatedInfoTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get order regulated info too many requests response has a 5xx status code
func (o *GetOrderRegulatedInfoTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get order regulated info too many requests response a status code equal to that given
func (o *GetOrderRegulatedInfoTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrderRegulatedInfoTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrderRegulatedInfoTooManyRequests) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrderRegulatedInfoTooManyRequests) GetPayload() *orders_v0_models.GetOrderRegulatedInfoResponse {
	return o.Payload
}

func (o *GetOrderRegulatedInfoTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderRegulatedInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderRegulatedInfoInternalServerError creates a GetOrderRegulatedInfoInternalServerError with default headers values
func NewGetOrderRegulatedInfoInternalServerError() *GetOrderRegulatedInfoInternalServerError {
	return &GetOrderRegulatedInfoInternalServerError{}
}

/*
GetOrderRegulatedInfoInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetOrderRegulatedInfoInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderRegulatedInfoResponse
}

// IsSuccess returns true when this get order regulated info internal server error response has a 2xx status code
func (o *GetOrderRegulatedInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order regulated info internal server error response has a 3xx status code
func (o *GetOrderRegulatedInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order regulated info internal server error response has a 4xx status code
func (o *GetOrderRegulatedInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order regulated info internal server error response has a 5xx status code
func (o *GetOrderRegulatedInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get order regulated info internal server error response a status code equal to that given
func (o *GetOrderRegulatedInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrderRegulatedInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderRegulatedInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderRegulatedInfoInternalServerError) GetPayload() *orders_v0_models.GetOrderRegulatedInfoResponse {
	return o.Payload
}

func (o *GetOrderRegulatedInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderRegulatedInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderRegulatedInfoServiceUnavailable creates a GetOrderRegulatedInfoServiceUnavailable with default headers values
func NewGetOrderRegulatedInfoServiceUnavailable() *GetOrderRegulatedInfoServiceUnavailable {
	return &GetOrderRegulatedInfoServiceUnavailable{}
}

/*
GetOrderRegulatedInfoServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetOrderRegulatedInfoServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderRegulatedInfoResponse
}

// IsSuccess returns true when this get order regulated info service unavailable response has a 2xx status code
func (o *GetOrderRegulatedInfoServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get order regulated info service unavailable response has a 3xx status code
func (o *GetOrderRegulatedInfoServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order regulated info service unavailable response has a 4xx status code
func (o *GetOrderRegulatedInfoServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order regulated info service unavailable response has a 5xx status code
func (o *GetOrderRegulatedInfoServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get order regulated info service unavailable response a status code equal to that given
func (o *GetOrderRegulatedInfoServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrderRegulatedInfoServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrderRegulatedInfoServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/regulatedInfo][%d] getOrderRegulatedInfoServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrderRegulatedInfoServiceUnavailable) GetPayload() *orders_v0_models.GetOrderRegulatedInfoResponse {
	return o.Payload
}

func (o *GetOrderRegulatedInfoServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderRegulatedInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
