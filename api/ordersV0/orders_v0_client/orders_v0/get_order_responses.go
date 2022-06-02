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

// GetOrderReader is a Reader for the GetOrder structure.
type GetOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrderOK creates a GetOrderOK with default headers values
func NewGetOrderOK() *GetOrderOK {
	return &GetOrderOK{}
}

/* GetOrderOK describes a response with status code 200, with default header values.

Success.
*/
type GetOrderOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderResponse
}

func (o *GetOrderOK) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}][%d] getOrderOK  %+v", 200, o.Payload)
}
func (o *GetOrderOK) GetPayload() *orders_v0_models.GetOrderResponse {
	return o.Payload
}

func (o *GetOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBadRequest creates a GetOrderBadRequest with default headers values
func NewGetOrderBadRequest() *GetOrderBadRequest {
	return &GetOrderBadRequest{}
}

/* GetOrderBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetOrderBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderResponse
}

func (o *GetOrderBadRequest) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}][%d] getOrderBadRequest  %+v", 400, o.Payload)
}
func (o *GetOrderBadRequest) GetPayload() *orders_v0_models.GetOrderResponse {
	return o.Payload
}

func (o *GetOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderForbidden creates a GetOrderForbidden with default headers values
func NewGetOrderForbidden() *GetOrderForbidden {
	return &GetOrderForbidden{}
}

/* GetOrderForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetOrderForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderResponse
}

func (o *GetOrderForbidden) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}][%d] getOrderForbidden  %+v", 403, o.Payload)
}
func (o *GetOrderForbidden) GetPayload() *orders_v0_models.GetOrderResponse {
	return o.Payload
}

func (o *GetOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderNotFound creates a GetOrderNotFound with default headers values
func NewGetOrderNotFound() *GetOrderNotFound {
	return &GetOrderNotFound{}
}

/* GetOrderNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetOrderNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderResponse
}

func (o *GetOrderNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}][%d] getOrderNotFound  %+v", 404, o.Payload)
}
func (o *GetOrderNotFound) GetPayload() *orders_v0_models.GetOrderResponse {
	return o.Payload
}

func (o *GetOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderTooManyRequests creates a GetOrderTooManyRequests with default headers values
func NewGetOrderTooManyRequests() *GetOrderTooManyRequests {
	return &GetOrderTooManyRequests{}
}

/* GetOrderTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetOrderTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderResponse
}

func (o *GetOrderTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}][%d] getOrderTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetOrderTooManyRequests) GetPayload() *orders_v0_models.GetOrderResponse {
	return o.Payload
}

func (o *GetOrderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderInternalServerError creates a GetOrderInternalServerError with default headers values
func NewGetOrderInternalServerError() *GetOrderInternalServerError {
	return &GetOrderInternalServerError{}
}

/* GetOrderInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetOrderInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderResponse
}

func (o *GetOrderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}][%d] getOrderInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOrderInternalServerError) GetPayload() *orders_v0_models.GetOrderResponse {
	return o.Payload
}

func (o *GetOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderServiceUnavailable creates a GetOrderServiceUnavailable with default headers values
func NewGetOrderServiceUnavailable() *GetOrderServiceUnavailable {
	return &GetOrderServiceUnavailable{}
}

/* GetOrderServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetOrderServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderResponse
}

func (o *GetOrderServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}][%d] getOrderServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetOrderServiceUnavailable) GetPayload() *orders_v0_models.GetOrderResponse {
	return o.Payload
}

func (o *GetOrderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
