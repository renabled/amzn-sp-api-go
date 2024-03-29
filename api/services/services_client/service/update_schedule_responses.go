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

// UpdateScheduleReader is a Reader for the UpdateSchedule structure.
type UpdateScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewUpdateScheduleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewUpdateScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateScheduleOK creates a UpdateScheduleOK with default headers values
func NewUpdateScheduleOK() *UpdateScheduleOK {
	return &UpdateScheduleOK{}
}

/*
UpdateScheduleOK describes a response with status code 200, with default header values.

Success response.
*/
type UpdateScheduleOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule o k response has a 2xx status code
func (o *UpdateScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update schedule o k response has a 3xx status code
func (o *UpdateScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule o k response has a 4xx status code
func (o *UpdateScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update schedule o k response has a 5xx status code
func (o *UpdateScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update schedule o k response a status code equal to that given
func (o *UpdateScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateScheduleOK) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleOK  %+v", 200, o.Payload)
}

func (o *UpdateScheduleOK) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleOK  %+v", 200, o.Payload)
}

func (o *UpdateScheduleOK) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleBadRequest creates a UpdateScheduleBadRequest with default headers values
func NewUpdateScheduleBadRequest() *UpdateScheduleBadRequest {
	return &UpdateScheduleBadRequest{}
}

/*
UpdateScheduleBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateScheduleBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule bad request response has a 2xx status code
func (o *UpdateScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule bad request response has a 3xx status code
func (o *UpdateScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule bad request response has a 4xx status code
func (o *UpdateScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update schedule bad request response has a 5xx status code
func (o *UpdateScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update schedule bad request response a status code equal to that given
func (o *UpdateScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateScheduleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateScheduleBadRequest) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateScheduleBadRequest) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleForbidden creates a UpdateScheduleForbidden with default headers values
func NewUpdateScheduleForbidden() *UpdateScheduleForbidden {
	return &UpdateScheduleForbidden{}
}

/*
UpdateScheduleForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateScheduleForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule forbidden response has a 2xx status code
func (o *UpdateScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule forbidden response has a 3xx status code
func (o *UpdateScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule forbidden response has a 4xx status code
func (o *UpdateScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update schedule forbidden response has a 5xx status code
func (o *UpdateScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update schedule forbidden response a status code equal to that given
func (o *UpdateScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateScheduleForbidden) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateScheduleForbidden) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateScheduleForbidden) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleNotFound creates a UpdateScheduleNotFound with default headers values
func NewUpdateScheduleNotFound() *UpdateScheduleNotFound {
	return &UpdateScheduleNotFound{}
}

/*
UpdateScheduleNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type UpdateScheduleNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule not found response has a 2xx status code
func (o *UpdateScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule not found response has a 3xx status code
func (o *UpdateScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule not found response has a 4xx status code
func (o *UpdateScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update schedule not found response has a 5xx status code
func (o *UpdateScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update schedule not found response a status code equal to that given
func (o *UpdateScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateScheduleNotFound) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateScheduleNotFound) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateScheduleNotFound) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleRequestEntityTooLarge creates a UpdateScheduleRequestEntityTooLarge with default headers values
func NewUpdateScheduleRequestEntityTooLarge() *UpdateScheduleRequestEntityTooLarge {
	return &UpdateScheduleRequestEntityTooLarge{}
}

/*
UpdateScheduleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type UpdateScheduleRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	**Note:** For this status code, the rate limit header is deprecated and no longer returned.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule request entity too large response has a 2xx status code
func (o *UpdateScheduleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule request entity too large response has a 3xx status code
func (o *UpdateScheduleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule request entity too large response has a 4xx status code
func (o *UpdateScheduleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this update schedule request entity too large response has a 5xx status code
func (o *UpdateScheduleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this update schedule request entity too large response a status code equal to that given
func (o *UpdateScheduleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *UpdateScheduleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateScheduleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateScheduleRequestEntityTooLarge) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleUnsupportedMediaType creates a UpdateScheduleUnsupportedMediaType with default headers values
func NewUpdateScheduleUnsupportedMediaType() *UpdateScheduleUnsupportedMediaType {
	return &UpdateScheduleUnsupportedMediaType{}
}

/*
UpdateScheduleUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type UpdateScheduleUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	**Note:** For this status code, the rate limit header is deprecated and no longer returned.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule unsupported media type response has a 2xx status code
func (o *UpdateScheduleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule unsupported media type response has a 3xx status code
func (o *UpdateScheduleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule unsupported media type response has a 4xx status code
func (o *UpdateScheduleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this update schedule unsupported media type response has a 5xx status code
func (o *UpdateScheduleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this update schedule unsupported media type response a status code equal to that given
func (o *UpdateScheduleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *UpdateScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateScheduleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateScheduleUnsupportedMediaType) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleTooManyRequests creates a UpdateScheduleTooManyRequests with default headers values
func NewUpdateScheduleTooManyRequests() *UpdateScheduleTooManyRequests {
	return &UpdateScheduleTooManyRequests{}
}

/*
UpdateScheduleTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateScheduleTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	**Note:** For this status code, the rate limit header is deprecated and no longer returned.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule too many requests response has a 2xx status code
func (o *UpdateScheduleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule too many requests response has a 3xx status code
func (o *UpdateScheduleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule too many requests response has a 4xx status code
func (o *UpdateScheduleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update schedule too many requests response has a 5xx status code
func (o *UpdateScheduleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update schedule too many requests response a status code equal to that given
func (o *UpdateScheduleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateScheduleTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateScheduleTooManyRequests) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleInternalServerError creates a UpdateScheduleInternalServerError with default headers values
func NewUpdateScheduleInternalServerError() *UpdateScheduleInternalServerError {
	return &UpdateScheduleInternalServerError{}
}

/*
UpdateScheduleInternalServerError describes a response with status code 500, with default header values.

Encountered an unexpected condition which prevented the server from fulfilling the request.
*/
type UpdateScheduleInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	**Note:** For this status code, the rate limit header is deprecated and no longer returned.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule internal server error response has a 2xx status code
func (o *UpdateScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule internal server error response has a 3xx status code
func (o *UpdateScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule internal server error response has a 4xx status code
func (o *UpdateScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update schedule internal server error response has a 5xx status code
func (o *UpdateScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update schedule internal server error response a status code equal to that given
func (o *UpdateScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateScheduleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateScheduleInternalServerError) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScheduleServiceUnavailable creates a UpdateScheduleServiceUnavailable with default headers values
func NewUpdateScheduleServiceUnavailable() *UpdateScheduleServiceUnavailable {
	return &UpdateScheduleServiceUnavailable{}
}

/*
UpdateScheduleServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateScheduleServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	**Note:** For this status code, the rate limit header is deprecated and no longer returned.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *services_models.UpdateScheduleResponse
}

// IsSuccess returns true when this update schedule service unavailable response has a 2xx status code
func (o *UpdateScheduleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update schedule service unavailable response has a 3xx status code
func (o *UpdateScheduleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update schedule service unavailable response has a 4xx status code
func (o *UpdateScheduleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update schedule service unavailable response has a 5xx status code
func (o *UpdateScheduleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update schedule service unavailable response a status code equal to that given
func (o *UpdateScheduleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *UpdateScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateScheduleServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /service/v1/serviceResources/{resourceId}/schedules][%d] updateScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateScheduleServiceUnavailable) GetPayload() *services_models.UpdateScheduleResponse {
	return o.Payload
}

func (o *UpdateScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(services_models.UpdateScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
