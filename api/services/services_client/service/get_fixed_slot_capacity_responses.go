// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/services/services_models"
)

// GetFixedSlotCapacityReader is a Reader for the GetFixedSlotCapacity structure.
type GetFixedSlotCapacityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFixedSlotCapacityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFixedSlotCapacityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFixedSlotCapacityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFixedSlotCapacityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFixedSlotCapacityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFixedSlotCapacityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFixedSlotCapacityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFixedSlotCapacityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFixedSlotCapacityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFixedSlotCapacityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFixedSlotCapacityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFixedSlotCapacityOK creates a GetFixedSlotCapacityOK with default headers values
func NewGetFixedSlotCapacityOK() *GetFixedSlotCapacityOK {
	return &GetFixedSlotCapacityOK{}
}

/*
GetFixedSlotCapacityOK describes a response with status code 200, with default header values.

Success response.
*/
type GetFixedSlotCapacityOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacity
}

// IsSuccess returns true when this get fixed slot capacity o k response has a 2xx status code
func (o *GetFixedSlotCapacityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fixed slot capacity o k response has a 3xx status code
func (o *GetFixedSlotCapacityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity o k response has a 4xx status code
func (o *GetFixedSlotCapacityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fixed slot capacity o k response has a 5xx status code
func (o *GetFixedSlotCapacityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity o k response a status code equal to that given
func (o *GetFixedSlotCapacityOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFixedSlotCapacityOK) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityOK  %+v", 200, o.Payload)
}

func (o *GetFixedSlotCapacityOK) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityOK  %+v", 200, o.Payload)
}

func (o *GetFixedSlotCapacityOK) GetPayload() *services_models.FixedSlotCapacity {
	return o.Payload
}

func (o *GetFixedSlotCapacityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.FixedSlotCapacity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityBadRequest creates a GetFixedSlotCapacityBadRequest with default headers values
func NewGetFixedSlotCapacityBadRequest() *GetFixedSlotCapacityBadRequest {
	return &GetFixedSlotCapacityBadRequest{}
}

/*
GetFixedSlotCapacityBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetFixedSlotCapacityBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity bad request response has a 2xx status code
func (o *GetFixedSlotCapacityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity bad request response has a 3xx status code
func (o *GetFixedSlotCapacityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity bad request response has a 4xx status code
func (o *GetFixedSlotCapacityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fixed slot capacity bad request response has a 5xx status code
func (o *GetFixedSlotCapacityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity bad request response a status code equal to that given
func (o *GetFixedSlotCapacityBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFixedSlotCapacityBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityBadRequest  %+v", 400, o.Payload)
}

func (o *GetFixedSlotCapacityBadRequest) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityBadRequest  %+v", 400, o.Payload)
}

func (o *GetFixedSlotCapacityBadRequest) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityUnauthorized creates a GetFixedSlotCapacityUnauthorized with default headers values
func NewGetFixedSlotCapacityUnauthorized() *GetFixedSlotCapacityUnauthorized {
	return &GetFixedSlotCapacityUnauthorized{}
}

/*
GetFixedSlotCapacityUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetFixedSlotCapacityUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity unauthorized response has a 2xx status code
func (o *GetFixedSlotCapacityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity unauthorized response has a 3xx status code
func (o *GetFixedSlotCapacityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity unauthorized response has a 4xx status code
func (o *GetFixedSlotCapacityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fixed slot capacity unauthorized response has a 5xx status code
func (o *GetFixedSlotCapacityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity unauthorized response a status code equal to that given
func (o *GetFixedSlotCapacityUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFixedSlotCapacityUnauthorized) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFixedSlotCapacityUnauthorized) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFixedSlotCapacityUnauthorized) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityForbidden creates a GetFixedSlotCapacityForbidden with default headers values
func NewGetFixedSlotCapacityForbidden() *GetFixedSlotCapacityForbidden {
	return &GetFixedSlotCapacityForbidden{}
}

/*
GetFixedSlotCapacityForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetFixedSlotCapacityForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity forbidden response has a 2xx status code
func (o *GetFixedSlotCapacityForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity forbidden response has a 3xx status code
func (o *GetFixedSlotCapacityForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity forbidden response has a 4xx status code
func (o *GetFixedSlotCapacityForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fixed slot capacity forbidden response has a 5xx status code
func (o *GetFixedSlotCapacityForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity forbidden response a status code equal to that given
func (o *GetFixedSlotCapacityForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFixedSlotCapacityForbidden) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityForbidden  %+v", 403, o.Payload)
}

func (o *GetFixedSlotCapacityForbidden) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityForbidden  %+v", 403, o.Payload)
}

func (o *GetFixedSlotCapacityForbidden) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityNotFound creates a GetFixedSlotCapacityNotFound with default headers values
func NewGetFixedSlotCapacityNotFound() *GetFixedSlotCapacityNotFound {
	return &GetFixedSlotCapacityNotFound{}
}

/*
GetFixedSlotCapacityNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetFixedSlotCapacityNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity not found response has a 2xx status code
func (o *GetFixedSlotCapacityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity not found response has a 3xx status code
func (o *GetFixedSlotCapacityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity not found response has a 4xx status code
func (o *GetFixedSlotCapacityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fixed slot capacity not found response has a 5xx status code
func (o *GetFixedSlotCapacityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity not found response a status code equal to that given
func (o *GetFixedSlotCapacityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFixedSlotCapacityNotFound) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityNotFound  %+v", 404, o.Payload)
}

func (o *GetFixedSlotCapacityNotFound) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityNotFound  %+v", 404, o.Payload)
}

func (o *GetFixedSlotCapacityNotFound) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityRequestEntityTooLarge creates a GetFixedSlotCapacityRequestEntityTooLarge with default headers values
func NewGetFixedSlotCapacityRequestEntityTooLarge() *GetFixedSlotCapacityRequestEntityTooLarge {
	return &GetFixedSlotCapacityRequestEntityTooLarge{}
}

/*
GetFixedSlotCapacityRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetFixedSlotCapacityRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity request entity too large response has a 2xx status code
func (o *GetFixedSlotCapacityRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity request entity too large response has a 3xx status code
func (o *GetFixedSlotCapacityRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity request entity too large response has a 4xx status code
func (o *GetFixedSlotCapacityRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fixed slot capacity request entity too large response has a 5xx status code
func (o *GetFixedSlotCapacityRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity request entity too large response a status code equal to that given
func (o *GetFixedSlotCapacityRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFixedSlotCapacityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFixedSlotCapacityRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFixedSlotCapacityRequestEntityTooLarge) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityUnsupportedMediaType creates a GetFixedSlotCapacityUnsupportedMediaType with default headers values
func NewGetFixedSlotCapacityUnsupportedMediaType() *GetFixedSlotCapacityUnsupportedMediaType {
	return &GetFixedSlotCapacityUnsupportedMediaType{}
}

/*
GetFixedSlotCapacityUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetFixedSlotCapacityUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity unsupported media type response has a 2xx status code
func (o *GetFixedSlotCapacityUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity unsupported media type response has a 3xx status code
func (o *GetFixedSlotCapacityUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity unsupported media type response has a 4xx status code
func (o *GetFixedSlotCapacityUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fixed slot capacity unsupported media type response has a 5xx status code
func (o *GetFixedSlotCapacityUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity unsupported media type response a status code equal to that given
func (o *GetFixedSlotCapacityUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFixedSlotCapacityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFixedSlotCapacityUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFixedSlotCapacityUnsupportedMediaType) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityTooManyRequests creates a GetFixedSlotCapacityTooManyRequests with default headers values
func NewGetFixedSlotCapacityTooManyRequests() *GetFixedSlotCapacityTooManyRequests {
	return &GetFixedSlotCapacityTooManyRequests{}
}

/*
GetFixedSlotCapacityTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetFixedSlotCapacityTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity too many requests response has a 2xx status code
func (o *GetFixedSlotCapacityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity too many requests response has a 3xx status code
func (o *GetFixedSlotCapacityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity too many requests response has a 4xx status code
func (o *GetFixedSlotCapacityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fixed slot capacity too many requests response has a 5xx status code
func (o *GetFixedSlotCapacityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get fixed slot capacity too many requests response a status code equal to that given
func (o *GetFixedSlotCapacityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFixedSlotCapacityTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFixedSlotCapacityTooManyRequests) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFixedSlotCapacityTooManyRequests) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityInternalServerError creates a GetFixedSlotCapacityInternalServerError with default headers values
func NewGetFixedSlotCapacityInternalServerError() *GetFixedSlotCapacityInternalServerError {
	return &GetFixedSlotCapacityInternalServerError{}
}

/*
GetFixedSlotCapacityInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetFixedSlotCapacityInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity internal server error response has a 2xx status code
func (o *GetFixedSlotCapacityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity internal server error response has a 3xx status code
func (o *GetFixedSlotCapacityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity internal server error response has a 4xx status code
func (o *GetFixedSlotCapacityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fixed slot capacity internal server error response has a 5xx status code
func (o *GetFixedSlotCapacityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fixed slot capacity internal server error response a status code equal to that given
func (o *GetFixedSlotCapacityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFixedSlotCapacityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFixedSlotCapacityInternalServerError) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFixedSlotCapacityInternalServerError) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFixedSlotCapacityServiceUnavailable creates a GetFixedSlotCapacityServiceUnavailable with default headers values
func NewGetFixedSlotCapacityServiceUnavailable() *GetFixedSlotCapacityServiceUnavailable {
	return &GetFixedSlotCapacityServiceUnavailable{}
}

/*
GetFixedSlotCapacityServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetFixedSlotCapacityServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.FixedSlotCapacityErrors
}

// IsSuccess returns true when this get fixed slot capacity service unavailable response has a 2xx status code
func (o *GetFixedSlotCapacityServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fixed slot capacity service unavailable response has a 3xx status code
func (o *GetFixedSlotCapacityServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fixed slot capacity service unavailable response has a 4xx status code
func (o *GetFixedSlotCapacityServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fixed slot capacity service unavailable response has a 5xx status code
func (o *GetFixedSlotCapacityServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get fixed slot capacity service unavailable response a status code equal to that given
func (o *GetFixedSlotCapacityServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFixedSlotCapacityServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFixedSlotCapacityServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/fixed][%d] getFixedSlotCapacityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFixedSlotCapacityServiceUnavailable) GetPayload() *services_models.FixedSlotCapacityErrors {
	return o.Payload
}

func (o *GetFixedSlotCapacityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.FixedSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
