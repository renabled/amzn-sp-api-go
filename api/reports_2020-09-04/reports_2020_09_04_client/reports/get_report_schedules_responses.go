// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renabled/amzn-sp-api-go/api/reports_2020-09-04/reports_2020_09_04_models"
)

// GetReportSchedulesReader is a Reader for the GetReportSchedules structure.
type GetReportSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReportSchedulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReportSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReportSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetReportSchedulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReportSchedulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetReportSchedulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportSchedulesOK creates a GetReportSchedulesOK with default headers values
func NewGetReportSchedulesOK() *GetReportSchedulesOK {
	return &GetReportSchedulesOK{}
}

/*
GetReportSchedulesOK describes a response with status code 200, with default header values.

Success.
*/
type GetReportSchedulesOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules o k response has a 2xx status code
func (o *GetReportSchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get report schedules o k response has a 3xx status code
func (o *GetReportSchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules o k response has a 4xx status code
func (o *GetReportSchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report schedules o k response has a 5xx status code
func (o *GetReportSchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedules o k response a status code equal to that given
func (o *GetReportSchedulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetReportSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesOK  %+v", 200, o.Payload)
}

func (o *GetReportSchedulesOK) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesOK  %+v", 200, o.Payload)
}

func (o *GetReportSchedulesOK) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesBadRequest creates a GetReportSchedulesBadRequest with default headers values
func NewGetReportSchedulesBadRequest() *GetReportSchedulesBadRequest {
	return &GetReportSchedulesBadRequest{}
}

/*
GetReportSchedulesBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetReportSchedulesBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules bad request response has a 2xx status code
func (o *GetReportSchedulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules bad request response has a 3xx status code
func (o *GetReportSchedulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules bad request response has a 4xx status code
func (o *GetReportSchedulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedules bad request response has a 5xx status code
func (o *GetReportSchedulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedules bad request response a status code equal to that given
func (o *GetReportSchedulesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetReportSchedulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *GetReportSchedulesBadRequest) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *GetReportSchedulesBadRequest) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesUnauthorized creates a GetReportSchedulesUnauthorized with default headers values
func NewGetReportSchedulesUnauthorized() *GetReportSchedulesUnauthorized {
	return &GetReportSchedulesUnauthorized{}
}

/*
GetReportSchedulesUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetReportSchedulesUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules unauthorized response has a 2xx status code
func (o *GetReportSchedulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules unauthorized response has a 3xx status code
func (o *GetReportSchedulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules unauthorized response has a 4xx status code
func (o *GetReportSchedulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedules unauthorized response has a 5xx status code
func (o *GetReportSchedulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedules unauthorized response a status code equal to that given
func (o *GetReportSchedulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetReportSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReportSchedulesUnauthorized) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReportSchedulesUnauthorized) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesForbidden creates a GetReportSchedulesForbidden with default headers values
func NewGetReportSchedulesForbidden() *GetReportSchedulesForbidden {
	return &GetReportSchedulesForbidden{}
}

/*
GetReportSchedulesForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetReportSchedulesForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules forbidden response has a 2xx status code
func (o *GetReportSchedulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules forbidden response has a 3xx status code
func (o *GetReportSchedulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules forbidden response has a 4xx status code
func (o *GetReportSchedulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedules forbidden response has a 5xx status code
func (o *GetReportSchedulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedules forbidden response a status code equal to that given
func (o *GetReportSchedulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetReportSchedulesForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *GetReportSchedulesForbidden) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *GetReportSchedulesForbidden) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesNotFound creates a GetReportSchedulesNotFound with default headers values
func NewGetReportSchedulesNotFound() *GetReportSchedulesNotFound {
	return &GetReportSchedulesNotFound{}
}

/*
GetReportSchedulesNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetReportSchedulesNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules not found response has a 2xx status code
func (o *GetReportSchedulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules not found response has a 3xx status code
func (o *GetReportSchedulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules not found response has a 4xx status code
func (o *GetReportSchedulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedules not found response has a 5xx status code
func (o *GetReportSchedulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedules not found response a status code equal to that given
func (o *GetReportSchedulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetReportSchedulesNotFound) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *GetReportSchedulesNotFound) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *GetReportSchedulesNotFound) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesUnsupportedMediaType creates a GetReportSchedulesUnsupportedMediaType with default headers values
func NewGetReportSchedulesUnsupportedMediaType() *GetReportSchedulesUnsupportedMediaType {
	return &GetReportSchedulesUnsupportedMediaType{}
}

/*
GetReportSchedulesUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type GetReportSchedulesUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules unsupported media type response has a 2xx status code
func (o *GetReportSchedulesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules unsupported media type response has a 3xx status code
func (o *GetReportSchedulesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules unsupported media type response has a 4xx status code
func (o *GetReportSchedulesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedules unsupported media type response has a 5xx status code
func (o *GetReportSchedulesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedules unsupported media type response a status code equal to that given
func (o *GetReportSchedulesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetReportSchedulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetReportSchedulesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetReportSchedulesUnsupportedMediaType) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesTooManyRequests creates a GetReportSchedulesTooManyRequests with default headers values
func NewGetReportSchedulesTooManyRequests() *GetReportSchedulesTooManyRequests {
	return &GetReportSchedulesTooManyRequests{}
}

/*
GetReportSchedulesTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetReportSchedulesTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules too many requests response has a 2xx status code
func (o *GetReportSchedulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules too many requests response has a 3xx status code
func (o *GetReportSchedulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules too many requests response has a 4xx status code
func (o *GetReportSchedulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedules too many requests response has a 5xx status code
func (o *GetReportSchedulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedules too many requests response a status code equal to that given
func (o *GetReportSchedulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetReportSchedulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetReportSchedulesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetReportSchedulesTooManyRequests) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesInternalServerError creates a GetReportSchedulesInternalServerError with default headers values
func NewGetReportSchedulesInternalServerError() *GetReportSchedulesInternalServerError {
	return &GetReportSchedulesInternalServerError{}
}

/*
GetReportSchedulesInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetReportSchedulesInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules internal server error response has a 2xx status code
func (o *GetReportSchedulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules internal server error response has a 3xx status code
func (o *GetReportSchedulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules internal server error response has a 4xx status code
func (o *GetReportSchedulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report schedules internal server error response has a 5xx status code
func (o *GetReportSchedulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get report schedules internal server error response a status code equal to that given
func (o *GetReportSchedulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetReportSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReportSchedulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReportSchedulesInternalServerError) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportSchedulesServiceUnavailable creates a GetReportSchedulesServiceUnavailable with default headers values
func NewGetReportSchedulesServiceUnavailable() *GetReportSchedulesServiceUnavailable {
	return &GetReportSchedulesServiceUnavailable{}
}

/*
GetReportSchedulesServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetReportSchedulesServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportSchedulesResponse
}

// IsSuccess returns true when this get report schedules service unavailable response has a 2xx status code
func (o *GetReportSchedulesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedules service unavailable response has a 3xx status code
func (o *GetReportSchedulesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedules service unavailable response has a 4xx status code
func (o *GetReportSchedulesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report schedules service unavailable response has a 5xx status code
func (o *GetReportSchedulesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get report schedules service unavailable response a status code equal to that given
func (o *GetReportSchedulesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetReportSchedulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReportSchedulesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules][%d] getReportSchedulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReportSchedulesServiceUnavailable) GetPayload() *reports_2020_09_04_models.GetReportSchedulesResponse {
	return o.Payload
}

func (o *GetReportSchedulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportSchedulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
