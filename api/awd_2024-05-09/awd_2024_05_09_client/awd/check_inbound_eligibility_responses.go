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

// CheckInboundEligibilityReader is a Reader for the CheckInboundEligibility structure.
type CheckInboundEligibilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckInboundEligibilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckInboundEligibilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckInboundEligibilityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckInboundEligibilityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckInboundEligibilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCheckInboundEligibilityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCheckInboundEligibilityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCheckInboundEligibilityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckInboundEligibilityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCheckInboundEligibilityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckInboundEligibilityOK creates a CheckInboundEligibilityOK with default headers values
func NewCheckInboundEligibilityOK() *CheckInboundEligibilityOK {
	return &CheckInboundEligibilityOK{}
}

/*
CheckInboundEligibilityOK describes a response with status code 200, with default header values.

The 200 response for `checkInboundEligibility`.
*/
type CheckInboundEligibilityOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.InboundEligibility
}

// IsSuccess returns true when this check inbound eligibility o k response has a 2xx status code
func (o *CheckInboundEligibilityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this check inbound eligibility o k response has a 3xx status code
func (o *CheckInboundEligibilityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility o k response has a 4xx status code
func (o *CheckInboundEligibilityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this check inbound eligibility o k response has a 5xx status code
func (o *CheckInboundEligibilityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this check inbound eligibility o k response a status code equal to that given
func (o *CheckInboundEligibilityOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckInboundEligibilityOK) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityOK  %+v", 200, o.Payload)
}

func (o *CheckInboundEligibilityOK) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityOK  %+v", 200, o.Payload)
}

func (o *CheckInboundEligibilityOK) GetPayload() *awd_2024_05_09_models.InboundEligibility {
	return o.Payload
}

func (o *CheckInboundEligibilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(awd_2024_05_09_models.InboundEligibility)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckInboundEligibilityBadRequest creates a CheckInboundEligibilityBadRequest with default headers values
func NewCheckInboundEligibilityBadRequest() *CheckInboundEligibilityBadRequest {
	return &CheckInboundEligibilityBadRequest{}
}

/*
CheckInboundEligibilityBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CheckInboundEligibilityBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility bad request response has a 2xx status code
func (o *CheckInboundEligibilityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility bad request response has a 3xx status code
func (o *CheckInboundEligibilityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility bad request response has a 4xx status code
func (o *CheckInboundEligibilityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this check inbound eligibility bad request response has a 5xx status code
func (o *CheckInboundEligibilityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this check inbound eligibility bad request response a status code equal to that given
func (o *CheckInboundEligibilityBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckInboundEligibilityBadRequest) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityBadRequest  %+v", 400, o.Payload)
}

func (o *CheckInboundEligibilityBadRequest) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityBadRequest  %+v", 400, o.Payload)
}

func (o *CheckInboundEligibilityBadRequest) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCheckInboundEligibilityForbidden creates a CheckInboundEligibilityForbidden with default headers values
func NewCheckInboundEligibilityForbidden() *CheckInboundEligibilityForbidden {
	return &CheckInboundEligibilityForbidden{}
}

/*
CheckInboundEligibilityForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CheckInboundEligibilityForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility forbidden response has a 2xx status code
func (o *CheckInboundEligibilityForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility forbidden response has a 3xx status code
func (o *CheckInboundEligibilityForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility forbidden response has a 4xx status code
func (o *CheckInboundEligibilityForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this check inbound eligibility forbidden response has a 5xx status code
func (o *CheckInboundEligibilityForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this check inbound eligibility forbidden response a status code equal to that given
func (o *CheckInboundEligibilityForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckInboundEligibilityForbidden) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityForbidden  %+v", 403, o.Payload)
}

func (o *CheckInboundEligibilityForbidden) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityForbidden  %+v", 403, o.Payload)
}

func (o *CheckInboundEligibilityForbidden) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCheckInboundEligibilityNotFound creates a CheckInboundEligibilityNotFound with default headers values
func NewCheckInboundEligibilityNotFound() *CheckInboundEligibilityNotFound {
	return &CheckInboundEligibilityNotFound{}
}

/*
CheckInboundEligibilityNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CheckInboundEligibilityNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility not found response has a 2xx status code
func (o *CheckInboundEligibilityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility not found response has a 3xx status code
func (o *CheckInboundEligibilityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility not found response has a 4xx status code
func (o *CheckInboundEligibilityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this check inbound eligibility not found response has a 5xx status code
func (o *CheckInboundEligibilityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this check inbound eligibility not found response a status code equal to that given
func (o *CheckInboundEligibilityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckInboundEligibilityNotFound) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityNotFound  %+v", 404, o.Payload)
}

func (o *CheckInboundEligibilityNotFound) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityNotFound  %+v", 404, o.Payload)
}

func (o *CheckInboundEligibilityNotFound) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCheckInboundEligibilityRequestEntityTooLarge creates a CheckInboundEligibilityRequestEntityTooLarge with default headers values
func NewCheckInboundEligibilityRequestEntityTooLarge() *CheckInboundEligibilityRequestEntityTooLarge {
	return &CheckInboundEligibilityRequestEntityTooLarge{}
}

/*
CheckInboundEligibilityRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CheckInboundEligibilityRequestEntityTooLarge struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility request entity too large response has a 2xx status code
func (o *CheckInboundEligibilityRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility request entity too large response has a 3xx status code
func (o *CheckInboundEligibilityRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility request entity too large response has a 4xx status code
func (o *CheckInboundEligibilityRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this check inbound eligibility request entity too large response has a 5xx status code
func (o *CheckInboundEligibilityRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this check inbound eligibility request entity too large response a status code equal to that given
func (o *CheckInboundEligibilityRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CheckInboundEligibilityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CheckInboundEligibilityRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CheckInboundEligibilityRequestEntityTooLarge) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCheckInboundEligibilityUnsupportedMediaType creates a CheckInboundEligibilityUnsupportedMediaType with default headers values
func NewCheckInboundEligibilityUnsupportedMediaType() *CheckInboundEligibilityUnsupportedMediaType {
	return &CheckInboundEligibilityUnsupportedMediaType{}
}

/*
CheckInboundEligibilityUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CheckInboundEligibilityUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility unsupported media type response has a 2xx status code
func (o *CheckInboundEligibilityUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility unsupported media type response has a 3xx status code
func (o *CheckInboundEligibilityUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility unsupported media type response has a 4xx status code
func (o *CheckInboundEligibilityUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this check inbound eligibility unsupported media type response has a 5xx status code
func (o *CheckInboundEligibilityUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this check inbound eligibility unsupported media type response a status code equal to that given
func (o *CheckInboundEligibilityUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CheckInboundEligibilityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CheckInboundEligibilityUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CheckInboundEligibilityUnsupportedMediaType) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCheckInboundEligibilityTooManyRequests creates a CheckInboundEligibilityTooManyRequests with default headers values
func NewCheckInboundEligibilityTooManyRequests() *CheckInboundEligibilityTooManyRequests {
	return &CheckInboundEligibilityTooManyRequests{}
}

/*
CheckInboundEligibilityTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CheckInboundEligibilityTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility too many requests response has a 2xx status code
func (o *CheckInboundEligibilityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility too many requests response has a 3xx status code
func (o *CheckInboundEligibilityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility too many requests response has a 4xx status code
func (o *CheckInboundEligibilityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this check inbound eligibility too many requests response has a 5xx status code
func (o *CheckInboundEligibilityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this check inbound eligibility too many requests response a status code equal to that given
func (o *CheckInboundEligibilityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CheckInboundEligibilityTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityTooManyRequests  %+v", 429, o.Payload)
}

func (o *CheckInboundEligibilityTooManyRequests) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityTooManyRequests  %+v", 429, o.Payload)
}

func (o *CheckInboundEligibilityTooManyRequests) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCheckInboundEligibilityInternalServerError creates a CheckInboundEligibilityInternalServerError with default headers values
func NewCheckInboundEligibilityInternalServerError() *CheckInboundEligibilityInternalServerError {
	return &CheckInboundEligibilityInternalServerError{}
}

/*
CheckInboundEligibilityInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CheckInboundEligibilityInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility internal server error response has a 2xx status code
func (o *CheckInboundEligibilityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility internal server error response has a 3xx status code
func (o *CheckInboundEligibilityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility internal server error response has a 4xx status code
func (o *CheckInboundEligibilityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this check inbound eligibility internal server error response has a 5xx status code
func (o *CheckInboundEligibilityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this check inbound eligibility internal server error response a status code equal to that given
func (o *CheckInboundEligibilityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckInboundEligibilityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckInboundEligibilityInternalServerError) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityInternalServerError  %+v", 500, o.Payload)
}

func (o *CheckInboundEligibilityInternalServerError) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCheckInboundEligibilityServiceUnavailable creates a CheckInboundEligibilityServiceUnavailable with default headers values
func NewCheckInboundEligibilityServiceUnavailable() *CheckInboundEligibilityServiceUnavailable {
	return &CheckInboundEligibilityServiceUnavailable{}
}

/*
CheckInboundEligibilityServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CheckInboundEligibilityServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *awd_2024_05_09_models.ErrorList
}

// IsSuccess returns true when this check inbound eligibility service unavailable response has a 2xx status code
func (o *CheckInboundEligibilityServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this check inbound eligibility service unavailable response has a 3xx status code
func (o *CheckInboundEligibilityServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check inbound eligibility service unavailable response has a 4xx status code
func (o *CheckInboundEligibilityServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this check inbound eligibility service unavailable response has a 5xx status code
func (o *CheckInboundEligibilityServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this check inbound eligibility service unavailable response a status code equal to that given
func (o *CheckInboundEligibilityServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CheckInboundEligibilityServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CheckInboundEligibilityServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /awd/2024-05-09/inboundEligibility][%d] checkInboundEligibilityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CheckInboundEligibilityServiceUnavailable) GetPayload() *awd_2024_05_09_models.ErrorList {
	return o.Payload
}

func (o *CheckInboundEligibilityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
