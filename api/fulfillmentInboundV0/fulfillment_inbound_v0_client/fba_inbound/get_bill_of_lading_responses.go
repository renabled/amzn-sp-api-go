// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/fulfillmentInboundV0/fulfillment_inbound_v0_models"
)

// GetBillOfLadingReader is a Reader for the GetBillOfLading structure.
type GetBillOfLadingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBillOfLadingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBillOfLadingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBillOfLadingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetBillOfLadingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBillOfLadingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBillOfLadingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetBillOfLadingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBillOfLadingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetBillOfLadingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBillOfLadingOK creates a GetBillOfLadingOK with default headers values
func NewGetBillOfLadingOK() *GetBillOfLadingOK {
	return &GetBillOfLadingOK{}
}

/*
GetBillOfLadingOK describes a response with status code 200, with default header values.

Success.
*/
type GetBillOfLadingOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading o k response has a 2xx status code
func (o *GetBillOfLadingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bill of lading o k response has a 3xx status code
func (o *GetBillOfLadingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading o k response has a 4xx status code
func (o *GetBillOfLadingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bill of lading o k response has a 5xx status code
func (o *GetBillOfLadingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bill of lading o k response a status code equal to that given
func (o *GetBillOfLadingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBillOfLadingOK) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingOK  %+v", 200, o.Payload)
}

func (o *GetBillOfLadingOK) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingOK  %+v", 200, o.Payload)
}

func (o *GetBillOfLadingOK) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillOfLadingBadRequest creates a GetBillOfLadingBadRequest with default headers values
func NewGetBillOfLadingBadRequest() *GetBillOfLadingBadRequest {
	return &GetBillOfLadingBadRequest{}
}

/*
GetBillOfLadingBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetBillOfLadingBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading bad request response has a 2xx status code
func (o *GetBillOfLadingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bill of lading bad request response has a 3xx status code
func (o *GetBillOfLadingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading bad request response has a 4xx status code
func (o *GetBillOfLadingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bill of lading bad request response has a 5xx status code
func (o *GetBillOfLadingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get bill of lading bad request response a status code equal to that given
func (o *GetBillOfLadingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetBillOfLadingBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingBadRequest  %+v", 400, o.Payload)
}

func (o *GetBillOfLadingBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingBadRequest  %+v", 400, o.Payload)
}

func (o *GetBillOfLadingBadRequest) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillOfLadingUnauthorized creates a GetBillOfLadingUnauthorized with default headers values
func NewGetBillOfLadingUnauthorized() *GetBillOfLadingUnauthorized {
	return &GetBillOfLadingUnauthorized{}
}

/*
GetBillOfLadingUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetBillOfLadingUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading unauthorized response has a 2xx status code
func (o *GetBillOfLadingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bill of lading unauthorized response has a 3xx status code
func (o *GetBillOfLadingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading unauthorized response has a 4xx status code
func (o *GetBillOfLadingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bill of lading unauthorized response has a 5xx status code
func (o *GetBillOfLadingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get bill of lading unauthorized response a status code equal to that given
func (o *GetBillOfLadingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetBillOfLadingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBillOfLadingUnauthorized) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBillOfLadingUnauthorized) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillOfLadingForbidden creates a GetBillOfLadingForbidden with default headers values
func NewGetBillOfLadingForbidden() *GetBillOfLadingForbidden {
	return &GetBillOfLadingForbidden{}
}

/*
GetBillOfLadingForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetBillOfLadingForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading forbidden response has a 2xx status code
func (o *GetBillOfLadingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bill of lading forbidden response has a 3xx status code
func (o *GetBillOfLadingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading forbidden response has a 4xx status code
func (o *GetBillOfLadingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bill of lading forbidden response has a 5xx status code
func (o *GetBillOfLadingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get bill of lading forbidden response a status code equal to that given
func (o *GetBillOfLadingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetBillOfLadingForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingForbidden  %+v", 403, o.Payload)
}

func (o *GetBillOfLadingForbidden) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingForbidden  %+v", 403, o.Payload)
}

func (o *GetBillOfLadingForbidden) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillOfLadingNotFound creates a GetBillOfLadingNotFound with default headers values
func NewGetBillOfLadingNotFound() *GetBillOfLadingNotFound {
	return &GetBillOfLadingNotFound{}
}

/*
GetBillOfLadingNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetBillOfLadingNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading not found response has a 2xx status code
func (o *GetBillOfLadingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bill of lading not found response has a 3xx status code
func (o *GetBillOfLadingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading not found response has a 4xx status code
func (o *GetBillOfLadingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bill of lading not found response has a 5xx status code
func (o *GetBillOfLadingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get bill of lading not found response a status code equal to that given
func (o *GetBillOfLadingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetBillOfLadingNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingNotFound  %+v", 404, o.Payload)
}

func (o *GetBillOfLadingNotFound) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingNotFound  %+v", 404, o.Payload)
}

func (o *GetBillOfLadingNotFound) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillOfLadingTooManyRequests creates a GetBillOfLadingTooManyRequests with default headers values
func NewGetBillOfLadingTooManyRequests() *GetBillOfLadingTooManyRequests {
	return &GetBillOfLadingTooManyRequests{}
}

/*
GetBillOfLadingTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetBillOfLadingTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading too many requests response has a 2xx status code
func (o *GetBillOfLadingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bill of lading too many requests response has a 3xx status code
func (o *GetBillOfLadingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading too many requests response has a 4xx status code
func (o *GetBillOfLadingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bill of lading too many requests response has a 5xx status code
func (o *GetBillOfLadingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get bill of lading too many requests response a status code equal to that given
func (o *GetBillOfLadingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetBillOfLadingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBillOfLadingTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetBillOfLadingTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillOfLadingInternalServerError creates a GetBillOfLadingInternalServerError with default headers values
func NewGetBillOfLadingInternalServerError() *GetBillOfLadingInternalServerError {
	return &GetBillOfLadingInternalServerError{}
}

/*
GetBillOfLadingInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetBillOfLadingInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading internal server error response has a 2xx status code
func (o *GetBillOfLadingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bill of lading internal server error response has a 3xx status code
func (o *GetBillOfLadingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading internal server error response has a 4xx status code
func (o *GetBillOfLadingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bill of lading internal server error response has a 5xx status code
func (o *GetBillOfLadingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get bill of lading internal server error response a status code equal to that given
func (o *GetBillOfLadingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetBillOfLadingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBillOfLadingInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetBillOfLadingInternalServerError) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBillOfLadingServiceUnavailable creates a GetBillOfLadingServiceUnavailable with default headers values
func NewGetBillOfLadingServiceUnavailable() *GetBillOfLadingServiceUnavailable {
	return &GetBillOfLadingServiceUnavailable{}
}

/*
GetBillOfLadingServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetBillOfLadingServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetBillOfLadingResponse
}

// IsSuccess returns true when this get bill of lading service unavailable response has a 2xx status code
func (o *GetBillOfLadingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bill of lading service unavailable response has a 3xx status code
func (o *GetBillOfLadingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bill of lading service unavailable response has a 4xx status code
func (o *GetBillOfLadingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bill of lading service unavailable response has a 5xx status code
func (o *GetBillOfLadingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get bill of lading service unavailable response a status code equal to that given
func (o *GetBillOfLadingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetBillOfLadingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetBillOfLadingServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/billOfLading][%d] getBillOfLadingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetBillOfLadingServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.GetBillOfLadingResponse {
	return o.Payload
}

func (o *GetBillOfLadingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetBillOfLadingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
