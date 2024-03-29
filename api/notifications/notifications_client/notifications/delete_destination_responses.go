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

// DeleteDestinationReader is a Reader for the DeleteDestination structure.
type DeleteDestinationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDestinationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDestinationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDestinationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDestinationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteDestinationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteDestinationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteDestinationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteDestinationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteDestinationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteDestinationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteDestinationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDestinationOK creates a DeleteDestinationOK with default headers values
func NewDeleteDestinationOK() *DeleteDestinationOK {
	return &DeleteDestinationOK{}
}

/*
DeleteDestinationOK describes a response with status code 200, with default header values.

Success.
*/
type DeleteDestinationOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination o k response has a 2xx status code
func (o *DeleteDestinationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete destination o k response has a 3xx status code
func (o *DeleteDestinationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination o k response has a 4xx status code
func (o *DeleteDestinationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete destination o k response has a 5xx status code
func (o *DeleteDestinationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination o k response a status code equal to that given
func (o *DeleteDestinationOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteDestinationOK) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationOK  %+v", 200, o.Payload)
}

func (o *DeleteDestinationOK) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationOK  %+v", 200, o.Payload)
}

func (o *DeleteDestinationOK) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationBadRequest creates a DeleteDestinationBadRequest with default headers values
func NewDeleteDestinationBadRequest() *DeleteDestinationBadRequest {
	return &DeleteDestinationBadRequest{}
}

/*
DeleteDestinationBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type DeleteDestinationBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination bad request response has a 2xx status code
func (o *DeleteDestinationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination bad request response has a 3xx status code
func (o *DeleteDestinationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination bad request response has a 4xx status code
func (o *DeleteDestinationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete destination bad request response has a 5xx status code
func (o *DeleteDestinationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination bad request response a status code equal to that given
func (o *DeleteDestinationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteDestinationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDestinationBadRequest) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDestinationBadRequest) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationForbidden creates a DeleteDestinationForbidden with default headers values
func NewDeleteDestinationForbidden() *DeleteDestinationForbidden {
	return &DeleteDestinationForbidden{}
}

/*
DeleteDestinationForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type DeleteDestinationForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination forbidden response has a 2xx status code
func (o *DeleteDestinationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination forbidden response has a 3xx status code
func (o *DeleteDestinationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination forbidden response has a 4xx status code
func (o *DeleteDestinationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete destination forbidden response has a 5xx status code
func (o *DeleteDestinationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination forbidden response a status code equal to that given
func (o *DeleteDestinationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteDestinationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDestinationForbidden) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteDestinationForbidden) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationNotFound creates a DeleteDestinationNotFound with default headers values
func NewDeleteDestinationNotFound() *DeleteDestinationNotFound {
	return &DeleteDestinationNotFound{}
}

/*
DeleteDestinationNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type DeleteDestinationNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination not found response has a 2xx status code
func (o *DeleteDestinationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination not found response has a 3xx status code
func (o *DeleteDestinationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination not found response has a 4xx status code
func (o *DeleteDestinationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete destination not found response has a 5xx status code
func (o *DeleteDestinationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination not found response a status code equal to that given
func (o *DeleteDestinationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteDestinationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDestinationNotFound) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDestinationNotFound) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationConflict creates a DeleteDestinationConflict with default headers values
func NewDeleteDestinationConflict() *DeleteDestinationConflict {
	return &DeleteDestinationConflict{}
}

/*
DeleteDestinationConflict describes a response with status code 409, with default header values.

The resource specified conflicts with the current state.
*/
type DeleteDestinationConflict struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination conflict response has a 2xx status code
func (o *DeleteDestinationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination conflict response has a 3xx status code
func (o *DeleteDestinationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination conflict response has a 4xx status code
func (o *DeleteDestinationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete destination conflict response has a 5xx status code
func (o *DeleteDestinationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination conflict response a status code equal to that given
func (o *DeleteDestinationConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteDestinationConflict) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationConflict  %+v", 409, o.Payload)
}

func (o *DeleteDestinationConflict) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationConflict  %+v", 409, o.Payload)
}

func (o *DeleteDestinationConflict) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationRequestEntityTooLarge creates a DeleteDestinationRequestEntityTooLarge with default headers values
func NewDeleteDestinationRequestEntityTooLarge() *DeleteDestinationRequestEntityTooLarge {
	return &DeleteDestinationRequestEntityTooLarge{}
}

/*
DeleteDestinationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type DeleteDestinationRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination request entity too large response has a 2xx status code
func (o *DeleteDestinationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination request entity too large response has a 3xx status code
func (o *DeleteDestinationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination request entity too large response has a 4xx status code
func (o *DeleteDestinationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete destination request entity too large response has a 5xx status code
func (o *DeleteDestinationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination request entity too large response a status code equal to that given
func (o *DeleteDestinationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteDestinationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteDestinationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteDestinationRequestEntityTooLarge) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationUnsupportedMediaType creates a DeleteDestinationUnsupportedMediaType with default headers values
func NewDeleteDestinationUnsupportedMediaType() *DeleteDestinationUnsupportedMediaType {
	return &DeleteDestinationUnsupportedMediaType{}
}

/*
DeleteDestinationUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type DeleteDestinationUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination unsupported media type response has a 2xx status code
func (o *DeleteDestinationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination unsupported media type response has a 3xx status code
func (o *DeleteDestinationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination unsupported media type response has a 4xx status code
func (o *DeleteDestinationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete destination unsupported media type response has a 5xx status code
func (o *DeleteDestinationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination unsupported media type response a status code equal to that given
func (o *DeleteDestinationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteDestinationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteDestinationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteDestinationUnsupportedMediaType) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationTooManyRequests creates a DeleteDestinationTooManyRequests with default headers values
func NewDeleteDestinationTooManyRequests() *DeleteDestinationTooManyRequests {
	return &DeleteDestinationTooManyRequests{}
}

/*
DeleteDestinationTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type DeleteDestinationTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination too many requests response has a 2xx status code
func (o *DeleteDestinationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination too many requests response has a 3xx status code
func (o *DeleteDestinationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination too many requests response has a 4xx status code
func (o *DeleteDestinationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete destination too many requests response has a 5xx status code
func (o *DeleteDestinationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete destination too many requests response a status code equal to that given
func (o *DeleteDestinationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteDestinationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteDestinationTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteDestinationTooManyRequests) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationInternalServerError creates a DeleteDestinationInternalServerError with default headers values
func NewDeleteDestinationInternalServerError() *DeleteDestinationInternalServerError {
	return &DeleteDestinationInternalServerError{}
}

/*
DeleteDestinationInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type DeleteDestinationInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination internal server error response has a 2xx status code
func (o *DeleteDestinationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination internal server error response has a 3xx status code
func (o *DeleteDestinationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination internal server error response has a 4xx status code
func (o *DeleteDestinationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete destination internal server error response has a 5xx status code
func (o *DeleteDestinationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete destination internal server error response a status code equal to that given
func (o *DeleteDestinationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteDestinationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDestinationInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDestinationInternalServerError) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDestinationServiceUnavailable creates a DeleteDestinationServiceUnavailable with default headers values
func NewDeleteDestinationServiceUnavailable() *DeleteDestinationServiceUnavailable {
	return &DeleteDestinationServiceUnavailable{}
}

/*
DeleteDestinationServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type DeleteDestinationServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.DeleteDestinationResponse
}

// IsSuccess returns true when this delete destination service unavailable response has a 2xx status code
func (o *DeleteDestinationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete destination service unavailable response has a 3xx status code
func (o *DeleteDestinationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete destination service unavailable response has a 4xx status code
func (o *DeleteDestinationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete destination service unavailable response has a 5xx status code
func (o *DeleteDestinationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete destination service unavailable response a status code equal to that given
func (o *DeleteDestinationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteDestinationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteDestinationServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /notifications/v1/destinations/{destinationId}][%d] deleteDestinationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteDestinationServiceUnavailable) GetPayload() *notifications_models.DeleteDestinationResponse {
	return o.Payload
}

func (o *DeleteDestinationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.DeleteDestinationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
