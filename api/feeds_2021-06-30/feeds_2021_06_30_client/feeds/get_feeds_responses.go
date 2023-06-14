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

// GetFeedsReader is a Reader for the GetFeeds structure.
type GetFeedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFeedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFeedsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFeedsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFeedsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFeedsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFeedsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFeedsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFeedsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFeedsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFeedsOK creates a GetFeedsOK with default headers values
func NewGetFeedsOK() *GetFeedsOK {
	return &GetFeedsOK{}
}

/*
GetFeedsOK describes a response with status code 200, with default header values.

Success.
*/
type GetFeedsOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.GetFeedsResponse
}

// IsSuccess returns true when this get feeds o k response has a 2xx status code
func (o *GetFeedsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get feeds o k response has a 3xx status code
func (o *GetFeedsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds o k response has a 4xx status code
func (o *GetFeedsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get feeds o k response has a 5xx status code
func (o *GetFeedsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get feeds o k response a status code equal to that given
func (o *GetFeedsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFeedsOK) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsOK  %+v", 200, o.Payload)
}

func (o *GetFeedsOK) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsOK  %+v", 200, o.Payload)
}

func (o *GetFeedsOK) GetPayload() *feeds_2021_06_30_models.GetFeedsResponse {
	return o.Payload
}

func (o *GetFeedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(feeds_2021_06_30_models.GetFeedsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFeedsBadRequest creates a GetFeedsBadRequest with default headers values
func NewGetFeedsBadRequest() *GetFeedsBadRequest {
	return &GetFeedsBadRequest{}
}

/*
GetFeedsBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetFeedsBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds bad request response has a 2xx status code
func (o *GetFeedsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds bad request response has a 3xx status code
func (o *GetFeedsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds bad request response has a 4xx status code
func (o *GetFeedsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feeds bad request response has a 5xx status code
func (o *GetFeedsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get feeds bad request response a status code equal to that given
func (o *GetFeedsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFeedsBadRequest) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFeedsBadRequest) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFeedsBadRequest) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFeedsUnauthorized creates a GetFeedsUnauthorized with default headers values
func NewGetFeedsUnauthorized() *GetFeedsUnauthorized {
	return &GetFeedsUnauthorized{}
}

/*
GetFeedsUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetFeedsUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds unauthorized response has a 2xx status code
func (o *GetFeedsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds unauthorized response has a 3xx status code
func (o *GetFeedsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds unauthorized response has a 4xx status code
func (o *GetFeedsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feeds unauthorized response has a 5xx status code
func (o *GetFeedsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get feeds unauthorized response a status code equal to that given
func (o *GetFeedsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFeedsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFeedsUnauthorized) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFeedsUnauthorized) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFeedsForbidden creates a GetFeedsForbidden with default headers values
func NewGetFeedsForbidden() *GetFeedsForbidden {
	return &GetFeedsForbidden{}
}

/*
GetFeedsForbidden describes a response with status code 403, with default header values.

Indicates access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetFeedsForbidden struct {

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds forbidden response has a 2xx status code
func (o *GetFeedsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds forbidden response has a 3xx status code
func (o *GetFeedsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds forbidden response has a 4xx status code
func (o *GetFeedsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feeds forbidden response has a 5xx status code
func (o *GetFeedsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get feeds forbidden response a status code equal to that given
func (o *GetFeedsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFeedsForbidden) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsForbidden  %+v", 403, o.Payload)
}

func (o *GetFeedsForbidden) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsForbidden  %+v", 403, o.Payload)
}

func (o *GetFeedsForbidden) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFeedsNotFound creates a GetFeedsNotFound with default headers values
func NewGetFeedsNotFound() *GetFeedsNotFound {
	return &GetFeedsNotFound{}
}

/*
GetFeedsNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetFeedsNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds not found response has a 2xx status code
func (o *GetFeedsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds not found response has a 3xx status code
func (o *GetFeedsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds not found response has a 4xx status code
func (o *GetFeedsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feeds not found response has a 5xx status code
func (o *GetFeedsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get feeds not found response a status code equal to that given
func (o *GetFeedsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFeedsNotFound) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsNotFound  %+v", 404, o.Payload)
}

func (o *GetFeedsNotFound) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsNotFound  %+v", 404, o.Payload)
}

func (o *GetFeedsNotFound) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFeedsUnsupportedMediaType creates a GetFeedsUnsupportedMediaType with default headers values
func NewGetFeedsUnsupportedMediaType() *GetFeedsUnsupportedMediaType {
	return &GetFeedsUnsupportedMediaType{}
}

/*
GetFeedsUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type GetFeedsUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds unsupported media type response has a 2xx status code
func (o *GetFeedsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds unsupported media type response has a 3xx status code
func (o *GetFeedsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds unsupported media type response has a 4xx status code
func (o *GetFeedsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feeds unsupported media type response has a 5xx status code
func (o *GetFeedsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get feeds unsupported media type response a status code equal to that given
func (o *GetFeedsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFeedsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFeedsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFeedsUnsupportedMediaType) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFeedsTooManyRequests creates a GetFeedsTooManyRequests with default headers values
func NewGetFeedsTooManyRequests() *GetFeedsTooManyRequests {
	return &GetFeedsTooManyRequests{}
}

/*
GetFeedsTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetFeedsTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds too many requests response has a 2xx status code
func (o *GetFeedsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds too many requests response has a 3xx status code
func (o *GetFeedsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds too many requests response has a 4xx status code
func (o *GetFeedsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get feeds too many requests response has a 5xx status code
func (o *GetFeedsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get feeds too many requests response a status code equal to that given
func (o *GetFeedsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFeedsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFeedsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFeedsTooManyRequests) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFeedsInternalServerError creates a GetFeedsInternalServerError with default headers values
func NewGetFeedsInternalServerError() *GetFeedsInternalServerError {
	return &GetFeedsInternalServerError{}
}

/*
GetFeedsInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetFeedsInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds internal server error response has a 2xx status code
func (o *GetFeedsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds internal server error response has a 3xx status code
func (o *GetFeedsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds internal server error response has a 4xx status code
func (o *GetFeedsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get feeds internal server error response has a 5xx status code
func (o *GetFeedsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get feeds internal server error response a status code equal to that given
func (o *GetFeedsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFeedsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFeedsInternalServerError) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFeedsInternalServerError) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFeedsServiceUnavailable creates a GetFeedsServiceUnavailable with default headers values
func NewGetFeedsServiceUnavailable() *GetFeedsServiceUnavailable {
	return &GetFeedsServiceUnavailable{}
}

/*
GetFeedsServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetFeedsServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference ID.
	 */
	XAmznRequestID string

	Payload *feeds_2021_06_30_models.ErrorList
}

// IsSuccess returns true when this get feeds service unavailable response has a 2xx status code
func (o *GetFeedsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get feeds service unavailable response has a 3xx status code
func (o *GetFeedsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get feeds service unavailable response has a 4xx status code
func (o *GetFeedsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get feeds service unavailable response has a 5xx status code
func (o *GetFeedsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get feeds service unavailable response a status code equal to that given
func (o *GetFeedsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFeedsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFeedsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /feeds/2021-06-30/feeds][%d] getFeedsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFeedsServiceUnavailable) GetPayload() *feeds_2021_06_30_models.ErrorList {
	return o.Payload
}

func (o *GetFeedsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
