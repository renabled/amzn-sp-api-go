// Code generated by go-swagger; DO NOT EDIT.

package feeds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/feeds_2021-06-30/feeds_2021_06_30_models"
)

// GetFeedReader is a Reader for the GetFeed structure.
type GetFeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFeedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFeedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFeedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFeedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFeedUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFeedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFeedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFeedServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFeedOK creates a GetFeedOK with default headers values
func NewGetFeedOK() *GetFeedOK {
	return &GetFeedOK{}
}

/*
GetFeedOK describes a response with status code 200, with default header values.

Success.
*/
type GetFeedOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.Feed
}

// IsSuccess returns true when this get feed o k response has a 2xx status code
func (o *GetFeedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get feed o k response has a 3xx status code
func (o *GetFeedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed o k response has a 4xx status code
func (o *GetFeedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get feed o k response has a 5xx status code
func (o *GetFeedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get feed o k response a status code equal to that given
func (o *GetFeedOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFeedOK) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedOK  %+v", 200, o.Payload)
}

func (o *GetFeedOK) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedOK  %+v", 200, o.Payload)
}

func (o *GetFeedOK) GetPayload() *feeds_2021_06_30_models.Feed {
	return o.Payload
}

func (o *GetFeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2021_06_30_models.Feed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedBadRequest creates a GetFeedBadRequest with default headers values
func NewGetFeedBadRequest() *GetFeedBadRequest {
	return &GetFeedBadRequest{}
}

/*
GetFeedBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetFeedBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed bad request response has a 2xx status code
func (o *GetFeedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed bad request response has a 3xx status code
func (o *GetFeedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed bad request response has a 4xx status code
func (o *GetFeedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feed bad request response has a 5xx status code
func (o *GetFeedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get feed bad request response a status code equal to that given
func (o *GetFeedBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFeedBadRequest) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedBadRequest  %+v", 400, o.Payload)
}

func (o *GetFeedBadRequest) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedBadRequest  %+v", 400, o.Payload)
}

func (o *GetFeedBadRequest) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedUnauthorized creates a GetFeedUnauthorized with default headers values
func NewGetFeedUnauthorized() *GetFeedUnauthorized {
	return &GetFeedUnauthorized{}
}

/*
GetFeedUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetFeedUnauthorized struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed unauthorized response has a 2xx status code
func (o *GetFeedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed unauthorized response has a 3xx status code
func (o *GetFeedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed unauthorized response has a 4xx status code
func (o *GetFeedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feed unauthorized response has a 5xx status code
func (o *GetFeedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get feed unauthorized response a status code equal to that given
func (o *GetFeedUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFeedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFeedUnauthorized) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFeedUnauthorized) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedForbidden creates a GetFeedForbidden with default headers values
func NewGetFeedForbidden() *GetFeedForbidden {
	return &GetFeedForbidden{}
}

/*
GetFeedForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetFeedForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed forbidden response has a 2xx status code
func (o *GetFeedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed forbidden response has a 3xx status code
func (o *GetFeedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed forbidden response has a 4xx status code
func (o *GetFeedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feed forbidden response has a 5xx status code
func (o *GetFeedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get feed forbidden response a status code equal to that given
func (o *GetFeedForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFeedForbidden) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedForbidden  %+v", 403, o.Payload)
}

func (o *GetFeedForbidden) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedForbidden  %+v", 403, o.Payload)
}

func (o *GetFeedForbidden) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedNotFound creates a GetFeedNotFound with default headers values
func NewGetFeedNotFound() *GetFeedNotFound {
	return &GetFeedNotFound{}
}

/*
GetFeedNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetFeedNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed not found response has a 2xx status code
func (o *GetFeedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed not found response has a 3xx status code
func (o *GetFeedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed not found response has a 4xx status code
func (o *GetFeedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feed not found response has a 5xx status code
func (o *GetFeedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get feed not found response a status code equal to that given
func (o *GetFeedNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFeedNotFound) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedNotFound  %+v", 404, o.Payload)
}

func (o *GetFeedNotFound) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedNotFound  %+v", 404, o.Payload)
}

func (o *GetFeedNotFound) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedUnsupportedMediaType creates a GetFeedUnsupportedMediaType with default headers values
func NewGetFeedUnsupportedMediaType() *GetFeedUnsupportedMediaType {
	return &GetFeedUnsupportedMediaType{}
}

/*
GetFeedUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type GetFeedUnsupportedMediaType struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed unsupported media type response has a 2xx status code
func (o *GetFeedUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed unsupported media type response has a 3xx status code
func (o *GetFeedUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed unsupported media type response has a 4xx status code
func (o *GetFeedUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feed unsupported media type response has a 5xx status code
func (o *GetFeedUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get feed unsupported media type response a status code equal to that given
func (o *GetFeedUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFeedUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFeedUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFeedUnsupportedMediaType) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedTooManyRequests creates a GetFeedTooManyRequests with default headers values
func NewGetFeedTooManyRequests() *GetFeedTooManyRequests {
	return &GetFeedTooManyRequests{}
}

/*
GetFeedTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetFeedTooManyRequests struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed too many requests response has a 2xx status code
func (o *GetFeedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed too many requests response has a 3xx status code
func (o *GetFeedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed too many requests response has a 4xx status code
func (o *GetFeedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feed too many requests response has a 5xx status code
func (o *GetFeedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get feed too many requests response a status code equal to that given
func (o *GetFeedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFeedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFeedTooManyRequests) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFeedTooManyRequests) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedInternalServerError creates a GetFeedInternalServerError with default headers values
func NewGetFeedInternalServerError() *GetFeedInternalServerError {
	return &GetFeedInternalServerError{}
}

/*
GetFeedInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetFeedInternalServerError struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed internal server error response has a 2xx status code
func (o *GetFeedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed internal server error response has a 3xx status code
func (o *GetFeedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed internal server error response has a 4xx status code
func (o *GetFeedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get feed internal server error response has a 5xx status code
func (o *GetFeedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get feed internal server error response a status code equal to that given
func (o *GetFeedInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFeedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFeedInternalServerError) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFeedInternalServerError) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedServiceUnavailable creates a GetFeedServiceUnavailable with default headers values
func NewGetFeedServiceUnavailable() *GetFeedServiceUnavailable {
	return &GetFeedServiceUnavailable{}
}

/*
GetFeedServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetFeedServiceUnavailable struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feed service unavailable response has a 2xx status code
func (o *GetFeedServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feed service unavailable response has a 3xx status code
func (o *GetFeedServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feed service unavailable response has a 4xx status code
func (o *GetFeedServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get feed service unavailable response has a 5xx status code
func (o *GetFeedServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get feed service unavailable response a status code equal to that given
func (o *GetFeedServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFeedServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFeedServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds/{feedId}][%d] getFeedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFeedServiceUnavailable) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(feeds_2021_06_30_models.ErrorList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
