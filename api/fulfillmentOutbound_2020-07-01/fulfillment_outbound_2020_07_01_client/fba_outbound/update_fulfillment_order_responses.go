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

// UpdateFulfillmentOrderReader is a Reader for the UpdateFulfillmentOrder structure.
type UpdateFulfillmentOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFulfillmentOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFulfillmentOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateFulfillmentOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateFulfillmentOrderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateFulfillmentOrderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFulfillmentOrderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateFulfillmentOrderTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateFulfillmentOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateFulfillmentOrderServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFulfillmentOrderOK creates a UpdateFulfillmentOrderOK with default headers values
func NewUpdateFulfillmentOrderOK() *UpdateFulfillmentOrderOK {
	return &UpdateFulfillmentOrderOK{}
}

/*
UpdateFulfillmentOrderOK describes a response with status code 200, with default header values.

Success.
*/
type UpdateFulfillmentOrderOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order o k response has a 2xx status code
func (o *UpdateFulfillmentOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update fulfillment order o k response has a 3xx status code
func (o *UpdateFulfillmentOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order o k response has a 4xx status code
func (o *UpdateFulfillmentOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update fulfillment order o k response has a 5xx status code
func (o *UpdateFulfillmentOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update fulfillment order o k response a status code equal to that given
func (o *UpdateFulfillmentOrderOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateFulfillmentOrderOK) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderOK  %+v", 200, o.Payload)
}

func (o *UpdateFulfillmentOrderOK) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderOK  %+v", 200, o.Payload)
}

func (o *UpdateFulfillmentOrderOK) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFulfillmentOrderBadRequest creates a UpdateFulfillmentOrderBadRequest with default headers values
func NewUpdateFulfillmentOrderBadRequest() *UpdateFulfillmentOrderBadRequest {
	return &UpdateFulfillmentOrderBadRequest{}
}

/*
UpdateFulfillmentOrderBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateFulfillmentOrderBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order bad request response has a 2xx status code
func (o *UpdateFulfillmentOrderBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update fulfillment order bad request response has a 3xx status code
func (o *UpdateFulfillmentOrderBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order bad request response has a 4xx status code
func (o *UpdateFulfillmentOrderBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update fulfillment order bad request response has a 5xx status code
func (o *UpdateFulfillmentOrderBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update fulfillment order bad request response a status code equal to that given
func (o *UpdateFulfillmentOrderBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateFulfillmentOrderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFulfillmentOrderBadRequest) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFulfillmentOrderBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFulfillmentOrderUnauthorized creates a UpdateFulfillmentOrderUnauthorized with default headers values
func NewUpdateFulfillmentOrderUnauthorized() *UpdateFulfillmentOrderUnauthorized {
	return &UpdateFulfillmentOrderUnauthorized{}
}

/*
UpdateFulfillmentOrderUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type UpdateFulfillmentOrderUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order unauthorized response has a 2xx status code
func (o *UpdateFulfillmentOrderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update fulfillment order unauthorized response has a 3xx status code
func (o *UpdateFulfillmentOrderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order unauthorized response has a 4xx status code
func (o *UpdateFulfillmentOrderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update fulfillment order unauthorized response has a 5xx status code
func (o *UpdateFulfillmentOrderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update fulfillment order unauthorized response a status code equal to that given
func (o *UpdateFulfillmentOrderUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateFulfillmentOrderUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateFulfillmentOrderUnauthorized) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateFulfillmentOrderUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFulfillmentOrderForbidden creates a UpdateFulfillmentOrderForbidden with default headers values
func NewUpdateFulfillmentOrderForbidden() *UpdateFulfillmentOrderForbidden {
	return &UpdateFulfillmentOrderForbidden{}
}

/*
UpdateFulfillmentOrderForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateFulfillmentOrderForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order forbidden response has a 2xx status code
func (o *UpdateFulfillmentOrderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update fulfillment order forbidden response has a 3xx status code
func (o *UpdateFulfillmentOrderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order forbidden response has a 4xx status code
func (o *UpdateFulfillmentOrderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update fulfillment order forbidden response has a 5xx status code
func (o *UpdateFulfillmentOrderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update fulfillment order forbidden response a status code equal to that given
func (o *UpdateFulfillmentOrderForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateFulfillmentOrderForbidden) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderForbidden  %+v", 403, o.Payload)
}

func (o *UpdateFulfillmentOrderForbidden) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderForbidden  %+v", 403, o.Payload)
}

func (o *UpdateFulfillmentOrderForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFulfillmentOrderNotFound creates a UpdateFulfillmentOrderNotFound with default headers values
func NewUpdateFulfillmentOrderNotFound() *UpdateFulfillmentOrderNotFound {
	return &UpdateFulfillmentOrderNotFound{}
}

/*
UpdateFulfillmentOrderNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type UpdateFulfillmentOrderNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order not found response has a 2xx status code
func (o *UpdateFulfillmentOrderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update fulfillment order not found response has a 3xx status code
func (o *UpdateFulfillmentOrderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order not found response has a 4xx status code
func (o *UpdateFulfillmentOrderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update fulfillment order not found response has a 5xx status code
func (o *UpdateFulfillmentOrderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update fulfillment order not found response a status code equal to that given
func (o *UpdateFulfillmentOrderNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateFulfillmentOrderNotFound) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFulfillmentOrderNotFound) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFulfillmentOrderNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFulfillmentOrderTooManyRequests creates a UpdateFulfillmentOrderTooManyRequests with default headers values
func NewUpdateFulfillmentOrderTooManyRequests() *UpdateFulfillmentOrderTooManyRequests {
	return &UpdateFulfillmentOrderTooManyRequests{}
}

/*
UpdateFulfillmentOrderTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateFulfillmentOrderTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order too many requests response has a 2xx status code
func (o *UpdateFulfillmentOrderTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update fulfillment order too many requests response has a 3xx status code
func (o *UpdateFulfillmentOrderTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order too many requests response has a 4xx status code
func (o *UpdateFulfillmentOrderTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update fulfillment order too many requests response has a 5xx status code
func (o *UpdateFulfillmentOrderTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update fulfillment order too many requests response a status code equal to that given
func (o *UpdateFulfillmentOrderTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateFulfillmentOrderTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateFulfillmentOrderTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateFulfillmentOrderTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFulfillmentOrderInternalServerError creates a UpdateFulfillmentOrderInternalServerError with default headers values
func NewUpdateFulfillmentOrderInternalServerError() *UpdateFulfillmentOrderInternalServerError {
	return &UpdateFulfillmentOrderInternalServerError{}
}

/*
UpdateFulfillmentOrderInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type UpdateFulfillmentOrderInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order internal server error response has a 2xx status code
func (o *UpdateFulfillmentOrderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update fulfillment order internal server error response has a 3xx status code
func (o *UpdateFulfillmentOrderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order internal server error response has a 4xx status code
func (o *UpdateFulfillmentOrderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update fulfillment order internal server error response has a 5xx status code
func (o *UpdateFulfillmentOrderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update fulfillment order internal server error response a status code equal to that given
func (o *UpdateFulfillmentOrderInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateFulfillmentOrderInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateFulfillmentOrderInternalServerError) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateFulfillmentOrderInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFulfillmentOrderServiceUnavailable creates a UpdateFulfillmentOrderServiceUnavailable with default headers values
func NewUpdateFulfillmentOrderServiceUnavailable() *UpdateFulfillmentOrderServiceUnavailable {
	return &UpdateFulfillmentOrderServiceUnavailable{}
}

/*
UpdateFulfillmentOrderServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateFulfillmentOrderServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse
}

// IsSuccess returns true when this update fulfillment order service unavailable response has a 2xx status code
func (o *UpdateFulfillmentOrderServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update fulfillment order service unavailable response has a 3xx status code
func (o *UpdateFulfillmentOrderServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update fulfillment order service unavailable response has a 4xx status code
func (o *UpdateFulfillmentOrderServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update fulfillment order service unavailable response has a 5xx status code
func (o *UpdateFulfillmentOrderServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update fulfillment order service unavailable response a status code equal to that given
func (o *UpdateFulfillmentOrderServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *UpdateFulfillmentOrderServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateFulfillmentOrderServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}][%d] updateFulfillmentOrderServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateFulfillmentOrderServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse {
	return o.Payload
}

func (o *UpdateFulfillmentOrderServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.UpdateFulfillmentOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
