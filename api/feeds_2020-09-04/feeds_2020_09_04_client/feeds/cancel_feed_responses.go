// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/feeds_2020-09-04/feeds_2020_09_04_models"
)

// CancelFeedReader is a Reader for the CancelFeed structure.
type CancelFeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelFeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelFeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelFeedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelFeedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelFeedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelFeedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCancelFeedUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCancelFeedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelFeedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCancelFeedServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCancelFeedOK creates a CancelFeedOK with default headers values
func NewCancelFeedOK() *CancelFeedOK {
	return &CancelFeedOK{}
}

/*
CancelFeedOK describes a response with status code 200, with default header values.

Success.
*/
type CancelFeedOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed o k response has a 2xx status code
func (o *CancelFeedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel feed o k response has a 3xx status code
func (o *CancelFeedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed o k response has a 4xx status code
func (o *CancelFeedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel feed o k response has a 5xx status code
func (o *CancelFeedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel feed o k response a status code equal to that given
func (o *CancelFeedOK) IsCode(code int) bool {
	return code == 200
}

func (o *CancelFeedOK) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedOK  %+v", 200, o.Payload)
}

func (o *CancelFeedOK) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedOK  %+v", 200, o.Payload)
}

func (o *CancelFeedOK) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedBadRequest creates a CancelFeedBadRequest with default headers values
func NewCancelFeedBadRequest() *CancelFeedBadRequest {
	return &CancelFeedBadRequest{}
}

/*
CancelFeedBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type CancelFeedBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed bad request response has a 2xx status code
func (o *CancelFeedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed bad request response has a 3xx status code
func (o *CancelFeedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed bad request response has a 4xx status code
func (o *CancelFeedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel feed bad request response has a 5xx status code
func (o *CancelFeedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel feed bad request response a status code equal to that given
func (o *CancelFeedBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CancelFeedBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedBadRequest  %+v", 400, o.Payload)
}

func (o *CancelFeedBadRequest) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedBadRequest  %+v", 400, o.Payload)
}

func (o *CancelFeedBadRequest) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedUnauthorized creates a CancelFeedUnauthorized with default headers values
func NewCancelFeedUnauthorized() *CancelFeedUnauthorized {
	return &CancelFeedUnauthorized{}
}

/*
CancelFeedUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type CancelFeedUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed unauthorized response has a 2xx status code
func (o *CancelFeedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed unauthorized response has a 3xx status code
func (o *CancelFeedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed unauthorized response has a 4xx status code
func (o *CancelFeedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel feed unauthorized response has a 5xx status code
func (o *CancelFeedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel feed unauthorized response a status code equal to that given
func (o *CancelFeedUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CancelFeedUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelFeedUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelFeedUnauthorized) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedForbidden creates a CancelFeedForbidden with default headers values
func NewCancelFeedForbidden() *CancelFeedForbidden {
	return &CancelFeedForbidden{}
}

/*
CancelFeedForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type CancelFeedForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed forbidden response has a 2xx status code
func (o *CancelFeedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed forbidden response has a 3xx status code
func (o *CancelFeedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed forbidden response has a 4xx status code
func (o *CancelFeedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel feed forbidden response has a 5xx status code
func (o *CancelFeedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel feed forbidden response a status code equal to that given
func (o *CancelFeedForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CancelFeedForbidden) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedForbidden  %+v", 403, o.Payload)
}

func (o *CancelFeedForbidden) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedForbidden  %+v", 403, o.Payload)
}

func (o *CancelFeedForbidden) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedNotFound creates a CancelFeedNotFound with default headers values
func NewCancelFeedNotFound() *CancelFeedNotFound {
	return &CancelFeedNotFound{}
}

/*
CancelFeedNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type CancelFeedNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed not found response has a 2xx status code
func (o *CancelFeedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed not found response has a 3xx status code
func (o *CancelFeedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed not found response has a 4xx status code
func (o *CancelFeedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel feed not found response has a 5xx status code
func (o *CancelFeedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel feed not found response a status code equal to that given
func (o *CancelFeedNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CancelFeedNotFound) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedNotFound  %+v", 404, o.Payload)
}

func (o *CancelFeedNotFound) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedNotFound  %+v", 404, o.Payload)
}

func (o *CancelFeedNotFound) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedUnsupportedMediaType creates a CancelFeedUnsupportedMediaType with default headers values
func NewCancelFeedUnsupportedMediaType() *CancelFeedUnsupportedMediaType {
	return &CancelFeedUnsupportedMediaType{}
}

/*
CancelFeedUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type CancelFeedUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed unsupported media type response has a 2xx status code
func (o *CancelFeedUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed unsupported media type response has a 3xx status code
func (o *CancelFeedUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed unsupported media type response has a 4xx status code
func (o *CancelFeedUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel feed unsupported media type response has a 5xx status code
func (o *CancelFeedUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel feed unsupported media type response a status code equal to that given
func (o *CancelFeedUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *CancelFeedUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelFeedUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CancelFeedUnsupportedMediaType) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedTooManyRequests creates a CancelFeedTooManyRequests with default headers values
func NewCancelFeedTooManyRequests() *CancelFeedTooManyRequests {
	return &CancelFeedTooManyRequests{}
}

/*
CancelFeedTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type CancelFeedTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed too many requests response has a 2xx status code
func (o *CancelFeedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed too many requests response has a 3xx status code
func (o *CancelFeedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed too many requests response has a 4xx status code
func (o *CancelFeedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cancel feed too many requests response has a 5xx status code
func (o *CancelFeedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel feed too many requests response a status code equal to that given
func (o *CancelFeedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *CancelFeedTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelFeedTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedTooManyRequests  %+v", 429, o.Payload)
}

func (o *CancelFeedTooManyRequests) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedInternalServerError creates a CancelFeedInternalServerError with default headers values
func NewCancelFeedInternalServerError() *CancelFeedInternalServerError {
	return &CancelFeedInternalServerError{}
}

/*
CancelFeedInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type CancelFeedInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed internal server error response has a 2xx status code
func (o *CancelFeedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed internal server error response has a 3xx status code
func (o *CancelFeedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed internal server error response has a 4xx status code
func (o *CancelFeedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel feed internal server error response has a 5xx status code
func (o *CancelFeedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel feed internal server error response a status code equal to that given
func (o *CancelFeedInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CancelFeedInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelFeedInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelFeedInternalServerError) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelFeedServiceUnavailable creates a CancelFeedServiceUnavailable with default headers values
func NewCancelFeedServiceUnavailable() *CancelFeedServiceUnavailable {
	return &CancelFeedServiceUnavailable{}
}

/*
CancelFeedServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type CancelFeedServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2020_09_04_models.CancelFeedResponse
}

// IsSuccess returns true when this cancel feed service unavailable response has a 2xx status code
func (o *CancelFeedServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cancel feed service unavailable response has a 3xx status code
func (o *CancelFeedServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel feed service unavailable response has a 4xx status code
func (o *CancelFeedServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel feed service unavailable response has a 5xx status code
func (o *CancelFeedServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this cancel feed service unavailable response a status code equal to that given
func (o *CancelFeedServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *CancelFeedServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelFeedServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /feeds/2020-09-04/feeds/{feedId}][%d] cancelFeedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *CancelFeedServiceUnavailable) GetPayload() *feeds_2020_09_04_models.CancelFeedResponse {
	return o.Payload
}

func (o *CancelFeedServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2020_09_04_models.CancelFeedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
