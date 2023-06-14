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

// GetTransportDetailsReader is a Reader for the GetTransportDetails structure.
type GetTransportDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransportDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransportDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTransportDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTransportDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTransportDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTransportDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTransportDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTransportDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTransportDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransportDetailsOK creates a GetTransportDetailsOK with default headers values
func NewGetTransportDetailsOK() *GetTransportDetailsOK {
	return &GetTransportDetailsOK{}
}

/*
GetTransportDetailsOK describes a response with status code 200, with default header values.

Success.
*/
type GetTransportDetailsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details o k response has a 2xx status code
func (o *GetTransportDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get transport details o k response has a 3xx status code
func (o *GetTransportDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details o k response has a 4xx status code
func (o *GetTransportDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get transport details o k response has a 5xx status code
func (o *GetTransportDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get transport details o k response a status code equal to that given
func (o *GetTransportDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTransportDetailsOK) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsOK  %+v", 200, o.Payload)
}

func (o *GetTransportDetailsOK) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsOK  %+v", 200, o.Payload)
}

func (o *GetTransportDetailsOK) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransportDetailsBadRequest creates a GetTransportDetailsBadRequest with default headers values
func NewGetTransportDetailsBadRequest() *GetTransportDetailsBadRequest {
	return &GetTransportDetailsBadRequest{}
}

/*
GetTransportDetailsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetTransportDetailsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details bad request response has a 2xx status code
func (o *GetTransportDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transport details bad request response has a 3xx status code
func (o *GetTransportDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details bad request response has a 4xx status code
func (o *GetTransportDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transport details bad request response has a 5xx status code
func (o *GetTransportDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get transport details bad request response a status code equal to that given
func (o *GetTransportDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTransportDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransportDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransportDetailsBadRequest) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransportDetailsUnauthorized creates a GetTransportDetailsUnauthorized with default headers values
func NewGetTransportDetailsUnauthorized() *GetTransportDetailsUnauthorized {
	return &GetTransportDetailsUnauthorized{}
}

/*
GetTransportDetailsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetTransportDetailsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details unauthorized response has a 2xx status code
func (o *GetTransportDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transport details unauthorized response has a 3xx status code
func (o *GetTransportDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details unauthorized response has a 4xx status code
func (o *GetTransportDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transport details unauthorized response has a 5xx status code
func (o *GetTransportDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get transport details unauthorized response a status code equal to that given
func (o *GetTransportDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTransportDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransportDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTransportDetailsUnauthorized) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransportDetailsForbidden creates a GetTransportDetailsForbidden with default headers values
func NewGetTransportDetailsForbidden() *GetTransportDetailsForbidden {
	return &GetTransportDetailsForbidden{}
}

/*
GetTransportDetailsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetTransportDetailsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details forbidden response has a 2xx status code
func (o *GetTransportDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transport details forbidden response has a 3xx status code
func (o *GetTransportDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details forbidden response has a 4xx status code
func (o *GetTransportDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transport details forbidden response has a 5xx status code
func (o *GetTransportDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get transport details forbidden response a status code equal to that given
func (o *GetTransportDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTransportDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetTransportDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetTransportDetailsForbidden) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransportDetailsNotFound creates a GetTransportDetailsNotFound with default headers values
func NewGetTransportDetailsNotFound() *GetTransportDetailsNotFound {
	return &GetTransportDetailsNotFound{}
}

/*
GetTransportDetailsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetTransportDetailsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details not found response has a 2xx status code
func (o *GetTransportDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transport details not found response has a 3xx status code
func (o *GetTransportDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details not found response has a 4xx status code
func (o *GetTransportDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transport details not found response has a 5xx status code
func (o *GetTransportDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get transport details not found response a status code equal to that given
func (o *GetTransportDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTransportDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetTransportDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetTransportDetailsNotFound) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransportDetailsTooManyRequests creates a GetTransportDetailsTooManyRequests with default headers values
func NewGetTransportDetailsTooManyRequests() *GetTransportDetailsTooManyRequests {
	return &GetTransportDetailsTooManyRequests{}
}

/*
GetTransportDetailsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetTransportDetailsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details too many requests response has a 2xx status code
func (o *GetTransportDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transport details too many requests response has a 3xx status code
func (o *GetTransportDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details too many requests response has a 4xx status code
func (o *GetTransportDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get transport details too many requests response has a 5xx status code
func (o *GetTransportDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get transport details too many requests response a status code equal to that given
func (o *GetTransportDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTransportDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTransportDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTransportDetailsTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransportDetailsInternalServerError creates a GetTransportDetailsInternalServerError with default headers values
func NewGetTransportDetailsInternalServerError() *GetTransportDetailsInternalServerError {
	return &GetTransportDetailsInternalServerError{}
}

/*
GetTransportDetailsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetTransportDetailsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details internal server error response has a 2xx status code
func (o *GetTransportDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transport details internal server error response has a 3xx status code
func (o *GetTransportDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details internal server error response has a 4xx status code
func (o *GetTransportDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get transport details internal server error response has a 5xx status code
func (o *GetTransportDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get transport details internal server error response a status code equal to that given
func (o *GetTransportDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTransportDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTransportDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTransportDetailsInternalServerError) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransportDetailsServiceUnavailable creates a GetTransportDetailsServiceUnavailable with default headers values
func NewGetTransportDetailsServiceUnavailable() *GetTransportDetailsServiceUnavailable {
	return &GetTransportDetailsServiceUnavailable{}
}

/*
GetTransportDetailsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetTransportDetailsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetTransportDetailsResponse
}

// IsSuccess returns true when this get transport details service unavailable response has a 2xx status code
func (o *GetTransportDetailsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get transport details service unavailable response has a 3xx status code
func (o *GetTransportDetailsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get transport details service unavailable response has a 4xx status code
func (o *GetTransportDetailsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get transport details service unavailable response has a 5xx status code
func (o *GetTransportDetailsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get transport details service unavailable response a status code equal to that given
func (o *GetTransportDetailsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTransportDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTransportDetailsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/shipments/{shipmentId}/transport][%d] getTransportDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTransportDetailsServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.GetTransportDetailsResponse {
	return o.Payload
}

func (o *GetTransportDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetTransportDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
