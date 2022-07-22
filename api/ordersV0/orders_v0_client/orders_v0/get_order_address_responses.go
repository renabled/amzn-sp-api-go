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

// GetOrderAddressReader is a Reader for the GetOrderAddress structure.
type GetOrderAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrderAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrderAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrderAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrderAddressTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrderAddressServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrderAddressOK creates a GetOrderAddressOK with default headers values
func NewGetOrderAddressOK() *GetOrderAddressOK {
	return &GetOrderAddressOK{}
}

/* GetOrderAddressOK describes a response with status code 200, with default header values.

Success.
*/
type GetOrderAddressOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderAddressResponse
}

func (o *GetOrderAddressOK) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/address][%d] getOrderAddressOK  %+v", 200, o.Payload)
}
func (o *GetOrderAddressOK) GetPayload() *orders_v0_models.GetOrderAddressResponse {
	return o.Payload
}

func (o *GetOrderAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderAddressBadRequest creates a GetOrderAddressBadRequest with default headers values
func NewGetOrderAddressBadRequest() *GetOrderAddressBadRequest {
	return &GetOrderAddressBadRequest{}
}

/* GetOrderAddressBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetOrderAddressBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderAddressResponse
}

func (o *GetOrderAddressBadRequest) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/address][%d] getOrderAddressBadRequest  %+v", 400, o.Payload)
}
func (o *GetOrderAddressBadRequest) GetPayload() *orders_v0_models.GetOrderAddressResponse {
	return o.Payload
}

func (o *GetOrderAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderAddressForbidden creates a GetOrderAddressForbidden with default headers values
func NewGetOrderAddressForbidden() *GetOrderAddressForbidden {
	return &GetOrderAddressForbidden{}
}

/* GetOrderAddressForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetOrderAddressForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderAddressResponse
}

func (o *GetOrderAddressForbidden) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/address][%d] getOrderAddressForbidden  %+v", 403, o.Payload)
}
func (o *GetOrderAddressForbidden) GetPayload() *orders_v0_models.GetOrderAddressResponse {
	return o.Payload
}

func (o *GetOrderAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderAddressNotFound creates a GetOrderAddressNotFound with default headers values
func NewGetOrderAddressNotFound() *GetOrderAddressNotFound {
	return &GetOrderAddressNotFound{}
}

/* GetOrderAddressNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetOrderAddressNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderAddressResponse
}

func (o *GetOrderAddressNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/address][%d] getOrderAddressNotFound  %+v", 404, o.Payload)
}
func (o *GetOrderAddressNotFound) GetPayload() *orders_v0_models.GetOrderAddressResponse {
	return o.Payload
}

func (o *GetOrderAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderAddressTooManyRequests creates a GetOrderAddressTooManyRequests with default headers values
func NewGetOrderAddressTooManyRequests() *GetOrderAddressTooManyRequests {
	return &GetOrderAddressTooManyRequests{}
}

/* GetOrderAddressTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetOrderAddressTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderAddressResponse
}

func (o *GetOrderAddressTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/address][%d] getOrderAddressTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetOrderAddressTooManyRequests) GetPayload() *orders_v0_models.GetOrderAddressResponse {
	return o.Payload
}

func (o *GetOrderAddressTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderAddressInternalServerError creates a GetOrderAddressInternalServerError with default headers values
func NewGetOrderAddressInternalServerError() *GetOrderAddressInternalServerError {
	return &GetOrderAddressInternalServerError{}
}

/* GetOrderAddressInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetOrderAddressInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderAddressResponse
}

func (o *GetOrderAddressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/address][%d] getOrderAddressInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOrderAddressInternalServerError) GetPayload() *orders_v0_models.GetOrderAddressResponse {
	return o.Payload
}

func (o *GetOrderAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderAddressServiceUnavailable creates a GetOrderAddressServiceUnavailable with default headers values
func NewGetOrderAddressServiceUnavailable() *GetOrderAddressServiceUnavailable {
	return &GetOrderAddressServiceUnavailable{}
}

/* GetOrderAddressServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetOrderAddressServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderAddressResponse
}

func (o *GetOrderAddressServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/address][%d] getOrderAddressServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetOrderAddressServiceUnavailable) GetPayload() *orders_v0_models.GetOrderAddressResponse {
	return o.Payload
}

func (o *GetOrderAddressServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
