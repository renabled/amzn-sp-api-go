// Code generated by go-swagger; DO NOT EDIT.

package fba_outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/fulfillmentOutbound_2020-07-01/fulfillment_outbound_2020_07_01_models"
)

// ListReturnReasonCodesReader is a Reader for the ListReturnReasonCodes structure.
type ListReturnReasonCodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReturnReasonCodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListReturnReasonCodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListReturnReasonCodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListReturnReasonCodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListReturnReasonCodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListReturnReasonCodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListReturnReasonCodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListReturnReasonCodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListReturnReasonCodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListReturnReasonCodesOK creates a ListReturnReasonCodesOK with default headers values
func NewListReturnReasonCodesOK() *ListReturnReasonCodesOK {
	return &ListReturnReasonCodesOK{}
}

/*
ListReturnReasonCodesOK describes a response with status code 200, with default header values.

Success.
*/
type ListReturnReasonCodesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes o k response has a 2xx status code
func (o *ListReturnReasonCodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list return reason codes o k response has a 3xx status code
func (o *ListReturnReasonCodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes o k response has a 4xx status code
func (o *ListReturnReasonCodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list return reason codes o k response has a 5xx status code
func (o *ListReturnReasonCodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list return reason codes o k response a status code equal to that given
func (o *ListReturnReasonCodesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListReturnReasonCodesOK) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesOK  %+v", 200, o.Payload)
}

func (o *ListReturnReasonCodesOK) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesOK  %+v", 200, o.Payload)
}

func (o *ListReturnReasonCodesOK) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReturnReasonCodesBadRequest creates a ListReturnReasonCodesBadRequest with default headers values
func NewListReturnReasonCodesBadRequest() *ListReturnReasonCodesBadRequest {
	return &ListReturnReasonCodesBadRequest{}
}

/*
ListReturnReasonCodesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type ListReturnReasonCodesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes bad request response has a 2xx status code
func (o *ListReturnReasonCodesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list return reason codes bad request response has a 3xx status code
func (o *ListReturnReasonCodesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes bad request response has a 4xx status code
func (o *ListReturnReasonCodesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list return reason codes bad request response has a 5xx status code
func (o *ListReturnReasonCodesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list return reason codes bad request response a status code equal to that given
func (o *ListReturnReasonCodesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListReturnReasonCodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesBadRequest  %+v", 400, o.Payload)
}

func (o *ListReturnReasonCodesBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesBadRequest  %+v", 400, o.Payload)
}

func (o *ListReturnReasonCodesBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReturnReasonCodesUnauthorized creates a ListReturnReasonCodesUnauthorized with default headers values
func NewListReturnReasonCodesUnauthorized() *ListReturnReasonCodesUnauthorized {
	return &ListReturnReasonCodesUnauthorized{}
}

/*
ListReturnReasonCodesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type ListReturnReasonCodesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes unauthorized response has a 2xx status code
func (o *ListReturnReasonCodesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list return reason codes unauthorized response has a 3xx status code
func (o *ListReturnReasonCodesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes unauthorized response has a 4xx status code
func (o *ListReturnReasonCodesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list return reason codes unauthorized response has a 5xx status code
func (o *ListReturnReasonCodesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list return reason codes unauthorized response a status code equal to that given
func (o *ListReturnReasonCodesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListReturnReasonCodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListReturnReasonCodesUnauthorized) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListReturnReasonCodesUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReturnReasonCodesForbidden creates a ListReturnReasonCodesForbidden with default headers values
func NewListReturnReasonCodesForbidden() *ListReturnReasonCodesForbidden {
	return &ListReturnReasonCodesForbidden{}
}

/*
ListReturnReasonCodesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type ListReturnReasonCodesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes forbidden response has a 2xx status code
func (o *ListReturnReasonCodesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list return reason codes forbidden response has a 3xx status code
func (o *ListReturnReasonCodesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes forbidden response has a 4xx status code
func (o *ListReturnReasonCodesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list return reason codes forbidden response has a 5xx status code
func (o *ListReturnReasonCodesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list return reason codes forbidden response a status code equal to that given
func (o *ListReturnReasonCodesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListReturnReasonCodesForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesForbidden  %+v", 403, o.Payload)
}

func (o *ListReturnReasonCodesForbidden) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesForbidden  %+v", 403, o.Payload)
}

func (o *ListReturnReasonCodesForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReturnReasonCodesNotFound creates a ListReturnReasonCodesNotFound with default headers values
func NewListReturnReasonCodesNotFound() *ListReturnReasonCodesNotFound {
	return &ListReturnReasonCodesNotFound{}
}

/*
ListReturnReasonCodesNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type ListReturnReasonCodesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes not found response has a 2xx status code
func (o *ListReturnReasonCodesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list return reason codes not found response has a 3xx status code
func (o *ListReturnReasonCodesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes not found response has a 4xx status code
func (o *ListReturnReasonCodesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list return reason codes not found response has a 5xx status code
func (o *ListReturnReasonCodesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list return reason codes not found response a status code equal to that given
func (o *ListReturnReasonCodesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListReturnReasonCodesNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesNotFound  %+v", 404, o.Payload)
}

func (o *ListReturnReasonCodesNotFound) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesNotFound  %+v", 404, o.Payload)
}

func (o *ListReturnReasonCodesNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReturnReasonCodesTooManyRequests creates a ListReturnReasonCodesTooManyRequests with default headers values
func NewListReturnReasonCodesTooManyRequests() *ListReturnReasonCodesTooManyRequests {
	return &ListReturnReasonCodesTooManyRequests{}
}

/*
ListReturnReasonCodesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type ListReturnReasonCodesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes too many requests response has a 2xx status code
func (o *ListReturnReasonCodesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list return reason codes too many requests response has a 3xx status code
func (o *ListReturnReasonCodesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes too many requests response has a 4xx status code
func (o *ListReturnReasonCodesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this list return reason codes too many requests response has a 5xx status code
func (o *ListReturnReasonCodesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this list return reason codes too many requests response a status code equal to that given
func (o *ListReturnReasonCodesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *ListReturnReasonCodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListReturnReasonCodesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListReturnReasonCodesTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReturnReasonCodesInternalServerError creates a ListReturnReasonCodesInternalServerError with default headers values
func NewListReturnReasonCodesInternalServerError() *ListReturnReasonCodesInternalServerError {
	return &ListReturnReasonCodesInternalServerError{}
}

/*
ListReturnReasonCodesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type ListReturnReasonCodesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes internal server error response has a 2xx status code
func (o *ListReturnReasonCodesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list return reason codes internal server error response has a 3xx status code
func (o *ListReturnReasonCodesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes internal server error response has a 4xx status code
func (o *ListReturnReasonCodesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list return reason codes internal server error response has a 5xx status code
func (o *ListReturnReasonCodesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list return reason codes internal server error response a status code equal to that given
func (o *ListReturnReasonCodesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListReturnReasonCodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListReturnReasonCodesInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListReturnReasonCodesInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReturnReasonCodesServiceUnavailable creates a ListReturnReasonCodesServiceUnavailable with default headers values
func NewListReturnReasonCodesServiceUnavailable() *ListReturnReasonCodesServiceUnavailable {
	return &ListReturnReasonCodesServiceUnavailable{}
}

/*
ListReturnReasonCodesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type ListReturnReasonCodesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse
}

// IsSuccess returns true when this list return reason codes service unavailable response has a 2xx status code
func (o *ListReturnReasonCodesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list return reason codes service unavailable response has a 3xx status code
func (o *ListReturnReasonCodesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list return reason codes service unavailable response has a 4xx status code
func (o *ListReturnReasonCodesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this list return reason codes service unavailable response has a 5xx status code
func (o *ListReturnReasonCodesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this list return reason codes service unavailable response a status code equal to that given
func (o *ListReturnReasonCodesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *ListReturnReasonCodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListReturnReasonCodesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/outbound/2020-07-01/returnReasonCodes][%d] listReturnReasonCodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ListReturnReasonCodesServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse {
	return o.Payload
}

func (o *ListReturnReasonCodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.ListReturnReasonCodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
