// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xamandar/amzn-sp-api-go/api/reports_2020-09-04/reports_2020_09_04_models"
)

// GetReportScheduleReader is a Reader for the GetReportSchedule structure.
type GetReportScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReportScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReportScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReportScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetReportScheduleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetReportScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetReportScheduleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportScheduleOK creates a GetReportScheduleOK with default headers values
func NewGetReportScheduleOK() *GetReportScheduleOK {
	return &GetReportScheduleOK{}
}

/*
GetReportScheduleOK describes a response with status code 200, with default header values.

Success.
*/
type GetReportScheduleOK struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule o k response has a 2xx status code
func (o *GetReportScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get report schedule o k response has a 3xx status code
func (o *GetReportScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule o k response has a 4xx status code
func (o *GetReportScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report schedule o k response has a 5xx status code
func (o *GetReportScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedule o k response a status code equal to that given
func (o *GetReportScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetReportScheduleOK) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleOK  %+v", 200, o.Payload)
}

func (o *GetReportScheduleOK) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleOK  %+v", 200, o.Payload)
}

func (o *GetReportScheduleOK) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleBadRequest creates a GetReportScheduleBadRequest with default headers values
func NewGetReportScheduleBadRequest() *GetReportScheduleBadRequest {
	return &GetReportScheduleBadRequest{}
}

/*
GetReportScheduleBadRequest describes a response with status code 400, with default header values.

Request has missing or invalid parameters and cannot be parsed.
*/
type GetReportScheduleBadRequest struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule bad request response has a 2xx status code
func (o *GetReportScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule bad request response has a 3xx status code
func (o *GetReportScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule bad request response has a 4xx status code
func (o *GetReportScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedule bad request response has a 5xx status code
func (o *GetReportScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedule bad request response a status code equal to that given
func (o *GetReportScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetReportScheduleBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetReportScheduleBadRequest) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *GetReportScheduleBadRequest) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleUnauthorized creates a GetReportScheduleUnauthorized with default headers values
func NewGetReportScheduleUnauthorized() *GetReportScheduleUnauthorized {
	return &GetReportScheduleUnauthorized{}
}

/*
GetReportScheduleUnauthorized describes a response with status code 401, with default header values.

The request's Authorization header is not formatted correctly or does not contain a valid token.
*/
type GetReportScheduleUnauthorized struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule unauthorized response has a 2xx status code
func (o *GetReportScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule unauthorized response has a 3xx status code
func (o *GetReportScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule unauthorized response has a 4xx status code
func (o *GetReportScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedule unauthorized response has a 5xx status code
func (o *GetReportScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedule unauthorized response a status code equal to that given
func (o *GetReportScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetReportScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReportScheduleUnauthorized) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReportScheduleUnauthorized) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleForbidden creates a GetReportScheduleForbidden with default headers values
func NewGetReportScheduleForbidden() *GetReportScheduleForbidden {
	return &GetReportScheduleForbidden{}
}

/*
GetReportScheduleForbidden describes a response with status code 403, with default header values.

Indicates that access to the resource is forbidden. Possible reasons include Access Denied, Unauthorized, Expired Token, or Invalid Signature.
*/
type GetReportScheduleForbidden struct {

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule forbidden response has a 2xx status code
func (o *GetReportScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule forbidden response has a 3xx status code
func (o *GetReportScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule forbidden response has a 4xx status code
func (o *GetReportScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedule forbidden response has a 5xx status code
func (o *GetReportScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedule forbidden response a status code equal to that given
func (o *GetReportScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetReportScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetReportScheduleForbidden) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetReportScheduleForbidden) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-amzn-RequestId
	hdrXAmznRequestID := response.GetHeader("x-amzn-RequestId")

	if hdrXAmznRequestID != "" {
		o.XAmznRequestID = hdrXAmznRequestID
	}

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleNotFound creates a GetReportScheduleNotFound with default headers values
func NewGetReportScheduleNotFound() *GetReportScheduleNotFound {
	return &GetReportScheduleNotFound{}
}

/*
GetReportScheduleNotFound describes a response with status code 404, with default header values.

The specified resource does not exist.
*/
type GetReportScheduleNotFound struct {

	/* Your rate limit (requests per second) for this operation.
	 */
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule not found response has a 2xx status code
func (o *GetReportScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule not found response has a 3xx status code
func (o *GetReportScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule not found response has a 4xx status code
func (o *GetReportScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedule not found response has a 5xx status code
func (o *GetReportScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedule not found response a status code equal to that given
func (o *GetReportScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetReportScheduleNotFound) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetReportScheduleNotFound) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleNotFound  %+v", 404, o.Payload)
}

func (o *GetReportScheduleNotFound) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleUnsupportedMediaType creates a GetReportScheduleUnsupportedMediaType with default headers values
func NewGetReportScheduleUnsupportedMediaType() *GetReportScheduleUnsupportedMediaType {
	return &GetReportScheduleUnsupportedMediaType{}
}

/*
GetReportScheduleUnsupportedMediaType describes a response with status code 415, with default header values.

The request's Content-Type header is invalid.
*/
type GetReportScheduleUnsupportedMediaType struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule unsupported media type response has a 2xx status code
func (o *GetReportScheduleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule unsupported media type response has a 3xx status code
func (o *GetReportScheduleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule unsupported media type response has a 4xx status code
func (o *GetReportScheduleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedule unsupported media type response has a 5xx status code
func (o *GetReportScheduleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedule unsupported media type response a status code equal to that given
func (o *GetReportScheduleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetReportScheduleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetReportScheduleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetReportScheduleUnsupportedMediaType) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleTooManyRequests creates a GetReportScheduleTooManyRequests with default headers values
func NewGetReportScheduleTooManyRequests() *GetReportScheduleTooManyRequests {
	return &GetReportScheduleTooManyRequests{}
}

/*
GetReportScheduleTooManyRequests describes a response with status code 429, with default header values.

The frequency of requests was greater than allowed.
*/
type GetReportScheduleTooManyRequests struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule too many requests response has a 2xx status code
func (o *GetReportScheduleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule too many requests response has a 3xx status code
func (o *GetReportScheduleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule too many requests response has a 4xx status code
func (o *GetReportScheduleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get report schedule too many requests response has a 5xx status code
func (o *GetReportScheduleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get report schedule too many requests response a status code equal to that given
func (o *GetReportScheduleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetReportScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetReportScheduleTooManyRequests) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetReportScheduleTooManyRequests) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleInternalServerError creates a GetReportScheduleInternalServerError with default headers values
func NewGetReportScheduleInternalServerError() *GetReportScheduleInternalServerError {
	return &GetReportScheduleInternalServerError{}
}

/*
GetReportScheduleInternalServerError describes a response with status code 500, with default header values.

An unexpected condition occurred that prevented the server from fulfilling the request.
*/
type GetReportScheduleInternalServerError struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule internal server error response has a 2xx status code
func (o *GetReportScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule internal server error response has a 3xx status code
func (o *GetReportScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule internal server error response has a 4xx status code
func (o *GetReportScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report schedule internal server error response has a 5xx status code
func (o *GetReportScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get report schedule internal server error response a status code equal to that given
func (o *GetReportScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetReportScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReportScheduleInternalServerError) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReportScheduleInternalServerError) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportScheduleServiceUnavailable creates a GetReportScheduleServiceUnavailable with default headers values
func NewGetReportScheduleServiceUnavailable() *GetReportScheduleServiceUnavailable {
	return &GetReportScheduleServiceUnavailable{}
}

/*
GetReportScheduleServiceUnavailable describes a response with status code 503, with default header values.

Temporary overloading or maintenance of the server.
*/
type GetReportScheduleServiceUnavailable struct {

	/* Your rate limit (requests per second) for this operation.
	_Note:_ For this status code, the rate limit header is deprecated and no longer returned.
	*/
	XAmznRateLimitLimit string

	/* Unique request reference identifier.
	 */
	XAmznRequestID string

	Payload *reports_2020_09_04_models.GetReportScheduleResponse
}

// IsSuccess returns true when this get report schedule service unavailable response has a 2xx status code
func (o *GetReportScheduleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get report schedule service unavailable response has a 3xx status code
func (o *GetReportScheduleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get report schedule service unavailable response has a 4xx status code
func (o *GetReportScheduleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get report schedule service unavailable response has a 5xx status code
func (o *GetReportScheduleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get report schedule service unavailable response a status code equal to that given
func (o *GetReportScheduleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetReportScheduleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReportScheduleServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /reports/2020-09-04/schedules/{reportScheduleId}][%d] getReportScheduleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetReportScheduleServiceUnavailable) GetPayload() *reports_2020_09_04_models.GetReportScheduleResponse {
	return o.Payload
}

func (o *GetReportScheduleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(reports_2020_09_04_models.GetReportScheduleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
