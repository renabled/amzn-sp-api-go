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

// CreateFulfillmentReturnReader is a Reader for the CreateFulfillmentReturn structure.
type CreateFulfillmentReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFulfillmentReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFulfillmentReturnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFulfillmentReturnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFulfillmentReturnUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFulfillmentReturnForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFulfillmentReturnNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateFulfillmentReturnTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateFulfillmentReturnInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateFulfillmentReturnServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFulfillmentReturnOK creates a CreateFulfillmentReturnOK with default headers values
func NewCreateFulfillmentReturnOK() *CreateFulfillmentReturnOK {
	return &CreateFulfillmentReturnOK{}
}

/*
CreateFulfillmentReturnOK describes a response with status code 200, with default header values.

Success.
*/
type CreateFulfillmentReturnOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return o k response has a 2xx status code
func (o *CreateFulfillmentReturnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create fulfillment return o k response has a 3xx status code
func (o *CreateFulfillmentReturnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return o k response has a 4xx status code
func (o *CreateFulfillmentReturnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create fulfillment return o k response has a 5xx status code
func (o *CreateFulfillmentReturnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create fulfillment return o k response a status code equal to that given
func (o *CreateFulfillmentReturnOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateFulfillmentReturnOK) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnOK  %+v", 200, o.Payload)
}

func (o *CreateFulfillmentReturnOK) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnOK  %+v", 200, o.Payload)
}

func (o *CreateFulfillmentReturnOK) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFulfillmentReturnBadRequest creates a CreateFulfillmentReturnBadRequest with default headers values
func NewCreateFulfillmentReturnBadRequest() *CreateFulfillmentReturnBadRequest {
	return &CreateFulfillmentReturnBadRequest{}
}

/*
CreateFulfillmentReturnBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateFulfillmentReturnBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return bad request response has a 2xx status code
func (o *CreateFulfillmentReturnBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create fulfillment return bad request response has a 3xx status code
func (o *CreateFulfillmentReturnBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return bad request response has a 4xx status code
func (o *CreateFulfillmentReturnBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create fulfillment return bad request response has a 5xx status code
func (o *CreateFulfillmentReturnBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create fulfillment return bad request response a status code equal to that given
func (o *CreateFulfillmentReturnBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateFulfillmentReturnBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnBadRequest  %+v", 400, o.Payload)
}

func (o *CreateFulfillmentReturnBadRequest) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnBadRequest  %+v", 400, o.Payload)
}

func (o *CreateFulfillmentReturnBadRequest) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFulfillmentReturnUnauthorized creates a CreateFulfillmentReturnUnauthorized with default headers values
func NewCreateFulfillmentReturnUnauthorized() *CreateFulfillmentReturnUnauthorized {
	return &CreateFulfillmentReturnUnauthorized{}
}

/*
CreateFulfillmentReturnUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CreateFulfillmentReturnUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return unauthorized response has a 2xx status code
func (o *CreateFulfillmentReturnUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create fulfillment return unauthorized response has a 3xx status code
func (o *CreateFulfillmentReturnUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return unauthorized response has a 4xx status code
func (o *CreateFulfillmentReturnUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create fulfillment return unauthorized response has a 5xx status code
func (o *CreateFulfillmentReturnUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create fulfillment return unauthorized response a status code equal to that given
func (o *CreateFulfillmentReturnUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateFulfillmentReturnUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateFulfillmentReturnUnauthorized) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateFulfillmentReturnUnauthorized) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFulfillmentReturnForbidden creates a CreateFulfillmentReturnForbidden with default headers values
func NewCreateFulfillmentReturnForbidden() *CreateFulfillmentReturnForbidden {
	return &CreateFulfillmentReturnForbidden{}
}

/*
CreateFulfillmentReturnForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateFulfillmentReturnForbidden struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return forbidden response has a 2xx status code
func (o *CreateFulfillmentReturnForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create fulfillment return forbidden response has a 3xx status code
func (o *CreateFulfillmentReturnForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return forbidden response has a 4xx status code
func (o *CreateFulfillmentReturnForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create fulfillment return forbidden response has a 5xx status code
func (o *CreateFulfillmentReturnForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create fulfillment return forbidden response a status code equal to that given
func (o *CreateFulfillmentReturnForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateFulfillmentReturnForbidden) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnForbidden  %+v", 403, o.Payload)
}

func (o *CreateFulfillmentReturnForbidden) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnForbidden  %+v", 403, o.Payload)
}

func (o *CreateFulfillmentReturnForbidden) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFulfillmentReturnNotFound creates a CreateFulfillmentReturnNotFound with default headers values
func NewCreateFulfillmentReturnNotFound() *CreateFulfillmentReturnNotFound {
	return &CreateFulfillmentReturnNotFound{}
}

/*
CreateFulfillmentReturnNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CreateFulfillmentReturnNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return not found response has a 2xx status code
func (o *CreateFulfillmentReturnNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create fulfillment return not found response has a 3xx status code
func (o *CreateFulfillmentReturnNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return not found response has a 4xx status code
func (o *CreateFulfillmentReturnNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create fulfillment return not found response has a 5xx status code
func (o *CreateFulfillmentReturnNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create fulfillment return not found response a status code equal to that given
func (o *CreateFulfillmentReturnNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateFulfillmentReturnNotFound) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnNotFound  %+v", 404, o.Payload)
}

func (o *CreateFulfillmentReturnNotFound) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnNotFound  %+v", 404, o.Payload)
}

func (o *CreateFulfillmentReturnNotFound) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFulfillmentReturnTooManyRequests creates a CreateFulfillmentReturnTooManyRequests with default headers values
func NewCreateFulfillmentReturnTooManyRequests() *CreateFulfillmentReturnTooManyRequests {
	return &CreateFulfillmentReturnTooManyRequests{}
}

/*
CreateFulfillmentReturnTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateFulfillmentReturnTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return too many requests response has a 2xx status code
func (o *CreateFulfillmentReturnTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create fulfillment return too many requests response has a 3xx status code
func (o *CreateFulfillmentReturnTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return too many requests response has a 4xx status code
func (o *CreateFulfillmentReturnTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create fulfillment return too many requests response has a 5xx status code
func (o *CreateFulfillmentReturnTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create fulfillment return too many requests response a status code equal to that given
func (o *CreateFulfillmentReturnTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateFulfillmentReturnTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateFulfillmentReturnTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateFulfillmentReturnTooManyRequests) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFulfillmentReturnInternalServerError creates a CreateFulfillmentReturnInternalServerError with default headers values
func NewCreateFulfillmentReturnInternalServerError() *CreateFulfillmentReturnInternalServerError {
	return &CreateFulfillmentReturnInternalServerError{}
}

/*
CreateFulfillmentReturnInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateFulfillmentReturnInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return internal server error response has a 2xx status code
func (o *CreateFulfillmentReturnInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create fulfillment return internal server error response has a 3xx status code
func (o *CreateFulfillmentReturnInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return internal server error response has a 4xx status code
func (o *CreateFulfillmentReturnInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create fulfillment return internal server error response has a 5xx status code
func (o *CreateFulfillmentReturnInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create fulfillment return internal server error response a status code equal to that given
func (o *CreateFulfillmentReturnInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateFulfillmentReturnInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateFulfillmentReturnInternalServerError) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateFulfillmentReturnInternalServerError) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFulfillmentReturnServiceUnavailable creates a CreateFulfillmentReturnServiceUnavailable with default headers values
func NewCreateFulfillmentReturnServiceUnavailable() *CreateFulfillmentReturnServiceUnavailable {
	return &CreateFulfillmentReturnServiceUnavailable{}
}

/*
CreateFulfillmentReturnServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateFulfillmentReturnServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse
}

// IsSuccess returns true when this create fulfillment return service unavailable response has a 2xx status code
func (o *CreateFulfillmentReturnServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create fulfillment return service unavailable response has a 3xx status code
func (o *CreateFulfillmentReturnServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create fulfillment return service unavailable response has a 4xx status code
func (o *CreateFulfillmentReturnServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create fulfillment return service unavailable response has a 5xx status code
func (o *CreateFulfillmentReturnServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create fulfillment return service unavailable response a status code equal to that given
func (o *CreateFulfillmentReturnServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateFulfillmentReturnServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateFulfillmentReturnServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /fba/outbound/2020-07-01/fulfillmentOrders/{sellerFulfillmentOrderId}/return][%d] createFulfillmentReturnServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateFulfillmentReturnServiceUnavailable) GetPayload() *fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse {
	return o.Payload
}

func (o *CreateFulfillmentReturnServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(fulfillment_outbound_2020_07_01_models.CreateFulfillmentReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
