// Code generated by go-swagger; DO NOT EDIT.

package fba_inbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/fulfillmentInboundV0/fulfillment_inbound_v0_models"
)

// GetPrepInstructionsReader is a Reader for the GetPrepInstructions structure.
type GetPrepInstructionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrepInstructionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrepInstructionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPrepInstructionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPrepInstructionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPrepInstructionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPrepInstructionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPrepInstructionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPrepInstructionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPrepInstructionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPrepInstructionsOK creates a GetPrepInstructionsOK with default headers values
func NewGetPrepInstructionsOK() *GetPrepInstructionsOK {
	return &GetPrepInstructionsOK{}
}

/* GetPrepInstructionsOK describes a response with status code 200, with default header values.

Success.
*/
type GetPrepInstructionsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsOK) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsOK  %+v", 200, o.Payload)
}
func (o *GetPrepInstructionsOK) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrepInstructionsBadRequest creates a GetPrepInstructionsBadRequest with default headers values
func NewGetPrepInstructionsBadRequest() *GetPrepInstructionsBadRequest {
	return &GetPrepInstructionsBadRequest{}
}

/* GetPrepInstructionsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetPrepInstructionsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsBadRequest  %+v", 400, o.Payload)
}
func (o *GetPrepInstructionsBadRequest) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrepInstructionsUnauthorized creates a GetPrepInstructionsUnauthorized with default headers values
func NewGetPrepInstructionsUnauthorized() *GetPrepInstructionsUnauthorized {
	return &GetPrepInstructionsUnauthorized{}
}

/* GetPrepInstructionsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetPrepInstructionsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPrepInstructionsUnauthorized) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrepInstructionsForbidden creates a GetPrepInstructionsForbidden with default headers values
func NewGetPrepInstructionsForbidden() *GetPrepInstructionsForbidden {
	return &GetPrepInstructionsForbidden{}
}

/* GetPrepInstructionsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetPrepInstructionsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsForbidden) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsForbidden  %+v", 403, o.Payload)
}
func (o *GetPrepInstructionsForbidden) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrepInstructionsNotFound creates a GetPrepInstructionsNotFound with default headers values
func NewGetPrepInstructionsNotFound() *GetPrepInstructionsNotFound {
	return &GetPrepInstructionsNotFound{}
}

/* GetPrepInstructionsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetPrepInstructionsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsNotFound) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsNotFound  %+v", 404, o.Payload)
}
func (o *GetPrepInstructionsNotFound) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrepInstructionsTooManyRequests creates a GetPrepInstructionsTooManyRequests with default headers values
func NewGetPrepInstructionsTooManyRequests() *GetPrepInstructionsTooManyRequests {
	return &GetPrepInstructionsTooManyRequests{}
}

/* GetPrepInstructionsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetPrepInstructionsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetPrepInstructionsTooManyRequests) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrepInstructionsInternalServerError creates a GetPrepInstructionsInternalServerError with default headers values
func NewGetPrepInstructionsInternalServerError() *GetPrepInstructionsInternalServerError {
	return &GetPrepInstructionsInternalServerError{}
}

/* GetPrepInstructionsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetPrepInstructionsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPrepInstructionsInternalServerError) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrepInstructionsServiceUnavailable creates a GetPrepInstructionsServiceUnavailable with default headers values
func NewGetPrepInstructionsServiceUnavailable() *GetPrepInstructionsServiceUnavailable {
	return &GetPrepInstructionsServiceUnavailable{}
}

/* GetPrepInstructionsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetPrepInstructionsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_inbound_v0_models.GetPrepInstructionsResponse
}

func (o *GetPrepInstructionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fba/inbound/v0/prepInstructions][%d] getPrepInstructionsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetPrepInstructionsServiceUnavailable) GetPayload() *fulfillment_inbound_v0_models.GetPrepInstructionsResponse {
	return o.Payload
}

func (o *GetPrepInstructionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_inbound_v0_models.GetPrepInstructionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
