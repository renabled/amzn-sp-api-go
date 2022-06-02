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

// GetOrderBuyerInfoReader is a Reader for the GetOrderBuyerInfo structure.
type GetOrderBuyerInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderBuyerInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderBuyerInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrderBuyerInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrderBuyerInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrderBuyerInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrderBuyerInfoTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrderBuyerInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrderBuyerInfoServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrderBuyerInfoOK creates a GetOrderBuyerInfoOK with default headers values
func NewGetOrderBuyerInfoOK() *GetOrderBuyerInfoOK {
	return &GetOrderBuyerInfoOK{}
}

/* GetOrderBuyerInfoOK describes a response with status code 200, with default header values.

Success.
*/
type GetOrderBuyerInfoOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderBuyerInfoResponse
}

func (o *GetOrderBuyerInfoOK) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/buyerInfo][%d] getOrderBuyerInfoOK  %+v", 200, o.Payload)
}
func (o *GetOrderBuyerInfoOK) GetPayload() *orders_v0_models.GetOrderBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderBuyerInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBuyerInfoBadRequest creates a GetOrderBuyerInfoBadRequest with default headers values
func NewGetOrderBuyerInfoBadRequest() *GetOrderBuyerInfoBadRequest {
	return &GetOrderBuyerInfoBadRequest{}
}

/* GetOrderBuyerInfoBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetOrderBuyerInfoBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderBuyerInfoResponse
}

func (o *GetOrderBuyerInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/buyerInfo][%d] getOrderBuyerInfoBadRequest  %+v", 400, o.Payload)
}
func (o *GetOrderBuyerInfoBadRequest) GetPayload() *orders_v0_models.GetOrderBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderBuyerInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBuyerInfoForbidden creates a GetOrderBuyerInfoForbidden with default headers values
func NewGetOrderBuyerInfoForbidden() *GetOrderBuyerInfoForbidden {
	return &GetOrderBuyerInfoForbidden{}
}

/* GetOrderBuyerInfoForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetOrderBuyerInfoForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderBuyerInfoResponse
}

func (o *GetOrderBuyerInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/buyerInfo][%d] getOrderBuyerInfoForbidden  %+v", 403, o.Payload)
}
func (o *GetOrderBuyerInfoForbidden) GetPayload() *orders_v0_models.GetOrderBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderBuyerInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(orders_v0_models.GetOrderBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBuyerInfoNotFound creates a GetOrderBuyerInfoNotFound with default headers values
func NewGetOrderBuyerInfoNotFound() *GetOrderBuyerInfoNotFound {
	return &GetOrderBuyerInfoNotFound{}
}

/* GetOrderBuyerInfoNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetOrderBuyerInfoNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderBuyerInfoResponse
}

func (o *GetOrderBuyerInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/buyerInfo][%d] getOrderBuyerInfoNotFound  %+v", 404, o.Payload)
}
func (o *GetOrderBuyerInfoNotFound) GetPayload() *orders_v0_models.GetOrderBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderBuyerInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBuyerInfoTooManyRequests creates a GetOrderBuyerInfoTooManyRequests with default headers values
func NewGetOrderBuyerInfoTooManyRequests() *GetOrderBuyerInfoTooManyRequests {
	return &GetOrderBuyerInfoTooManyRequests{}
}

/* GetOrderBuyerInfoTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetOrderBuyerInfoTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderBuyerInfoResponse
}

func (o *GetOrderBuyerInfoTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/buyerInfo][%d] getOrderBuyerInfoTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetOrderBuyerInfoTooManyRequests) GetPayload() *orders_v0_models.GetOrderBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderBuyerInfoTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBuyerInfoInternalServerError creates a GetOrderBuyerInfoInternalServerError with default headers values
func NewGetOrderBuyerInfoInternalServerError() *GetOrderBuyerInfoInternalServerError {
	return &GetOrderBuyerInfoInternalServerError{}
}

/* GetOrderBuyerInfoInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetOrderBuyerInfoInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderBuyerInfoResponse
}

func (o *GetOrderBuyerInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/buyerInfo][%d] getOrderBuyerInfoInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOrderBuyerInfoInternalServerError) GetPayload() *orders_v0_models.GetOrderBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderBuyerInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderBuyerInfoServiceUnavailable creates a GetOrderBuyerInfoServiceUnavailable with default headers values
func NewGetOrderBuyerInfoServiceUnavailable() *GetOrderBuyerInfoServiceUnavailable {
	return &GetOrderBuyerInfoServiceUnavailable{}
}

/* GetOrderBuyerInfoServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetOrderBuyerInfoServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *orders_v0_models.GetOrderBuyerInfoResponse
}

func (o *GetOrderBuyerInfoServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /orders/v0/orders/{orderId}/buyerInfo][%d] getOrderBuyerInfoServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetOrderBuyerInfoServiceUnavailable) GetPayload() *orders_v0_models.GetOrderBuyerInfoResponse {
	return o.Payload
}

func (o *GetOrderBuyerInfoServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(orders_v0_models.GetOrderBuyerInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
