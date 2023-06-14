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

// GetFulfillmentOrderReader is a Reader for the GetFulfillmentOrder structure.
type GetFulfillmentOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFulfillmentOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFulfillmentOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFulfillmentOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFulfillmentOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFulfillmentOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFulfillmentOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFulfillmentOrderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFulfillmentOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFulfillmentOrderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFulfillmentOrderOK creates a GetFulfillmentOrderOK with default headers values
func NewGetFulfillmentOrderOK() *GetFulfillmentOrderOK {
	return &GetFulfillmentOrderOK{}
}

/*
GetFulfillmentOrderOK describes a response with status code 200, with default header values.

Success.
*/
type GetFulfillmentOrderOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order o k response has a 2xx status code
func (o *GetFulfillmentOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fulfillment order o k response has a 3xx status code
func (o *GetFulfillmentOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order o k response has a 4xx status code
func (o *GetFulfillmentOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fulfillment order o k response has a 5xx status code
func (o *GetFulfillmentOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fulfillment order o k response a status code equal to that given
func (o *GetFulfillmentOrderOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFulfillmentOrderOK) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderOK  %+v", 200, o.Payload)
}

func (o *GetFulfillmentOrderOK) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderOK  %+v", 200, o.Payload)
}

func (o *GetFulfillmentOrderOK) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderBadRequest creates a GetFulfillmentOrderBadRequest with default headers values
func NewGetFulfillmentOrderBadRequest() *GetFulfillmentOrderBadRequest {
	return &GetFulfillmentOrderBadRequest{}
}

/*
GetFulfillmentOrderBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetFulfillmentOrderBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order bad request response has a 2xx status code
func (o *GetFulfillmentOrderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fulfillment order bad request response has a 3xx status code
func (o *GetFulfillmentOrderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order bad request response has a 4xx status code
func (o *GetFulfillmentOrderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fulfillment order bad request response has a 5xx status code
func (o *GetFulfillmentOrderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fulfillment order bad request response a status code equal to that given
func (o *GetFulfillmentOrderBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFulfillmentOrderBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderBadRequest  %+v", 400, o.Payload)
}

func (o *GetFulfillmentOrderBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderBadRequest  %+v", 400, o.Payload)
}

func (o *GetFulfillmentOrderBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderUnauthorized creates a GetFulfillmentOrderUnauthorized with default headers values
func NewGetFulfillmentOrderUnauthorized() *GetFulfillmentOrderUnauthorized {
	return &GetFulfillmentOrderUnauthorized{}
}

/*
GetFulfillmentOrderUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetFulfillmentOrderUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order unauthorized response has a 2xx status code
func (o *GetFulfillmentOrderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fulfillment order unauthorized response has a 3xx status code
func (o *GetFulfillmentOrderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order unauthorized response has a 4xx status code
func (o *GetFulfillmentOrderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fulfillment order unauthorized response has a 5xx status code
func (o *GetFulfillmentOrderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get fulfillment order unauthorized response a status code equal to that given
func (o *GetFulfillmentOrderUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFulfillmentOrderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFulfillmentOrderUnauthorized) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFulfillmentOrderUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderForbidden creates a GetFulfillmentOrderForbidden with default headers values
func NewGetFulfillmentOrderForbidden() *GetFulfillmentOrderForbidden {
	return &GetFulfillmentOrderForbidden{}
}

/*
GetFulfillmentOrderForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetFulfillmentOrderForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order forbidden response has a 2xx status code
func (o *GetFulfillmentOrderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fulfillment order forbidden response has a 3xx status code
func (o *GetFulfillmentOrderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order forbidden response has a 4xx status code
func (o *GetFulfillmentOrderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fulfillment order forbidden response has a 5xx status code
func (o *GetFulfillmentOrderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fulfillment order forbidden response a status code equal to that given
func (o *GetFulfillmentOrderForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFulfillmentOrderForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderForbidden  %+v", 403, o.Payload)
}

func (o *GetFulfillmentOrderForbidden) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderForbidden  %+v", 403, o.Payload)
}

func (o *GetFulfillmentOrderForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderNotFound creates a GetFulfillmentOrderNotFound with default headers values
func NewGetFulfillmentOrderNotFound() *GetFulfillmentOrderNotFound {
	return &GetFulfillmentOrderNotFound{}
}

/*
GetFulfillmentOrderNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetFulfillmentOrderNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order not found response has a 2xx status code
func (o *GetFulfillmentOrderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fulfillment order not found response has a 3xx status code
func (o *GetFulfillmentOrderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order not found response has a 4xx status code
func (o *GetFulfillmentOrderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fulfillment order not found response has a 5xx status code
func (o *GetFulfillmentOrderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fulfillment order not found response a status code equal to that given
func (o *GetFulfillmentOrderNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFulfillmentOrderNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderNotFound  %+v", 404, o.Payload)
}

func (o *GetFulfillmentOrderNotFound) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderNotFound  %+v", 404, o.Payload)
}

func (o *GetFulfillmentOrderNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderTooManyRequests creates a GetFulfillmentOrderTooManyRequests with default headers values
func NewGetFulfillmentOrderTooManyRequests() *GetFulfillmentOrderTooManyRequests {
	return &GetFulfillmentOrderTooManyRequests{}
}

/*
GetFulfillmentOrderTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetFulfillmentOrderTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order too many requests response has a 2xx status code
func (o *GetFulfillmentOrderTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fulfillment order too many requests response has a 3xx status code
func (o *GetFulfillmentOrderTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order too many requests response has a 4xx status code
func (o *GetFulfillmentOrderTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fulfillment order too many requests response has a 5xx status code
func (o *GetFulfillmentOrderTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get fulfillment order too many requests response a status code equal to that given
func (o *GetFulfillmentOrderTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFulfillmentOrderTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFulfillmentOrderTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFulfillmentOrderTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderInternalServerError creates a GetFulfillmentOrderInternalServerError with default headers values
func NewGetFulfillmentOrderInternalServerError() *GetFulfillmentOrderInternalServerError {
	return &GetFulfillmentOrderInternalServerError{}
}

/*
GetFulfillmentOrderInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetFulfillmentOrderInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order internal server error response has a 2xx status code
func (o *GetFulfillmentOrderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fulfillment order internal server error response has a 3xx status code
func (o *GetFulfillmentOrderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order internal server error response has a 4xx status code
func (o *GetFulfillmentOrderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fulfillment order internal server error response has a 5xx status code
func (o *GetFulfillmentOrderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fulfillment order internal server error response a status code equal to that given
func (o *GetFulfillmentOrderInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFulfillmentOrderInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFulfillmentOrderInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFulfillmentOrderInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFulfillmentOrderServiceUnavailable creates a GetFulfillmentOrderServiceUnavailable with default headers values
func NewGetFulfillmentOrderServiceUnavailable() *GetFulfillmentOrderServiceUnavailable {
	return &GetFulfillmentOrderServiceUnavailable{}
}

/*
GetFulfillmentOrderServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetFulfillmentOrderServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse
}

// IsSuccess returns true when this get fulfillment order service unavailable response has a 2xx status code
func (o *GetFulfillmentOrderServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fulfillment order service unavailable response has a 3xx status code
func (o *GetFulfillmentOrderServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fulfillment order service unavailable response has a 4xx status code
func (o *GetFulfillmentOrderServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fulfillment order service unavailable response has a 5xx status code
func (o *GetFulfillmentOrderServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get fulfillment order service unavailable response a status code equal to that given
func (o *GetFulfillmentOrderServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFulfillmentOrderServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFulfillmentOrderServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] getFulfillmentOrderServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFulfillmentOrderServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse {
	return o.Payload
}

func (o *GetFulfillmentOrderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.GetFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
