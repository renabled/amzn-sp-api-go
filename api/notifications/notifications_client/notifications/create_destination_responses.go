// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/notifications/notifications_models"
)

// CreateDestinationReader is a Reader for the CreateDestination structure.
type CreateDestinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDestinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDestinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDestinationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDestinationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDestinationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDestinationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewCreateDestinationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateDestinationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateDestinationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDestinationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCreateDestinationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDestinationOK creates a CreateDestinationOK with default headers values
func NewCreateDestinationOK() *CreateDestinationOK {
	return &CreateDestinationOK{}
}

/*
CreateDestinationOK describes a response with status code 200, with default header values.

Success.
*/
type CreateDestinationOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination o k response has a 2xx status code
func (o *CreateDestinationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create destination o k response has a 3xx status code
func (o *CreateDestinationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination o k response has a 4xx status code
func (o *CreateDestinationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create destination o k response has a 5xx status code
func (o *CreateDestinationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination o k response a status code equal to that given
func (o *CreateDestinationOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateDestinationOK) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationOK  %+v", 200, o.Payload)
}

func (o *CreateDestinationOK) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationOK  %+v", 200, o.Payload)
}

func (o *CreateDestinationOK) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationBadRequest creates a CreateDestinationBadRequest with default headers values
func NewCreateDestinationBadRequest() *CreateDestinationBadRequest {
	return &CreateDestinationBadRequest{}
}

/*
CreateDestinationBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CreateDestinationBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination bad request response has a 2xx status code
func (o *CreateDestinationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination bad request response has a 3xx status code
func (o *CreateDestinationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination bad request response has a 4xx status code
func (o *CreateDestinationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create destination bad request response has a 5xx status code
func (o *CreateDestinationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination bad request response a status code equal to that given
func (o *CreateDestinationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateDestinationBadRequest) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDestinationBadRequest) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDestinationBadRequest) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationForbidden creates a CreateDestinationForbidden with default headers values
func NewCreateDestinationForbidden() *CreateDestinationForbidden {
	return &CreateDestinationForbidden{}
}

/*
CreateDestinationForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CreateDestinationForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination forbidden response has a 2xx status code
func (o *CreateDestinationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination forbidden response has a 3xx status code
func (o *CreateDestinationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination forbidden response has a 4xx status code
func (o *CreateDestinationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create destination forbidden response has a 5xx status code
func (o *CreateDestinationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination forbidden response a status code equal to that given
func (o *CreateDestinationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateDestinationForbidden) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationForbidden  %+v", 403, o.Payload)
}

func (o *CreateDestinationForbidden) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationForbidden  %+v", 403, o.Payload)
}

func (o *CreateDestinationForbidden) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationNotFound creates a CreateDestinationNotFound with default headers values
func NewCreateDestinationNotFound() *CreateDestinationNotFound {
	return &CreateDestinationNotFound{}
}

/*
CreateDestinationNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type CreateDestinationNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination not found response has a 2xx status code
func (o *CreateDestinationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination not found response has a 3xx status code
func (o *CreateDestinationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination not found response has a 4xx status code
func (o *CreateDestinationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create destination not found response has a 5xx status code
func (o *CreateDestinationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination not found response a status code equal to that given
func (o *CreateDestinationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateDestinationNotFound) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationNotFound  %+v", 404, o.Payload)
}

func (o *CreateDestinationNotFound) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationNotFound  %+v", 404, o.Payload)
}

func (o *CreateDestinationNotFound) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationConflict creates a CreateDestinationConflict with default headers values
func NewCreateDestinationConflict() *CreateDestinationConflict {
	return &CreateDestinationConflict{}
}

/*
CreateDestinationConflict describes a response with status code 409, with default header values.

The resource specified conflicts with the current state.
*/
type CreateDestinationConflict struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination conflict response has a 2xx status code
func (o *CreateDestinationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination conflict response has a 3xx status code
func (o *CreateDestinationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination conflict response has a 4xx status code
func (o *CreateDestinationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create destination conflict response has a 5xx status code
func (o *CreateDestinationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination conflict response a status code equal to that given
func (o *CreateDestinationConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateDestinationConflict) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationConflict  %+v", 409, o.Payload)
}

func (o *CreateDestinationConflict) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationConflict  %+v", 409, o.Payload)
}

func (o *CreateDestinationConflict) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationRequestEntityTooLarge creates a CreateDestinationRequestEntityTooLarge with default headers values
func NewCreateDestinationRequestEntityTooLarge() *CreateDestinationRequestEntityTooLarge {
	return &CreateDestinationRequestEntityTooLarge{}
}

/*
CreateDestinationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type CreateDestinationRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination request entity too large response has a 2xx status code
func (o *CreateDestinationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination request entity too large response has a 3xx status code
func (o *CreateDestinationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination request entity too large response has a 4xx status code
func (o *CreateDestinationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this create destination request entity too large response has a 5xx status code
func (o *CreateDestinationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination request entity too large response a status code equal to that given
func (o *CreateDestinationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *CreateDestinationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateDestinationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *CreateDestinationRequestEntityTooLarge) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationUnsupportedMediaType creates a CreateDestinationUnsupportedMediaType with default headers values
func NewCreateDestinationUnsupportedMediaType() *CreateDestinationUnsupportedMediaType {
	return &CreateDestinationUnsupportedMediaType{}
}

/*
CreateDestinationUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type CreateDestinationUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination unsupported media type response has a 2xx status code
func (o *CreateDestinationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination unsupported media type response has a 3xx status code
func (o *CreateDestinationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination unsupported media type response has a 4xx status code
func (o *CreateDestinationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this create destination unsupported media type response has a 5xx status code
func (o *CreateDestinationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination unsupported media type response a status code equal to that given
func (o *CreateDestinationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CreateDestinationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateDestinationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateDestinationUnsupportedMediaType) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationTooManyRequests creates a CreateDestinationTooManyRequests with default headers values
func NewCreateDestinationTooManyRequests() *CreateDestinationTooManyRequests {
	return &CreateDestinationTooManyRequests{}
}

/*
CreateDestinationTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CreateDestinationTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination too many requests response has a 2xx status code
func (o *CreateDestinationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination too many requests response has a 3xx status code
func (o *CreateDestinationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination too many requests response has a 4xx status code
func (o *CreateDestinationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create destination too many requests response has a 5xx status code
func (o *CreateDestinationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create destination too many requests response a status code equal to that given
func (o *CreateDestinationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CreateDestinationTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateDestinationTooManyRequests) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateDestinationTooManyRequests) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationInternalServerError creates a CreateDestinationInternalServerError with default headers values
func NewCreateDestinationInternalServerError() *CreateDestinationInternalServerError {
	return &CreateDestinationInternalServerError{}
}

/*
CreateDestinationInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CreateDestinationInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination internal server error response has a 2xx status code
func (o *CreateDestinationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination internal server error response has a 3xx status code
func (o *CreateDestinationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination internal server error response has a 4xx status code
func (o *CreateDestinationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create destination internal server error response has a 5xx status code
func (o *CreateDestinationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create destination internal server error response a status code equal to that given
func (o *CreateDestinationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateDestinationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDestinationInternalServerError) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDestinationInternalServerError) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDestinationServiceUnavailable creates a CreateDestinationServiceUnavailable with default headers values
func NewCreateDestinationServiceUnavailable() *CreateDestinationServiceUnavailable {
	return &CreateDestinationServiceUnavailable{}
}

/*
CreateDestinationServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CreateDestinationServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.CreateDestinationResponse
}

// IsSuccess returns true when this create destination service unavailable response has a 2xx status code
func (o *CreateDestinationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create destination service unavailable response has a 3xx status code
func (o *CreateDestinationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create destination service unavailable response has a 4xx status code
func (o *CreateDestinationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this create destination service unavailable response has a 5xx status code
func (o *CreateDestinationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this create destination service unavailable response a status code equal to that given
func (o *CreateDestinationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CreateDestinationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateDestinationServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /notifications/v1/destinations][%d] createDestinationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CreateDestinationServiceUnavailable) GetPayload() *notifications_models.CreateDestinationResponse {
	return o.Payload
}

func (o *CreateDestinationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.CreateDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
