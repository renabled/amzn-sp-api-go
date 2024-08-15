// Code generated by go-swagger; DO NOT EDIT.

package awd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/awd_2024-05-09/awd_2024_05_09_models"
)

// UpdateInboundReader is a Reader for the UpdateInbound structure.
type UpdateInboundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInboundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInboundOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateInboundBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateInboundForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInboundNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateInboundConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewUpdateInboundRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewUpdateInboundUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateInboundTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateInboundInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateInboundServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInboundOK creates a UpdateInboundOK with default headers values
func NewUpdateInboundOK() *UpdateInboundOK {
	return &UpdateInboundOK{}
}

/*
UpdateInboundOK describes a response with status code 200, with default header values.

The 200 response for `updateInbound`.
*/
type UpdateInboundOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.InboundOrderReference
}

// IsSuccess returns true when this update inbound o k response has a 2xx status code
func (o *UpdateInboundOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update inbound o k response has a 3xx status code
func (o *UpdateInboundOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound o k response has a 4xx status code
func (o *UpdateInboundOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update inbound o k response has a 5xx status code
func (o *UpdateInboundOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound o k response a status code equal to that given
func (o *UpdateInboundOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateInboundOK) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundOK  %+v", 200, o.Payload)
}

func (o *UpdateInboundOK) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundOK  %+v", 200, o.Payload)
}

func (o *UpdateInboundOK) GetPayload() *awd_2024_05_09_models.InboundOrderReference {
	return o.Payload
}

func (o *UpdateInboundOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(awd_2024_05_09_models.InboundOrderReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundBadRequest creates a UpdateInboundBadRequest with default headers values
func NewUpdateInboundBadRequest() *UpdateInboundBadRequest {
	return &UpdateInboundBadRequest{}
}

/*
UpdateInboundBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type UpdateInboundBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound bad request response has a 2xx status code
func (o *UpdateInboundBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound bad request response has a 3xx status code
func (o *UpdateInboundBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound bad request response has a 4xx status code
func (o *UpdateInboundBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound bad request response has a 5xx status code
func (o *UpdateInboundBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound bad request response a status code equal to that given
func (o *UpdateInboundBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateInboundBadRequest) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInboundBadRequest) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateInboundBadRequest) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundForbidden creates a UpdateInboundForbidden with default headers values
func NewUpdateInboundForbidden() *UpdateInboundForbidden {
	return &UpdateInboundForbidden{}
}

/*
UpdateInboundForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type UpdateInboundForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound forbidden response has a 2xx status code
func (o *UpdateInboundForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound forbidden response has a 3xx status code
func (o *UpdateInboundForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound forbidden response has a 4xx status code
func (o *UpdateInboundForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound forbidden response has a 5xx status code
func (o *UpdateInboundForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound forbidden response a status code equal to that given
func (o *UpdateInboundForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateInboundForbidden) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInboundForbidden) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInboundForbidden) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundNotFound creates a UpdateInboundNotFound with default headers values
func NewUpdateInboundNotFound() *UpdateInboundNotFound {
	return &UpdateInboundNotFound{}
}

/*
UpdateInboundNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type UpdateInboundNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound not found response has a 2xx status code
func (o *UpdateInboundNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound not found response has a 3xx status code
func (o *UpdateInboundNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound not found response has a 4xx status code
func (o *UpdateInboundNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound not found response has a 5xx status code
func (o *UpdateInboundNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound not found response a status code equal to that given
func (o *UpdateInboundNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateInboundNotFound) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInboundNotFound) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInboundNotFound) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundConflict creates a UpdateInboundConflict with default headers values
func NewUpdateInboundConflict() *UpdateInboundConflict {
	return &UpdateInboundConflict{}
}

/*
UpdateInboundConflict describes a response with status code 409, with default header values.

`ConflictException` 409 response.
*/
type UpdateInboundConflict struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound conflict response has a 2xx status code
func (o *UpdateInboundConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound conflict response has a 3xx status code
func (o *UpdateInboundConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound conflict response has a 4xx status code
func (o *UpdateInboundConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound conflict response has a 5xx status code
func (o *UpdateInboundConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound conflict response a status code equal to that given
func (o *UpdateInboundConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateInboundConflict) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundConflict  %+v", 409, o.Payload)
}

func (o *UpdateInboundConflict) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundConflict  %+v", 409, o.Payload)
}

func (o *UpdateInboundConflict) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundRequestEntityTooLarge creates a UpdateInboundRequestEntityTooLarge with default headers values
func NewUpdateInboundRequestEntityTooLarge() *UpdateInboundRequestEntityTooLarge {
	return &UpdateInboundRequestEntityTooLarge{}
}

/*
UpdateInboundRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type UpdateInboundRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound request entity too large response has a 2xx status code
func (o *UpdateInboundRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound request entity too large response has a 3xx status code
func (o *UpdateInboundRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound request entity too large response has a 4xx status code
func (o *UpdateInboundRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound request entity too large response has a 5xx status code
func (o *UpdateInboundRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound request entity too large response a status code equal to that given
func (o *UpdateInboundRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *UpdateInboundRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateInboundRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateInboundRequestEntityTooLarge) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundUnsupportedMediaType creates a UpdateInboundUnsupportedMediaType with default headers values
func NewUpdateInboundUnsupportedMediaType() *UpdateInboundUnsupportedMediaType {
	return &UpdateInboundUnsupportedMediaType{}
}

/*
UpdateInboundUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type UpdateInboundUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound unsupported media type response has a 2xx status code
func (o *UpdateInboundUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound unsupported media type response has a 3xx status code
func (o *UpdateInboundUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound unsupported media type response has a 4xx status code
func (o *UpdateInboundUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound unsupported media type response has a 5xx status code
func (o *UpdateInboundUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound unsupported media type response a status code equal to that given
func (o *UpdateInboundUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *UpdateInboundUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateInboundUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *UpdateInboundUnsupportedMediaType) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundTooManyRequests creates a UpdateInboundTooManyRequests with default headers values
func NewUpdateInboundTooManyRequests() *UpdateInboundTooManyRequests {
	return &UpdateInboundTooManyRequests{}
}

/*
UpdateInboundTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type UpdateInboundTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound too many requests response has a 2xx status code
func (o *UpdateInboundTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound too many requests response has a 3xx status code
func (o *UpdateInboundTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound too many requests response has a 4xx status code
func (o *UpdateInboundTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update inbound too many requests response has a 5xx status code
func (o *UpdateInboundTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update inbound too many requests response a status code equal to that given
func (o *UpdateInboundTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *UpdateInboundTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateInboundTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateInboundTooManyRequests) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundInternalServerError creates a UpdateInboundInternalServerError with default headers values
func NewUpdateInboundInternalServerError() *UpdateInboundInternalServerError {
	return &UpdateInboundInternalServerError{}
}

/*
UpdateInboundInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type UpdateInboundInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound internal server error response has a 2xx status code
func (o *UpdateInboundInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound internal server error response has a 3xx status code
func (o *UpdateInboundInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound internal server error response has a 4xx status code
func (o *UpdateInboundInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update inbound internal server error response has a 5xx status code
func (o *UpdateInboundInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update inbound internal server error response a status code equal to that given
func (o *UpdateInboundInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateInboundInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInboundInternalServerError) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateInboundInternalServerError) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInboundServiceUnavailable creates a UpdateInboundServiceUnavailable with default headers values
func NewUpdateInboundServiceUnavailable() *UpdateInboundServiceUnavailable {
	return &UpdateInboundServiceUnavailable{}
}

/*
UpdateInboundServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type UpdateInboundServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this update inbound service unavailable response has a 2xx status code
func (o *UpdateInboundServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update inbound service unavailable response has a 3xx status code
func (o *UpdateInboundServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update inbound service unavailable response has a 4xx status code
func (o *UpdateInboundServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update inbound service unavailable response has a 5xx status code
func (o *UpdateInboundServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update inbound service unavailable response a status code equal to that given
func (o *UpdateInboundServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *UpdateInboundServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateInboundServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /awd/2024-05-09/inboundOrders/{orderId}][%d] updateInboundServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateInboundServiceUnavailable) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *UpdateInboundServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(awd_2024_05_09_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
