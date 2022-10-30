// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/notifications/notifications_models"
)

// GetDestinationsReader is a Reader for the GetDestinations structure.
type GetDestinationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDestinationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDestinationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDestinationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDestinationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDestinationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetDestinationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetDestinationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDestinationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDestinationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDestinationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDestinationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDestinationsOK creates a GetDestinationsOK with default headers values
func NewGetDestinationsOK() *GetDestinationsOK {
	return &GetDestinationsOK{}
}

/*
GetDestinationsOK describes a response with status code 200, with default header values.

Success.
*/
type GetDestinationsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations o k response has a 2xx status code
func (o *GetDestinationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get destinations o k response has a 3xx status code
func (o *GetDestinationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations o k response has a 4xx status code
func (o *GetDestinationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get destinations o k response has a 5xx status code
func (o *GetDestinationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations o k response a status code equal to that given
func (o *GetDestinationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDestinationsOK) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsOK  %+v", 200, o.Payload)
}

func (o *GetDestinationsOK) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsOK  %+v", 200, o.Payload)
}

func (o *GetDestinationsOK) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsBadRequest creates a GetDestinationsBadRequest with default headers values
func NewGetDestinationsBadRequest() *GetDestinationsBadRequest {
	return &GetDestinationsBadRequest{}
}

/*
GetDestinationsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetDestinationsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations bad request response has a 2xx status code
func (o *GetDestinationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations bad request response has a 3xx status code
func (o *GetDestinationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations bad request response has a 4xx status code
func (o *GetDestinationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get destinations bad request response has a 5xx status code
func (o *GetDestinationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations bad request response a status code equal to that given
func (o *GetDestinationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDestinationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetDestinationsBadRequest) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetDestinationsBadRequest) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsForbidden creates a GetDestinationsForbidden with default headers values
func NewGetDestinationsForbidden() *GetDestinationsForbidden {
	return &GetDestinationsForbidden{}
}

/*
GetDestinationsForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetDestinationsForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations forbidden response has a 2xx status code
func (o *GetDestinationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations forbidden response has a 3xx status code
func (o *GetDestinationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations forbidden response has a 4xx status code
func (o *GetDestinationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get destinations forbidden response has a 5xx status code
func (o *GetDestinationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations forbidden response a status code equal to that given
func (o *GetDestinationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDestinationsForbidden) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsForbidden  %+v", 403, o.Payload)
}

func (o *GetDestinationsForbidden) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsForbidden  %+v", 403, o.Payload)
}

func (o *GetDestinationsForbidden) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsNotFound creates a GetDestinationsNotFound with default headers values
func NewGetDestinationsNotFound() *GetDestinationsNotFound {
	return &GetDestinationsNotFound{}
}

/*
GetDestinationsNotFound describes a response with status code 404, with default header values.

The resource specified does not exist.
*/
type GetDestinationsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations not found response has a 2xx status code
func (o *GetDestinationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations not found response has a 3xx status code
func (o *GetDestinationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations not found response has a 4xx status code
func (o *GetDestinationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get destinations not found response has a 5xx status code
func (o *GetDestinationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations not found response a status code equal to that given
func (o *GetDestinationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDestinationsNotFound) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsNotFound  %+v", 404, o.Payload)
}

func (o *GetDestinationsNotFound) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsNotFound  %+v", 404, o.Payload)
}

func (o *GetDestinationsNotFound) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsConflict creates a GetDestinationsConflict with default headers values
func NewGetDestinationsConflict() *GetDestinationsConflict {
	return &GetDestinationsConflict{}
}

/*
GetDestinationsConflict describes a response with status code 409, with default header values.

The resource specified conflicts with the current state.
*/
type GetDestinationsConflict struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations conflict response has a 2xx status code
func (o *GetDestinationsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations conflict response has a 3xx status code
func (o *GetDestinationsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations conflict response has a 4xx status code
func (o *GetDestinationsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get destinations conflict response has a 5xx status code
func (o *GetDestinationsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations conflict response a status code equal to that given
func (o *GetDestinationsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetDestinationsConflict) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsConflict  %+v", 409, o.Payload)
}

func (o *GetDestinationsConflict) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsConflict  %+v", 409, o.Payload)
}

func (o *GetDestinationsConflict) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsRequestEntityTooLarge creates a GetDestinationsRequestEntityTooLarge with default headers values
func NewGetDestinationsRequestEntityTooLarge() *GetDestinationsRequestEntityTooLarge {
	return &GetDestinationsRequestEntityTooLarge{}
}

/*
GetDestinationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request size exceeded the maximum accepted size.
*/
type GetDestinationsRequestEntityTooLarge struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations request entity too large response has a 2xx status code
func (o *GetDestinationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations request entity too large response has a 3xx status code
func (o *GetDestinationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations request entity too large response has a 4xx status code
func (o *GetDestinationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get destinations request entity too large response has a 5xx status code
func (o *GetDestinationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations request entity too large response a status code equal to that given
func (o *GetDestinationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetDestinationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDestinationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDestinationsRequestEntityTooLarge) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsUnsupportedMediaType creates a GetDestinationsUnsupportedMediaType with default headers values
func NewGetDestinationsUnsupportedMediaType() *GetDestinationsUnsupportedMediaType {
	return &GetDestinationsUnsupportedMediaType{}
}

/*
GetDestinationsUnsupportedMediaType describes a response with status code 415, with default header values.

The request payload is in an unsupported format.
*/
type GetDestinationsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations unsupported media type response has a 2xx status code
func (o *GetDestinationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations unsupported media type response has a 3xx status code
func (o *GetDestinationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations unsupported media type response has a 4xx status code
func (o *GetDestinationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get destinations unsupported media type response has a 5xx status code
func (o *GetDestinationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations unsupported media type response a status code equal to that given
func (o *GetDestinationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetDestinationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDestinationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDestinationsUnsupportedMediaType) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsTooManyRequests creates a GetDestinationsTooManyRequests with default headers values
func NewGetDestinationsTooManyRequests() *GetDestinationsTooManyRequests {
	return &GetDestinationsTooManyRequests{}
}

/*
GetDestinationsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetDestinationsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations too many requests response has a 2xx status code
func (o *GetDestinationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations too many requests response has a 3xx status code
func (o *GetDestinationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations too many requests response has a 4xx status code
func (o *GetDestinationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get destinations too many requests response has a 5xx status code
func (o *GetDestinationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get destinations too many requests response a status code equal to that given
func (o *GetDestinationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetDestinationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDestinationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDestinationsTooManyRequests) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsInternalServerError creates a GetDestinationsInternalServerError with default headers values
func NewGetDestinationsInternalServerError() *GetDestinationsInternalServerError {
	return &GetDestinationsInternalServerError{}
}

/*
GetDestinationsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetDestinationsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations internal server error response has a 2xx status code
func (o *GetDestinationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations internal server error response has a 3xx status code
func (o *GetDestinationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations internal server error response has a 4xx status code
func (o *GetDestinationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get destinations internal server error response has a 5xx status code
func (o *GetDestinationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get destinations internal server error response a status code equal to that given
func (o *GetDestinationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDestinationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDestinationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDestinationsInternalServerError) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDestinationsServiceUnavailable creates a GetDestinationsServiceUnavailable with default headers values
func NewGetDestinationsServiceUnavailable() *GetDestinationsServiceUnavailable {
	return &GetDestinationsServiceUnavailable{}
}

/*
GetDestinationsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetDestinationsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *notifications_models.GetDestinationsResponse
}

// IsSuccess returns true when this get destinations service unavailable response has a 2xx status code
func (o *GetDestinationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get destinations service unavailable response has a 3xx status code
func (o *GetDestinationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get destinations service unavailable response has a 4xx status code
func (o *GetDestinationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get destinations service unavailable response has a 5xx status code
func (o *GetDestinationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get destinations service unavailable response a status code equal to that given
func (o *GetDestinationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetDestinationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDestinationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /notifications/v1/destinations][%d] getDestinationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDestinationsServiceUnavailable) GetPayload() *notifications_models.GetDestinationsResponse {
	return o.Payload
}

func (o *GetDestinationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(notifications_models.GetDestinationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
