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

// GetRangeSlotCapacityReader is a Reader for the GetRangeSlotCapacity structure.
type GetRangeSlotCapacityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRangeSlotCapacityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRangeSlotCapacityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRangeSlotCapacityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRangeSlotCapacityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRangeSlotCapacityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRangeSlotCapacityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRangeSlotCapacityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRangeSlotCapacityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRangeSlotCapacityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRangeSlotCapacityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRangeSlotCapacityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRangeSlotCapacityOK creates a GetRangeSlotCapacityOK with default headers values
func NewGetRangeSlotCapacityOK() *GetRangeSlotCapacityOK {
	return &GetRangeSlotCapacityOK{}
}

/*
GetRangeSlotCapacityOK describes a response with status code 200, with default header values.

Success response.
*/
type GetRangeSlotCapacityOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacity
}

// IsSuccess returns true when this get range slot capacity o k response has a 2xx status code
func (o *GetRangeSlotCapacityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get range slot capacity o k response has a 3xx status code
func (o *GetRangeSlotCapacityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity o k response has a 4xx status code
func (o *GetRangeSlotCapacityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get range slot capacity o k response has a 5xx status code
func (o *GetRangeSlotCapacityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity o k response a status code equal to that given
func (o *GetRangeSlotCapacityOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRangeSlotCapacityOK) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityOK  %+v", 200, o.Payload)
}

func (o *GetRangeSlotCapacityOK) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityOK  %+v", 200, o.Payload)
}

func (o *GetRangeSlotCapacityOK) GetPayload() *services_models.RangeSlotCapacity {
	return o.Payload
}

func (o *GetRangeSlotCapacityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.RangeSlotCapacity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityBadRequest creates a GetRangeSlotCapacityBadRequest with default headers values
func NewGetRangeSlotCapacityBadRequest() *GetRangeSlotCapacityBadRequest {
	return &GetRangeSlotCapacityBadRequest{}
}

/*
GetRangeSlotCapacityBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetRangeSlotCapacityBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity bad request response has a 2xx status code
func (o *GetRangeSlotCapacityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity bad request response has a 3xx status code
func (o *GetRangeSlotCapacityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity bad request response has a 4xx status code
func (o *GetRangeSlotCapacityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get range slot capacity bad request response has a 5xx status code
func (o *GetRangeSlotCapacityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity bad request response a status code equal to that given
func (o *GetRangeSlotCapacityBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRangeSlotCapacityBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityBadRequest  %+v", 400, o.Payload)
}

func (o *GetRangeSlotCapacityBadRequest) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityBadRequest  %+v", 400, o.Payload)
}

func (o *GetRangeSlotCapacityBadRequest) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityUnauthorized creates a GetRangeSlotCapacityUnauthorized with default headers values
func NewGetRangeSlotCapacityUnauthorized() *GetRangeSlotCapacityUnauthorized {
	return &GetRangeSlotCapacityUnauthorized{}
}

/*
GetRangeSlotCapacityUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetRangeSlotCapacityUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity unauthorized response has a 2xx status code
func (o *GetRangeSlotCapacityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity unauthorized response has a 3xx status code
func (o *GetRangeSlotCapacityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity unauthorized response has a 4xx status code
func (o *GetRangeSlotCapacityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get range slot capacity unauthorized response has a 5xx status code
func (o *GetRangeSlotCapacityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity unauthorized response a status code equal to that given
func (o *GetRangeSlotCapacityUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRangeSlotCapacityUnauthorized) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRangeSlotCapacityUnauthorized) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRangeSlotCapacityUnauthorized) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityForbidden creates a GetRangeSlotCapacityForbidden with default headers values
func NewGetRangeSlotCapacityForbidden() *GetRangeSlotCapacityForbidden {
	return &GetRangeSlotCapacityForbidden{}
}

/*
GetRangeSlotCapacityForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetRangeSlotCapacityForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity forbidden response has a 2xx status code
func (o *GetRangeSlotCapacityForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity forbidden response has a 3xx status code
func (o *GetRangeSlotCapacityForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity forbidden response has a 4xx status code
func (o *GetRangeSlotCapacityForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get range slot capacity forbidden response has a 5xx status code
func (o *GetRangeSlotCapacityForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity forbidden response a status code equal to that given
func (o *GetRangeSlotCapacityForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRangeSlotCapacityForbidden) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityForbidden  %+v", 403, o.Payload)
}

func (o *GetRangeSlotCapacityForbidden) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityForbidden  %+v", 403, o.Payload)
}

func (o *GetRangeSlotCapacityForbidden) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityNotFound creates a GetRangeSlotCapacityNotFound with default headers values
func NewGetRangeSlotCapacityNotFound() *GetRangeSlotCapacityNotFound {
	return &GetRangeSlotCapacityNotFound{}
}

/*
GetRangeSlotCapacityNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetRangeSlotCapacityNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity not found response has a 2xx status code
func (o *GetRangeSlotCapacityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity not found response has a 3xx status code
func (o *GetRangeSlotCapacityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity not found response has a 4xx status code
func (o *GetRangeSlotCapacityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get range slot capacity not found response has a 5xx status code
func (o *GetRangeSlotCapacityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity not found response a status code equal to that given
func (o *GetRangeSlotCapacityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRangeSlotCapacityNotFound) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityNotFound  %+v", 404, o.Payload)
}

func (o *GetRangeSlotCapacityNotFound) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityNotFound  %+v", 404, o.Payload)
}

func (o *GetRangeSlotCapacityNotFound) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityRequestEntityTooLarge creates a GetRangeSlotCapacityRequestEntityTooLarge with default headers values
func NewGetRangeSlotCapacityRequestEntityTooLarge() *GetRangeSlotCapacityRequestEntityTooLarge {
	return &GetRangeSlotCapacityRequestEntityTooLarge{}
}

/*
GetRangeSlotCapacityRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetRangeSlotCapacityRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity request entity too large response has a 2xx status code
func (o *GetRangeSlotCapacityRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity request entity too large response has a 3xx status code
func (o *GetRangeSlotCapacityRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity request entity too large response has a 4xx status code
func (o *GetRangeSlotCapacityRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get range slot capacity request entity too large response has a 5xx status code
func (o *GetRangeSlotCapacityRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity request entity too large response a status code equal to that given
func (o *GetRangeSlotCapacityRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRangeSlotCapacityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRangeSlotCapacityRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRangeSlotCapacityRequestEntityTooLarge) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityUnsupportedMediaType creates a GetRangeSlotCapacityUnsupportedMediaType with default headers values
func NewGetRangeSlotCapacityUnsupportedMediaType() *GetRangeSlotCapacityUnsupportedMediaType {
	return &GetRangeSlotCapacityUnsupportedMediaType{}
}

/*
GetRangeSlotCapacityUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetRangeSlotCapacityUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity unsupported media type response has a 2xx status code
func (o *GetRangeSlotCapacityUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity unsupported media type response has a 3xx status code
func (o *GetRangeSlotCapacityUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity unsupported media type response has a 4xx status code
func (o *GetRangeSlotCapacityUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get range slot capacity unsupported media type response has a 5xx status code
func (o *GetRangeSlotCapacityUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity unsupported media type response a status code equal to that given
func (o *GetRangeSlotCapacityUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRangeSlotCapacityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRangeSlotCapacityUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRangeSlotCapacityUnsupportedMediaType) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityTooManyRequests creates a GetRangeSlotCapacityTooManyRequests with default headers values
func NewGetRangeSlotCapacityTooManyRequests() *GetRangeSlotCapacityTooManyRequests {
	return &GetRangeSlotCapacityTooManyRequests{}
}

/*
GetRangeSlotCapacityTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetRangeSlotCapacityTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity too many requests response has a 2xx status code
func (o *GetRangeSlotCapacityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity too many requests response has a 3xx status code
func (o *GetRangeSlotCapacityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity too many requests response has a 4xx status code
func (o *GetRangeSlotCapacityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get range slot capacity too many requests response has a 5xx status code
func (o *GetRangeSlotCapacityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get range slot capacity too many requests response a status code equal to that given
func (o *GetRangeSlotCapacityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRangeSlotCapacityTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRangeSlotCapacityTooManyRequests) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRangeSlotCapacityTooManyRequests) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityInternalServerError creates a GetRangeSlotCapacityInternalServerError with default headers values
func NewGetRangeSlotCapacityInternalServerError() *GetRangeSlotCapacityInternalServerError {
	return &GetRangeSlotCapacityInternalServerError{}
}

/*
GetRangeSlotCapacityInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetRangeSlotCapacityInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity internal server error response has a 2xx status code
func (o *GetRangeSlotCapacityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity internal server error response has a 3xx status code
func (o *GetRangeSlotCapacityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity internal server error response has a 4xx status code
func (o *GetRangeSlotCapacityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get range slot capacity internal server error response has a 5xx status code
func (o *GetRangeSlotCapacityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get range slot capacity internal server error response a status code equal to that given
func (o *GetRangeSlotCapacityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRangeSlotCapacityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRangeSlotCapacityInternalServerError) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRangeSlotCapacityInternalServerError) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRangeSlotCapacityServiceUnavailable creates a GetRangeSlotCapacityServiceUnavailable with default headers values
func NewGetRangeSlotCapacityServiceUnavailable() *GetRangeSlotCapacityServiceUnavailable {
	return &GetRangeSlotCapacityServiceUnavailable{}
}

/*
GetRangeSlotCapacityServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetRangeSlotCapacityServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.RangeSlotCapacityErrors
}

// IsSuccess returns true when this get range slot capacity service unavailable response has a 2xx status code
func (o *GetRangeSlotCapacityServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get range slot capacity service unavailable response has a 3xx status code
func (o *GetRangeSlotCapacityServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get range slot capacity service unavailable response has a 4xx status code
func (o *GetRangeSlotCapacityServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get range slot capacity service unavailable response has a 5xx status code
func (o *GetRangeSlotCapacityServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get range slot capacity service unavailable response a status code equal to that given
func (o *GetRangeSlotCapacityServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRangeSlotCapacityServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRangeSlotCapacityServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /service/v1/serviceResources/{resourceId}/capacity/range][%d] getRangeSlotCapacityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRangeSlotCapacityServiceUnavailable) GetPayload() *services_models.RangeSlotCapacityErrors {
	return o.Payload
}

func (o *GetRangeSlotCapacityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.RangeSlotCapacityErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
